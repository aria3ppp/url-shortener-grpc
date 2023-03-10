package gapi_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/aria3ppp/url-shortener-grpc/internal/core/usecase"
	"github.com/aria3ppp/url-shortener-grpc/internal/gapi"
	"github.com/aria3ppp/url-shortener-grpc/internal/generator"
	"github.com/aria3ppp/url-shortener-grpc/internal/pb"
	"github.com/aria3ppp/url-shortener-grpc/internal/repository"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

func setup(
	ctx context.Context,
) (client pb.UrlShortenerClient, teardown func()) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panicf(
			"gapi_test.setup: postgres.WithInstance error: %s",
			err,
		)
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Panicf(
			"gapi_test.setup: migrate.NewWithDatabaseInstance error: %s",
			err,
		)
	}

	// run migrations
	err = migrator.Up()
	if err != nil {
		log.Panicf("gapi_test.setup: migrator.Up error: %s", err)
	}

	repository := repository.NewRepository(db)
	generator := generator.NewRandomStringGenerator(6)
	serviceUseCases := usecase.NewService(repository, generator)

	listener := bufconn.Listen(101024 * 1024)

	server := grpc.NewServer()
	pb.RegisterUrlShortenerServer(server, gapi.NewServer(serviceUseCases))

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal("gapi_test.setup: failed to serve server:", err)
		}
	}()

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return listener.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("gapi_test.setup: failed to dial server:", err)
	}

	// prepare teardown
	teardown = func() {
		// drop migrations
		if err = migrator.Drop(); err != nil {
			log.Panicf("gapi_test.teardown: migrator.Drop error: %s", err)
		}
		// close connection
		if err := conn.Close(); err != nil {
			log.Fatal("gapi_test.teardown: failed to close connection")
		}
		// close listener
		if err := listener.Close(); err != nil {
			log.Fatal("gapi_test.teardown: failed to close listener:", err)
		}
		// stop server
		server.Stop()
	}

	return pb.NewUrlShortenerClient(conn), teardown
}

var db *sql.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf(
			"gapi_test.TestMain: could not connect to database %q: %s",
			dsn,
			err,
		)
	}
	err = db.Ping()
	if err != nil {
		log.Panicf(
			"gapi_test.TestMain: could not ping database %q: %s",
			dsn,
			err,
		)
	}

	// Run tests
	code := m.Run()

	// close db
	err = db.Close()
	if err != nil {
		log.Panicf("gapi_test.TestMain: db.Close error: %s", err)
	}

	os.Exit(code)
}

func TestGetLink(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	client, teardown := setup(ctx)
	t.Cleanup(teardown)

	// create helper user
	createUserRequest := &pb.CreateUserRequest{
		Username: "username",
		Password: "password",
	}
	_, err := client.CreateUser(ctx, createUserRequest)
	require.NoError(err)

	linkShortenedString := "LaLiLuLeLo"
	url := "https://example.com"

	// invalid link
	getLinkResponse, err := client.GetLink(ctx, &pb.GetLinkRequest{})
	require.Equal(
		status.Error(codes.InvalidArgument, validation.Errors{
			"shortened_string": validation.ErrRequired,
		}.Error()),
		err,
	)
	require.Nil(getLinkResponse)

	// link not found
	getLinkResponse, err = client.GetLink(
		ctx,
		&pb.GetLinkRequest{ShortenedString: linkShortenedString},
	)
	require.Equal(status.Error(codes.NotFound, "link not found"), err)
	require.Nil(getLinkResponse)

	// create a new link
	createLinkResponse, err := client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{ShortenedString: linkShortenedString, Url: url},
	)
	require.NoError(err)
	require.True(
		proto.Equal(
			&pb.CreateLinkResponse{ShortenedString: linkShortenedString},
			createLinkResponse,
		),
	)

	// get link
	getLinkResponse, err = client.GetLink(
		ctx,
		&pb.GetLinkRequest{ShortenedString: linkShortenedString},
	)
	require.NoError(err)
	require.True(
		proto.Equal(&pb.GetLinkResponse{Url: url}, getLinkResponse),
	)
}

func TestCreateLink(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	client, teardown := setup(ctx)
	t.Cleanup(teardown)

	// create helper user
	createUserRequest := &pb.CreateUserRequest{
		Username: "username",
		Password: "password",
	}
	_, err := client.CreateUser(ctx, createUserRequest)
	require.NoError(err)

	linkShortenedString := "LaLiLuLeLo"
	url := "https://example.com"

	// unauthenticated - invallid username and password
	createLinkResponse, err := client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", "",
				"password", "",
			),
		),
		&pb.CreateLinkRequest{Url: url, ShortenedString: linkShortenedString},
	)
	require.Equal(
		status.Error(codes.Unauthenticated, validation.Errors{
			"username": validation.ErrRequired,
			"password": validation.ErrRequired,
		}.Error()),
		err,
	)
	require.Nil(createLinkResponse)

	// invalid url and shortened string
	createLinkResponse, err = client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{},
	)
	require.Equal(
		status.Error(codes.InvalidArgument, validation.Errors{
			"url": validation.ErrRequired,
		}.Error()),
		err,
	)
	require.Nil(createLinkResponse)

	// unauthenticated - user not found
	createLinkResponse, err = client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", "undefined_username",
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{Url: url, ShortenedString: linkShortenedString},
	)
	require.Equal(
		status.Error(codes.Unauthenticated, "invalid username or password"),
		err,
	)
	require.Nil(createLinkResponse)

	// unauthenticated - incorrect password
	createLinkResponse, err = client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", "incorrect_password",
			),
		),
		&pb.CreateLinkRequest{Url: url, ShortenedString: linkShortenedString},
	)
	require.Equal(
		status.Error(codes.Unauthenticated, "invalid username or password"),
		err,
	)
	require.Nil(createLinkResponse)

	// create link
	createLinkResponse, err = client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{ShortenedString: linkShortenedString, Url: url},
	)
	require.NoError(err)
	require.True(
		proto.Equal(
			&pb.CreateLinkResponse{ShortenedString: linkShortenedString},
			createLinkResponse,
		),
	)

	// ensure link is created
	getLinkResponse, err := client.GetLink(
		ctx,
		&pb.GetLinkRequest{ShortenedString: linkShortenedString},
	)
	require.NoError(err)
	require.True(
		proto.Equal(&pb.GetLinkResponse{Url: url}, getLinkResponse),
	)

	// shortened string have used
	createLinkResponse, err = client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{Url: url, ShortenedString: linkShortenedString},
	)
	require.Equal(
		status.Error(codes.AlreadyExists, "shortened string have used"),
		err,
	)
	require.Nil(createLinkResponse)
}

func TestGetLinkUser(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	client, teardown := setup(ctx)
	t.Cleanup(teardown)

	// create helper user
	createUserRequest := &pb.CreateUserRequest{
		Username: "username",
		Password: "password",
	}
	_, err := client.CreateUser(ctx, createUserRequest)
	require.NoError(err)

	linkShortenedString := "LaLiLuLeLo"
	url := "https://example.com"

	// invalid link
	getLinkUserResponse, err := client.GetLinkUser(
		ctx,
		&pb.GetLinkUserRequest{},
	)
	require.Equal(
		status.Error(codes.InvalidArgument, validation.Errors{
			"shortened_string": validation.ErrRequired,
		}.Error()),
		err,
	)
	require.Nil(getLinkUserResponse)

	// link not found
	getLinkUserResponse, err = client.GetLinkUser(
		ctx,
		&pb.GetLinkUserRequest{ShortenedString: linkShortenedString},
	)
	require.Equal(status.Error(codes.NotFound, "link not found"), err)
	require.Nil(getLinkUserResponse)

	// create a new link
	createLinkResponse, err := client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{ShortenedString: linkShortenedString, Url: url},
	)
	require.NoError(err)
	require.True(
		proto.Equal(
			&pb.CreateLinkResponse{ShortenedString: linkShortenedString},
			createLinkResponse,
		),
	)

	// get link's user
	getLinkUserResponse, err = client.GetLinkUser(
		ctx,
		&pb.GetLinkUserRequest{ShortenedString: linkShortenedString},
	)
	require.NoError(err)
	require.True(
		proto.Equal(
			&pb.GetLinkUserResponse{Username: createUserRequest.Username},
			getLinkUserResponse,
		),
	)
}

func TestCreateUser(t *testing.T) {
	require := require.New(t)
	ctx := context.Background()

	client, teardown := setup(ctx)
	t.Cleanup(teardown)

	// invalid username and password
	createUserResponse, err := client.CreateUser(ctx, &pb.CreateUserRequest{})
	require.Equal(
		status.Error(codes.InvalidArgument, validation.Errors{
			"username": validation.ErrRequired,
			"password": validation.ErrRequired,
		}.Error()),
		err,
	)
	require.Nil(createUserResponse)

	// create user
	createUserRequest := &pb.CreateUserRequest{
		Username: "username",
		Password: "password",
	}
	_, err = client.CreateUser(ctx, createUserRequest)
	require.NoError(err)

	// create a link
	linkShortenedString := "LaLiLuLeLo"
	url := "https://example.com"

	createLinkResponse, err := client.CreateLink(
		metadata.NewOutgoingContext(
			ctx,
			metadata.Pairs(
				"username", createUserRequest.Username,
				"password", createUserRequest.Password,
			),
		),
		&pb.CreateLinkRequest{ShortenedString: linkShortenedString, Url: url},
	)
	require.NoError(err)
	require.True(
		proto.Equal(
			&pb.CreateLinkResponse{ShortenedString: linkShortenedString},
			createLinkResponse,
		),
	)

	// ensure user is created
	getLinkUserResponse, err := client.GetLinkUser(
		ctx,
		&pb.GetLinkUserRequest{ShortenedString: linkShortenedString},
	)
	require.NoError(err)
	require.True(
		proto.Equal(
			&pb.GetLinkUserResponse{Username: createUserRequest.Username},
			getLinkUserResponse,
		),
	)

	// username have taken
	createUserResponse, err = client.CreateUser(ctx, &pb.CreateUserRequest{
		Username: createUserRequest.Username,
		Password: createUserRequest.Password,
	})
	require.Equal(
		status.Error(codes.AlreadyExists, "username have taken"),
		err,
	)
	require.Nil(createUserResponse)
}

package gapi

import (
	"context"
	"errors"

	"github.com/aria3ppp/url-shortener-grpc/internal/core/domain"
	domain_errors "github.com/aria3ppp/url-shortener-grpc/internal/core/errors"
	"github.com/aria3ppp/url-shortener-grpc/internal/core/port"
	"github.com/aria3ppp/url-shortener-grpc/internal/pb"
	"github.com/aria3ppp/url-shortener-grpc/internal/validate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnsafeUrlShortenerServer
	serviceUseCases port.ServiceUseCases
}

var _ pb.UrlShortenerServer = &Server{}

func NewServer(serviceUseCases port.ServiceUseCases) *Server {
	return &Server{serviceUseCases: serviceUseCases}
}

func (s *Server) GetLink(
	ctx context.Context,
	req *pb.GetLinkRequest,
) (*pb.GetLinkResponse, error) {
	// validate request
	if err := validate.GetLinkRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	link, err := s.serviceUseCases.GetLink(req.ShortenedString)
	if err != nil {
		if errors.Is(err, domain_errors.ErrLinkNotFound) {
			return nil, status.Error(codes.NotFound, "link not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetLinkResponse{Url: link.URL}, nil
}

func parseUserFromMetadata(ctx context.Context, user *domain.User) {
	usernameSlc := metadata.ValueFromIncomingContext(ctx, "username")
	passwordSlc := metadata.ValueFromIncomingContext(ctx, "password")

	if len(usernameSlc) > 0 {
		user.Username = usernameSlc[0]
	}
	if len(passwordSlc) > 0 {
		user.Password = passwordSlc[0]
	}
}

func (s *Server) CreateLink(
	ctx context.Context,
	req *pb.CreateLinkRequest,
) (*pb.CreateLinkResponse, error) {
	// parse and validate user from header
	var user domain.User
	parseUserFromMetadata(ctx, &user)
	if err := user.Validate(); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// validate request
	if err := validate.CreateLinkRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	link, err := s.serviceUseCases.CreateLink(
		req.Url,
		req.ShortenedString,
		&user,
	)
	if err != nil {
		if errors.Is(err, domain_errors.ErrUserNotFound) ||
			errors.Is(err, domain_errors.ErrIncorrectPassword) {
			return nil, status.Error(
				codes.Unauthenticated,
				"invalid username or password",
			)
		}
		if errors.Is(err, domain_errors.ErrUsedShortenedString) {
			return nil, status.Error(
				codes.AlreadyExists,
				"shortened string have used",
			)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateLinkResponse{ShortenedString: link.ShortenedString}, nil
}

func (s *Server) GetLinkUser(
	ctx context.Context,
	req *pb.GetLinkUserRequest,
) (*pb.GetLinkUserResponse, error) {
	// validate request
	if err := validate.GetLinkUserRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.serviceUseCases.GetLinkUser(req.ShortenedString)
	if err != nil {
		if errors.Is(err, domain_errors.ErrLinkNotFound) {
			return nil, status.Error(codes.NotFound, "link not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetLinkUserResponse{Username: user.Username}, nil
}

func (s *Server) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	// validate request
	if err := validate.CreateUserRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.serviceUseCases.CreateUser(&domain.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, domain_errors.ErrUsernameTaken) {
			return nil, status.Error(codes.AlreadyExists, "username have taken")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUserResponse{}, nil
}

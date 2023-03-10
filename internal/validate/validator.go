package validate

import (
	"github.com/aria3ppp/url-shortener-grpc/internal/pb"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func shortenedStringRules(required bool) []validation.Rule {
	var rules []validation.Rule
	if required {
		rules = append(rules, validation.Required)
	}
	rules = append(
		rules,
		validation.Length(6, 32),
		is.Alphanumeric,
	)
	return rules
}

func GetLinkRequest(r *pb.GetLinkRequest) error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.ShortenedString,
			shortenedStringRules(true)...,
		),
	)
}

func CreateLinkRequest(r *pb.CreateLinkRequest) error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Url,
			validation.Required,
		),
		validation.Field(
			&r.ShortenedString,
			shortenedStringRules(false)...,
		),
	)
}

func GetLinkUserRequest(r *pb.GetLinkUserRequest) error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.ShortenedString,
			shortenedStringRules(true)...,
		),
	)
}

func CreateUserRequest(r *pb.CreateUserRequest) error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Username,
			validation.Required,
			validation.Length(8, 40),
		),
		validation.Field(
			&r.Password,
			validation.Required,
			validation.Length(8, 40),
		),
	)
}

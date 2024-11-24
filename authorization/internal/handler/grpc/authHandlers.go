package grpchandler

import (
	"context"
	"errors"

	authv1 "github.com/DavidG9999/DMS/api/grpc/auth_api/gen/go"
	"github.com/DavidG9999/DMS/authorization/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	authv1.UnimplementedAuthServer
	Service service.Service
}

func (h *Handler) SignUp(ctx context.Context, req *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {
	if err := validateSignUp(req); err != nil {
		return nil, err
	}
	userId, err := h.Service.SignUp(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrUserExist) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authv1.SignUpResponse{
		UserId: userId,
	}, nil
}

func (h *Handler) SignIn(ctx context.Context, req *authv1.SignInRequest) (*authv1.SignInResponse, error) {
	if err := validateSignIn(req); err != nil {
		return nil, err
	}
	token, err := h.Service.SignIn(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "user in user-client not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &authv1.SignInResponse{
		Token: token,
	}, nil
}

func validateSignUp(req *authv1.SignUpRequest) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "username is required")
	}
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func validateSignIn(req *authv1.SignInRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password id required")
	}
	return nil
}

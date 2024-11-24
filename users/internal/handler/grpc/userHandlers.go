package handler

import (
	"context"
	"errors"

	userv1 "github.com/DavidG9999/DMS/api/grpc/user_api/gen/go"
	"github.com/DavidG9999/DMS/users/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	userv1.UnimplementedUserServer
	Service service.Service
}

func (h *Handler) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	if err := validateCreateUser(req); err != nil {
		return nil, err
	}

	userId, err := h.Service.CreateUser(ctx, req.GetName(), req.GetEmail(), req.GetPasswordHash())
	if err != nil {
		if errors.Is(err, service.ErrUserExist) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &userv1.CreateUserResponse{
		UserId: userId,
	}, nil

}

func (h *Handler) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	if err := validateGetUser(req); err != nil {
		return nil, err
	}

	user, err := h.Service.GetUser(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &userv1.GetUserResponse{
		UserId:       user.ID,
		Name:         user.Name,
		Email: user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (h *Handler) GetUserById(ctx context.Context, req *userv1.GetUserByIdRequest) (*userv1.GetUserByIdResponse, error) {
	if err := validateGetUserById(req); err != nil {
		return nil, err
	}

	user, err := h.Service.GetUserById(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &userv1.GetUserByIdResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *Handler) UpdateName(ctx context.Context, req *userv1.UpdateNameRequest) (*userv1.UpdateNameResponse, error) {
	if err := validateUpdateName(req); err != nil {
		return nil, err
	}

	err := h.Service.UpdateName(ctx, req.GetUserId(), req.GetUpdateName())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &userv1.UpdateNameResponse{
		Message: "updated name",
	}, nil
}

func (h *Handler) UpdatePassword(ctx context.Context, req *userv1.UpdatePasswordRequest) (*userv1.UpdatePasswordResponse, error) {
	if err := validateUpdatePassword(req); err != nil {
		return nil, err
	}

	err := h.Service.UpdatePassword(ctx, req.GetUserId(), req.GetUpdatePassword())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &userv1.UpdatePasswordResponse{
		Message: "updated password",
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userv1.DeleteUserRequest) (*userv1.DeleteUserResponse, error) {
	if err := validateDeleteUser(req); err != nil {
		return nil, err
	}

	err := h.Service.DeleteUser(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &userv1.DeleteUserResponse{
		Message: "deleted",
	}, nil
}

func validateCreateUser(req *userv1.CreateUserRequest) error {
	if req.GetName() == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPasswordHash() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}

func validateGetUser(req *userv1.GetUserRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	return nil
}

func validateGetUserById(req *userv1.GetUserByIdRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	return nil
}

func validateUpdateName(req *userv1.UpdateNameRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetUpdateName() == "" {
		return status.Error(codes.InvalidArgument, "update name is required")
	}
	return nil
}

func validateUpdatePassword(req *userv1.UpdatePasswordRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	if req.GetUpdatePassword() == "" {
		return status.Error(codes.InvalidArgument, "update password is required")
	}
	return nil
}

func validateDeleteUser(req *userv1.DeleteUserRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "user ID is required")
	}
	return nil
}

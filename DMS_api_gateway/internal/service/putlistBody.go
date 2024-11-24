package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PutlistBodyCreator interface {
	CreatePutlistBody(ctx context.Context, putlistNumber int64, putlistBody entity.PutlistBody) (int64, error)
}

type PutlistBodyProvider interface {
	GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error)
}

type PutlistBodyEditor interface {
	UpdatePutlistBody(ctx context.Context, putlistBodyId int64, updateData entity.UpdatePutlistBodyInput) (string, error)
	DeletePutlistBody(ctx context.Context, putlistBodyId int64) (string, error)
}

func (p *PutlistService) CreatePutlistBody(ctx context.Context, putlistNumber int64, putlistBody entity.PutlistBody) (int64, error) {
	const op = "service.CreatePutlistBody"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("creating putlist body")

	putlistBodyId, err := p.putlistClient.CreatePutlistBody(ctx, putlistNumber, putlistBody.Number, putlistBody.ContragentId, putlistBody.Item, putlistBody.TimeWith, putlistBody.TimeFor)
	if err != nil {
		if errors.Is(err, status.Error(codes.InvalidArgument, "invalid datetime format")) {
			logger.Error("invalid datetime format")
			return 0, fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		logger.Error("failed to create putlist body")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist body created")

	return putlistBodyId, nil
}

func (p *PutlistService) GetPutlistBodies(ctx context.Context, putlistNumber int64) ([]entity.PutlistBody, error) {
	const op = "service.GetPutlists"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("getting putlist bodies")

	putlistBodies, err := p.putlistClient.GetPutlistBodies(ctx, putlistNumber)
	if err != nil {
		logger.Error("failed to get putlist bodies")
		return []entity.PutlistBody{}, fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist bodies received")

	return putlistBodies, nil
}

func (p *PutlistService) UpdatePutlistBody(ctx context.Context, putlistBodyId int64, updateData entity.UpdatePutlistBodyInput) (string, error) {
	const op = "service.UpdatePutlistBody"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("updating putlist body")

	err := updateData.Validate()
	if err != nil {
		return "", err
	}

	message, err := p.putlistClient.UpdatePutlistBody(ctx, putlistBodyId, updateData.Number, updateData.ContragentId, updateData.Item, updateData.TimeWith, updateData.TimeFor)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "putlist body not found")) {
			logger.Error("putlist body not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		if errors.Is(err, status.Error(codes.InvalidArgument, "invalid datetime format")) {
			logger.Error("invalid datetime format")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidDateTimeFormat)
		}
		logger.Error("failed to update putlist body")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist body updated")

	return message, nil
}

func (p *PutlistService) DeletePutlistBody(ctx context.Context, putlistBodyId int64) (string, error) {
	const op = "service.DeletePutlistBody"

	logger := p.logger.With(slog.String("op", op))
	logger.Info("deleting putlist body")

	message, err := p.putlistClient.DeletePutlistBody(ctx, putlistBodyId)
	if err != nil {
		if errors.Is(err, status.Error(codes.NotFound, "putlist body not found")) {
			logger.Error("putlist body not found")
			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}
		logger.Error("failed to delete putlist body")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	logger.Info("putlist body deleted")

	return message, nil
}

package putlistgrpc

import (
		"github.com/DavidG9999/DMS/DMS_api_gateway/internal/entity"
	"context"
	"fmt"

	putlistv1 "github.com/DavidG9999/DMS/api/grpc/document_api/putlist/gen/go"
)

func (pc *PutlistClient) CreatePutlist(ctx context.Context, userId, number, bankAccountId int64, dateWith, dateFor string, autoId, driverId, dispetcherId, mehanicId int64) (int64, error) {
	const op = "grpc.ClientCreatePutlist"

	resp, err := pc.apiPutlist.CreatePutlist(ctx, &putlistv1.CreatePutlistRequest{UserId: userId, Number: number, BankAccountId: bankAccountId, DateWith: dateWith, DateFor: dateFor, AutoId: autoId, DriverId: driverId, DispetcherId: dispetcherId, MehanicId: mehanicId})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetPutlistId(), nil
}

func (pc *PutlistClient) GetPutlists(ctx context.Context, userId int64) ([]entity.Putlist, error) {
	const op = "grpc.ClientGetPutlists"

	resp, err := pc.apiPutlist.GetPutlists(ctx, &putlistv1.GetPutlistsRequest{UserId: userId})
	if err != nil {
		return []entity.Putlist{}, fmt.Errorf("%s: %w", op, err)
	}

	putlists := make([]entity.Putlist, len(resp.Putlists))
	for id, putlist := range resp.GetPutlists() {
		putlists[id].Id = putlist.PutlistId
		putlists[id].UserId = putlist.UserId
		putlists[id].Number = putlist.Number
		putlists[id].BankAccountId = putlist.BankAccountId
		putlists[id].DateWith = putlist.DateWith
		putlists[id].DateFor = putlist.DateFor
		putlists[id].AutoId = putlist.AutoId
		putlists[id].DriverId = putlist.DriverId
		putlists[id].DispetcherId = putlist.DispetcherId
		putlists[id].MehanicId = putlist.MehanicId
	}
	return putlists, nil
}

func (pc *PutlistClient) GetPutlistByNumber(ctx context.Context, userId int64, number int64) (entity.Putlist, error) {
	const op = "grpc.ClientGetPutlistByNumber"

	resp, err := pc.apiPutlist.GetPutlistByNumber(ctx, &putlistv1.GetPutlistByNumberRequest{UserId: userId, Number: number})
	if err != nil {
		return entity.Putlist{}, fmt.Errorf("%s: %w", op, err)
	}

	putlist := entity.Putlist{
		Id:            resp.Putlist.GetPutlistId(),
		UserId:        resp.Putlist.GetUserId(),
		Number:        resp.Putlist.GetNumber(),
		BankAccountId: resp.Putlist.GetBankAccountId(),
		DateWith:      resp.Putlist.GetDateWith(),
		DateFor:       resp.Putlist.GetDateFor(),
		AutoId:        resp.Putlist.GetAutoId(),
		DriverId:      resp.Putlist.GetDriverId(),
		DispetcherId:  resp.Putlist.GetDispetcherId(),
		MehanicId:    resp.Putlist.GetMehanicId(),
	}

	return putlist, nil
}

func (pc *PutlistClient) UpdatePutlist(ctx context.Context, userId, number int64, bankAccountId *int64, dateWith, dateFor *string, autoId, driverId, dispetcherId, mehanicId *int64) (string, error) {
	const op = "grpc.ClientUpdatePutlist"

	resp, err := pc.apiPutlist.UpdatePutlist(ctx, &putlistv1.UpdatePutlistRequest{UserId: userId, Number: number, BankAccountId: bankAccountId, DateWith: dateWith, DateFor: dateFor, AutoId: autoId, DriverId: driverId, DispetcherId: dispetcherId, MehanicId: mehanicId})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

func (pc *PutlistClient) DeletePutlist(ctx context.Context, userId, number int64) (string, error) {
	const op = "grpc.ClientDeletePutlist"

	resp, err := pc.apiPutlist.DeletePutlist(ctx, &putlistv1.DeletePutlistRequest{UserId: userId, Number: number})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetMessage(), nil
}

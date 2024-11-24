package suite

import (
	"context"
	"net"
	"strconv"
	"testing"

	docv1 "github.com/DavidG9999/DMS/api/grpc/document_api/gen/go"
	"github.com/DavidG9999/DMS/documents/putlists/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcHost = "localhost"
)

type Suite struct {
	*testing.T
	Cfg                *config.Config
	AutoClient         docv1.AutoClient
	BankAccountClient  docv1.BankAccountClient
	ContragentClient   docv1.ContragentClient
	DriverClient       docv1.DriverClient
	DispetcherClient   docv1.DispetcherClient
	OrganizationClient docv1.OrganizationClient
	MehanicClient      docv1.MehanicClient
	PutlistClient      docv1.PutlistClient
	PutlistBodyClient  docv1.PutlistBodyClient
}

func NewSuite(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("../configs/config_test.yml")

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.NewClient(grpcAddress(cfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}
	return ctx, &Suite{
		T:                  t,
		Cfg:                cfg,
		AutoClient:         docv1.NewAutoClient(cc),
		BankAccountClient:  docv1.NewBankAccountClient(cc),
		ContragentClient:   docv1.NewContragentClient(cc),
		DriverClient:       docv1.NewDriverClient(cc),
		DispetcherClient:   docv1.NewDispetcherClient(cc),
		OrganizationClient: docv1.NewOrganizationClient(cc),
		MehanicClient:      docv1.NewMehanicClient(cc),
		PutlistClient:      docv1.NewPutlistClient(cc),
		PutlistBodyClient:  docv1.NewPutlistBodyClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}

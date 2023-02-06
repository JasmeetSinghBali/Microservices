package server

import (
	"context"
	"log"

	protos "github.com/Jasmeet-1998/Microservices/grpc_golang_snippet/protos/fiat"
)

type Fiat struct {
	tracer *log.Logger
}

// implements GetRate of the fiatServer interface from .pb.go
func (f *Fiat) GetRate(context.Context, *FiatRateRequest) (*FiatRateResponse, error) {
	f.tracer.Println("Handle GetRate base", rr.GetBase())
	f.tracer.Println("Handle GetRate Destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}

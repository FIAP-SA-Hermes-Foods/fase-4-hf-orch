package client

import (
	"context"
	"fmt"

	"fase-4-hf-orch/internal/core/domain/entity/dto"
	"fase-4-hf-orch/internal/core/domain/rpc"
	op "fase-4-hf-orch/orch_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ rpc.ClientRPC = (*clientRPC)(nil)

type clientRPC struct {
	ctx  context.Context
	host string
	port string
}

func NewClientRPC(ctx context.Context, host, port string) rpc.ClientRPC {
	return clientRPC{ctx: ctx, host: host, port: port}
}

func (c clientRPC) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", c.host, c.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.CreateClientRequest{
		Name:  client.Name,
		Cpf:   client.CPF,
		Email: client.Email,
	}

	cc := op.NewClientClient(conn)

	resp, err := cc.CreateClient(c.ctx, &input)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, err
	}

	var out = dto.OutputClient{
		UUID:      resp.Uuid,
		Name:      resp.Name,
		CPF:       resp.Cpf,
		Email:     resp.Email,
		CreatedAt: resp.CreatedAt,
	}

	return &out, nil
}

func (c clientRPC) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", c.host, c.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	input := op.GetClientByCPFRequest{
		Cpf: cpf,
	}

	cc := op.NewClientClient(conn)

	resp, err := cc.GetClientByCPF(c.ctx, &input)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, err
	}

	var out = dto.OutputClient{
		UUID:      resp.Uuid,
		Name:      resp.Name,
		CPF:       resp.Cpf,
		Email:     resp.Email,
		CreatedAt: resp.CreatedAt,
	}

	return &out, nil
}

func (c clientRPC) GetClientByID(uuid string) (*dto.OutputClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", c.host, c.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	input := op.GetClientByIDRequest{
		Uuid: uuid,
	}

	cc := op.NewClientClient(conn)

	resp, err := cc.GetClientByID(c.ctx, &input)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, err
	}

	var out = dto.OutputClient{
		UUID:      resp.Uuid,
		Name:      resp.Name,
		CPF:       resp.Cpf,
		Email:     resp.Email,
		CreatedAt: resp.CreatedAt,
	}

	return &out, nil
}

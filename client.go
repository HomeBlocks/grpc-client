package grpcclient

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address string
}

type Client struct {
	config Config
	conn   *grpc.ClientConn
}

func (c *Client) OnStart(_ context.Context) error {
	conn, err := grpc.NewClient(c.config.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "failed to create grpc client")
	}

	c.conn = conn

	return nil
}

func (c *Client) OnStop(_ context.Context) error {
	err := c.conn.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close grpc client")
	}

	return nil
}

func (c *Client) Connection() *grpc.ClientConn {
	return c.conn
}

func NewClient(c Config) *Client {
	return &Client{
		config: c,
	}
}

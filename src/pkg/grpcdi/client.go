package grpcdi

import "google.golang.org/grpc"

type Client struct {
	conn *grpc.ClientConn
}

func (v *Client) Shutdown() error {
	_ = v.conn.Close()
	return nil
}

func NewClient(conn *grpc.ClientConn) (*Client, error) {
	return &Client{
		conn: conn,
	}, nil
}

// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package wireguard

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	vpe "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2106/vpe"
)

// RPCService defines RPC service wireguard.
type RPCService interface {
	WireguardInterfaceCreate(ctx context.Context, in *WireguardInterfaceCreate) (*WireguardInterfaceCreateReply, error)
	WireguardInterfaceDelete(ctx context.Context, in *WireguardInterfaceDelete) (*WireguardInterfaceDeleteReply, error)
	WireguardInterfaceDump(ctx context.Context, in *WireguardInterfaceDump) (RPCService_WireguardInterfaceDumpClient, error)
	WireguardPeerAdd(ctx context.Context, in *WireguardPeerAdd) (*WireguardPeerAddReply, error)
	WireguardPeerRemove(ctx context.Context, in *WireguardPeerRemove) (*WireguardPeerRemoveReply, error)
	WireguardPeersDump(ctx context.Context, in *WireguardPeersDump) (RPCService_WireguardPeersDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) WireguardInterfaceCreate(ctx context.Context, in *WireguardInterfaceCreate) (*WireguardInterfaceCreateReply, error) {
	out := new(WireguardInterfaceCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WireguardInterfaceDelete(ctx context.Context, in *WireguardInterfaceDelete) (*WireguardInterfaceDeleteReply, error) {
	out := new(WireguardInterfaceDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WireguardInterfaceDump(ctx context.Context, in *WireguardInterfaceDump) (RPCService_WireguardInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_WireguardInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_WireguardInterfaceDumpClient interface {
	Recv() (*WireguardInterfaceDetails, error)
	api.Stream
}

type serviceClient_WireguardInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_WireguardInterfaceDumpClient) Recv() (*WireguardInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *WireguardInterfaceDetails:
		return m, nil
	case *vpe.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) WireguardPeerAdd(ctx context.Context, in *WireguardPeerAdd) (*WireguardPeerAddReply, error) {
	out := new(WireguardPeerAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WireguardPeerRemove(ctx context.Context, in *WireguardPeerRemove) (*WireguardPeerRemoveReply, error) {
	out := new(WireguardPeerRemoveReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WireguardPeersDump(ctx context.Context, in *WireguardPeersDump) (RPCService_WireguardPeersDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_WireguardPeersDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_WireguardPeersDumpClient interface {
	Recv() (*WireguardPeersDetails, error)
	api.Stream
}

type serviceClient_WireguardPeersDumpClient struct {
	api.Stream
}

func (c *serviceClient_WireguardPeersDumpClient) Recv() (*WireguardPeersDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *WireguardPeersDetails:
		return m, nil
	case *vpe.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

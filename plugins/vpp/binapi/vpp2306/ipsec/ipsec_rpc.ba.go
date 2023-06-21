// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ipsec

import (
	"context"
	"fmt"
	"io"

	api "go.fd.io/govpp/api"
	memclnt "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/memclnt"
)

// RPCService defines RPC service ipsec.
type RPCService interface {
	IpsecBackendDump(ctx context.Context, in *IpsecBackendDump) (RPCService_IpsecBackendDumpClient, error)
	IpsecInterfaceAddDelSpd(ctx context.Context, in *IpsecInterfaceAddDelSpd) (*IpsecInterfaceAddDelSpdReply, error)
	IpsecItfCreate(ctx context.Context, in *IpsecItfCreate) (*IpsecItfCreateReply, error)
	IpsecItfDelete(ctx context.Context, in *IpsecItfDelete) (*IpsecItfDeleteReply, error)
	IpsecItfDump(ctx context.Context, in *IpsecItfDump) (RPCService_IpsecItfDumpClient, error)
	IpsecSaDump(ctx context.Context, in *IpsecSaDump) (RPCService_IpsecSaDumpClient, error)
	IpsecSaV2Dump(ctx context.Context, in *IpsecSaV2Dump) (RPCService_IpsecSaV2DumpClient, error)
	IpsecSaV3Dump(ctx context.Context, in *IpsecSaV3Dump) (RPCService_IpsecSaV3DumpClient, error)
	IpsecSadEntryAdd(ctx context.Context, in *IpsecSadEntryAdd) (*IpsecSadEntryAddReply, error)
	IpsecSadEntryAddDel(ctx context.Context, in *IpsecSadEntryAddDel) (*IpsecSadEntryAddDelReply, error)
	IpsecSadEntryAddDelV2(ctx context.Context, in *IpsecSadEntryAddDelV2) (*IpsecSadEntryAddDelV2Reply, error)
	IpsecSadEntryAddDelV3(ctx context.Context, in *IpsecSadEntryAddDelV3) (*IpsecSadEntryAddDelV3Reply, error)
	IpsecSadEntryDel(ctx context.Context, in *IpsecSadEntryDel) (*IpsecSadEntryDelReply, error)
	IpsecSadEntryUpdate(ctx context.Context, in *IpsecSadEntryUpdate) (*IpsecSadEntryUpdateReply, error)
	IpsecSelectBackend(ctx context.Context, in *IpsecSelectBackend) (*IpsecSelectBackendReply, error)
	IpsecSetAsyncMode(ctx context.Context, in *IpsecSetAsyncMode) (*IpsecSetAsyncModeReply, error)
	IpsecSpdAddDel(ctx context.Context, in *IpsecSpdAddDel) (*IpsecSpdAddDelReply, error)
	IpsecSpdDump(ctx context.Context, in *IpsecSpdDump) (RPCService_IpsecSpdDumpClient, error)
	IpsecSpdEntryAddDel(ctx context.Context, in *IpsecSpdEntryAddDel) (*IpsecSpdEntryAddDelReply, error)
	IpsecSpdEntryAddDelV2(ctx context.Context, in *IpsecSpdEntryAddDelV2) (*IpsecSpdEntryAddDelV2Reply, error)
	IpsecSpdInterfaceDump(ctx context.Context, in *IpsecSpdInterfaceDump) (RPCService_IpsecSpdInterfaceDumpClient, error)
	IpsecSpdsDump(ctx context.Context, in *IpsecSpdsDump) (RPCService_IpsecSpdsDumpClient, error)
	IpsecTunnelProtectDel(ctx context.Context, in *IpsecTunnelProtectDel) (*IpsecTunnelProtectDelReply, error)
	IpsecTunnelProtectDump(ctx context.Context, in *IpsecTunnelProtectDump) (RPCService_IpsecTunnelProtectDumpClient, error)
	IpsecTunnelProtectUpdate(ctx context.Context, in *IpsecTunnelProtectUpdate) (*IpsecTunnelProtectUpdateReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IpsecBackendDump(ctx context.Context, in *IpsecBackendDump) (RPCService_IpsecBackendDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecBackendDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecBackendDumpClient interface {
	Recv() (*IpsecBackendDetails, error)
	api.Stream
}

type serviceClient_IpsecBackendDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecBackendDumpClient) Recv() (*IpsecBackendDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecBackendDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecInterfaceAddDelSpd(ctx context.Context, in *IpsecInterfaceAddDelSpd) (*IpsecInterfaceAddDelSpdReply, error) {
	out := new(IpsecInterfaceAddDelSpdReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecItfCreate(ctx context.Context, in *IpsecItfCreate) (*IpsecItfCreateReply, error) {
	out := new(IpsecItfCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecItfDelete(ctx context.Context, in *IpsecItfDelete) (*IpsecItfDeleteReply, error) {
	out := new(IpsecItfDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecItfDump(ctx context.Context, in *IpsecItfDump) (RPCService_IpsecItfDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecItfDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecItfDumpClient interface {
	Recv() (*IpsecItfDetails, error)
	api.Stream
}

type serviceClient_IpsecItfDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecItfDumpClient) Recv() (*IpsecItfDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecItfDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSaDump(ctx context.Context, in *IpsecSaDump) (RPCService_IpsecSaDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSaDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSaDumpClient interface {
	Recv() (*IpsecSaDetails, error)
	api.Stream
}

type serviceClient_IpsecSaDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSaDumpClient) Recv() (*IpsecSaDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSaDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSaV2Dump(ctx context.Context, in *IpsecSaV2Dump) (RPCService_IpsecSaV2DumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSaV2DumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSaV2DumpClient interface {
	Recv() (*IpsecSaV2Details, error)
	api.Stream
}

type serviceClient_IpsecSaV2DumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSaV2DumpClient) Recv() (*IpsecSaV2Details, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSaV2Details:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSaV3Dump(ctx context.Context, in *IpsecSaV3Dump) (RPCService_IpsecSaV3DumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSaV3DumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSaV3DumpClient interface {
	Recv() (*IpsecSaV3Details, error)
	api.Stream
}

type serviceClient_IpsecSaV3DumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSaV3DumpClient) Recv() (*IpsecSaV3Details, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSaV3Details:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSadEntryAdd(ctx context.Context, in *IpsecSadEntryAdd) (*IpsecSadEntryAddReply, error) {
	out := new(IpsecSadEntryAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSadEntryAddDel(ctx context.Context, in *IpsecSadEntryAddDel) (*IpsecSadEntryAddDelReply, error) {
	out := new(IpsecSadEntryAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSadEntryAddDelV2(ctx context.Context, in *IpsecSadEntryAddDelV2) (*IpsecSadEntryAddDelV2Reply, error) {
	out := new(IpsecSadEntryAddDelV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSadEntryAddDelV3(ctx context.Context, in *IpsecSadEntryAddDelV3) (*IpsecSadEntryAddDelV3Reply, error) {
	out := new(IpsecSadEntryAddDelV3Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSadEntryDel(ctx context.Context, in *IpsecSadEntryDel) (*IpsecSadEntryDelReply, error) {
	out := new(IpsecSadEntryDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSadEntryUpdate(ctx context.Context, in *IpsecSadEntryUpdate) (*IpsecSadEntryUpdateReply, error) {
	out := new(IpsecSadEntryUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSelectBackend(ctx context.Context, in *IpsecSelectBackend) (*IpsecSelectBackendReply, error) {
	out := new(IpsecSelectBackendReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSetAsyncMode(ctx context.Context, in *IpsecSetAsyncMode) (*IpsecSetAsyncModeReply, error) {
	out := new(IpsecSetAsyncModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdAddDel(ctx context.Context, in *IpsecSpdAddDel) (*IpsecSpdAddDelReply, error) {
	out := new(IpsecSpdAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdDump(ctx context.Context, in *IpsecSpdDump) (RPCService_IpsecSpdDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSpdDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSpdDumpClient interface {
	Recv() (*IpsecSpdDetails, error)
	api.Stream
}

type serviceClient_IpsecSpdDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSpdDumpClient) Recv() (*IpsecSpdDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSpdDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSpdEntryAddDel(ctx context.Context, in *IpsecSpdEntryAddDel) (*IpsecSpdEntryAddDelReply, error) {
	out := new(IpsecSpdEntryAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdEntryAddDelV2(ctx context.Context, in *IpsecSpdEntryAddDelV2) (*IpsecSpdEntryAddDelV2Reply, error) {
	out := new(IpsecSpdEntryAddDelV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecSpdInterfaceDump(ctx context.Context, in *IpsecSpdInterfaceDump) (RPCService_IpsecSpdInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSpdInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSpdInterfaceDumpClient interface {
	Recv() (*IpsecSpdInterfaceDetails, error)
	api.Stream
}

type serviceClient_IpsecSpdInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSpdInterfaceDumpClient) Recv() (*IpsecSpdInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSpdInterfaceDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecSpdsDump(ctx context.Context, in *IpsecSpdsDump) (RPCService_IpsecSpdsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecSpdsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecSpdsDumpClient interface {
	Recv() (*IpsecSpdsDetails, error)
	api.Stream
}

type serviceClient_IpsecSpdsDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecSpdsDumpClient) Recv() (*IpsecSpdsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecSpdsDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecTunnelProtectDel(ctx context.Context, in *IpsecTunnelProtectDel) (*IpsecTunnelProtectDelReply, error) {
	out := new(IpsecTunnelProtectDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpsecTunnelProtectDump(ctx context.Context, in *IpsecTunnelProtectDump) (RPCService_IpsecTunnelProtectDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpsecTunnelProtectDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpsecTunnelProtectDumpClient interface {
	Recv() (*IpsecTunnelProtectDetails, error)
	api.Stream
}

type serviceClient_IpsecTunnelProtectDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpsecTunnelProtectDumpClient) Recv() (*IpsecTunnelProtectDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpsecTunnelProtectDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IpsecTunnelProtectUpdate(ctx context.Context, in *IpsecTunnelProtectUpdate) (*IpsecTunnelProtectUpdateReply, error) {
	out := new(IpsecTunnelProtectUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

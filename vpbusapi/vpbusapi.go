// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package vpbusapi

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/ufoot/vapor/vpcommonapi"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = vpcommonapi.GoUnusedProtection__

type VpBusApi interface {
	vpcommonapi.VpCommonApi
	//VpBusApi is used to communicate between Vapor and Fumes.
	//Vapor is the Golang server and Fumes the C++ client.

	// Halt stops the server.
	Halt() (err error)
}

//VpBusApi is used to communicate between Vapor and Fumes.
//Vapor is the Golang server and Fumes the C++ client.
type VpBusApiClient struct {
	*vpcommonapi.VpCommonApiClient
}

func NewVpBusApiClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *VpBusApiClient {
	return &VpBusApiClient{VpCommonApiClient: vpcommonapi.NewVpCommonApiClientFactory(t, f)}
}

func NewVpBusApiClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *VpBusApiClient {
	return &VpBusApiClient{VpCommonApiClient: vpcommonapi.NewVpCommonApiClientProtocol(t, iprot, oprot)}
}

// Halt stops the server.
func (p *VpBusApiClient) Halt() (err error) {
	if err = p.sendHalt(); err != nil {
		return
	}
	return
}

func (p *VpBusApiClient) sendHalt() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("halt", thrift.ONEWAY, p.SeqId); err != nil {
		return
	}
	args := VpBusApiHaltArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

type VpBusApiProcessor struct {
	*vpcommonapi.VpCommonApiProcessor
}

func NewVpBusApiProcessor(handler VpBusApi) *VpBusApiProcessor {
	self0 := &VpBusApiProcessor{vpcommonapi.NewVpCommonApiProcessor(handler)}
	self0.AddToProcessorMap("halt", &vpBusApiProcessorHalt{handler: handler})
	return self0
}

type vpBusApiProcessorHalt struct {
	handler VpBusApi
}

func (p *vpBusApiProcessorHalt) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := VpBusApiHaltArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	if err2 = p.handler.Halt(); err2 != nil {
		return true, err2
	}
	return true, nil
}

// HELPER FUNCTIONS AND STRUCTURES

type VpBusApiHaltArgs struct {
}

func NewVpBusApiHaltArgs() *VpBusApiHaltArgs {
	return &VpBusApiHaltArgs{}
}

func (p *VpBusApiHaltArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VpBusApiHaltArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("halt_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VpBusApiHaltArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VpBusApiHaltArgs(%+v)", *p)
}

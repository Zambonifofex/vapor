// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package vpp2papi

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

type VpP2pApi interface {
	vpcommonapi.VpCommonApi
	//VpP2pApi is used to communicate between 2 Vapor nodes
	//in peer-to-peer mode.

	// Parameters:
	//  - Key
	//  - KeyShift
	//  - ImaginaryNode
	IAmPrev(key []byte, keyShift []byte, imaginaryNode []byte) (r *LookupData, err error)
	// Parameters:
	//  - Key
	//  - KeyShift
	//  - ImaginaryNode
	Lookup(key []byte, keyShift []byte, imaginaryNode []byte) (r *LookupData, err error)
}

//VpP2pApi is used to communicate between 2 Vapor nodes
//in peer-to-peer mode.
type VpP2pApiClient struct {
	*vpcommonapi.VpCommonApiClient
}

func NewVpP2pApiClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *VpP2pApiClient {
	return &VpP2pApiClient{VpCommonApiClient: vpcommonapi.NewVpCommonApiClientFactory(t, f)}
}

func NewVpP2pApiClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *VpP2pApiClient {
	return &VpP2pApiClient{VpCommonApiClient: vpcommonapi.NewVpCommonApiClientProtocol(t, iprot, oprot)}
}

// Parameters:
//  - Key
//  - KeyShift
//  - ImaginaryNode
func (p *VpP2pApiClient) IAmPrev(key []byte, keyShift []byte, imaginaryNode []byte) (r *LookupData, err error) {
	if err = p.sendIAmPrev(key, keyShift, imaginaryNode); err != nil {
		return
	}
	return p.recvIAmPrev()
}

func (p *VpP2pApiClient) sendIAmPrev(key []byte, keyShift []byte, imaginaryNode []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("IAmPrev", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := VpP2pApiIAmPrevArgs{
		Key:           key,
		KeyShift:      keyShift,
		ImaginaryNode: imaginaryNode,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *VpP2pApiClient) recvIAmPrev() (value *LookupData, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "IAmPrev" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "IAmPrev failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "IAmPrev failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "IAmPrev failed: invalid message type")
		return
	}
	result := VpP2pApiIAmPrevResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Key
//  - KeyShift
//  - ImaginaryNode
func (p *VpP2pApiClient) Lookup(key []byte, keyShift []byte, imaginaryNode []byte) (r *LookupData, err error) {
	if err = p.sendLookup(key, keyShift, imaginaryNode); err != nil {
		return
	}
	return p.recvLookup()
}

func (p *VpP2pApiClient) sendLookup(key []byte, keyShift []byte, imaginaryNode []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Lookup", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := VpP2pApiLookupArgs{
		Key:           key,
		KeyShift:      keyShift,
		ImaginaryNode: imaginaryNode,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *VpP2pApiClient) recvLookup() (value *LookupData, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "Lookup" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "Lookup failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Lookup failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "Lookup failed: invalid message type")
		return
	}
	result := VpP2pApiLookupResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type VpP2pApiProcessor struct {
	*vpcommonapi.VpCommonApiProcessor
}

func NewVpP2pApiProcessor(handler VpP2pApi) *VpP2pApiProcessor {
	self6 := &VpP2pApiProcessor{vpcommonapi.NewVpCommonApiProcessor(handler)}
	self6.AddToProcessorMap("IAmPrev", &vpP2pApiProcessorIAmPrev{handler: handler})
	self6.AddToProcessorMap("Lookup", &vpP2pApiProcessorLookup{handler: handler})
	return self6
}

type vpP2pApiProcessorIAmPrev struct {
	handler VpP2pApi
}

func (p *vpP2pApiProcessorIAmPrev) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := VpP2pApiIAmPrevArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("IAmPrev", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := VpP2pApiIAmPrevResult{}
	var retval *LookupData
	var err2 error
	if retval, err2 = p.handler.IAmPrev(args.Key, args.KeyShift, args.ImaginaryNode); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing IAmPrev: "+err2.Error())
		oprot.WriteMessageBegin("IAmPrev", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("IAmPrev", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type vpP2pApiProcessorLookup struct {
	handler VpP2pApi
}

func (p *vpP2pApiProcessorLookup) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := VpP2pApiLookupArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Lookup", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := VpP2pApiLookupResult{}
	var retval *LookupData
	var err2 error
	if retval, err2 = p.handler.Lookup(args.Key, args.KeyShift, args.ImaginaryNode); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Lookup: "+err2.Error())
		oprot.WriteMessageBegin("Lookup", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Lookup", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Key
//  - KeyShift
//  - ImaginaryNode
type VpP2pApiIAmPrevArgs struct {
	Key           []byte `thrift:"key,1" json:"key"`
	KeyShift      []byte `thrift:"keyShift,2" json:"keyShift"`
	ImaginaryNode []byte `thrift:"imaginaryNode,3" json:"imaginaryNode"`
}

func NewVpP2pApiIAmPrevArgs() *VpP2pApiIAmPrevArgs {
	return &VpP2pApiIAmPrevArgs{}
}

func (p *VpP2pApiIAmPrevArgs) GetKey() []byte {
	return p.Key
}

func (p *VpP2pApiIAmPrevArgs) GetKeyShift() []byte {
	return p.KeyShift
}

func (p *VpP2pApiIAmPrevArgs) GetImaginaryNode() []byte {
	return p.ImaginaryNode
}
func (p *VpP2pApiIAmPrevArgs) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
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

func (p *VpP2pApiIAmPrevArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *VpP2pApiIAmPrevArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.KeyShift = v
	}
	return nil
}

func (p *VpP2pApiIAmPrevArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.ImaginaryNode = v
	}
	return nil
}

func (p *VpP2pApiIAmPrevArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IAmPrev_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VpP2pApiIAmPrevArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteBinary(p.Key); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *VpP2pApiIAmPrevArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("keyShift", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:keyShift: ", p), err)
	}
	if err := oprot.WriteBinary(p.KeyShift); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.keyShift (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:keyShift: ", p), err)
	}
	return err
}

func (p *VpP2pApiIAmPrevArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("imaginaryNode", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:imaginaryNode: ", p), err)
	}
	if err := oprot.WriteBinary(p.ImaginaryNode); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.imaginaryNode (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:imaginaryNode: ", p), err)
	}
	return err
}

func (p *VpP2pApiIAmPrevArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VpP2pApiIAmPrevArgs(%+v)", *p)
}

// Attributes:
//  - Success
type VpP2pApiIAmPrevResult struct {
	Success *LookupData `thrift:"success,0" json:"success,omitempty"`
}

func NewVpP2pApiIAmPrevResult() *VpP2pApiIAmPrevResult {
	return &VpP2pApiIAmPrevResult{}
}

var VpP2pApiIAmPrevResult_Success_DEFAULT *LookupData

func (p *VpP2pApiIAmPrevResult) GetSuccess() *LookupData {
	if !p.IsSetSuccess() {
		return VpP2pApiIAmPrevResult_Success_DEFAULT
	}
	return p.Success
}
func (p *VpP2pApiIAmPrevResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VpP2pApiIAmPrevResult) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
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

func (p *VpP2pApiIAmPrevResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &LookupData{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *VpP2pApiIAmPrevResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IAmPrev_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VpP2pApiIAmPrevResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *VpP2pApiIAmPrevResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VpP2pApiIAmPrevResult(%+v)", *p)
}

// Attributes:
//  - Key
//  - KeyShift
//  - ImaginaryNode
type VpP2pApiLookupArgs struct {
	Key           []byte `thrift:"key,1" json:"key"`
	KeyShift      []byte `thrift:"keyShift,2" json:"keyShift"`
	ImaginaryNode []byte `thrift:"imaginaryNode,3" json:"imaginaryNode"`
}

func NewVpP2pApiLookupArgs() *VpP2pApiLookupArgs {
	return &VpP2pApiLookupArgs{}
}

func (p *VpP2pApiLookupArgs) GetKey() []byte {
	return p.Key
}

func (p *VpP2pApiLookupArgs) GetKeyShift() []byte {
	return p.KeyShift
}

func (p *VpP2pApiLookupArgs) GetImaginaryNode() []byte {
	return p.ImaginaryNode
}
func (p *VpP2pApiLookupArgs) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
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

func (p *VpP2pApiLookupArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *VpP2pApiLookupArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.KeyShift = v
	}
	return nil
}

func (p *VpP2pApiLookupArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.ImaginaryNode = v
	}
	return nil
}

func (p *VpP2pApiLookupArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Lookup_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VpP2pApiLookupArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err)
	}
	if err := oprot.WriteBinary(p.Key); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err)
	}
	return err
}

func (p *VpP2pApiLookupArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("keyShift", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:keyShift: ", p), err)
	}
	if err := oprot.WriteBinary(p.KeyShift); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.keyShift (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:keyShift: ", p), err)
	}
	return err
}

func (p *VpP2pApiLookupArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("imaginaryNode", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:imaginaryNode: ", p), err)
	}
	if err := oprot.WriteBinary(p.ImaginaryNode); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.imaginaryNode (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:imaginaryNode: ", p), err)
	}
	return err
}

func (p *VpP2pApiLookupArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VpP2pApiLookupArgs(%+v)", *p)
}

// Attributes:
//  - Success
type VpP2pApiLookupResult struct {
	Success *LookupData `thrift:"success,0" json:"success,omitempty"`
}

func NewVpP2pApiLookupResult() *VpP2pApiLookupResult {
	return &VpP2pApiLookupResult{}
}

var VpP2pApiLookupResult_Success_DEFAULT *LookupData

func (p *VpP2pApiLookupResult) GetSuccess() *LookupData {
	if !p.IsSetSuccess() {
		return VpP2pApiLookupResult_Success_DEFAULT
	}
	return p.Success
}
func (p *VpP2pApiLookupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VpP2pApiLookupResult) Read(iprot thrift.TProtocol) error {
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
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
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

func (p *VpP2pApiLookupResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &LookupData{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *VpP2pApiLookupResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Lookup_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VpP2pApiLookupResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *VpP2pApiLookupResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VpP2pApiLookupResult(%+v)", *p)
}

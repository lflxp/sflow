package sflow

import (
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"encoding/binary"
)


//origin datagram info
type Datagram struct {
	IPLength         int
	SrcIP, DstIP     net.IP
	SrcPort, DstPort uint16
}

type Data struct {
	Type		string
	Datagram	Datagram
	DatagramVersion uint32
	AgentAddress    net.IP
	SubAgentID      uint32
	SequenceNumber  uint32
	AgentUptime     uint32
	SampleCount     uint32
}

func NewData() *Data {
	return &Data{}
}

func (this *Data) InitDatagram(tuple *Datagram) error {
	this.Datagram.IPLength = tuple.IPLength
	this.Datagram.SrcIP = tuple.SrcIP
	this.Datagram.DstIP = tuple.DstIP
	this.Datagram.SrcPort = tuple.SrcPort
	this.Datagram.DstPort = tuple.DstPort
	return nil
}

func (this *Data) Init(payload []byte) error {

	pp := gopacket.NewPacket(payload, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		//fmt.Println(pp.Data())
		this.DecodeDataFromBytes(pp.Data())
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		this.DatagramVersion = got.DatagramVersion
		this.AgentAddress = got.AgentAddress
		this.SubAgentID = got.SubAgentID
		this.SequenceNumber = got.SequenceNumber
		this.AgentUptime = got.AgentUptime
		this.SampleCount = got.SampleCount
	}
	return nil
}

func (this *Data) DecodeDataFromBytes(data []byte) error {
	var agentAddressType layers.SFlowIPType

	data ,this.DatagramVersion = data[4:],binary.BigEndian.Uint32(data[:4])
	data, agentAddressType = data[4:], layers.SFlowIPType(binary.BigEndian.Uint32(data[:4]))
	data, this.AgentAddress = data[agentAddressType.Length():], data[:agentAddressType.Length()]
	data, this.SubAgentID = data[4:], binary.BigEndian.Uint32(data[:4])
	data, this.SequenceNumber = data[4:], binary.BigEndian.Uint32(data[:4])
	data, this.AgentUptime = data[4:], binary.BigEndian.Uint32(data[:4])
	data, this.SampleCount = data[4:], binary.BigEndian.Uint32(data[:4])
	return nil
}
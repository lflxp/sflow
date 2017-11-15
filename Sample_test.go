package sflow

import (
	"testing"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TestNewFlowSamples(t *testing.T) {
	data := NewFlowSamples()
	if data == nil {
		t.Fatal("no data new")
	} else {
		t.Log("ok")
	}
}

func TestFlowSamples_SendUdp(t *testing.T) {
	data := NewFlowSamples()
	//counter collector
	err := data.SendUdp("{'json':'yes','name':'test'}","127.0.0.1:6789","127.0.0.1:9876",true)
	if err != nil {
		t.Fatal("SendUdp Counter collector error",err.Error())
	}
	t.Log("SendUdp Counter ok")

	//sample collector
	err = data.SendUdp("{'json':'yes','name':'test'}","127.0.0.1:6789","127.0.0.1:9876",false)
	if err != nil {
		t.Fatal("SendUdp sample collector error",err.Error())
	}
	t.Log("SendUdp sample collector ok")
}

func TestFlowSamples_InitFlowSampleData(t *testing.T) {
	Header := NewData()
	Header.InitDatagram(&datagram)
	Header.Init(SFlow)

	pp := gopacket.NewPacket(SFlow, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		t.Fatal(pp.ErrorLayer().Error())
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		if len(got.FlowSamples) > 0 {
			for _, y := range got.FlowSamples {
				tmp := NewFlowSamples()
				Header.Type = "sample"
				tmp.Data = Header
				err := tmp.InitFlowSampleData(y)
				if err != nil {
					t.Fatal(err.Error())
				}
			}
		}
	}
	t.Log("InitFlowSampleData ok")
}

func TestFlowSamples_ParseLayers(t *testing.T) {
	Header := NewData()
	Header.InitDatagram(&datagram)
	Header.Init(SFlow)

	pp := gopacket.NewPacket(SFlow, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		t.Fatal(pp.ErrorLayer().Error())
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		if len(got.FlowSamples) > 0 {
			for _, y := range got.FlowSamples {
				tmp := NewFlowSamples()
				for _,yy := range y.Records {
					if g1,ok1 := yy.(layers.SFlowRawPacketFlowRecord); ok1 {
						err := tmp.ParseLayers(g1.Header)
						if err != nil {
							t.Fatal(err.Error())
						}
					}
				}
			}
		}
	}
	t.Log("ParseLayers ok")
}
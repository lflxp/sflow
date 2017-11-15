package sflow

import (
	"testing"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TestNewCounterFlow(t *testing.T) {
	data := NewCounterFlow()
	if data == nil {
		t.Fatal("NewCounterFlow return nil")
	}
	t.Log("NewCounterFlow ok", data)
}

func TestSFlowCounterSample_InitDatagram(t *testing.T) {
	data := NewCounterFlow()
	err := data.InitDatagram(&datagram)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("TestSFlowCounterSample_InitDatagram ok", data)
}

func TestSFlowCounterSample_InitOriginData(t *testing.T) {
	data := NewCounterFlow()
	err := data.InitDatagram(&datagram)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = data.InitOriginData(SFlow)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("TestSFlowCounterSample_InitOriginData ok", data)
}

func TestSFlowCounterSample_InitCounterSample(t *testing.T) {
	Header := NewData()
	Header.InitDatagram(&datagram)
	Header.Init(SFlow)

	pp := gopacket.NewPacket(SFlow, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		t.Fatal(pp.ErrorLayer().Error())
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		if len(got.CounterSamples) > 0 {
			tmp := NewCounterFlow()
			tmp.InitOriginData(SFlow)
			tmp.InitDatagram(&datagram)
			for _, y := range got.CounterSamples {
				err := tmp.InitCounterSample(y)
				if err != nil {
					t.Fatal("InitCounterSample error", err.Error())
				}
			}
		}
	}
	t.Log("TestSFlowCounterSample_InitCounterSample ok")
}

func TestSFlowCounterSample_InitCounterSampleStruct(t *testing.T) {
	Header := NewData()
	Header.InitDatagram(&datagram)
	Header.Init(SFlow)

	pp := gopacket.NewPacket(SFlow, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		t.Fatal(pp.ErrorLayer().Error())
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		tmp := NewCounterFlow()
		tmp.InitOriginData(SFlow)
		err := tmp.InitCounterSampleStruct(got)
		if err != nil {
			t.Fatal("InitCounterSampleStruct error:", err.Error())
		}
	}
	t.Log("TestSFlowCounterSample_InitCounterSampleStruct ok")
}
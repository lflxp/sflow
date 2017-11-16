package sflow

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

//decode sflow V5 struct
//contains sample and counter info
func DecodeSflow(tuple *Datagram, payload []byte) (*[]FlowSamples,*[]SFlowCounterSample, error) {
	samples := []FlowSamples{}
	counters := []SFlowCounterSample{}
	Header := NewData()
	Header.InitDatagram(tuple)
	err := Header.Init(payload)
	if err != nil {
		return nil,nil,err
	}

	pp := gopacket.NewPacket(payload, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		return nil,nil, pp.ErrorLayer().Error()
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		if len(got.FlowSamples) > 0 {
			for _, y := range got.FlowSamples {
				tmp := NewFlowSamples()
				Header.Type = "sample"
				tmp.Data = Header
				tmp.InitFlowSampleData(y)
				samples = append(samples,*tmp)
			}
		}

		if len(got.CounterSamples) > 0 {
			tmp := NewCounterFlow()
			tmp.InitOriginData(payload)
			tmp.InitDatagram(tuple)
			for _, y := range got.CounterSamples {
				tmp.InitCounterSample(y)
			}

			counters = append(counters,*tmp)
		}
	}
	return &samples,&counters, nil
}

//only decode and return sample
func DecodeSample(tuple *Datagram, payload []byte) ([]*FlowSamples, error) {
	samples := []*FlowSamples{}
	Header := NewData()
	Header.InitDatagram(tuple)
	err := Header.Init(payload)
	if err != nil {
		return nil,err
	}

	pp := gopacket.NewPacket(payload, layers.LayerTypeSFlow, gopacket.Default)
	if pp.ErrorLayer() != nil {
		return nil, pp.ErrorLayer().Error()
	}
	if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
		if len(got.FlowSamples) > 0 {
			for _, y := range got.FlowSamples {
				tmp := NewFlowSamples()
				Header.Type = "sample"
				tmp.Data = Header
				tmp.InitFlowSampleData(y)
				samples = append(samples,tmp)
			}
		}

	}
	return samples, nil
}

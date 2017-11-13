package sflow

import (
	//"fmt"
	//"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func DecodeSflow(tuple *Datagram, payload []byte) ([]*FlowSamples,[]*SFlowCounterSample, error) {
	samples := []*FlowSamples{}
	counters := []*SFlowCounterSample{}
	Header := NewData()
	Header.InitDatagram(tuple)
	Header.Init(payload)

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
				//b, err := json.Marshal(tmp)
				//if err != nil {
				//	fmt.Println(err.Error())
				//}
				//
				//samples = append(samples, string(b))
				samples = append(samples,tmp)
			}
		}

		if len(got.CounterSamples) > 0 {
			tmp := NewCounterFlow()
			tmp.InitOriginData(payload)
			tmp.InitDatagram(tuple)
			for _, y := range got.CounterSamples {
				tmp.InitCounterSample(y)
			}

			//b, err := json.Marshal(tmp)
			//if err != nil {
			//	fmt.Println(err.Error())
			//}
			//
			//counters = append(counters, string(b))
			counters = append(counters,tmp)
		}
	}
	return samples,counters, nil
}

func DecodeSample(tuple *Datagram, payload []byte) ([]*FlowSamples, error) {
	samples := []*FlowSamples{}
	Header := NewData()
	Header.InitDatagram(tuple)
	Header.Init(payload)

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
				//b, err := json.Marshal(tmp)
				//if err != nil {
				//	fmt.Println(err.Error())
				//}
				//
				//samples = append(samples, string(b))
				samples = append(samples,tmp)
			}
		}

	}
	return samples, nil
}

package sflow

import (
	"fmt"
	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func DecodeSflow(tuple *Datagram, payload []byte) ([]string, error) {
	result := []string{}
	Header := NewData()
	Header.InitDatagram(tuple)
	Header.Init(payload)

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
				b, err := json.Marshal(tmp)
				if err != nil {
					fmt.Println(err.Error())
				}

				result = append(result, string(b))
			}
		}

		if len(got.CounterSamples) > 0 {
			tmp := NewCounterFlow()
			tmp.InitOriginData(payload)
			tmp.InitDatagram(tuple)
			for _, y := range got.CounterSamples {
				tmp.InitCounterSample(y)
			}

			b, err := json.Marshal(tmp)
			if err != nil {
				fmt.Println(err.Error())
			}

			result = append(result, string(b))
		}
	}
	return result, nil
}

# SFlow [![Circle CI](https://circleci.com/gh/lflxp/sflow.svg)](https://circleci.com/gh/lflxp) [![GoDoc](https://godoc.org/github.com/lflxp/sflow?status.svg)](https://godoc.org/github.com/lflxp/sflow) [![Coverage Status](https://coveralls.io/repos/github/lflxp/sflow/badge.svg?branch=master)](https://coveralls.io/github/lflxp/sflow?branch=master)

# Info
Parsing sFlow for packetbeat plug-ins

This is struct to decode sflow 5 layers and datagram info together  
Fork on part of gopacket's sFlow data structure


# Sample

Only decoder five layers of sflow sample

```
SFlowRawPacketFlowRecord
SFlowExtendedSwitchFlowRecord
SFlowExtendedRouterFlowRecord
SFlowExtendedGatewayFlowRecord
SFlowExtendedUserFlow
```

# Counter

decoder all of layers

```
SFlowGenericInterfaceCounters
SFlowEthernetCounters
SFlowProcessorCounters
```

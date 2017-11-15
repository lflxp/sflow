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
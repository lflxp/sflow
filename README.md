# sflow
Decode sflow struct like this

This is struct to decode sflow 5 layers and datagram info together  
Borrowed part of gopacket's sFlow data structure


# Useage

DecodeSflow(tuple *common.IPPortTuple, payload []byte) ([]string, error) 
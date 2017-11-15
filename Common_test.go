package sflow

import "testing"

func TestNewData(t *testing.T) {
	data := NewData()
	if data == nil {
		t.Fatal("NewData func return nil")
	}
	t.Log("NewData ok",data)
}

func TestData_InitDatagram(t *testing.T) {
	data := NewData()
	err := data.InitDatagram(&datagram)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("InitDatagram ok",data)
}

func TestData_Init(t *testing.T) {
	data := NewData()
	err := data.Init(SFlow)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("TestData_Init ok",data)
}

func TestData_DecodeDataFromBytes(t *testing.T) {
	data := NewData()
	err := data.DecodeDataFromBytes(SFlow)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("TestData_DecodeDataFromBytes ok",data)
}
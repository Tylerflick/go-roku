package client

import (
	"fmt"
	"testing"
)

func TestQueryDeviceInfo(t *testing.T) {
	c := NewClient("192.168.1.126")
	info, err := c.GetDeviceInfo()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(info)
}

func TestGetChannels(t *testing.T) {
	c := NewClient("192.168.1.126")
	info := c.GetChannels()
	fmt.Println(info)
}

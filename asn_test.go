package main

import (
	"fmt"
	"testing"
)

func TestGetASN(t *testing.T) {
	fmt.Println("Testing ASN...")
	p := SetupParam("81.2.69.142")
	data, err := GetASN(p)
	if err != nil {
		fmt.Println(err)
	}
	PPrint(data)
}

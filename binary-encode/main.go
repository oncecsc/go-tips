package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Protocol struct {
	Version  uint8
	BodyLen  uint16
	Reserved [2]byte
	Unit     uint8
	Value    uint32
	A        uint64
}

func main() {

	p := Protocol{
		Version: 1,
		BodyLen: 2,
		Unit:    3,
		Value:   5,
		A:       67,
	}
	buffer := new(bytes.Buffer)

	err := binary.Write(buffer, binary.LittleEndian, p)
	if err != nil {
		panic(err)
	}

	bin := buffer.Bytes()
	fmt.Printf("serialize protocol: %+v\n", bin)

	var b Protocol
	err = binary.Read(bytes.NewReader(bin), binary.LittleEndian, &b)
	if err != nil {
		panic(err)
	}

	fmt.Printf("un serialize protocol: %+v\n", b)
}

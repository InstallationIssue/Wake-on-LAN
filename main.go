package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

const headerByte byte = 0xFF

var (
	// ipTarget    string
	macTarget   string
	ipBroadcast string
	port        string
)

type MACAddress [6]byte

type MagicPacket struct {
	header  [6]byte
	payload [16]MACAddress
}

// Constructing Magic Packet
func create(hw net.HardwareAddr) (bytes.Buffer, error) {
	var mac MACAddress
	var mp MagicPacket

	for b := range mp.header {
		mp.header[b] = headerByte
	}

	for add := range mac {
		mac[add] = hw[add]
	}

	for m := range mp.payload {
		mp.payload[m] = mac
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, mp); err != nil {
		return buf, err
	}

	return buf, nil
}

// Sending Magic Packet
func send(mp bytes.Buffer) error {
	conn, err := net.Dial("udp", ipBroadcast+":"+port)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Write(mp.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// Main Execution
func main() {

	hw, err := net.ParseMAC(macTarget)
	if err != nil {
		panic(err)
	}

	mp, err := create(hw)
	if err != nil {
		panic(err)
	}

	err = send(mp)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sent Magic Packet")
}

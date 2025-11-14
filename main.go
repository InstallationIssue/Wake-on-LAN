package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
)

const headerByte byte = 0xFF

var (
	macTarget          string
	ipBroadcast        string
	port               string
	defaultMacTarget   string
	defaultIpBroadcast string
	defaultPort        = "8829"
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
func send(mp bytes.Buffer, ipBroadcast string, port string) error {
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

	flag.StringVar(&macTarget, "mac", defaultMacTarget, "MAC Target for WoL Magic Packet")
	flag.StringVar(&ipBroadcast, "bc", defaultIpBroadcast, "Broadcast Address on local network")
	flag.StringVar(&port, "port", defaultPort, "WoL Port (default 8829)")
	flag.Parse()

	if macTarget == "" || ipBroadcast == "" || port == "" {
		log.Fatalf("A Default Variable is undefined:\n\tdefaultMacTarget = %s\n\tdefaultIpBroadcast = %s\n\tdefaultPort = %s\n", defaultMacTarget, defaultIpBroadcast, defaultPort)
	}

	hw, err := net.ParseMAC(macTarget)
	if err != nil {
		log.Fatalln(err)
	}

	mp, err := create(hw)
	if err != nil {
		log.Fatalln(err)
	}

	err = send(mp, ipBroadcast, port)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sent Magic Packet")
}

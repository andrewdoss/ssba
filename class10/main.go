package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

type CaptureParser struct {
	f              *os.File
	majorVersion   uint16
	minorVersion   uint16
	snapshotLength uint32
	linkLayerType  uint32
	decoder        binary.ByteOrder
}

func checkErr(msg string, err error) {
	if err != nil {
		err = fmt.Errorf("%s: %w", msg, err)
		log.Fatal(err)
	}
}

func (p *CaptureParser) readNBytes(numBytes int) ([]byte, error) {
	buf := make([]byte, numBytes)
	n, err := p.f.Read(buf)
	// Fail if EOF reached mid-read
	if n != 0 && n != numBytes {
		log.Fatalf("read %d bytes, expected %d bytes", n, numBytes)
	}
	return buf, err
}

// ParseCaptureConf parses a pcap capture config
func (p *CaptureParser) ParseCaptureConf() {
	_, err := p.f.Seek(0, 0)
	checkErr("failed to seek to start of file", err)

	buf, err := p.readNBytes(4)
	checkErr("failed to read magic number bytes", err)
	// Use magic number to set byte ordering for field decoder
	switch binary.LittleEndian.Uint32(buf) {
	case 0xa1b2c3d4:
		p.decoder = binary.LittleEndian
	case 0xd4c3b2a1:
		p.decoder = binary.BigEndian
	// For now, ignoring variations 0xa1b23c4d and 0x4d3cb2a1 which
	// indicate higher-precision packet time stamps (nano vs. micro)
	default:
		err = fmt.Errorf("unrecognized magic number, unable to infer byte order")
		log.Fatal(err)
	}

	// Parse the rest of the header
	buf, err = p.readNBytes(2)
	checkErr("failed to read major version", err)
	p.majorVersion = p.decoder.Uint16(buf)

	buf, err = p.readNBytes(2)
	checkErr("failed to read minor version", err)
	p.minorVersion = p.decoder.Uint16(buf)

	// Seek past time zone and time stamp fields
	_, err = p.f.Seek(8, 1)
	checkErr("failed to seek past unused fields", err)

	buf, err = p.readNBytes(4)
	checkErr("failed to read snapshot length", err)
	p.snapshotLength = p.decoder.Uint32(buf)

	buf, err = p.readNBytes(4)
	checkErr("failed to read link-layer header type", err)
	p.linkLayerType = p.decoder.Uint32(buf)

	log.Printf("Capture config: version: %d.%d, snapshot length: %d, link layer type: %d",
		p.majorVersion,
		p.minorVersion,
		p.snapshotLength,
		p.linkLayerType,
	)
}

// ParsePacketHeader consumes one pcap packet header and returns
// packet length and any errors
func (p *CaptureParser) ParsePacketHeader() (uint32, error) {
	buf, err := p.readNBytes(16)
	capLength := p.decoder.Uint32(buf[8:12])
	fullLength := p.decoder.Uint32(buf[12:])
	if capLength != fullLength {
		log.Fatalf("encountered truncated packet")
	}
	return capLength, err
}

// ParsePacket consumes one packet of the specified length
func (p *CaptureParser) ParsePacket(pLength uint32) error {
	_, err := p.readNBytes(int(pLength))
	return err
}

// ParseFile parses a pcap capture file
func (p *CaptureParser) ParseFile() {
	p.ParseCaptureConf()

	// The remainder of the file consists of 16 byte per-packet capture
	// headers, and the packets themselves. The next step is to find the total
	// number of packets.
	nPackets := 0
	for {
		pLength, err := p.ParsePacketHeader()
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr("failed to parse packet header", err)
		}
		nPackets++
		err = p.ParsePacket(pLength)
		if err != nil {
			checkErr("failed to parse packet", err)
		}
	}
	log.Printf("Found %d packets.", nPackets)
	// Confirm that entire file has been consumed

}

func main() {
	parser := CaptureParser{}
	switch len(os.Args) {
	case 1:
		parser.f = os.Stdin
	case 2:
		f, err := os.Open(os.Args[1])
		if err != nil {
			err = fmt.Errorf("failed to open specified file: %w", err)
			log.Fatal(err)
		}
		parser.f = f
	default:
		err := fmt.Errorf("expected at most one argument")
		log.Fatal(err)
	}
	parser.ParseFile()
}

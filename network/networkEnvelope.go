package network

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"gochain/utils"
	"log"
)


type NetworkEnvelope struct {
	NetworkMagic []byte
	Command      []byte
	Payload      []byte
}

func (ne *NetworkEnvelope) Parse(s []byte) *NetworkEnvelope {
	magic := s[:4]
	if magic == nil {
		log.Panic("ERROR: Connection reset!")
	}
	if !bytes.Equal(magic, NETWORK_MAGIC[:]) {
		log.Panic("magic is not right")
	}
	command := s[4:16]
	payloadLength := binary.LittleEndian.Uint16(utils.ToLittleEndian(s[16:20], 4))
	checksum := s[20:24]
	payload := s[24 : 24+payloadLength]
	payloadHash := sha256.Sum256(payload)
	calculatedChecksum := payloadHash[:4]
	if !bytes.Equal(calculatedChecksum[:], checksum) {
		log.Panic("checksum does not match")
	}

	return &NetworkEnvelope{NETWORK_MAGIC, command, payload}
}

func (ne *NetworkEnvelope) Serialize() []byte {
	result := ne.NetworkMagic
	result = append(result, ne.Command...)
	result = append(result, utils.ToLittleEndian(utils.ToHex(int64(len(ne.Payload))), 4)...)
	checksum := sha256.Sum256(ne.Payload)
	result = append(result, checksum[:4]...)
	result = append(result, ne.Payload...)
	return result
}

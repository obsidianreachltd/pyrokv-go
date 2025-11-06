package frame

import (
	"errors"
	"pyrokv-go/internal/header"
)

type Frame struct {
	Header  *header.Header
	Payload []byte
}

func NewFrame(header *header.Header, payload []byte) *Frame {
	return &Frame{
		Header:  header,
		Payload: payload,
	}
}

func (f *Frame) ToBytes() []byte {
	headerBytes := f.Header.ToBytes()
	return append(headerBytes, f.Payload...)
}

func FrameFromBytes(b []byte) (*Frame, error) {
	if len(b) < header.HeaderSize {
		return nil, errors.New("byte slice too short to contain header")
	}
	head := header.HeaderFromBytes(b[:header.HeaderSize])
	if len(b) < header.HeaderSize+int(head.PayloadLen) {
		return nil, errors.New("byte slice too short to contain payload")
	}
	payload := b[header.HeaderSize : header.HeaderSize+int(head.PayloadLen)]
	return &Frame{
		Header:  head,
		Payload: payload,
	}, nil
}

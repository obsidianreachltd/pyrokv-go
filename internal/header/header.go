package header

const MAGIC = 0x4D51
const VERSION = 1

const (
	FrameType_Request      = 0
	FrameType_Response     = 1
	FrameType_Notification = 2
)

const (
	OpCode_Set    = 0x01
	OpCode_Get    = 0x02
	OpCode_Del    = 0x03
	OpCode_MSet   = 0x04
	OpCode_MGet   = 0x05
	OpCode_Exists = 0x06
	OpCode_Ping   = 0x10
	OpCode_Info   = 0x11
)

const (
	Flag_Error = 1 << iota
	Flag_Batch
	Flag_Compressed
)

type Header struct {
	Magic      uint16
	Version    uint8
	Operation  uint8
	FrameType  uint8
	Flags      uint16
	ReqID      uint32
	PayloadLen uint32
}

const HeaderSize = 15

func HeaderFromBytes(b []byte) *Header {
	return &Header{
		Magic:      uint16(b[0])<<8 | uint16(b[1]),
		Version:    b[2],
		Operation:  b[3],
		FrameType:  b[4],
		Flags:      uint16(b[5])<<8 | uint16(b[6]),
		ReqID:      uint32(b[7])<<24 | uint32(b[8])<<16 | uint32(b[9])<<8 | uint32(b[10]),
		PayloadLen: uint32(b[11])<<24 | uint32(b[12])<<16 | uint32(b[13])<<8 | uint32(b[14]),
	}
}

func (h *Header) ToBytes() []byte {
	b := make([]byte, HeaderSize)
	b[0] = byte(h.Magic >> 8)
	b[1] = byte(h.Magic & 0xFF)
	b[2] = h.Version
	b[3] = h.Operation
	b[4] = h.FrameType
	b[5] = byte(h.Flags >> 8)
	b[6] = byte(h.Flags & 0xFF)
	b[7] = byte(h.ReqID >> 24)
	b[8] = byte((h.ReqID >> 16) & 0xFF)
	b[9] = byte((h.ReqID >> 8) & 0xFF)
	b[10] = byte(h.ReqID & 0xFF)
	b[11] = byte(h.PayloadLen >> 24)
	b[12] = byte((h.PayloadLen >> 16) & 0xFF)
	b[13] = byte((h.PayloadLen >> 8) & 0xFF)
	b[14] = byte(h.PayloadLen & 0xFF)
	return b
}

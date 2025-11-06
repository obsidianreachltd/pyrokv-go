package payload

func NewSetRequestPayload(key string, value []byte, expiry int64) []byte {
	keyLenU32 := uint32(len(key))
	valueLenU32 := uint32(len(value))
	expiryU32 := uint32(expiry)

	payload := make([]byte, 4+keyLenU32+4+valueLenU32+4)
	// Expiry
	payload[0] = byte(expiryU32 >> 24)
	payload[1] = byte((expiryU32 >> 16) & 0xFF)
	payload[2] = byte((expiryU32 >> 8) & 0xFF)
	payload[3] = byte(expiryU32 & 0xFF)
	// Key Length
	payload[4] = byte(keyLenU32 >> 24)
	payload[5] = byte((keyLenU32 >> 16) & 0xFF)
	payload[6] = byte((keyLenU32 >> 8) & 0xFF)
	payload[7] = byte(keyLenU32 & 0xFF)
	// Key
	copy(payload[8:8+keyLenU32], key)
	// Value Length
	valueLenStart := 8 + keyLenU32
	payload[valueLenStart] = byte(valueLenU32 >> 24)
	payload[valueLenStart+1] = byte((valueLenU32 >> 16) & 0xFF)
	payload[valueLenStart+2] = byte((valueLenU32 >> 8) & 0xFF)
	payload[valueLenStart+3] = byte(valueLenU32 & 0xFF)
	// Value
	copy(payload[valueLenStart+4:], value)
	return payload
}

func NewGetRequestPayload(key string) []byte {
	keyLenU32 := uint32(len(key))
	payload := make([]byte, 4+keyLenU32)
	// Key Length
	payload[0] = byte(keyLenU32 >> 24)
	payload[1] = byte((keyLenU32 >> 16) & 0xFF)
	payload[2] = byte((keyLenU32 >> 8) & 0xFF)
	payload[3] = byte(keyLenU32 & 0xFF)
	// Key
	copy(payload[4:], key)
	return payload
}

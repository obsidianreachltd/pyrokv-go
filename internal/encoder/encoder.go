package encoder

import (
	"encoding/json"
)

func encodeIntegerToBytes[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](value T, byteSize int) []byte {
	buf := make([]byte, byteSize)
	for i := 0; i < byteSize; i++ {
		buf[i] = byte(value >> (8 * i))
	}
	return buf
}

func decodeBytesToInteger[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](data []byte) T {
	var value T
	for i := 0; i < len(data); i++ {
		value |= T(data[i]) << (8 * i)
	}
	return value
}

func Marshal(value any) ([]byte, error) {
	// Check type
	switch v := value.(type) {
	case string:
		return []byte(v), nil
	case int:
		return encodeIntegerToBytes(v, 8), nil
	case int64:
		return encodeIntegerToBytes(v, 8), nil
	case uint:
		return encodeIntegerToBytes(v, 8), nil
	case uint64:
		return encodeIntegerToBytes(v, 8), nil
	case int32:
		return encodeIntegerToBytes(v, 4), nil
	case uint32:
		return encodeIntegerToBytes(v, 4), nil
	case int16:
		return encodeIntegerToBytes(v, 2), nil
	case uint16:
		return encodeIntegerToBytes(v, 2), nil
	case int8:
		return []byte{byte(v)}, nil
	case uint8:
		return []byte{byte(v)}, nil
	case bool:
		if v {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case []byte:
		return v, nil
	default:
		return json.Marshal(v)
	}
}

func Unmarshal(data []byte, out any) error {
	// Check type of out
	switch v := out.(type) {
	case *string:
		*v = string(data)
		return nil
	case *int:
		var val int
		val = decodeBytesToInteger[int](data)
		*v = val
		return nil
	case *int64:
		var val int64
		val = decodeBytesToInteger[int64](data)
		*v = val
		return nil
	case *uint:
		var val uint
		val = decodeBytesToInteger[uint](data)
		*v = val
		return nil
	case *uint64:
		var val uint64
		val = decodeBytesToInteger[uint64](data)
		*v = val
		return nil
	case *int32:
		var val int32
		val = decodeBytesToInteger[int32](data)
		*v = val
		return nil
	case *uint32:
		var val uint32
		val = decodeBytesToInteger[uint32](data)
		*v = val
		return nil
	case *int16:
		var val int16
		val = decodeBytesToInteger[int16](data)
		*v = val
		return nil
	case *uint16:
		var val uint16
		val = decodeBytesToInteger[uint16](data)
		*v = val
		return nil
	case *int8:
		var val int8
		val = decodeBytesToInteger[int8](data)
		*v = val
		return nil
	case *uint8:
		var val uint8
		val = decodeBytesToInteger[uint8](data)
		*v = val
		return nil
	case *bool:
		if len(data) == 0 {
			*v = false
		} else {
			*v = data[0] != 0
		}
		return nil
	case *[]byte:
		*v = data
		return nil
	default:
		return json.Unmarshal(data, v)
	}
}

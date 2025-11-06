package kvstore

import "fmt"

type kvstore struct {
	Expiry int32 // in seconds
	Key    string
	Value  []byte
}

func KVStoreFromBytes(data []byte) (*kvstore, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf("invalid data length")
	}
	kv := &kvstore{}
	kv.Expiry = int32(data[0])<<24 | int32(data[1])<<16 | int32(data[2])<<8 | int32(data[3])
	keyLen := int(data[4])<<24 | int(data[5])<<16 | int(data[6])<<8 | int(data[7])
	if len(data) < 8+keyLen+4 {
		return nil, fmt.Errorf("invalid data length")
	}
	kv.Key = string(data[8 : 8+keyLen])
	valueLen := int(data[8+keyLen])<<24 | int(data[8+keyLen+1])<<16 | int(data[8+keyLen+2])<<8 | int(data[8+keyLen+3])
	if len(data) < 12+keyLen+valueLen {
		return nil, fmt.Errorf("invalid data length")
	}
	kv.Value = data[12+keyLen : 12+keyLen+valueLen]
	return kv, nil
}

package encoder

import "testing"

type CustomStruct struct {
	Field1 string
	Field2 int
}

func Test_Marshal_Int(t *testing.T) {
	val := int64(42)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled int bytes: %v", b)
	if len(b) != 8 {
		t.Fatalf("Expected byte length 8, got %d", len(b))
	}
	expected := []byte{42, 0, 0, 0, 0, 0, 0, 0}
	for i := range expected {
		if b[i] != expected[i] {
			t.Fatalf("Byte %d: expected %d, got %d", i, expected[i], b[i])
		}
	}
	var decoded int64
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled int value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_UInt(t *testing.T) {
	val := uint64(123456789)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled uint64 bytes: %v", b)
	if len(b) != 8 {
		t.Fatalf("Expected byte length 8, got %d", len(b))
	}
	var decoded uint64
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled uint64 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_Int32(t *testing.T) {
	val := int32(-12345)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled int32 bytes: %v", b)
	if len(b) != 4 {
		t.Fatalf("Expected byte length 4, got %d", len(b))
	}
	var decoded int32
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled int32 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_UInt32(t *testing.T) {
	val := uint32(987654321)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled uint32 bytes: %v", b)
	if len(b) != 4 {
		t.Fatalf("Expected byte length 4, got %d", len(b))
	}
	var decoded uint32
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled uint32 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_Int16(t *testing.T) {
	val := int16(-1234)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled int16 bytes: %v", b)
	if len(b) != 2 {
		t.Fatalf("Expected byte length 2, got %d", len(b))
	}
	var decoded int16
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled int16 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_UInt16(t *testing.T) {
	val := uint16(43210)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled uint16 bytes: %v", b)
	if len(b) != 2 {
		t.Fatalf("Expected byte length 2, got %d", len(b))
	}
	var decoded uint16
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled uint16 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_Int8(t *testing.T) {
	val := int8(-100)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled int8 bytes: %v", b)
	if len(b) != 1 {
		t.Fatalf("Expected byte length 1, got %d", len(b))
	}
	var decoded int8
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled int8 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_UInt8(t *testing.T) {
	val := uint8(200)
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled uint8 bytes: %v", b)
	if len(b) != 1 {
		t.Fatalf("Expected byte length 1, got %d", len(b))
	}
	var decoded uint8
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled uint8 value: %d", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %d, got %d", val, decoded)
	}
}

func Test_Marshal_Bool(t *testing.T) {
	val := true
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled bool bytes: %v", b)
	if len(b) != 1 {
		t.Fatalf("Expected byte length 1, got %d", len(b))
	}
	var decoded bool
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled bool value: %v", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %v, got %v", val, decoded)
	}
}

func Test_Marshal_String(t *testing.T) {
	val := "hello"
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled string bytes: %v", b)
	if string(b) != val {
		t.Fatalf("Expected string %s, got %s", val, string(b))
	}
	var decoded string
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled string value: %s", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %s, got %s", val, decoded)
	}
}

func Test_Marshal_Struct(t *testing.T) {
	val := CustomStruct{
		Field1: "test",
		Field2: 123,
	}
	b, err := Marshal(val)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	t.Logf("Marshaled struct JSON: %s", string(b))
	expectedPrefix := `{"Field1":"test","Field2":123}`
	if string(b[:len(expectedPrefix)]) != expectedPrefix {
		t.Fatalf("Expected JSON prefix %s, got %s", expectedPrefix, string(b))
	}
	var decoded CustomStruct
	if err := Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	t.Logf("Unmarshaled struct value: %+v", decoded)
	if decoded != val {
		t.Fatalf("Expected decoded value %+v, got %+v", val, decoded)
	}
}

func Test_Marshal_UnsupportedType(t *testing.T) {
	val := make(chan int)
	_, err := Marshal(val)
	if err == nil {
		t.Fatalf("Expected error for unsupported type, got nil")
	}
	t.Logf("Received expected error: %v", err)
}

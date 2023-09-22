package utils

import (
	"bytes"
	"encoding/gob"
)

// Serialize 将interface{} 序列化成 byte slice
func Serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)
	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Deserialize 将byte slice 反序列化成 interface{}
func Deserialize(valueBytes []byte) interface{} {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&value)
	if err != nil {
		return nil
	}
	return value
}

// Decode将Encode的结果重新赋值给ptr指向的类型
func Decode(b []byte, ptr interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(b))

	err := decoder.Decode(ptr)
	if err != nil {
		return err
	}

	return nil
}

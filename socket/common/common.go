package common

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	len := int32(len(message))
	pkg := new(bytes.Buffer)

	if err := binary.Write(pkg, binary.LittleEndian, len); err != nil {
		return nil, err
	}
	if err := binary.Write(pkg, binary.LittleEndian, []byte(message)); err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil

}

func Decode(reader *bufio.Reader) (string, error) {
	lenByte, _ := reader.Peek(4)
	lenBuff := bytes.NewBuffer(lenByte)
	var len int32
	err := binary.Read(lenBuff, binary.LittleEndian, &len)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < len+4 {
		return "", err

	}

	//读取数据
	pack := make([]byte, int(4+len))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[4:]), nil

}

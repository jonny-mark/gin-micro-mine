/**
 * @author jiangshangfang
 * @date 2022/1/17 2:33 PM
 **/
package encoding

import (
	"encoding/json"
	"compress/gzip"
	"bytes"
	"fmt"
	"io/ioutil"
	"github.com/golang/snappy"
)

type JSONEncoding struct{}

func (j JSONEncoding) Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	return buf, err
}

func (j JSONEncoding) Unmarshal(data []byte, value interface{}) error {
	err := json.Unmarshal(data, value)
	if err != nil {
		return err
	}
	return nil
}

//Gzip压缩
type JSONGzipEncoding struct{}

func (jz JSONGzipEncoding) Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	buf, err = GzipEncode(buf)
	return buf, err
}

func (jz JSONGzipEncoding) Unmarshal(data []byte, value interface{}) error {
	jsonData, err := GzipDecode(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, value)
	if err != nil {
		return err
	}
	return nil
}

// GzipEncode 编码
func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer, err := gzip.NewWriterLevel(&buffer, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(in)
	if err != nil {
		err = writer.Close()
		if err != nil {
			return out, err
		}
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

// GzipDecode 解码
func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer func() {
		err = reader.Close()
		if err != nil {
			fmt.Printf("reader close err: %+v", err)
		}
	}()

	return ioutil.ReadAll(reader)
}

//snappy压缩
type JSONSnappyEncoding struct{}

func (s JSONSnappyEncoding) Marshal(v interface{}) (data []byte, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	d := snappy.Encode(nil, b)
	return d, nil
}

func (s JSONSnappyEncoding) Unmarshal(data []byte, value interface{}) error {
	b, err := snappy.Decode(nil, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, value)
}

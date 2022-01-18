/**
 * @author jiangshangfang
 * @date 2021/12/12 4:17 PM
 **/
package encoding

import (
	"encoding"
	"reflect"
	"errors"
)

var (
	ErrNotAPointer = errors.New("v argument must be a pointer")
)

type Encoding interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

func isPointer(data interface{}) bool {
	switch reflect.ValueOf(data).Kind() {
	case reflect.Ptr, reflect.Interface:
		return true
	default:
		return false
	}
}

func Marshal(e Encoding, v interface{}) (data []byte, err error) {
	if !isPointer(v) {
		return data, ErrNotAPointer
	}
	bm, ok := v.(encoding.BinaryMarshaler)
	if ok && e == nil {
		data, err = bm.MarshalBinary()
		return
	}

	data, err = e.Marshal(v)
	if err == nil {
		return
	}
	if ok {
		data, err = bm.MarshalBinary()
	}

	return
}

func Unmarshal(e Encoding, data []byte, v interface{}) (err error) {
	if !isPointer(v) {
		return ErrNotAPointer
	}
	bm, ok := v.(encoding.BinaryUnmarshaler)
	if ok && e == nil {
		err = bm.UnmarshalBinary(data)
		return err
	}
	err = e.Unmarshal(data, v)
	if err == nil {
		return
	}
	if ok {
		return bm.UnmarshalBinary(data)
	}
	return
}

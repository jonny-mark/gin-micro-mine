/**
 * @author jiangshangfang
 * @date 2021/12/12 4:17 PM
 **/
package encoding

type Encoding interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}


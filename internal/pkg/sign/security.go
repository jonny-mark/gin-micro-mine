/**
 * @author jiangshangfang
 * @date 2022/4/27 5:47 PM
 **/
package sign

type Security struct {
	Sign      string
	Key       string
	KeepEmpty bool
	data      interface{}
}

// 校验签名
func (s *Security) VerifySign() {

}

// 生成签名
func  (s *Security) GenerateSign() {

}

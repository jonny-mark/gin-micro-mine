/**
 * @author jiangshangfang
 * @date 2022/1/24 6:28 PM
 **/
package app

// Response define a response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Details []string    `json:"details,omitempty"`
}

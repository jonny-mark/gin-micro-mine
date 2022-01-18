/**
 * @author jiangshangfang
 * @date 2021/12/1 4:35 PM
 **/
package lock

import "github.com/google/uuid"

func getToken() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

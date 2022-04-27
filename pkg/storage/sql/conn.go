/**
 * @author jiangshangfang
 * @date 2022/4/12 11:06 PM
 **/
package sql

import (
	"database/sql"
)

type conn struct {
	*sql.DB
}

func (*conn) begin() (*Tx, error) {

}

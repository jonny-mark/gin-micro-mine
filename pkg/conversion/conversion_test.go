/**
 * @author jiangshangfang
 * @date 2022/1/26 11:38 AM
 **/
package conversion

import (
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type Server struct {
		Name        string `json:"name,omitempty"`
		ID          int
		Enabled     bool
		users       []string // not exported
		http.Server          // embedded
	}

	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}

	maps := StructToMap(server)
	spew.Dump(maps)
}

func TestArrayToString(t *testing.T) {
	slice := []interface{}{"a", "b", "c", "d", "e", "f"}
	spew.Dump(slice)
	spew.Dump(SliceToString(slice))
}

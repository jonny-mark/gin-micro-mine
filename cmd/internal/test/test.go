/**
 * @author jiangshangfang
 * @date 2022/4/5 9:39 PM
 **/
package test

import (
	"fmt"
	"github.com/jonny-mark/gin-micro-mine/cmd/internal/test/test"
	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdTest = &cobra.Command{
	Use:   "test",
	Short: "Generate the test file",
	Long:  "Generate the test file.",
	Run:   run,
}

func init() {
	CmdTest.AddCommand(test.CmdAdd)
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Please enter the cache filename")
}

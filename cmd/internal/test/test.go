/**
 * @author jiangshangfang
 * @date 2022/4/5 9:39 PM
 **/
package test

import (
	"github.com/spf13/cobra"
	"gin/cmd/internal/test/test"
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
}

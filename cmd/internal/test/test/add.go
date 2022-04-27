/**
 * @author jiangshangfang
 * @date 2022/4/5 9:40 PM
 **/
package test

import "github.com/spf13/cobra"

// CmdCache represents the new command.
var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Create a test file by template",
	Long:  "Create a test file using the cache template.",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	print("99999")
}
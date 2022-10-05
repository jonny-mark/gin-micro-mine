/**
 * @author jiangshangfang
 * @date 2022/3/20 4:42 PM
 **/
package main

import (
	"github.com/jonny-mark/gin-micro-mine/cmd/internal/test"
	"github.com/spf13/cobra"
	"log"
)

var (
	// Version is the version of the compiled software.
	Version = "v0.15.0"

	rootCmd = &cobra.Command{
		Use:     "gin-micro",
		Short:   "Gin: An develop kit for Go microservices.",
		Long:    `Gin: An develop kit for Go microservices.`,
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(test.CmdTest)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

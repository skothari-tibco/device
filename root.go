package device

import (
	"fmt"

	"github.com/project-flogo/cli/common" // Flogo CLI support code
	"github.com/spf13/cobra"
)

func init() {
	common.RegisterPlugin(DeviceCmd)

}

var verbose bool

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "flogo device",
	Long:  `flogo command line plugin for developing flogo device apps`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root Device Command")
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "flogo device build",
	Long:  `flogo command line for building flogo device app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Command")
	},
}

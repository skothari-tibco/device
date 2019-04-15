package device

import (
	"fmt"

	"github.com/project-flogo/cli/common" // Flogo CLI support code
	"github.com/skothari-tibco/device/api"
	"github.com/spf13/cobra"
)

func init() {
	common.RegisterPlugin(deviceCmd)
	deviceCmd.AddCommand()
}

var deviceCmd = &cobra.Command{
	Use:              "device",
	Short:            "flogo device",
	Long:             `flogo command line plugin for developing flogo device apps`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
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

var createCmd = &cobra.Command{
	Use:              "create",
	Short:            "flogo device create",
	Long:             `flogo command line for building flogo create app`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Args:             cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create Command")
		api.Create("Get Current")
	},
}

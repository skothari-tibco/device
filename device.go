package device

import (
	"fmt"
	"os"

	cliApi "github.com/project-flogo/cli/api"
	"github.com/project-flogo/cli/common" // Flogo CLI support code
	"github.com/skothari-tibco/device/api"
	"github.com/spf13/cobra"
)

func init() {
	common.RegisterPlugin(deviceCmd)
	deviceCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&flogoJsonPath, "file", "f", "", "specify a flogo.json to create project from")
}

var verbose bool

var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "flogo device",
	Long:  `flogo command line plugin for developing flogo device apps`,
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

var flogoJsonPath string

var createCmd = &cobra.Command{
	Use:              "create [flags] [appName]",
	Short:            "create a flogo device application project",
	Long:             `Creates a flogo device application project.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Args:             cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		cliApi.SetVerbose(verbose)

		appName := ""
		if len(args) > 0 {
			appName = args[0]
		}

		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error determining working directory: %v\n", err)
			os.Exit(1)
		}

		api.Create(currentDir, appName, flogoJsonPath)
	},
}

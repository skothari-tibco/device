package device

import (
	"fmt"
	"os"

	cliApi "github.com/project-flogo/cli/api"

	"github.com/skothari-tibco/device/api"
	"github.com/spf13/cobra"
)

func init() {
	DeviceCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&flogoJSONPath, "file", "f", "", "specify a flogo.json to create device project from")
	createCmd.Flags().StringVarP(&boardName, "board", "b", "", "specify the board name for flogo device")

}

var verbose bool

var flogoJSONPath string
var boardName string

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

		err = api.Create(currentDir, appName, flogoJSONPath, boardName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating device project: %v\n", err)
			os.Exit(1)
		}
	},
}

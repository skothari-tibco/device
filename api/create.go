package api

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/project-flogo/cli/util"

	"github.com/project-flogo/cli/common"
)

func Create(basePath, appName, appCfgPath, boardName string) error {
	var err error
	var appJson string

	path, err := common.GetPluginPath("github.com/skothari-tibco/device")
	fmt.Println("Plugin Version.. ", path, err)

	if appCfgPath != "" {

		if util.IsRemote(appCfgPath) {

			appJson, err = util.LoadRemoteFile(appCfgPath)
			if err != nil {
				return fmt.Errorf("unable to load remote app file '%s' - %s", appCfgPath, err.Error())
			}
		} else {
			appJson, err = util.LoadLocalFile(appCfgPath)
			if err != nil {
				return fmt.Errorf("unable to load app file '%s' - %s", appCfgPath, err.Error())
			}
		}
	} else {
		if len(appName) == 0 {
			return fmt.Errorf("app name not specified")
		}
		if len(boardName) == 0 {
			return fmt.Errorf("board name not specified")
		}
	}

	appName, err = getAppName(appName, appJson)
	if err != nil {
		return err
	}

	fmt.Printf("Creating Flogo Device App: %s\n", appName)

	appDir, err := createAppDirectory(basePath, appName)
	if err != nil {
		return err
	}

	deviceProject := NewDeviceProject(appDir)

	err = util.ExecCmd(exec.Command("platformio", "init", "--board", boardName), deviceProject.AppDir())

	if err != nil {
		return err
	}

	return nil
}

func getAppName(appName, appJson string) (string, error) {

	if appJson != "" && appName == "" {
		descriptor, err := util.ParseAppDescriptor(appJson)
		if err != nil {
			return "", err
		}

		return descriptor.Name, nil
	}

	return appName, nil
}

func createAppDirectory(basePath, appName string) (string, error) {

	var err error

	_, err = exec.LookPath("platformio")

	if err != nil {
		return "", errors.New("platformio not installed")
	}

	if basePath == "." {
		basePath, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	appPath := filepath.Join(basePath, appName)
	err = os.Mkdir(appPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return appPath, nil
}

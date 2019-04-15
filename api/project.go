package api

import (
	"path/filepath"

	"github.com/skothari-tibco/device/common"
)

type deviceProjectImpl struct {
	appDir  string
	appName string
}

func NewDeviceProject(appDir string) common.DeviceProject {

	project := &deviceProjectImpl{appDir: appDir}
	project.appName = filepath.Base(appDir)
	return project
}

func (project *deviceProjectImpl) AppDir() string {
	return project.appDir
}

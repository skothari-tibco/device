package api

import (
	"path/filepath"

	"github.com/skothari-tibco/device/common"
)

type deviceProjectImpl struct {
	appDir     string
	appName    string
	srcDir     string
	libDir     string
	includeDir string
}

func NewDeviceProject(appDir string) common.DeviceProject {

	project := &deviceProjectImpl{appDir: appDir}
	project.appName = filepath.Base(appDir)
	return project
}

func (project *deviceProjectImpl) AppDir() string {
	return project.appDir
}

func (project *deviceProjectImpl) SrcDir() string {
	project.srcDir = filepath.Join(project.appDir, "src")
	return project.srcDir
}

func (project *deviceProjectImpl) LibDir() string {
	project.libDir = filepath.Join(project.appDir, "lib")
	return project.libDir
}

func (project *deviceProjectImpl) IncludeDir() string {
	project.includeDir = filepath.Join(project.appDir, "include")
	return project.includeDir
}

func (project *deviceProjectImpl) TestDir() string {
	return filepath.Join(project.appDir, "test")
}

package common

type DeviceProject interface {
	AppDir() string
	SrcDir() string
	LibDir() string
	IncludeDir() string
	TestDir() string
}

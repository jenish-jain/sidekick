package version

var version string

func Get() string {
	return getVersion()
}

func getVersion() string {
	// The value of the version comes from -ldflags that is set during `go build`
	return version
}

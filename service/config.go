package service

// Args passed via build in the `Makefile`
// -ldflags="-X 'pkg/service.BuildDate=$(name)' -X 'pkg/service.Branch=$(version)...'"
var (
	Name        string
	Description string
	Version     string
	BuildDate   string
	Branch      string
	Hash        string
	BuildMode   string
)

const (
	ServerPort = ":8080" // serve command on this port
)

// Command args
var (
	DebouncedTimeout = int64(300) // git timeout in seconds, 5m
	DebouncedCmd     = ""         //"git add .; git commit -m 'Sync cms'; git push;"
)

package plugin

import (
	"net/http"
)

type Plugin interface {
	Start(config string) error
	Get(req *http.Request) ([]byte, error)
	Post(req *http.Request) ([]byte, error)
}

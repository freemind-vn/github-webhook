package plugin

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(p string, v any) error {
	var (
		buf []byte
		err error
	)

	buf, err = os.ReadFile(p)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(buf, v)
}

func ReadBodyJson(req *http.Request) (map[string]any, error) {
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var v map[string]any
	err = json.Unmarshal(buf, &v)
	return v, err
}

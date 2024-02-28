package service

import (
	"fmt"
	"log/slog"
	"net/http"
	"plugin"

	base "freemind.com/webhook/plugin"
	"freemind.com/webhook/service/health"
	"freemind.com/webhook/service/index"
)

type Hook struct {
	Plugin string `yaml:"plugin"` // path/to/plugin.so
	Config string // path/to/config.yaml
	Secret string // secret key
}

const (
	ServerPort = ":8080" // serve command on this port
)

var opts map[string]*Hook

func ServeHTTP(config string) error {
	if err := load(config); err != nil {
		return err
	}

	http.HandleFunc("/", index.Get)
	http.HandleFunc("/health", health.Get)

	return http.ListenAndServe(ServerPort, nil)
}

func load(config string) error {
	slog.Info("load config", "path", config)
	if err := base.ReadConfig(config, &opts); err != nil {
		return err
	}

	for k, v := range opts {
		if err := addHook(k, v.Plugin, v.Config); err != nil {
			return err
		}
	}

	return nil
}

func addHook(id, so, config string) error {
	var (
		p       *plugin.Plugin
		startFn plugin.Symbol
		getFn   plugin.Symbol
		postFn  plugin.Symbol
		err     error
	)

	slog.Info("add hook", "id", id, "so", so, "config", config)
	p, err = plugin.Open(so)
	if err != nil {
		return err
	}

	startFn, err = p.Lookup("Start")
	if err != nil {
		return err
	}

	if err = startFn.(func(string) error)(config); err != nil {
		return err
	}

	getFn, err = p.Lookup("Get")
	if err != nil {
		return err
	}

	postFn, err = p.Lookup("Post")
	if err != nil {
		return err
	}

	pattern := fmt.Sprintf("/hooks/%s", id)
	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		slog.Info(fmt.Sprintf("%s %s", req.Method, pattern), "q", req.URL.Query().Encode())
		switch req.Method {
		case http.MethodGet:
			buf, err := getFn.(func(*http.Request) ([]byte, error))(req)
			if err != nil {
				handleError(w, err)
				return
			}

			w.Write(buf)

		case http.MethodPost:
			defer req.Body.Close()

			res, err := postFn.(func(*http.Request) ([]byte, error))(req)
			if err != nil {
				handleError(w, err)
				return
			}

			w.Write(res)

		case http.MethodDelete:
			http.HandleFunc(pattern, nil)
			w.Write([]byte("TODO: Implement delete a endpoint"))
		}
	})

	return nil
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

package health

import "net/http"

type HealthResponse struct {
	Version string
	Uptime  string
	Status  string
}

type HealthStatus string

const (
	HealthStatusOK HealthStatus = "OK"
)

// Handle `/health` endpoint
func Get(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

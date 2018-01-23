// healthz is a liveness probe.
package handlers

import "net/http"

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

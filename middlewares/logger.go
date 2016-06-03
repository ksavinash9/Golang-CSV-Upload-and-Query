package middlewares

import "net/http"

//HTTPLogger logs the http requests to stdout
type HTTPLogger struct {
}

func (h *HTTPLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

package infrastructure

import (
	"github.com/gorilla/mux"
)

func NewRouter(controller *ControllHandler) (root *mux.Router) {
	root = mux.NewRouter()
	// root.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)
	// eh := errorRoutingDetected
	// PathPrefix
	// api := root.PathPrefix("/api/v1/").Subrouter()
	// SamplePath
	// common := controller.Common
	// api.HandleFunc("/ping", eh(common.SampleHandler)).Methods(http.MethodGet, "OPTIONS")
	// api.HandleFunc("/taglist", eh(common.TagList)).Methods(http.MethodGet, "OPTIONS")
	return
}

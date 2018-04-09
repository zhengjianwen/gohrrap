package handler

import (
	"github.com/gorilla/mux"
)

func ConfigRouter(r *mux.Router) {
	r.HandleFunc("/version", Version).Methods("GET")

	cmdb := r.PathPrefix("/cmdb").Subrouter()
	// 资源
	cmdb.HandleFunc("/resource/list", GetModelList).Methods("GET")


}

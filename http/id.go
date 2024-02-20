package http

import (
	"errors"
	"net/http"
	"wb-l0/cache"
	"wb-l0/domain"
	"wb-l0/logger"

	"github.com/gorilla/mux"
)

type Handler struct {
	Cache cache.Cache
}

func NewHandler(router *mux.Router, ch cache.Cache) {
	handler := &Handler{
		Cache: ch,
	}

	router.HandleFunc("/api/orders/uid/{uid}", handler.GetById).Methods(http.MethodGet, http.MethodOptions)
}

func (ch *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, bool := vars["uid"]
	if bool != true {
		domain.WriteError(w, "no data", http.StatusBadRequest)
		logs.LogError(logs.Logger, "http", "GetById", errors.New("no data"), "error")
		return
	}

	orders, bool := ch.Cache.GetOrder(id)
	if bool != true {
		domain.WriteError(w, "no cache", http.StatusNotFound)
		logs.LogError(logs.Logger, "http", "GetById", errors.New("no cache"), "error")
		return
	}

	domain.WriteResponse(
		w,
		map[string]interface{}{
			"order": orders,
		},
		http.StatusOK,
	)

	logs.Logger.Info(logs.Logger, " http ", " send response ")
}

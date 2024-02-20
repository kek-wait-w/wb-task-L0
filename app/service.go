package app

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"wb-l0/cache"
	http_ "wb-l0/http"
	logs "wb-l0/logger"
	"wb-l0/nats"
	"wb-l0/postgres"
)

func Start() {

	ctx := context.Background()
	dbParams := postgres.GetParamsForDB()
	pc := postgres.Connect(ctx, dbParams)
	data, err := postgres.GetOrdersToCache(pc, ctx)
	if err != nil {
		logs.Logger.Fatal(logs.Logger, " main ", err, " Failed to get data from db")
	}
	defer pc.Close()

	router := mux.NewRouter()
	ch := cache.NewCache()
	ch.SetOrder(data.OrderUID, data)

	go nats.SubscribeToNATS(ch)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})

	http_.NewHandler(router, *ch)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	logs.Logger.Info("starting server at :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logs.Logger.Fatal(logs.Logger, " main ", err, " Failed to start server")
	}
	logs.Logger.Info("server stopped")
}

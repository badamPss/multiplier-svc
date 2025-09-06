package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// обработчик запроса
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		multiplier := generateMultiplier(rtp)

		resp := map[string]float64{
			"result": multiplier,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: only GET method is supported"))
	}
}

// поднимаем сервер на порту 64333
func startServer(rtp float64) {
	http.HandleFunc("/get", handler)

	log.Printf("server started at :64333 with RTP = %.2f\n", rtp)

	err := http.ListenAndServe(":64333", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

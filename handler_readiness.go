package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responseFormatter(w,200, struct{}{})
}
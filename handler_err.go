package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request){
	RespondWithError(w, 400, "Capim na palheta na rebinboca da parafuseta")
}
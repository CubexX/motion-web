package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/do_motion", DoMotionHandler)
	http.HandleFunc("/check_motion", CheckMotionHandler)

	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fs1 := http.FileServer(http.Dir(getTargetDir()))
	http.Handle("/shots/", http.StripPrefix("/shots/", fs1))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Info struct {
	WebcamPort  string
	ControlPort string
	//TargetDir   string
	Shots template.HTML
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err)
	}

	files, _ := ioutil.ReadDir(getTargetDir())
	var shots string

	for _, f := range files {
		file := f.Name()
		if file[len(file)-3:] == "jpg" {
			shots += `<a href="/shots/` + file + `" target="_blank"><img src="/shots/` + file + `" class="shot"></a> `
		}
	}

	t.Execute(w, &Info{WebcamPort: getWebcamPort(), ControlPort: getControlPort(), Shots: template.HTML(shots)})
}

func DoMotionHandler(w http.ResponseWriter, r *http.Request) {
	if len(checkMotion()) != 0 {
		log.Println("Stop motion")
		lexec("killall motion -9")
	} else {
		log.Println("Run motion")
		lexec("motion -c motion.conf")
	}
}

func CheckMotionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(checkMotion()))
}

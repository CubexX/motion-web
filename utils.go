package main

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

func lexec(cmd string) string {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, _ := exec.Command(head, parts...).Output()

	return string(out)
}

func checkMotion() string {
	return lexec("pgrep motion")
}

func preg_match(match string, str string) []string {
	r, _ := regexp.Compile(match)
	return r.FindStringSubmatch(str)
}

func getWebcamPort() string {
	data, _ := ioutil.ReadFile("motion.conf")
	return preg_match("webcam_port (.*)\n", string(data))[1]
}

func getControlPort() string {
	data, _ := ioutil.ReadFile("motion.conf")
	return preg_match("control_port (.*)\n", string(data))[1]
}

func getTargetDir() string {
	data, _ := ioutil.ReadFile("motion.conf")
	return preg_match("target_dir (.*)\n", string(data))[1]
}

package main

import (
	"flag"
	"fmt"
	"gomscode/src/kaau"
	"gomscode/src/services"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Starting web server for KAAU... ")
	var kubeconfig string
	kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()

	kaau.Kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	r := services.NewRouter()
	if err := http.ListenAndServe(":3333", r); err != nil {
		panic(err)
	}

}

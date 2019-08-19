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
	fmt.Println("hello main")
	//	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	//	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	//	flag.Parse()
	//	kaau.GetKubeClient(kubeconfig)

	var ns, label, field, maxClaims, kubeconfig string
	kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	flag.StringVar(&ns, "namespace", "", "namespace")
	flag.StringVar(&label, "l", "", "Label selector")
	flag.StringVar(&field, "f", "", "Field selector")
	flag.StringVar(&maxClaims, "max-claims", "100Gi", "Maximum total claims to watch")
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()

	kaau.Kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")

	r := services.NewRouter()
	if err := http.ListenAndServe(":3333", r); err != nil {
		panic(err)
	}

}

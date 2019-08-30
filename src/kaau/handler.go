package kaau

import (
	"fmt"
	"gomscode/src/utility"
	"html/template"

	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Kubeconfig  defining  as global variable
var Kubeconfig string

// K8sClientSet exported. TODO: Save client set so no need to get client set everythibg
var K8sClientSet kubernetes.Clientset

//Username defining as global variable
var Username string

// NameSpaceDetails exported
type NameSpaceDetails struct {
	Name string
}

// DataIndexPage exported
type DataIndexPage struct {
	UserName string
}

// GetKubeClient Exported
func GetKubeClient(kubeconfig string) *kubernetes.Clientset {

//	fmt.Println(kubeconfig)
	inClusterConfig, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		config, err1 := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err1 != nil {
			fmt.Println(err1)
			fmt.Println("Could not found Kube config also..")
		} else {
			fmt.Println("Found Kube config using same")
			clientset, err2 := kubernetes.NewForConfig(config)
			if err2 != nil {
				fmt.Println(err2)
			} else {
				K8sClientSet = *clientset
				return clientset
			}
		}

	} else {
		fmt.Println("Found In cluster conf.. Using same.")
		clientset, err := kubernetes.NewForConfig(inClusterConfig)
		if err != nil {
			fmt.Println(err)
		}
		K8sClientSet = *clientset
		return clientset
	}
	return nil
}

// GetNameSpaces exported
func GetNameSpaces(clntset *kubernetes.Clientset) NameSpaceDetails {
	var NameSpaces NameSpaceDetails
	namespace, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, namespace := range namespace.Items {
		NameSpaces.Name = namespace.Name
	}
	return NameSpaces
}

// HomeHandler export
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	var indexdata DataIndexPage
	if !utility.IsEmpty(userName) {
		indexdata.UserName = userName
		tmpl := template.Must(template.ParseFiles("web/templetes/index.html"))
		tmpl.Execute(w, indexdata)

	} else {
		http.Redirect(w, r, "/", 302)
	}

}

// AppHandler export
func AppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you have requested app: %s\n", r.URL.Path)
}

// LoginHandler export
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if !utility.IsEmpty(name) && !utility.IsEmpty(pass) {
		// Database check for user data!
		IsValidUser := utility.UserIsValid(name, pass)
		if IsValidUser {
			utility.SetCookie(name, w)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

// LoginPageHandler exported
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		redirectTarget := "/index"
		http.Redirect(w, r, redirectTarget, 302)
	} else {
		tmpl := template.Must(template.ParseFiles("web/templetes/login.html"))
		var data string
		tmpl.Execute(w, data)
	}
}

// LogoutPageHandler exported
func LogoutPageHandler(w http.ResponseWriter, r *http.Request) {
	utility.ClearCookie(w)
	redirectTarget := "/"
	http.Redirect(w, r, redirectTarget, 302)
}

package kaau

import (
	"fmt"
	"gomscode/src/utility"
	"html/template"
	"net/http"
)

// Kubeconfig defining  as global variable
var Kubeconfig string

// HomeHandler export
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	fmt.Println(userName)
	if !utility.IsEmpty(userName) {
		//fmt.Fprintf(w, "Hello, %s you logged to the site.. continue broswing \n", userName)
		kubeclient := GetKubeClient(Kubeconfig)
		namespaces := GetNameSpaces(kubeclient)
		namespaces.UserName = userName
		namespaces.Host = r.Host
		fmt.Println(namespaces)
		//		clusterRole := GetClusterRole(kubeclient)
		//		fmt.Println(clusterRole)
		tmpl := template.Must(template.ParseFiles("web/templetes/index.html"))
		tmpl.Execute(w, namespaces)

	} else {
		http.Redirect(w, r, "/", 302)
	}

}

//ViewRolePageHandler exported
func ViewRolePageHandler(w http.ResponseWriter, r *http.Request) {
	QueryNamespace := r.URL.Query().Get("namespace")
	kubeclient := GetKubeClient(Kubeconfig)
	RolesDetails := GetRoles(kubeclient, QueryNamespace)
	fmt.Println(RolesDetails)
	tmpl := template.Must(template.ParseFiles("web/templetes/viewrole.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, RolesDetails)
}
// ViewClusterRolePageHandler Exported
func ViewClusterRolePageHandler(w http.ResponseWriter, r *http.Request) {
	kubeclient := GetKubeClient(Kubeconfig)
	ClusterRolesDetails := GetClusterRole(kubeclient)
	fmt.Println(ClusterRolesDetails)
	tmpl := template.Must(template.ParseFiles("web/templetes/viewclusterrole.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, ClusterRolesDetails)
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

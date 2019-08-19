package kaau

import (
	"fmt"
	"gomscode/src/utility"
	"html/template"
	"net/http"
)

// Kubeconfig  defining  as global variable
var Kubeconfig string

//Username defining as global variable
var Username string

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

//ViewRolePageHandler exported
func ViewRolePageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		RolesDetails := GetRoles(kubeclient)
		RolesDetails.UserName = userName
		fmt.Println(RolesDetails)
		tmpl := template.Must(template.ParseFiles("web/templetes/roles.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, RolesDetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ViewClusterRolePageHandler Exported
func ViewClusterRolePageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		ClusterRolesDetails := GetClusterRole(kubeclient)
		ClusterRolesDetails.UserName = userName
		tmpl := template.Must(template.ParseFiles("web/templetes/clusterrole.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ClusterRolesDetails)
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


// ViewSAPageHandler exported
func ViewSAPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		SADetails := GetServiceAccount(kubeclient)
		SADetails.UserName = userName
		fmt.Println(SADetails)
		tmpl := template.Must(template.ParseFiles("web/templetes/serviceaccount.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, SADetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
// ViewRoleBindingPageHandler exported
func ViewRoleBindingPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		RBDetails := GetRoleBinding(kubeclient)
		RBDetails.UserName = userName
		fmt.Println(RBDetails)
		tmpl := template.Must(template.ParseFiles("web/templetes/rolesbinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, RBDetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}


// ViewClusterRoleBindingPageHandler exported
func ViewClusterRoleBindingPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		ClusterRBDetails := GetClusterRoleBinding(kubeclient)
		ClusterRBDetails.UserName = userName
		fmt.Println(ClusterRBDetails)
		tmpl := template.Must(template.ParseFiles("web/templetes/clusterrolebinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ClusterRBDetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

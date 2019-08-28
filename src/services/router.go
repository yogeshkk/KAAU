package services

import (
	"fmt"
	"gomscode/src/kaau"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	staticDir = "/static/"
)

// Test1 is tset
func Test1() {
	fmt.Println("hello test")
}

// StartServer exported
func StartServer(w http.ResponseWriter, r *http.Request) {
	/*	message := r.URL.Path
			message = strings.TrimPrefix(message, "/")

		logger.LogOut("INFO", r.URL.Path)
		http.ServeFile(w, r, r.URL.Path)
		r := mux.NewRouter()
	*/
}

// NewRouter exported
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	//	router.HandleFunc("/web",http.FileServer(http.Dir(staticDir))
	//	router.HandleFunc("/", kubernetes.HomeHandler)
	router.HandleFunc("/login", Middleware(kaau.LoginHandler)).Methods("POST")
	router.HandleFunc("/logout", Middleware(kaau.LogoutPageHandler))
	router.HandleFunc("/index", Middleware(kaau.HomeHandler))
	router.HandleFunc("/", Middleware(kaau.LoginPageHandler))
	router.HandleFunc("/app", Middleware(kaau.AppHandler))
	//	router.PathPrefix("/css").Handler(http.FileServer(http.Dir("/web/css")))
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	router.HandleFunc("/role", Middleware(kaau.ViewRolePageHandler))
	router.HandleFunc("/clusterrole", Middleware(kaau.ViewClusterRolePageHandler))
	router.HandleFunc("/serviceaccount", Middleware(kaau.ViewSAPageHandler))
	router.HandleFunc("/rolebinding", Middleware(kaau.ViewRoleBindingPageHandler))
	router.HandleFunc("/clusterrolebinding", Middleware(kaau.ViewClusterRoleBindingPageHandler))
	router.HandleFunc("/managerole", Middleware(kaau.MangeRolePOSTHandler)).Methods("POST")
	router.HandleFunc("/managerole", Middleware(kaau.ManageRolePageHandler))
	router.HandleFunc("/manageclusterrole", Middleware(kaau.MangeClusterRolePOSTHandler)).Methods("POST")
	router.HandleFunc("/manageclusterrole", Middleware(kaau.ManageClusterRolePageHandler))
	router.HandleFunc("/managerolebinding", Middleware(kaau.ManageRolebBindingPOSTHandler)).Methods("POST")
	router.HandleFunc("/managerolebinding", Middleware(kaau.ManageRoleBindingPageHandler))
	router.HandleFunc("/manageclusterrolebinding", Middleware(kaau.ManageClusterRolebBindingPOSTHandler)).Methods("POST")
	router.HandleFunc("/manageclusterrolebinding", Middleware(kaau.ManageClusterRoleBindingPageHandler))
	router.HandleFunc("/manageserviceaccount", Middleware(kaau.ManageServiceAccountPOSTHandler)).Methods("POST")
	router.HandleFunc("/manageserviceaccount", Middleware(kaau.ManageServiceAccountPageHandler))

	return router

}

package services

import (
	"gomscode/src/kaau"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	staticDir = "/static/"
)

// NewRouter exported
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", Middleware(kaau.LoginHandler)).Methods("POST")
	router.HandleFunc("/logout", Middleware(kaau.LogoutPageHandler))
	router.HandleFunc("/index", Middleware(kaau.HomeHandler))
	router.HandleFunc("/", Middleware(kaau.LoginPageHandler))
	router.HandleFunc("/app", Middleware(kaau.AppHandler))
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

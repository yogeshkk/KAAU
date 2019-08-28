package kaau

import (
	"fmt"
	"gomscode/src/utility"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ghodss/yaml"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ServiceAccountDetails exported
type ServiceAccountDetails struct {
	Name      string
	NameSpace string
}

// DataServiceAccountPage exported
type DataServiceAccountPage struct {
	UserName       string
	ServiceAccount []ServiceAccountDetails
}

// DataManageServiceAccountPage exported
type DataManageServiceAccountPage struct {
	UserName   string
	Action     string
	ErrMessage string
}

// ViewSAPageHandler exported
func ViewSAPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		SADetails := GetServiceAccount(kubeclient)
		SADetails.UserName = userName

		tmpl := template.Must(template.ParseFiles("web/templetes/serviceaccount.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, SADetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// GetServiceAccount exported
func GetServiceAccount(clntset *kubernetes.Clientset) DataServiceAccountPage {
	var DataSA DataServiceAccountPage
	AllNamespaces, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range AllNamespaces.Items {
		ServiceAccount, err := clntset.CoreV1().ServiceAccounts(namespace.Name).List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		for _, items := range ServiceAccount.Items {
			/*		Rule := items.Rules
					y, err := yaml.Marshal(Rule)
					if err != nil {
						log.Fatal(err)
					}
			*/
			DataSA.ServiceAccount = append(DataSA.ServiceAccount, ServiceAccountDetails{
				Name:      items.Name,
				NameSpace: items.Namespace,
			})
		}
	}
	return DataSA
}

// router.HandleFunc("/manageserviceaccount", Middleware(kaau.ManageServiceAccountPOSTHandler)).Methods("POST")
// router.HandleFunc("/manageserviceaccount", Middleware(kaau.ManageServiceAccountPageHandler))

// ManageServiceAccountPageHandler  exported
func ManageServiceAccountPageHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageServiceAccountPage
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		ManagePage.UserName = userName
		ManagePage.Action = r.URL.Query().Get("action")
		tmpl := template.Must(template.ParseFiles("web/templetes/manageserviceaccount.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageServiceAccountPOSTHandler exported
func ManageServiceAccountPOSTHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageServiceAccountPage
	userName := utility.GetUserName(r)
	action := r.FormValue("submit")
	kubeclient := GetKubeClient(Kubeconfig)
	name := r.FormValue("name")
	name = strings.TrimSpace(name)
	namespace := r.FormValue("namespaces")
	namespace = strings.TrimSpace(namespace)
	roles := r.FormValue("rule")
	roles = strings.TrimSpace(roles)
	fmt.Println(name, namespace, roles, action)
	ManagePage.UserName = userName
	ManagePage.Action = r.URL.Query().Get("action")
	if !utility.IsEmpty(userName) {
		ManagePage.ErrMessage = ManageServiceAccount(kubeclient, name, namespace, roles, action)
		tmpl := template.Must(template.ParseFiles("web/templetes/manageserviceaccount.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageServiceAccount exported
func ManageServiceAccount(clntset *kubernetes.Clientset, name, namespace, roles, action string) string {
	///Create(*v1.RoleBinding) (*v1.RoleBinding, error)
	success := "Service Account Created Successfully"

	yamlString := `
apiVersion: v1
kind: ServiceAccount
metadata:
  name: ASECUBHKOMEH
  namespace: LKJOIHJJBWDJHB
`
	yamlString = strings.Replace(yamlString, "ASECUBHKOMEH", name, 1)
	yamlString = strings.Replace(yamlString, "LKJOIHJJBWDJHB", namespace, 1)

	fmt.Println(yamlString)

	v1ServiceAccount := new(v1.ServiceAccount)
	err := yaml.Unmarshal([]byte(yamlString), &v1ServiceAccount)
	if err != nil {
		fmt.Println(err)
		return "ERROR: Can not Unmarshall YAML Please Check the YAML syntax"
	}
	fmt.Println(v1ServiceAccount)

	if action == "create" {
		// roleOut, err := clntset.RbacV1().Roles(namespace).Create(v1ServiceAccount)
		roleOut, err := clntset.CoreV1().ServiceAccounts(namespace).Create(v1ServiceAccount)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)

	} else if action == "update" {
		roleOut, err := clntset.CoreV1().ServiceAccounts(namespace).Update(v1ServiceAccount)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)
		success = "Service Account Modified Successfully"
	} else if action == "delete" {
		RoleOptions := new(metav1.DeleteOptions)
		roleOut := clntset.CoreV1().ServiceAccounts(namespace).Delete(name, RoleOptions)
		if roleOut != nil {
			fmt.Println(roleOut)
			success = "ERROR: " + roleOut.Error()
			return success
		}
		success = "Service Account Deleted Successfully"
	}
	return success
}

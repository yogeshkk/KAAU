package kaau

import (
	"fmt"
	"gomscode/src/utility"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ghodss/yaml"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClusterRoleBindingDetails exported
type ClusterRoleBindingDetails struct {
	Name      string
	NameSpace string
	Role      string
	Account   string
}

// DataClusterRoleBindingPage exported
type DataClusterRoleBindingPage struct {
	UserName     string
	RoleBindings []ClusterRoleBindingDetails
}

// DataManageClusterRoleBindingPage exported
type DataManageClusterRoleBindingPage struct {
	UserName   string
	Action     string
	ErrMessage string
}

// ViewClusterRoleBindingPageHandler exported
func ViewClusterRoleBindingPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		ClusterRBDetails := GetClusterRoleBinding(kubeclient)
		ClusterRBDetails.UserName = userName

		tmpl := template.Must(template.ParseFiles("web/templetes/clusterrolebinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ClusterRBDetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// GetClusterRoleBinding exported
func GetClusterRoleBinding(clntset *kubernetes.Clientset) DataClusterRoleBindingPage {
	var DataClusterRoleBinding DataClusterRoleBindingPage
	ClusteRoleBindinag, err := clntset.RbacV1().ClusterRoleBindings().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, items := range ClusteRoleBindinag.Items {
		Role := items.RoleRef
		r, err := yaml.Marshal(Role)
		if err != nil {
			log.Fatal(err)
		}
		Account := items.Subjects
		a, err := yaml.Marshal(Account)
		if err != nil {
			log.Fatal(err)
		}
		DataClusterRoleBinding.RoleBindings = append(DataClusterRoleBinding.RoleBindings, ClusterRoleBindingDetails{
			Name:    items.Name,
			Role:    string(r),
			Account: string(a),
		})
	}
	return DataClusterRoleBinding
}

// router.HandleFunc("/managerolebinding", Middleware(kaau.ManageClusterRolebBindingPOSTHandler)).Methods("POST")
// router.HandleFunc("/managerolebinding", Middleware(kaau.ManageClusterRoleBindingPageHandler))

// ManageClusterRoleBindingPageHandler exported
func ManageClusterRoleBindingPageHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageClusterRoleBindingPage
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		ManagePage.UserName = userName
		ManagePage.Action = r.URL.Query().Get("action")
		tmpl := template.Must(template.ParseFiles("web/templetes/manageclusterrolebinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageClusterRolebBindingPOSTHandler exported
func ManageClusterRolebBindingPOSTHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageClusterRoleBindingPage
	userName := utility.GetUserName(r)
	action := r.FormValue("submit")
	kubeclient := GetKubeClient(Kubeconfig)
	name := r.FormValue("name")
	name = strings.TrimSpace(name)
	roleref := r.FormValue("rule")
	roleref = strings.TrimSpace(roleref)
	account := r.FormValue("account")
	account = strings.TrimSpace(account)
	fmt.Println(name, roleref, account)
	ManagePage.UserName = userName
	ManagePage.Action = r.URL.Query().Get("action")
	if !utility.IsEmpty(userName) {

		ManagePage.ErrMessage = ManageClusterRoleBinding(kubeclient, name, roleref, account, action)
		tmpl := template.Must(template.ParseFiles("web/templetes/manageclusterrolebinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageClusterRoleBinding exported
func ManageClusterRoleBinding(clntset *kubernetes.Clientset, name, roleref, account, action string) string {
	///Create(*v1.RoleBinding) (*v1.RoleBinding, error)
	success := "Cluster Role Binding Created Successfully"

	yamlString := `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: ASECUBHKOMEH
`
	yamlString = strings.Replace(yamlString, "ASECUBHKOMEH", name, 1)
	yamlString = yamlString + roleref
	//	yamlString = yamlString + "\nsubjects:"
	yamlString = yamlString + "\n" + account

	fmt.Println(yamlString)

	//v1Role := new(v1.Role)
	v1ClusterRoleBinding := new(v1.ClusterRoleBinding)
	err := yaml.Unmarshal([]byte(yamlString), &v1ClusterRoleBinding)
	if err != nil {
		fmt.Println(err)
		return "ERROR: Can not Unmarshall YAML Please Check the YAML syntax"
	}
	fmt.Println(v1ClusterRoleBinding)

	if action == "create" {
		roleOut, err := clntset.RbacV1().ClusterRoleBindings().Create(v1ClusterRoleBinding)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)

	} else if action == "update" {
		roleOut, err := clntset.RbacV1().ClusterRoleBindings().Update(v1ClusterRoleBinding)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)
		success = "Cluster Role Binding Modified Successfully"
	} else if action == "delete" {
		RoleOptions := new(metav1.DeleteOptions)
		roleOut := clntset.RbacV1().ClusterRoleBindings().Delete(name, RoleOptions)
		if roleOut != nil {
			fmt.Println(roleOut)
			success = "ERROR: " + roleOut.Error()
			return success
		}
		success = "Cluster Role Binding Deleted Successfully"
	}
	return success
}

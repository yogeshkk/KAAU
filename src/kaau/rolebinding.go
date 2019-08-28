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

// RoleBindingDetails exported
type RoleBindingDetails struct {
	Name      string
	NameSpace string
	Role      string
	Account   string
}

// DataRoleBindingPage exported
type DataRoleBindingPage struct {
	UserName     string
	RoleBindings []RoleBindingDetails
}

// DataManageRoleBindingPage exported
type DataManageRoleBindingPage struct {
	UserName   string
	Action     string
	ErrMessage string
}

// ViewRoleBindingPageHandler exported
func ViewRoleBindingPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		RBDetails := GetRoleBinding(kubeclient)
		RBDetails.UserName = userName

		tmpl := template.Must(template.ParseFiles("web/templetes/rolesbinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, RBDetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// GetRoleBinding exported
func GetRoleBinding(clntset *kubernetes.Clientset) DataRoleBindingPage {
	var DataRolesBinding DataRoleBindingPage
	AllNamespaces, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range AllNamespaces.Items {
		NamespaceRoleBinding, err := clntset.RbacV1().RoleBindings(namespace.Name).List(metav1.ListOptions{})

		if err != nil {
			log.Fatal(err)
		}
		for _, items := range NamespaceRoleBinding.Items {
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
			DataRolesBinding.RoleBindings = append(DataRolesBinding.RoleBindings, RoleBindingDetails{
				Name:      items.Name,
				NameSpace: namespace.Name,
				Role:      string(r),
				Account:   string(a),
			})
		}
	}
	return DataRolesBinding
}

// ManageRoleBindingPageHandler exported
func ManageRoleBindingPageHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageRoleBindingPage
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		ManagePage.UserName = userName
		ManagePage.Action = r.URL.Query().Get("action")
		tmpl := template.Must(template.ParseFiles("web/templetes/managerolebinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageRolebBindingPOSTHandler exported
func ManageRolebBindingPOSTHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageRoleBindingPage
	userName := utility.GetUserName(r)
	action := r.FormValue("submit")
	kubeclient := GetKubeClient(Kubeconfig)
	name := r.FormValue("name")
	name = strings.TrimSpace(name)
	namespace := r.FormValue("namespaces")
	roleref := r.FormValue("rule")
	roleref = strings.TrimSpace(roleref)
	account := r.FormValue("account")
	account = strings.TrimSpace(account)
	fmt.Println(name, roleref, account)
	ManagePage.UserName = userName
	ManagePage.Action = r.URL.Query().Get("action")
	if !utility.IsEmpty(userName) {

		ManagePage.ErrMessage = ManageRoleBinding(kubeclient, name, namespace, roleref, account, action)
		tmpl := template.Must(template.ParseFiles("web/templetes/managerolebinding.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageRoleBinding exported
func ManageRoleBinding(clntset *kubernetes.Clientset, name, namespace, roleref, account, action string) string {
	///Create(*v1.RoleBinding) (*v1.RoleBinding, error)
	success := "Role Binding Created Successfully"

	yamlString := `
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: ASECUBHKOMEH
  namespace: LKJOIHJJBWDJHB
`
	yamlString = strings.Replace(yamlString, "ASECUBHKOMEH", name, 1)
	yamlString = strings.Replace(yamlString, "LKJOIHJJBWDJHB", namespace, 1)
	yamlString = yamlString + roleref
	//	yamlString = yamlString + "\nsubjects:"
	yamlString = yamlString + "\n" + account

	fmt.Println(yamlString)

	//v1Role := new(v1.Role)
	v1RoleBinding := new(v1.RoleBinding)
	err := yaml.Unmarshal([]byte(yamlString), &v1RoleBinding)
	if err != nil {
		fmt.Println(err)
		return "ERROR: Can not Unmarshall YAML Please Check the YAML syntax"
	}
	fmt.Println(v1RoleBinding)

	if action == "create" {
		roleOut, err := clntset.RbacV1().RoleBindings(namespace).Create(v1RoleBinding)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)

	} else if action == "update" {
		roleOut, err := clntset.RbacV1().RoleBindings(namespace).Update(v1RoleBinding)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)
		success = "Role Binding Modified Successfully"
	} else if action == "delete" {
		RoleOptions := new(metav1.DeleteOptions)
		roleOut := clntset.RbacV1().RoleBindings(namespace).Delete(name, RoleOptions)
		if roleOut != nil {
			fmt.Println(roleOut)
			success = "ERROR: " + roleOut.Error()
			return success
		}
		success = "Role Binding Deleted Successfully"
	}
	return success
}

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

// DataManageRolePage exported
type DataManageRolePage struct {
	UserName   string
	Action     string
	ErrMessage string
}

// RoleDetails exported
type RoleDetails struct {
	Name      string
	Rule      string
	NameSpace string
}

// DataRolePage  exported
type DataRolePage struct {
	UserName   string
	NameSpeces []NameSpaceDetails
	Roles      []RoleDetails
}

//ViewRolePageHandler exported
func ViewRolePageHandler(w http.ResponseWriter, r *http.Request) {
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		kubeclient := GetKubeClient(Kubeconfig)
		RolesDetails := GetRoles(kubeclient)
		RolesDetails.UserName = userName
		tmpl := template.Must(template.ParseFiles("web/templetes/roles.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, RolesDetails)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// ManageRolePageHandler  exported
func ManageRolePageHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageRolePage
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		ManagePage.UserName = userName
		ManagePage.Action = r.URL.Query().Get("action")
		tmpl := template.Must(template.ParseFiles("web/templetes/managerole.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// MangeRolePOSTHandler exported
func MangeRolePOSTHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageRolePage
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

		ManagePage.ErrMessage = ManageRoles(kubeclient, name, namespace, roles, action)
		tmpl := template.Must(template.ParseFiles("web/templetes/managerole.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// GetRoles exported
func GetRoles(clntset *kubernetes.Clientset) DataRolePage {
	var DataRoles DataRolePage
	AllNamespaces, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range AllNamespaces.Items {
		NamespaceRole, err := clntset.RbacV1().Roles(namespace.Name).List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		for _, items := range NamespaceRole.Items {
			Rule := items.Rules
			y, err := yaml.Marshal(Rule)
			if err != nil {
				log.Fatal(err)
			}
			DataRoles.Roles = append(DataRoles.Roles, RoleDetails{
				Name:      items.Name,
				NameSpace: namespace.Name,
				Rule:      string(y),
			})
		}
	}
	return DataRoles
}

// ManageRoles exported
func ManageRoles(clntset *kubernetes.Clientset, name, namespace, roles, action string) string {
	///Create(*v1.RoleBinding) (*v1.RoleBinding, error)
	success := "Role Created Successfully"

	yamlString := `
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: ASECUBHKOMEH
  namespace: LKJOIHJJBWDJHB
`
	yamlString = strings.Replace(yamlString, "ASECUBHKOMEH", name, 1)
	yamlString = strings.Replace(yamlString, "LKJOIHJJBWDJHB", namespace, 1)
	yamlString = yamlString + roles

	fmt.Println(yamlString)

	v1Role := new(v1.Role)
	err := yaml.Unmarshal([]byte(yamlString), &v1Role)
	if err != nil {
		fmt.Println(err)
		return "ERROR: Can not Unmarshall YAML Please Check the YAML syntax"
	}
	fmt.Println(v1Role)

	if action == "create" {
		roleOut, err := clntset.RbacV1().Roles(namespace).Create(v1Role)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)

	} else if action == "update" {
		roleOut, err := clntset.RbacV1().Roles(namespace).Update(v1Role)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(roleOut)
		success = "Role Modified Successfully"
	} else if action == "delete" {
		RoleOptions := new(metav1.DeleteOptions)
		roleOut := clntset.RbacV1().Roles(namespace).Delete(name, RoleOptions)
		if roleOut != nil {
			fmt.Println(roleOut)
			success = "ERROR: " + roleOut.Error()
			return success
		}
		success = "Role Deleted Successfully"
	}
	return success
}

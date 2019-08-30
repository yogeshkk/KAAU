package kaau

import (
	"html/template"
	"fmt"
	"strings"
	"gomscode/src/utility"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"github.com/ghodss/yaml"
	"k8s.io/client-go/kubernetes"
)
// DataManageClusterRolePage exported
type DataManageClusterRolePage struct {
	UserName   string
	Action     string
	ErrMessage string
}
// ClusterRoleDetails exported
type ClusterRoleDetails struct {
	Name string
	Rule string
}

// DataClusterRolePage exported
type DataClusterRolePage struct {
	UserName   string
	NameSpeces []NameSpaceDetails
	Roles      []ClusterRoleDetails
}


// GetClusterRole exported
func GetClusterRole(clntset *kubernetes.Clientset) DataClusterRolePage {
	var DataClusterRole DataClusterRolePage
	ClusteRole, err := clntset.RbacV1().ClusterRoles().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, items := range ClusteRole.Items {
		Rule := items.Rules
		y, err := yaml.Marshal(Rule)
		if err != nil {
			fmt.Println(err)
		}
		DataClusterRole.Roles = append(DataClusterRole.Roles, ClusterRoleDetails{
			Name: items.Name,
			Rule: string(y),
		})
	}
	return DataClusterRole
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
// ManageClusterRoles exported
func ManageClusterRoles(clntset *kubernetes.Clientset, name, roles, action string) string {
	///Create(*v1.RoleBinding) (*v1.RoleBinding, error)
	success := "Cluster Role Created Successfully"

	yamlString := `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: ASECUBHKOMEH
`
	yamlString = strings.Replace(yamlString, "ASECUBHKOMEH", name, 1)

	yamlString = yamlString + roles

//	fmt.Println(yamlString)

	v1ClusterRole := new(v1.ClusterRole)
	err := yaml.Unmarshal([]byte(yamlString), &v1ClusterRole)
	if err != nil {
		fmt.Println(err)
		return "ERROR: Can not Unmarshall YAML Please Check the YAML syntax"
	}
//	fmt.Println(v1ClusterRole)

	if action == "create" {
//		roleOut, err := clntset.RbacV1().Roles(namespace).Create(v1ClusterRole)
		ClusteRoleout , err := clntset.RbacV1().ClusterRoles().Create(v1ClusterRole)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(ClusteRoleout)

	} else if action == "update" {
		ClusteRoleout , err := clntset.RbacV1().ClusterRoles().Update(v1ClusterRole)
		if err != nil {
			fmt.Println(err)
			success = "ERROR: " + err.Error()
			return success
		}
		fmt.Println(ClusteRoleout)
		success = "Cluster Role Modified Successfully"
	} else if action == "delete" {
		RoleOptions := new(metav1.DeleteOptions)
		ClusteRoleout := clntset.RbacV1().ClusterRoles().Delete(name, RoleOptions)
		if ClusteRoleout != nil {
			fmt.Println(ClusteRoleout)
			success = "ERROR: " + ClusteRoleout.Error()
			return success
		}
		success = "Cluster Role Deleted Successfully"
	}
	return success
}


// ManageClusterRolePageHandler exported
func ManageClusterRolePageHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageClusterRolePage
	userName := utility.GetUserName(r)
	if !utility.IsEmpty(userName) {
		ManagePage.UserName = userName
		ManagePage.Action = r.URL.Query().Get("action")
		tmpl := template.Must(template.ParseFiles("web/templetes/manageclusterrole.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

// MangeClusterRolePOSTHandler exported
func MangeClusterRolePOSTHandler(w http.ResponseWriter, r *http.Request) {
	var ManagePage DataManageRolePage
	userName := utility.GetUserName(r)
	action := r.FormValue("submit")
	kubeclient := GetKubeClient(Kubeconfig)
	name := r.FormValue("name")
	name = strings.TrimSpace(name)
	roles := r.FormValue("rule")
	roles = strings.TrimSpace(roles)
//	fmt.Println(name,  roles, action)
	ManagePage.UserName = userName
	ManagePage.Action = r.URL.Query().Get("action")
	if !utility.IsEmpty(userName) {

		ManagePage.ErrMessage = ManageClusterRoles(kubeclient, name, roles, action)
		tmpl := template.Must(template.ParseFiles("web/templetes/manageclusterrole.html"))
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, ManagePage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

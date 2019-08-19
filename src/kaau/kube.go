package kaau

import (
	"log"

	"github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// NameSpaceDetails exported
type NameSpaceDetails struct {
	Name string
}

// RoleDetails exported
type RoleDetails struct {
	Name      string
	Rule      string
	NameSpace string
}

// ServiceAccountDetails exported
type ServiceAccountDetails struct {
	Name      string
	NameSpace string
}

// RoleBindingDetails exported
type RoleBindingDetails struct {
	Name      string
	NameSpace string
	Role	string
	Account string
}

// ClusterRoleBindingDetails exported
type ClusterRoleBindingDetails struct {
	Name      string
	NameSpace string
	Role	string
	Account string
}

// ClusterRoleDetails exported
type ClusterRoleDetails struct {
	Name string
	Rule string

}

// DataIndexPage exported
type DataIndexPage struct {
	UserName string
}

// DataRolePage  exported
type DataRolePage struct {
	UserName   string
	NameSpeces []NameSpaceDetails
	Roles      []RoleDetails
}

// DataClusterRolePage exported
type DataClusterRolePage struct {
	UserName   string
	NameSpeces []NameSpaceDetails
	Roles      []ClusterRoleDetails
}

// DataServiceAccountPage exported
type DataServiceAccountPage struct {
	UserName       string
	ServiceAccount []ServiceAccountDetails
}

// DataRoleBindingPage exported
type DataRoleBindingPage struct {
	UserName     string
	RoleBindings []RoleBindingDetails
}

// DataClusterRoleBindingPage exported
type DataClusterRoleBindingPage struct {
	UserName     string
	RoleBindings []ClusterRoleBindingDetails
}

// GetKubeClient Exported
func GetKubeClient(kubeconfig string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return clientset
}

// GetNameSpaces exported
func GetNameSpaces(clntset *kubernetes.Clientset) NameSpaceDetails {
	var NameSpaces NameSpaceDetails
	namespace, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range namespace.Items {
		NameSpaces.Name = namespace.Name
	}
	return NameSpaces
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

// GetClusterRole exported
func GetClusterRole(clntset *kubernetes.Clientset) DataClusterRolePage {
	var DataClusterRole DataClusterRolePage
	ClusteRole, err := clntset.RbacV1().ClusterRoles().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, items := range ClusteRole.Items {
		Rule := items.Rules
		y, err := yaml.Marshal(Rule)
		if err != nil {
			log.Fatal(err)
		}
		DataClusterRole.Roles = append(DataClusterRole.Roles, ClusterRoleDetails{
			Name: items.Name,
			Rule: string(y),
		})
	}
	return DataClusterRole
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
				Name: items.Name,
				NameSpace: namespace.Name,
				Role: string(r),
				Account: string(a),
			})
		}
	}
	return DataRolesBinding
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
			Name: items.Name,
			Role: string(r),
			Account: string(a),
		})
	}
	return DataClusterRoleBinding
}


// GetUserAccountBinding exported
func GetUserAccount(clntset *kubernetes.Clientset) DataClusterRoleBindingPage {
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
			Name: items.Name,
			Role: string(r),
			Account: string(a),
		})
	}
	return DataClusterRoleBinding
}
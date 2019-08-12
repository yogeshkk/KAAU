package kaau

import (
	"fmt"
	"log"

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
	Name string
}

// ClusterRoleDetails exported
type ClusterRoleDetails struct {
	Name string
}

// DataIndexPage exported
type DataIndexPage struct {
	UserName    string
	Host        string
	NameSpeces  []NameSpaceDetails
	ClusterRole []ClusterRoleDetails
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
func GetNameSpaces(clntset *kubernetes.Clientset) DataIndexPage {
	var DataIndex DataIndexPage
	namespace, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, namespace := range namespace.Items {
		NamespaceRole, err := clntset.RbacV1().Roles(namespace.Name).List(metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		for _, role := range NamespaceRole.Items {
			DataIndex.NameSpeces = append(DataIndex.NameSpeces, NameSpaceDetails{

				Name: namespace.Name,
			})
			DataIndex.ClusterRole = append(DataIndex.ClusterRole, ClusterRoleDetails{

				Name: role.Name,
			})
			fmt.Println(namespace.Name, role.Name)
		}
	}
	return DataIndex
}

// GetRoles exported
func GetRoles(clntset *kubernetes.Clientset, Namespace string) DataIndexPage {
	var DataRoles DataIndexPage
	AllNamespace, err := clntset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, items := range AllNamespace.Items {
		DataRoles.NameSpeces = append(DataRoles.NameSpeces, NameSpaceDetails{
			Name: items.Name,
		})
	}
	NamespaceRole, err := clntset.RbacV1().Roles(Namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, items := range NamespaceRole.Items {
		DataRoles.ClusterRole = append(DataRoles.ClusterRole, ClusterRoleDetails{
			Name: items.Name,
		})
	}
	return DataRoles
}

// GetClusterRole exported
func GetClusterRole(clntset *kubernetes.Clientset) DataIndexPage {
	var DataClusterRole DataIndexPage
	ClusteRole, err := clntset.RbacV1().ClusterRoles().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, items := range ClusteRole.Items {
		DataClusterRole.ClusterRole = append(DataClusterRole.ClusterRole, ClusterRoleDetails{
			Name: items.Name,
		})
	}
	return DataClusterRole
}

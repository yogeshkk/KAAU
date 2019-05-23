package main

import (
	"reflect"
	"k8s.io/client-go/tools/clientcmd/api"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var ns, label, field, maxClaims string
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	flag.StringVar(&ns, "namespace", "", "namespace")
	flag.StringVar(&label, "l", "", "Label selector")
	flag.StringVar(&field, "f", "", "Field selector")
	flag.StringVar(&maxClaims, "max-claims", "100Gi", "Maximum total claims to watch")
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()

	// total resource quantities
	//	var totalClaimedQuant resource.Quantity
	//	maxClaimedQuant := resource.MustParse(maxClaims)

	// bootstrap config
	fmt.Println()
	fmt.Println("Using kubeconfig: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	api := clientset.CoreV1()
	fmt.Println(reflect.TypeOf(api))
	// initial list
	listOptions := metav1.ListOptions{LabelSelector: label, FieldSelector: field}

	namespace, err := api.Namespaces().List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range namespace.Items {
		fmt.Printf("%s\n", v.Name)
	//	getpods(api)

	}

	/*pvcs, err := api.PersistentVolumeClaims(ns).List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	*/

	//	printPVCs(pvcs)
//	fmt.Println()

	// watch future changes to PVCs
	/*	watcher, err := clientset.CoreV1().PersistentVolumeClaims(ns).Watch(listOptions)
		if err != nil {
			log.Fatal(err)
		}
		ch := watcher.ResultChan()

		fmt.Printf("--- PVC Watch (max claims %v) ----\n", maxClaimedQuant.String())
		for event := range ch {
			pvc, ok := event.Object.(*v1.PersistentVolumeClaim)
			if !ok {
				log.Fatal("unexpected type")
			}
			quant := pvc.Spec.Resources.Requests[v1.ResourceStorage]

			switch event.Type {
			case watch.Added:
				totalClaimedQuant.Add(quant)
				log.Printf("PVC %s added, claim size %s\n", pvc.Name, quant.String())

				// is claim overage?
				if totalClaimedQuant.Cmp(maxClaimedQuant) == 1 {
					log.Printf("\nClaim overage reached: max %s at %s",
						maxClaimedQuant.String(),
						totalClaimedQuant.String(),
					)
					// trigger action
					log.Println("*** Taking action ***")
				}

			case watch.Modified:
				//log.Printf("Pod %s modified\n", pod.GetName())
			case watch.Deleted:
				quant := pvc.Spec.Resources.Requests[v1.ResourceStorage]
				totalClaimedQuant.Sub(quant)
				log.Printf("PVC %s removed, size %s\n", pvc.Name, quant.String())

				if totalClaimedQuant.Cmp(maxClaimedQuant) <= 0 {
					log.Printf("Claim usage normal: max %s at %s",
						maxClaimedQuant.String(),
						totalClaimedQuant.String(),
					)
					// trigger action
					log.Println("*** Taking action ***")
				}
			case watch.Error:
				//log.Printf("watcher error encountered\n", pod.GetName())
			}

			log.Printf("\nAt %3.1f%% claim capcity (%s/%s)\n",
				float64(totalClaimedQuant.Value())/float64(maxClaimedQuant.Value())*100,
				totalClaimedQuant.String(),
				maxClaimedQuant.String(),
			)
		}
	*/

}

func getpods(api string) {
	
	//	pods, err := clientset.CoreV1().Pods(v.Name).List(metav1.ListOptions{})
	
	pods, err := api.Pods(v.Name).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, p := range pods.Items {
		fmt.Printf("%s\n", p.Name)
	}
	fmt.Printf("There are %d pods in the %s\n", len(pods.Items), v.Name)

}

// printPVCs prints a list of PersistentVolumeClaim on console
func printPVCs(pvcs *v1.PersistentVolumeClaimList) {
	if len(pvcs.Items) == 0 {
		log.Println("No claims found")
		return
	}
	template := "%-32s%-8s%-8s\n"
	fmt.Println("--- PVCs ----")
	fmt.Printf(template, "NAME", "STATUS", "CAPACITY")
	var cap resource.Quantity
	for _, pvc := range pvcs.Items {
		quant := pvc.Spec.Resources.Requests[v1.ResourceStorage]
		cap.Add(quant)
		fmt.Printf(template, pvc.Name, string(pvc.Status.Phase), quant.String())
	}

	fmt.Println("-----------------------------")
	fmt.Printf("Total capacity claimed: %s\n", cap.String())
	fmt.Println("-----------------------------")
}

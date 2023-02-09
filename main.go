package main

import (
	"context"
	"fmt"
	// "reflect"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Create a new Kubernetes client using ClusterInConfig
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("Error Connecting to config:", err)
		return
	}
  
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating clientset:", err)
		return
	}

	// // Get the list of PSPs in the cluster
	// psps, err := clientset.PolicyV1beta1().PodSecurityPolicies().List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println("Error getting PSPs:", err)
	// 	return
	// }

	// fmt.Println(psps)

	privileged := true
	hostPID := true
	hostNetwork := true

	// Attempt to create a pod that violates the PSP rules
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "test-container",
					Image: "test-image",
					SecurityContext: &corev1.SecurityContext{
						// Attempt to violate the PSP rules by running as privileged
						Privileged: &privileged,
					},
				},
			},
			// Attempt to violate the PSP rules by running as a host PID
			HostPID: hostPID,
			// Attempt to violate the PSP rules by using a host network
			HostNetwork: hostNetwork,
		},
	}

	psp, err := clientset.CoreV1().Pods("beta").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		// Check if the error is due to PSP violation
		// status, ok := err.
		// if ok && status.Code == 403 && status.Reason == "Forbidden" &&
		// 	reflect.DeepEqual(status.Details.Kind, "podsecuritypolicies") {
		// 	fmt.Println("PSP rule", psp.Name, "is enforced.")

		// } else {
		// 	// If no error, then the PSP rule was not enforced
		// 	fmt.Println("PSP rule", psp.Name, "is not enforced.")
		// 	// Delete the pod to clean up
		// 	clientset.CoreV1().Pods("*").Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
		// }
	}

	// Report on the current security level for multi-tenancy
	fmt.Println("Current security level for multi-tenancy: UNKNOWN", psp)
	fmt.Println("More analysis and testing is required to determine the exact security level.")
}

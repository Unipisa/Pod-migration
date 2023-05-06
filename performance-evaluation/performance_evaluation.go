package performance_evaluation

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"os"
)

func getResourceUsage(clientset *kubernetes.Clientset, podName string) (cpu, memory, disk float64, err error) {
	// Get pod's namespace
	pod, err := clientset.CoreV1().Pods("").Get(podName, metav1.GetOptions{})
	if err != nil {
		return 0, 0, 0, err
	}
	namespace := pod.Namespace

	// Get pod's containers
	podMetrics, err := clientset.MetricsV1beta1().PodMetricses(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return 0, 0, 0, err
	}

	// Sum resource usage across all containers
	for _, containerMetrics := range podMetrics.Containers {
		cpu += containerMetrics.Usage.Cpu().MilliValue()
		memory += float64(containerMetrics.Usage.Memory().Value()) / 1024 / 1024
	}

	// Get pod's disk usage by checking disk usage of its nodes
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return 0, 0, 0, err
	}
	for _, node := range nodes.Items {
		nodeDiskUsage, err := clientset.CoreV1().Nodes().ProxyGet("https", node.Name, "/stats/summary", nil).DoRaw()
		if err != nil {
			return 0, 0, 0, err
		}
		// TODO: Parse nodeDiskUsage for disk usage of the pod's containers
	}

	return cpu / 1000, memory, disk, nil
}

func main() {
	// Load Kubernetes configuration
	var config *rest.Config
	if os.Getenv("KUBECONFIG") != "" {
		kubeconfig := os.Getenv("KUBECONFIG")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create Metrics clientset
	metricsClientset, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Test getResourceUsage function
	podName := "test-pod-1-container"
	cpu, memory, disk, err := getResourceUsage(clientset, podName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pod %s CPU usage: %f\n", podName, cpu)
	fmt.Printf("Pod %s memory usage: %f MB\n", podName, memory)
	fmt.Printf("Pod %s disk usage: %f MB\n", podName, disk)
}

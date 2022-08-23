# Easily create Kubernetes clients in Go
This package provides an easy way to create a [Kubernetes client](https://pkg.go.dev/k8s.io/client-go/kubernetes#Clientset), whether with the default in-cluster configuration or by using `~/.kube/config` with the given context.

## Usage
```go
package main

import "github.com/inkel/kubeclient"

func main() {
	c, cfg, err := kubeclient.New(false, "my-context")
	if err != nil {
		println(err)
		return
	}

	// Do something with c and cfg
	_ = c
	_ = cfg
}
```

The first parameter indicates if you want to use the in-cluster configuration or not.
The second parameter only works when the first one is false, and allows to indicate which context from `~/.kube/config` to use.

## License
[MIT](LICENSE).

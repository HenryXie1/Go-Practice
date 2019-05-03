package main

import (
  "fmt"

  "k8s.io/api/extensions/v1beta1"
  "k8s.io/client-go/kubernetes/scheme"
)

var yml = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata: 
 name: testnginx
replicas: 1
spec: 
template:
  metadata:
    labels:
      run: testnginx
  spec:
    containers:
    - image: nginx
      name: testnginx
      ports:
      - containerPort: 8080
`

func main() {
	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, err := decode([]byte(yml), nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}

  //fmt.Printf("%#v\n", obj)
  deployment := obj.(*v1beta1.Deployment)
  
	fmt.Printf("%#v\n", deployment.ObjectMeta.Name)
}

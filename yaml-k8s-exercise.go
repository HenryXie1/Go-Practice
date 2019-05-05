package main
import (
  "fmt"
  "github.com/ghodss/yaml"
  "k8s.io/api/apps/v1"
)

func main() {
  var yml = `
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
name: testnginx
spec:
replicas: 2
template:
  metadata:
    labels:
      run: my-nginx
  spec:
    containers:
    - name: my-nginx
      image: nginx
      ports:
      - containerPort: 8080
`

  var spec v1.Deployment
  err := yaml.Unmarshal([]byte(yml), &spec)
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf("%T\n,%#v\n",spec.ObjectMeta.Name,spec.ObjectMeta.Name)
}

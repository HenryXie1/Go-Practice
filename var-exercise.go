package main

import (
	"fmt"
)

func main() {

var (
	namespaceExample = `
	# view the current namespace in your KUBECONFIG
	%[1]s setns

	# view all of the namespaces in use by contexts in your KUBECONFIG
	%[1]s setns --list

	# switch your current-context to one that contains the desired namespace
	%[1]s setns foo
	
	# test 2nd parameter
	%[2]s 2ndparameter
`
)


	a := fmt.Sprintf(namespaceExample, "kubectl","mycmd")
        fmt.Println(a)
}

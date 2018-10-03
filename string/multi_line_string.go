package main

import "fmt"

func main() {
	usage := `Open Container Initiative runtime

	runc is a command line client for running applications packaged according to
	the Open Container Initiative (OCI) format and is a compliant implementation of the
	Open Container Initiative specification.

	runc integrates well with existing process supervisors to provide a production
	container runtime environment for applications. It can be used with your
	existing process monitoring tools and the container will be spawned as a
	direct child of the process supervisor.

	Containers are configured using bundles. A bundle for a container is a directory
	that includes a specification file named "`
	fmt.Println(usage)
}

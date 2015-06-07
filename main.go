// ProcessKill project main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func runCommand(cmd string) (string, error) {
	res, err := exec.Command("/bin/sh", "-c", cmd).Output()
	return string(res), err
}

func getListeningProcessID(port string) string {
	if port == "" {
		fmt.Println("No port specified. Usage --port=8080")
		os.Exit(0)
	}
	cmd := fmt.Sprintf("lsof -nP -iTCP:%v -sTCP:LISTEN | awk -v i=2 -v j=2 'FNR == i {print $j}'", port)
	pid, err := runCommand(cmd)
	if err != nil || pid == "" {
		// Not found a process so exit
		fmt.Println("No process found listening on PORT=" + port)
		os.Exit(0)
	}
	return pid
}

func killProcess(pid string) {
	cmd := fmt.Sprintf("kill %v", pid)
	runCommand(cmd)
}

func main() {

	var (
		port string
		kill bool
	)

	// Parse port flag into variable
	flag.StringVar(&port, "p", "", "Listening port of the target process")
	flag.BoolVar(&kill, "k", false, "Wanna kill the process? [true/false]")

	// Parse flags
	flag.Parse()

	// Get the ID of the process listening on that port
	pid := getListeningProcessID(port)

	if kill {
		// Kill the process once we have the PID
		killProcess(pid)
	} else {
		// Just looking for the process PID
		fmt.Print(pid)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {
	bunny := []string{
		`         `,
		`(\__/)   `,
		`(='.'=)  `,
		`(")_(")  `,
	}
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	output := strings.Builder{}
	output.WriteString(bunny[0])
	output.WriteString(currentUser.Username)
	output.WriteString("@")
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	output.WriteString(hostname)
	output.WriteString("\n")

	output.WriteString(bunny[1])
	output.WriteString(runtime.GOARCH)
	output.WriteString("\n")

	output.WriteString(bunny[2])
	output.WriteString(runtime.GOOS)
	output.WriteString("\n")

	var info syscall.Sysinfo_t
	if err := syscall.Sysinfo(&info); err != nil {
		fmt.Printf("Error getting system uptime: %v\n", err)
		return
	}
	output.WriteString(bunny[3] + fmt.Sprintf("%v\n", time.Duration(info.Uptime)*time.Second))

	// output.WriteString(fmt.Sprintf("%v", info.Totalram))
	fmt.Print(output.String())
}

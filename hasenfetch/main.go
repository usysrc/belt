package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/mackerelio/go-osstat/uptime"
)

func getSystemUptime() (time.Duration, error) {
	uptimeDuration, err := uptime.Get()
	if err != nil {
		return 0, err
	}
	return uptimeDuration, nil
}

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

	uptime, err := getSystemUptime()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	output.WriteString(bunny[3] + fmt.Sprintf("%v\n", uptime))

	// output.WriteString(fmt.Sprintf("%v", info.Totalram))
	fmt.Print(output.String())
}

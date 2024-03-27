package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`\d+\s`)
	cmd := exec.Command("xinput")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Multi-Touch") {
			id := strings.TrimSpace(re.FindString(line))
			if err != nil {
				log.Fatal(err)
			}
			cmd2 := exec.Command("xinput", "disable", id)
			if err = cmd2.Run(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Se desactivó el dispositivo con id = %s exitosamente", id)
			}
		}
	}

}

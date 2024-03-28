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
			err := touchOff(id, false)
			if err != nil {
				log.Fatal(err)
			}

		}
	}

}

func touchOff(id string, revert bool) error {
	if !revert {
		cmd := exec.Command("xinput", "disable", id)

		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("Se desactivó el dispositivo con id = %s exitosamente\n", id)
		}
	} else {
		cmd := exec.Command("xinput", "enable", id)

		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("Se reactivó el dispositivo con id = %s exitosamente\n", id)
		}
	}

}

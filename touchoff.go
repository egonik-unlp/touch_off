package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`\d+\s`)
	cmd := exec.Command("xinput")
	out, err := cmd.Output()
	args := os.Args
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
			if len(args) > 1 {
				if args[1] == "enable" {
					err = touchOff(id, true)
					fmt.Println("Se reactivo la pantalla táctil")
				} else if args[1] == "disable" {
					err = touchOff(id, false)
					fmt.Println("Se desactivó la pantalla táctil")
				} else {
					fmt.Printf("El comando %s no existe\n", args[1])

				}
			} else {
				err = touchOff(id, false)
				fmt.Println("Se desactivo la pantalla táctil")
			}

			if err != nil {
				log.Fatal(err)
			}

		}
	}

}

func touchOff(id string, revert bool) error {
	argArray := []string{id}
	if !revert {
		field := []string{"disable"}
		argArray = append(field, argArray...)
	} else {
		field := []string{"enable"}
		argArray = append(field, argArray...)
	}
	cmd := exec.Command("xinput", argArray...)
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

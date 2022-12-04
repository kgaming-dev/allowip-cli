package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "allowip",
		Usage: "Defines ufw allow rules from a list of ip addresses",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Text file contains a list of Ip Addresses in form of 'ipaddress:?comment' per line. ':' and 'comment' can be omitted."},
		},
		Action: func(cCtx *cli.Context) error {
			filePath := cCtx.String("file")
			if filePath == "" {
				displayError("File is required -f --file")
				cli.Exit("File is required.", 0)
				return nil
			} else {
				ipsFile, err := os.Open(cCtx.String("file"))
				if err != nil {
					msg := fmt.Sprintf("File %s is not exists", filePath)
					displayError(msg)
					return nil
				} else {
					scanner := bufio.NewScanner(ipsFile)
					scanner.Split(bufio.ScanLines)

					var fileLines []string

					for scanner.Scan() {
						fileLines = append(fileLines, scanner.Text())
					}

					var line int
					for key, value := range fileLines {
						if key == 0 {
							line = 1
						}
						if value != "" {
							var ipAddr string
							var comment string
							split := strings.Split(value, ":")
							if len(split) > 1 {
								ipAddr = split[0]
								comment = split[1]
								allowMsg := fmt.Sprintf("%s Allowing %s with comment %s", strconv.Itoa(line), ipAddr, comment)
								displaySuccess(allowMsg)
							} else {
								ipAddr = value
								comment = ""
								allowMsg := fmt.Sprintf("%s Allowing %s", strconv.Itoa(line), ipAddr)
								displaySuccess(allowMsg)
							}
							_, allowErr := AllowIpAddr(ipAddr, comment)
							if allowErr != nil {
								cli.Exit("Error", 0)
								return nil
							}
						}
						line++
					}
				}

				if err = ipsFile.Close(); err != nil {
					msg := fmt.Sprintf("Could not close the file due to this %s error \n", err)
					displayError(msg)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func AllowIpAddr(ip string, comment string) (bool, error) {
	ufwCmd, lookErr := exec.LookPath("ufw")

	if lookErr != nil {
		displayError("Command ufw is not found.")
		return false, errors.New("Command is not found")
	}

	arg := "ufw allow to any from " + ip

	if comment != "" {
		arg += " comment " + comment
	}

	args := []string{arg}
	args.append(arg)
	env := os.Environ()
	execErr := syscall.Exec(ufwCmd, args, env)
	if execErr != nil {
		displayError("Error while executing the comand.")
		return false, errors.New("Execution error.")
	}
	return true, nil
}

func displayError(msg string) {
	fg := color.New(color.FgWhite)
	redBg := fg.Add(color.BgRed)
	redBg.Println(msg)
}

func displaySuccess(msg string) {
	fg := color.New(color.FgHiGreen)
	fg.Println(msg)
}

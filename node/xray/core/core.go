package core

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Core struct {
}

func New() (*Core, error) {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configPath := fmt.Sprintf("%s/../bin/config.json", cwd)
	cmd := exec.Command("xray", "-c", configPath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd:", err)
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting Cmd:", err)
		return nil, err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from pipe:", err)
		return nil, err
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Println("Error waiting for Cmd:", err)
		}
	}()

	return &Core{}, nil
}

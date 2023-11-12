package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadEnvs() {
	file, err := os.Open("local.env")
	defer file.Close()

	if err != nil {
		fmt.Println("Reading from docker", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key, value := parts[0], parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Reading from docker", err)
		return
	}

	fmt.Println("Reading from local")
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

package jb

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// StdinAvailable here
func StdinAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// GetHomeDir here
func GetHomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// FileExists here
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadFile here
func ReadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	return file
}

// WriteFile here
func WriteFile(path string, file []byte) error {
	return ioutil.WriteFile(path, file, 0644)
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

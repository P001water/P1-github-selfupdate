package main

import (
	"bufio"
	"fmt"
	"github.com/P001water/P1-github-selfupdate/selfupdate"
	"github.com/blang/semver"
	"log"
	"os"
)

const version = "0.0.5"

func confirmAndSelfUpdate(version string, repo string) {
	latest, found, err := selfupdate.DetectLatest(repo)
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	v := semver.MustParse(version)
	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	fmt.Printf("Current version: %v\n", version)
	fmt.Printf("latest version: %v, Whether to update ? (yes or no)\n", latest.Version)
	// 创建一个 Scanner 用于读取用户输入
	fmt.Printf("Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	switch userInput {
	case "yes":
		fmt.Println("Updating...")
		exe, err := os.Executable()
		if err != nil {
			log.Println("Could not locate executable path")
			return
		}
		if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
			log.Println("Error occurred while updating binary:", err)
			return
		}
		log.Println("Successfully updated to version", latest.Version)
	case "no":
		fmt.Println("Update cancelled.")
	default:
		fmt.Println("Invalid input. Update cancelled.")
	}
}

func main() {
	confirmAndSelfUpdate(version, "P001water/P1finger")
}

package selfupdate

import (
	"bufio"
	"fmt"
	"github.com/blang/semver"
	"os"
)

// 用户输入验证是否更新, repo参数输入用户/仓库名, eg: P001water/P1finger
func ConfirmAndSelfUpdate(version string, repo string) {
	latest, found, err := DetectLatest(repo)
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	v := semver.MustParse(version)
	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return
	}

	fmt.Printf("Current version: %v", version)
	fmt.Printf("latest version: %v, Whether to update ? (yes or no)\n", latest.Version)
	//读取用户输入
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
		if err := UpdateTo(latest.AssetURL, exe); err != nil {
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

func NoticeUpdate(version string, latest *Release) {
	fmt.Printf("Current version: %v\n", version)
	fmt.Printf("latest version: %v, Whether to update ? (yes or no)\n", latest.Version)
	//读取用户输入
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
		if err := UpdateTo(latest.AssetURL, exe); err != nil {
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

// 检查是否是最新版本
func CheckVersionIsLatest(version string, repo string) (bool, *Release, error) {
	latest, found, err := DetectLatest(repo)
	if err != nil || !found {
		return false, latest, err
	}

	v := semver.MustParse(version)
	if latest.Version.LTE(v) {
		return true, latest, nil
	} else {
		return false, latest, nil
	}
}

package main

import (
	"fmt"
	"github.com/P001water/P1-github-selfupdate/selfupdate"
	"github.com/blang/semver"
	"log"
)

const version = "0.0.6"

func checkVersionUpdate(version string, repo string) (bool, error) {
	latest, found, err := selfupdate.DetectLatest(repo)
	if err != nil || !found {
		log.Println("Error occurred while detecting version:", err)
		return false, err
	}

	v := semver.MustParse(version)
	if latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	isLatest, err := checkVersionUpdate(version, "P001water/P1finger")
	if err != nil {
		fmt.Println(err)
		return
	}

	if isLatest {
		fmt.Println("tools is Latest")
	} else {
		fmt.Println("tools is Outdated")
	}

}

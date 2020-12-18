package main

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/google/go-github/v33/github"
	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
)

const (
	githubOrg  string = "serverhd"
	githubRepo string = "serverhd"
)

var (
	versionString string
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Infoln("-- Server Home Dashboard  👋")

	if err := versionCheck(); err != nil {
		os.Exit(1)
	}

}

func versionCheck() error {

	currentVersion, err := version.NewVersion(versionString)
	if err != nil {
		log.Errorf("❌ Could not determine current application version\n")
		return err
	}
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), githubOrg, githubRepo)
	if err != nil {
		log.Errorf("❌ Could not check for the latest release via GitHub\n")
		return err
	}
	latestVersion, err := version.NewVersion(*release.TagName)
	if err != nil {
		log.Errorf("❌ Could not determine latest application version\n")
		return err
	}

	if currentVersion.GreaterThan(latestVersion) {
		log.Debugf("🥳 Go you! Using a pre-release version\n!")
	} else if currentVersion.LessThan(latestVersion) {
		bold := color.New(color.Bold).SprintFunc()
		log.Warnf("⬆️  This app is not up to date\n\t\t%s\n", bold("Latest Version: ", latestVersion.String()))
	} else {
		log.Debugf("✅ ServerHD is up to date!\n")
	}

	return nil
}

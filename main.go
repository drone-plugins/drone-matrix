// Copyright (c) 2017 Paul Tötterman <ptman@iki.fi>. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matrix-org/gomatrix"
)

func main() {
	// Secrets
	password := os.Getenv("MATRIX_PASSWORD")
	accessToken := os.Getenv("MATRIX_ACCESSTOKEN")
	// Not sure if these are secrets or nice to have close to them
	userName := os.Getenv("MATRIX_USERNAME")
	userID := os.Getenv("MATRIX_USERID")

	// Override secrets if present
	if pw := os.Getenv("PLUGIN_PASSWORD"); pw != "" {
		password = pw
	}
	if at := os.Getenv("PLUGIN_ACCESSTOKEN"); at != "" {
		accessToken = at
	}
	if un := os.Getenv("PLUGIN_USERNAME"); un != "" {
		userName = un
	}
	if ui := os.Getenv("PLUGIN_USERID"); ui != "" {
		userID = ui
	}

	homeServer := os.Getenv("PLUGIN_HOMESERVER")
	if homeServer == "" {
		homeServer = "https://matrix.org"
	}

	// TODO: resolve room aliases
	roomID := os.Getenv("PLUGIN_ROOMID")
	message := os.Getenv("PLUGIN_MESSAGE")

	repoOwner := os.Getenv("DRONE_REPO_OWNER")
	repoName := os.Getenv("DRONE_REPO_NAME")

	buildStatus := os.Getenv("DRONE_BUILD_STATUS")
	buildLink := os.Getenv("DRONE_BUILD_LINK")
	buildBranch := os.Getenv("DRONE_BRANCH")
	buildAuthor := os.Getenv("DRONE_COMMIT_AUTHOR")
	buildCommit := os.Getenv("DRONE_COMMIT")

	m, err := gomatrix.NewClient(homeServer, userID, accessToken)
	if err != nil {
		log.Fatal(err)
	}

	if userID == "" || accessToken == "" {
		r, err := m.Login(&gomatrix.ReqLogin{
			Type:                     "m.login.password",
			User:                     userName,
			Password:                 password,
			InitialDeviceDisplayName: "Drone",
		})
		if err != nil {
			log.Fatal(err)
		}
		m.SetCredentials(r.UserID, r.AccessToken)
	}

	if message == "" {
		message = fmt.Sprintf("Build %s <%s> %s/%s#%s (%s) by %s",
			buildStatus,
			buildLink,
			repoOwner,
			repoName,
			buildCommit[:8],
			buildBranch,
			buildAuthor)
	}

	if _, err := m.SendNotice(roomID, message); err != nil {
		log.Fatal(err)
	}
}

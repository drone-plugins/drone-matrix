// Copyright (c) 2017 Paul Tötterman <ptman@iki.fi>. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/matrix-org/gomatrix"
)

func main() {
	homeServer := os.Getenv("PLUGIN_HOMESERVER")
	userName := os.Getenv("PLUGIN_USERNAME")
	password := os.Getenv("PLUGIN_PASSWORD")

	userID := os.Getenv("PLUGIN_USERID")
	accessToken := os.Getenv("PLUGIN_ACCESSTOKEN")

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

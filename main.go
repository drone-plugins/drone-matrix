package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "matrix plugin"
	app.Usage = "matrix plugin"
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username",
			Usage:  "username for authentication",
			EnvVar: "PLUGIN_USERNAME,MATRIX_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "password for authentication",
			EnvVar: "PLUGIN_PASSWORD,MATRIX_PASSWORD",
		},
		cli.StringFlag{
			Name:   "userid",
			Usage:  "userid for authentication",
			EnvVar: "PLUGIN_USERID,PLUGIN_USER_ID,MATRIX_USERID,MATRIX_USER_ID",
		},
		cli.StringFlag{
			Name:   "accesstoken",
			Usage:  "accesstoken for authentication",
			EnvVar: "PLUGIN_ACCESSTOKEN,PLUGIN_ACCESS_TOKEN,MATRIX_ACCESSTOKEN,MATRIX_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "homeserver",
			Usage:  "matrix home server",
			EnvVar: "PLUGIN_HOMESERVER,MATRIX_HOMESERVER",
			Value:  "https://matrix.org",
		},
		cli.StringFlag{
			Name:   "roomid",
			Usage:  "roomid to send messages",
			EnvVar: "PLUGIN_ROOMID,MATRIX_ROOMID",
		},
		cli.StringFlag{
			Name:   "template",
			Usage:  "template for the message",
			EnvVar: "PLUGIN_TEMPLATE,MATRIX_TEMPLATE",
			Value:  "Build {{ build.status }} <{{ build.link }}|{{ repo.Owner }}/{{ repo.Name }}#{{ truncate build.commit 8 }}> ({{ build.branch }}) by {{ build.author }}",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Commit:   c.String("commit.sha"),
			Ref:      c.String("commit.ref"),
			Branch:   c.String("commit.branch"),
			Author:   c.String("commit.author"),
			Message:  c.String("commit.message"),
			DeployTo: c.String("build.deployTo"),
			Link:     c.String("build.link"),
			Started:  c.Int64("build.started"),
			Created:  c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Username:    c.String("username"),
			Password:    c.String("password"),
			UserID:      c.String("userid"),
			AccessToken: c.String("accesstoken"),
			Homeserver:  c.String("homeserver"),
			RoomID:      c.String("roomid"),
			Template:    c.String("template"),
		},
	}

	return plugin.Exec()
}

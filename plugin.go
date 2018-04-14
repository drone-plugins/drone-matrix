package main

import (
	"strings"

	"github.com/drone/drone-template-lib/template"
	"github.com/matrix-org/gomatrix"
	"github.com/pkg/errors"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag      string
		Event    string
		Number   int
		Commit   string
		Ref      string
		Branch   string
		Author   string
		Message  string
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	Job struct {
		Started int64
	}

	Config struct {
		Username    string
		Password    string
		UserID      string
		AccessToken string
		Homeserver  string
		RoomID      string
		Template    string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Job    Job
		Config Config
	}
)

func (p Plugin) Exec() error {
	m, err := gomatrix.NewClient(p.Config.Homeserver, prepend("@", p.Config.UserID), p.Config.AccessToken)

	if err != nil {
		return errors.Wrap(err, "failed to initialize client")
	}

	if p.Config.UserID == "" || p.Config.AccessToken == "" {
		r, err := m.Login(&gomatrix.ReqLogin{
			Type:                     "m.login.password",
			User:                     p.Config.Username,
			Password:                 p.Config.Password,
			InitialDeviceDisplayName: "Drone",
		})

		if err != nil {
			return errors.Wrap(err, "failed to authenticate user")
		}

		m.SetCredentials(r.UserID, r.AccessToken)
	}

	joined, err := m.JoinRoom(p.Config.RoomID, "", nil)

	if err != nil {
		return errors.Wrap(err, "failed to join room")
	}

	message, err := template.RenderTrim(p.Config.Template, p)

	if err != nil {
		return errors.Wrap(err, "failed to render template")
	}

	if _, err := m.SendNotice(joined.RoomID, message); err != nil {
		return errors.Wrap(err, "failed to submit message")
	}

	return nil
}

func prepend(prefix, s string) string {
	if s == "" {
		return s
	}

	if strings.HasPrefix(s, prefix) {
		return s
	}

	return prefix + s
}

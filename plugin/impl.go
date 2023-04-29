// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"fmt"
	"strings"

	"github.com/drone/drone-template-lib/template"
	"github.com/matrix-org/gomatrix"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// Settings for the plugin.
type Settings struct {
	Username    string
	Password    string
	UserID      string
	AccessToken string
	Homeserver  string
	RoomID      string
	Template    string
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Currently there's no validation
	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	m, err := gomatrix.NewClient(p.settings.Homeserver, prepend("@", p.settings.UserID), p.settings.AccessToken)

	if err != nil {
		return fmt.Errorf("failed to initialize client: %w", err)
	}

	if p.settings.UserID == "" || p.settings.AccessToken == "" {
		r, err := m.Login(&gomatrix.ReqLogin{
			Type:                     "m.login.password",
			User:                     p.settings.Username,
			Password:                 p.settings.Password,
			InitialDeviceDisplayName: "Drone",
		})

		if err != nil {
			return fmt.Errorf("failed to authenticate user: %w", err)
		}

		m.SetCredentials(r.UserID, r.AccessToken)
	}

	joined, err := m.JoinRoom(prepend("!", p.settings.RoomID), "", nil)

	if err != nil {
		return fmt.Errorf("failed to join room: %w", err)
	}

	message, err := template.RenderTrim(p.settings.Template, p.pipeline)

	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	formatted := bluemonday.UGCPolicy().SanitizeBytes(
		blackfriday.Run([]byte(message)),
	)

	content := gomatrix.HTMLMessage{
		Body:          message,
		MsgType:       "m.notice",
		Format:        "org.matrix.custom.html",
		FormattedBody: string(formatted),
	}

	if _, err := m.SendMessageEvent(joined.RoomID, "m.room.message", content); err != nil {
		return fmt.Errorf("failed to submit message: %w", err)
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

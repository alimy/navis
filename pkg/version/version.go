// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package version

import (
	"fmt"
)

var version, commitID, buildDate string

type BuildInfo struct {
	Version   string `json:"version"`
	Sum       string `json:"sum"`
	BuildDate string `json:"build_date"`
}

func VersionInfo() string {
	return fmt.Sprintf("navis %s (build:%s %s)", version, commitID, buildDate)
}

func ReadBuildInfo() *BuildInfo {
	return &BuildInfo{
		Version:   version,
		Sum:       commitID,
		BuildDate: buildDate,
	}
}

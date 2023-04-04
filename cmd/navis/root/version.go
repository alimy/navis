// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package root

import (
	"fmt"

	"github.com/alimy/navis/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "print version information",
		Long:  "print version information",
		Run:   runVersion,
	}
	rootCmd.AddCommand(versionCmd)
}

func runVersion(_cmd *cobra.Command, _args []string) {
	fmt.Println(version.VersionInfo())
}

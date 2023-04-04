// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package root

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "navis",
		Short: "navis cli cmd",
		Long:  `navis cli cmd`,
	}

	NoDefaultFeaturesArg bool
	FeaturesArg          []string
)

// AddCommand add sub-command
func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute start application
func Execute() {
	rootCmd.Execute()
}

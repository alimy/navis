// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package serv

import (
	"fmt"

	"github.com/alimy/navis/cmd/navis/conf"
	"github.com/alimy/navis/cmd/navis/root"
	"github.com/spf13/cobra"
)

var (
	cfgFileArg string
)

func init() {
	servCmd := &cobra.Command{
		Use:   "serv",
		Short: "start a object storage service",
		Long:  "start a object storage service",
		Run:   runServ,
	}
	servCmd.Flags().StringVarP(&cfgFileArg, "conf", "c", "app.yaml", "config file")
	servCmd.Flags().BoolVar(&root.NoDefaultFeaturesArg, "no-default-feature", false, "whether not use default features")
	servCmd.Flags().StringSliceVar(&root.FeaturesArg, "features", []string{}, "use special features")
	root.AddCommand(servCmd)
}

func runServ(cmd *cobra.Command, args []string) {
	conf.Initial(root.FeaturesArg, root.NoDefaultFeaturesArg)
	// TODO
	fmt.Printf("appConf:%+v\nserverConf:%+v\n", conf.AppConf, conf.ServerConf)
}

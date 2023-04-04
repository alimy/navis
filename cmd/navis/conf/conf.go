// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	_ "embed"
	"log"

	"github.com/alimy/cfg"
)

var (
	AppConf    *appConf
	ServerConf *serverConf
)

type appConf struct {
	RunMode string
}

type serverConf struct {
	Addr string
}

func setupConf(suite []string, noDefault bool) error {
	setting, err := newSetting()
	if err != nil {
		return err
	}

	// initialize features configure
	ss, kv := setting.featuresInfoFrom("Features")
	cfg.Initial(ss, kv)
	if len(suite) > 0 {
		cfg.Use(suite, noDefault)
	}

	objects := map[string]any{
		"App":    &AppConf,
		"Server": &ServerConf,
	}
	if err = setting.Unmarshal(objects); err != nil {
		return err
	}
	return nil
}

func Initial(suite []string, noDefault bool) {
	err := setupConf(suite, noDefault)
	if err != nil {
		log.Fatalf("configure setup err: %v", err)
	}
}

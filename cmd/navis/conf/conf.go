// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"github.com/alimy/cfg"
	"github.com/sirupsen/logrus"
)

var (
	AppConf    *appConf
	ServerConf *serverConf
)

func setupConf(suite []string, noDefault bool) error {
	vp, err := newViper()
	if err != nil {
		return err
	}

	// initialize features configure
	ss, kv := featuresFrom(vp, "Features")
	cfg.Initial(ss, kv)
	if len(suite) > 0 {
		cfg.Use(suite, noDefault)
	}

	objects := map[string]any{
		"App":    &AppConf,
		"Server": &ServerConf,
	}
	for k, v := range objects {
		err := vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Initial(suite []string, noDefault bool) {
	err := setupConf(suite, noDefault)
	if err != nil {
		logrus.Fatalf("configure setup err: %v", err)
	}
}

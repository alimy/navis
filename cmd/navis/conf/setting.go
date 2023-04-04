// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"bytes"
	_ "embed"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

//go:embed app.yaml
var fileBytes []byte

type appConf struct {
	RunMode string
}

type serverConf struct {
	Addr string
}

func featuresFrom(vp *viper.Viper, k string) (map[string][]string, map[string]string) {
	sub := vp.Sub(k)
	keys := sub.AllKeys()

	suites := make(map[string][]string)
	kv := make(map[string]string, len(keys))
	for _, key := range sub.AllKeys() {
		val := sub.Get(key)
		switch v := val.(type) {
		case string:
			kv[key] = v
		case []any:
			suites[key] = sub.GetStringSlice(key)
		}
	}
	return suites, kv
}

func newViper() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("app")
	vp.AddConfigPath(".")
	vp.AddConfigPath("custom/")
	vp.SetConfigType("yaml")
	err := vp.ReadConfig(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, errors.Wrap(err, "create viper")
	}
	if err = vp.MergeInConfig(); err != nil {
		return nil, err
	}
	return vp, nil
}

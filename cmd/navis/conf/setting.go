// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"bytes"
	_ "embed"

	"github.com/ory/viper"
)

//go:embed app.yaml
var fileBytes []byte

type setting struct {
	vp *viper.Viper
}

func (s *setting) ReadSection(k string, v any) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *setting) Unmarshal(objects map[string]any) error {
	for k, v := range objects {
		err := s.vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *setting) featuresInfoFrom(k string) (map[string][]string, map[string]string) {
	sub := s.vp.Sub(k)
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

func newSetting() (*setting, error) {
	vp := viper.New()
	vp.SetConfigName("app")
	vp.AddConfigPath(".")
	vp.AddConfigPath("custom/")
	vp.SetConfigType("yaml")
	err := vp.ReadConfig(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, err
	}
	if err = vp.MergeInConfig(); err != nil {
		return nil, err
	}
	return &setting{vp}, nil
}

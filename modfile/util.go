/*
 * Copyright 2023 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package modfile

import (
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"gopkg.in/yaml.v3"
)

func New(decoders Decoders, generators Generators) *ModFile {
	return &ModFile{decoders: decoders, generators: generators}
}

func (mf *ModFile) UnmarshalYAML(yn *yaml.Node) error {
	if len(mf.decoders) < 1 {
		return errors.New("no decoders")
	}
	var mfb Base
	if err := yn.Decode(&mfb); err != nil {
		return err
	}
	if mfb.Version == "" {
		return errors.New("no version")
	}
	d, ok := mf.decoders[mfb.Version]
	if !ok {
		return fmt.Errorf("no decoder for version '%s'", mfb.Version)
	}
	m, err := d(yn)
	if err != nil {
		return err
	}
	mf.version = mfb.Version
	mf.modFile = m
	return nil
}

func (mf *ModFile) GetModule() (*module.Module, error) {
	if len(mf.generators) < 1 {
		return nil, errors.New("no generators")
	}
	g, ok := mf.generators[mf.version]
	if !ok {
		return nil, fmt.Errorf("no generator for version '%s'", mf.version)
	}
	return g(mf.modFile)
}

func (d Decoders) Add(f func() (string, func(*yaml.Node) (any, error))) {
	key, df := f()
	d[key] = df
}

func (g Generators) Add(f func() (string, func(any) (*module.Module, error))) {
	key, gf := f()
	g[key] = gf
}

/*
 * Copyright 2026 InfAI (CC SES)
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

package modfile_lib

import (
	"errors"
	"io"

	v1_generator "github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator"
	v1_model "github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	module_lib "github.com/SENERGY-Platform/mgw-module-lib/model"
	"gopkg.in/yaml.v3"
)

func Unmarshal(b []byte) (module_lib.Module, error) {
	var nw nodeWrapper
	err := yaml.Unmarshal(b, &nw)
	if err != nil {
		return module_lib.Module{}, err
	}
	return getModule(nw.Version, nw.Node)
}

func Decode(r io.Reader) (module_lib.Module, error) {
	var nw nodeWrapper
	err := yaml.NewDecoder(r).Decode(&nw)
	if err != nil {
		return module_lib.Module{}, err
	}
	return getModule(nw.Version, nw.Node)
}

func getModule(version string, yn *yaml.Node) (module_lib.Module, error) {
	switch version {
	case v1_model.Version:
		return v1_generator.GetModule(yn)
	default:
		return module_lib.Module{}, errors.New("unknown modfile version: " + version)
	}
}

type modfileBase struct {
	Version string `yaml:"modfileVersion"`
}

type nodeWrapper struct {
	Node    *yaml.Node
	Version string
}

func (w *nodeWrapper) UnmarshalYAML(yn *yaml.Node) error {
	var mb modfileBase
	if err := yn.Decode(&mb); err != nil {
		return err
	}
	w.Version = mb.Version
	w.Node = yn
	return nil
}

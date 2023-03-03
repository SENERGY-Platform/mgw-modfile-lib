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
	"github.com/SENERGY-Platform/mgw-modfile-lib/itf"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"gopkg.in/yaml.v3"
)

func NewModFile(decoders Decoders) *ModFile {
	return &ModFile{decoders: decoders}
}

func (mf *ModFile) UnmarshalYAML(yn *yaml.Node) error {
	if len(mf.decoders) == 0 {
		return errors.New("no decoders")
	}
	var b base
	if err := yn.Decode(&b); err != nil {
		return err
	}
	if b.Version == "" {
		return errors.New("no version")
	}
	d, ok := mf.decoders[b.Version]
	if !ok {
		return fmt.Errorf("no decoder for version '%s'", b.Version)
	}
	modFile, err := d(yn)
	if err != nil {
		return err
	}
	mf.base = b
	mf.modFile = modFile
	return nil
}

func (mf *ModFile) GetModule() (*module.Module, error) {
	return mf.modFile.GenModule()
}

func (d Decoders) Add(gf func() (string, func(*yaml.Node) (itf.ModFile, error))) {
	key, decoder := gf()
	d[key] = decoder
}

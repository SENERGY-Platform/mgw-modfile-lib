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

package v1

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/itf"
	"gopkg.in/yaml.v3"
)

var ver = "v1"

func decode(yn *yaml.Node) (itf.ModFile, error) {
	var mf ModFile
	if err := yn.Decode(&mf); err != nil {
		return mf, err
	}
	return &mf, nil
}

func GetDecoder() (string, func(*yaml.Node) (itf.ModFile, error)) {
	return ver, decode
}
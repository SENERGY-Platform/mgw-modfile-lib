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

package configs

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	module_lib "github.com/SENERGY-Platform/mgw-module-lib/model"
)

func GenConfigs(mfCVs map[string]model.ConfigValue) (module_lib.Configs, error) {
	mCs := make(module_lib.Configs)
	for ref, mfCV := range mfCVs {
		if mfCV.IsList {
			if err := SetSlice(ref, mfCV, mCs); err != nil {
				return nil, err
			}
		} else {
			if err := SetValue(ref, mfCV, mCs); err != nil {
				return nil, err
			}
		}
	}
	return mCs, nil
}

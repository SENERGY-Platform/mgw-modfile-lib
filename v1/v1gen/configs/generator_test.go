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
	"testing"
)

func TestGenConfigs(t *testing.T) {
	strType := module_lib.StringType
	var mfCVs map[string]model.ConfigValue
	if mCs, err := GenConfigs(mfCVs); err != nil {
		t.Error("err != nil")
	} else if len(mCs) != 0 {
		t.Errorf("len(%v) != 0", mCs)
	}
	mfCVs = make(map[string]model.ConfigValue)
	if mCs, err := GenConfigs(mfCVs); err != nil {
		t.Error("err != nil")
	} else if len(mCs) != 0 {
		t.Errorf("len(%v) != 0", mCs)
	}
	str := "test"
	mfCVs[str] = model.ConfigValue{}
	if _, err := GenConfigs(mfCVs); err != nil {
		t.Error("err != nil")
	}
	mfCVs[str] = model.ConfigValue{IsList: true}
	if _, err := GenConfigs(mfCVs); err != nil {
		t.Error("err != nil")
	}
	mfCVs[str] = model.ConfigValue{DataType: &strType}
	if mCs, err := GenConfigs(mfCVs); err != nil {
		t.Error("err != nil")
	} else if mC, ok := mCs[str]; !ok {
		t.Errorf("mC, ok := mCs[%s]; !ok", str)
	} else if mC.IsSlice == true {
		t.Error("mC.IsSlice == true")
	}
	mfCVs[str] = model.ConfigValue{DataType: &strType, IsList: true}
	if mCs, err := GenConfigs(mfCVs); err != nil {
		t.Error("err != nil")
	} else if mC, ok := mCs[str]; !ok {
		t.Errorf("mC, ok := mCs[%s]; !ok", str)
	} else if mC.IsSlice == false {
		t.Error("mC.IsSlice == false")
	}
}

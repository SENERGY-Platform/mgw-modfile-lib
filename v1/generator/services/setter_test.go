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

package services

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"reflect"
	"testing"
)

func TestSetSrvReferences(t *testing.T) {
	ref := "a"
	var mfSRs map[string][]model.DependencyTarget
	mSs := map[string]*module.Service{ref: {}}
	if err := SetSrvReferences(mfSRs, mSs); err != nil {
		t.Error("err != nil")
	}
	// --------------------------------
	mfSRs = make(map[string][]model.DependencyTarget)
	dRef := "b"
	rVar := "var"
	mfSRs[dRef] = []model.DependencyTarget{
		{
			RefVar:   rVar,
			Services: []string{ref},
		},
	}
	a := map[string]string{rVar: dRef}
	if err := SetSrvReferences(mfSRs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[ref]; reflect.DeepEqual(a, ms.SrvReferences) == false {
		t.Errorf("%v != %v", a, ms.SrvReferences)
	}
	// --------------------------------
	mfSRs = make(map[string][]model.DependencyTarget)
	mfSRs[dRef] = []model.DependencyTarget{
		{
			RefVar:   rVar,
			Services: []string{ref},
		},
		{
			RefVar:   rVar,
			Services: []string{ref},
		},
	}
	if err := SetSrvReferences(mfSRs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[ref]; reflect.DeepEqual(a, ms.SrvReferences) == false {
		t.Errorf("%v != %v", a, ms.SrvReferences)
	}
	// --------------------------------
	mfSRs = make(map[string][]model.DependencyTarget)
	mfSRs[dRef] = []model.DependencyTarget{
		{
			RefVar:   rVar,
			Services: []string{ref},
		},
	}
	mfSRs["c"] = []model.DependencyTarget{
		{
			RefVar:   rVar,
			Services: []string{ref},
		},
	}
	if err := SetSrvReferences(mfSRs, mSs); err == nil {
		t.Error("err != nil")
	}
	// --------------------------------
	mfSRs = make(map[string][]model.DependencyTarget)
	mfSRs[dRef] = []model.DependencyTarget{
		{
			RefVar:   rVar,
			Services: []string{"c"},
		},
	}
	if err := SetSrvReferences(mfSRs, mSs); err == nil {
		t.Error("err != nil")
	}
}
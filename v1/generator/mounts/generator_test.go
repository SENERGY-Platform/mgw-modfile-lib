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

package mounts

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"testing"
)

func TestGenVolumes(t *testing.T) {
	var mfVs map[string][]model.VolumeTarget
	if vtm := GenVolumes(mfVs); len(vtm) != 0 {
		t.Errorf("vtm := GenVolumes(%v); len(%v) != 0", mfVs, vtm)
	}
	mfVs = make(map[string][]model.VolumeTarget)
	str := "test"
	mfVs[str] = []model.VolumeTarget{}
	if vtm := GenVolumes(mfVs); len(vtm) != 1 {
		t.Errorf("vtm := GenVolumes(%v); len(%v) != 1", mfVs, vtm)
	} else if _, ok := vtm[str]; !ok {
		t.Errorf("_, ok := vtm[%s]; !ok", str)
	}
}

func TestGenDependencies(t *testing.T) {
	var mfMDs map[string]model.ModuleDependency
	if m := GenDependencies(mfMDs); len(m) != 0 {
		t.Errorf("m := GenDependencies(%v); len(%v) != 0", mfMDs, m)
	}
	mfMDs = make(map[string]model.ModuleDependency)
	str := "test"
	mfMDs[str] = model.ModuleDependency{}
	if m := GenDependencies(mfMDs); len(m) != 1 {
		t.Errorf("m := GenDependencies(%v); len(%v) != 1", mfMDs, m)
	} else if _, ok := m[str]; !ok {
		t.Errorf("_, ok := m[%s]; !ok", str)
	}
}

func TestGenResources(t *testing.T) {
	var mfRs map[string]model.Resource
	if m := GenResources(mfRs); len(m) != 0 {
		t.Errorf("m := GenResources(%v); len(%v) != 0", mfRs, m)
	}
	mfRs = make(map[string]model.Resource)
	str := "test"
	mfRs[str] = model.Resource{ResourceBase: model.ResourceBase{Tags: []string{str}}}
	if m := GenResources(mfRs); len(m) != 1 {
		t.Errorf("m := GenResources(%v); len(%v) != 1", mfRs, m)
	} else if set, ok := m[str]; !ok {
		t.Errorf("set, ok := m[%s]; !ok", str)
	} else if len(set) != 1 {
		t.Errorf("len(%v) != 1", set)
	} else if _, ok := set[str]; !ok {
		t.Errorf("_, ok := set[%s]; !ok", str)
	}
}

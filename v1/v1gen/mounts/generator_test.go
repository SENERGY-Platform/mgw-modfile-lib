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
		t.Errorf("len(%v) != 0", vtm)
	}
	mfVs = make(map[string][]model.VolumeTarget)
	str := "test"
	mfVs[str] = []model.VolumeTarget{}
	if vtm := GenVolumes(mfVs); len(vtm) != 1 {
		t.Errorf("len(%v) != 1", vtm)
	} else if _, ok := vtm[str]; !ok {
		t.Errorf("_, ok := vtm[%s]; !ok", str)
	}
}

func TestGenDependencies(t *testing.T) {
	var mfMDs map[string]model.ModuleDependency
	if m := GenDependencies(mfMDs); len(m) != 0 {
		t.Errorf("len(%v) != 0", m)
	}
	mfMDs = make(map[string]model.ModuleDependency)
	str := "test"
	mfMDs[str] = model.ModuleDependency{}
	if m := GenDependencies(mfMDs); len(m) != 1 {
		t.Errorf("len(%v) != 1", m)
	} else if _, ok := m[str]; !ok {
		t.Errorf("_, ok := m[%s]; !ok", str)
	}
}

func TestGenResources(t *testing.T) {
	var mfRs map[string]model.HostResource
	if m := GenHostResources(mfRs); len(m) != 0 {
		t.Errorf("len(%v) != 0", m)
	}
	mfRs = make(map[string]model.HostResource)
	str := "test"
	mfRs[str] = model.HostResource{
		Resource: model.Resource{
			Tags:     []string{str},
			Optional: false,
		},
	}
	if m := GenHostResources(mfRs); len(m) != 1 {
		t.Errorf("len(%v) != 1", m)
	} else if mR, ok := m[str]; !ok {
		t.Errorf("mR, ok := m[%s]; !ok", str)
	} else if len(mR.Tags) != 1 {
		t.Errorf("len(%v) != 1", mR.Tags)
	} else if _, ok := mR.Tags[str]; !ok {
		t.Errorf("_, ok := mR.Tags[%s]; !ok", str)
	} else if mR.Required == false {
		t.Error("mR.Required == false")
	}
}

func TestGenSecrets(t *testing.T) {
	var mfSs map[string]model.Secret
	if sm := GenSecrets(mfSs); len(sm) != 0 {
		t.Errorf("len(%v) != 0", sm)
	}
	mfSs = make(map[string]model.Secret)
	str := "test"
	mfSs[str] = model.Secret{
		Resource: model.Resource{
			Tags:     []string{str},
			Optional: false,
		},
		Type: str,
	}
	if sm := GenSecrets(mfSs); len(sm) != 1 {
		t.Errorf("len(%v) != 1", sm)
	} else if s, ok := sm[str]; !ok {
		t.Errorf("s, ok := sm[%s]; !ok", str)
	} else if len(s.Tags) != 1 {
		t.Errorf("len(%v) != 1", s.Tags)
	} else if _, ok := s.Tags[str]; !ok {
		t.Errorf("_, ok := s.Tags[%s]; !ok", str)
	} else if s.Type != str {
		t.Errorf("%s != %s", s.Type, str)
	} else if s.Required == false {
		t.Error("s.Required == false")
	}
}

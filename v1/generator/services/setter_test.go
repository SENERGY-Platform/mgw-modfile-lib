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

func TestSetVolumes(t *testing.T) {
	sRef := "a"
	var mfVs map[string][]model.VolumeTarget
	mSs := map[string]*module.Service{sRef: {}}
	if err := SetVolumes(mfVs, mSs); err != nil {
		t.Error("err != nil")
	}
	// --------------------------------
	mfVs = make(map[string][]model.VolumeTarget)
	vl := "vl"
	mp := "mp"
	mfVs[vl] = []model.VolumeTarget{
		{
			MountPoint: mp,
			Services:   []string{sRef},
		},
	}
	a := map[string]string{mp: vl}
	if err := SetVolumes(mfVs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[sRef]; reflect.DeepEqual(a, ms.Volumes) == false {
		t.Errorf("%v != %v", a, ms.Volumes)
	}
	// --------------------------------
	mfVs[vl] = []model.VolumeTarget{
		{
			MountPoint: mp,
			Services:   []string{sRef},
		},
		{
			MountPoint: mp,
			Services:   []string{sRef},
		},
	}
	if err := SetVolumes(mfVs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[sRef]; reflect.DeepEqual(a, ms.Volumes) == false {
		t.Errorf("%v != %v", a, ms.Volumes)
	}
	// --------------------------------
	mfVs[vl] = []model.VolumeTarget{
		{
			MountPoint: mp,
			Services:   []string{sRef},
		},
	}
	mfVs["vl2"] = []model.VolumeTarget{
		{
			MountPoint: mp,
			Services:   []string{sRef},
		},
	}
	if err := SetVolumes(mfVs, mSs); err == nil {
		t.Error("err != nil")
	}
	// --------------------------------
	mfVs[vl] = []model.VolumeTarget{
		{
			MountPoint: mp,
			Services:   []string{"b"},
		},
	}
	if err := SetVolumes(mfVs, mSs); err == nil {
		t.Error("err != nil")
	}
}

func TestSetExtDependencies(t *testing.T) {
	ref := "a"
	var mfMDs map[string]model.ModuleDependency
	mSs := map[string]*module.Service{ref: {}}
	if err := SetExtDependencies(mfMDs, mSs); err != nil {
		t.Error("err != nil")
	}
	// --------------------------------
	mfMDs = make(map[string]model.ModuleDependency)
	mID := "mid"
	mVer := "ver"
	dRef := "b"
	rVar := "var"
	mfMDs[mID] = model.ModuleDependency{
		Version: mVer,
		RequiredServices: map[string][]model.DependencyTarget{
			dRef: {
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
			},
		},
	}
	a := map[string]module.ExtDependencyTarget{
		rVar: {
			ID:      mID,
			Service: dRef,
		},
	}
	if err := SetExtDependencies(mfMDs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[ref]; reflect.DeepEqual(a, ms.ExtDependencies) == false {
		t.Errorf("%v != %v", a, ms.ExtDependencies)
	}
	// --------------------------------
	mfMDs[mID] = model.ModuleDependency{
		Version: mVer,
		RequiredServices: map[string][]model.DependencyTarget{
			dRef: {
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
			},
		},
	}
	if err := SetExtDependencies(mfMDs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[ref]; reflect.DeepEqual(a, ms.ExtDependencies) == false {
		t.Errorf("%v != %v", a, ms.ExtDependencies)
	}
	// --------------------------------
	mfMDs[mID] = model.ModuleDependency{
		Version: mVer,
		RequiredServices: map[string][]model.DependencyTarget{
			dRef: {
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
			},
			"c": {
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
			},
		},
	}
	if err := SetExtDependencies(mfMDs, mSs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mfMDs[mID] = model.ModuleDependency{
		Version: mVer,
		RequiredServices: map[string][]model.DependencyTarget{
			dRef: {
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
			},
		},
	}
	mfMDs["test"] = model.ModuleDependency{
		Version: mVer,
		RequiredServices: map[string][]model.DependencyTarget{
			dRef: {
				{
					RefVar:   rVar,
					Services: []string{ref},
				},
			},
		},
	}
	if err := SetExtDependencies(mfMDs, mSs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mfMDs[mID] = model.ModuleDependency{
		Version: mVer,
		RequiredServices: map[string][]model.DependencyTarget{
			dRef: {
				{
					RefVar:   rVar,
					Services: []string{"c"},
				},
			},
		},
	}
	if err := SetExtDependencies(mfMDs, mSs); err == nil {
		t.Error("err == nil")
	}
}

func TestSetResources(t *testing.T) {
	ref := "a"
	var mfRs map[string]model.Resource
	mSs := map[string]*module.Service{ref: {}}
	if err := SetResources(mfRs, mSs); err != nil {
		t.Error("err != nil")
	}
	// --------------------------------
	mfRs = make(map[string]model.Resource)
	res := "res"
	mp := "mp"
	mfRs[res] = model.Resource{
		Targets: []model.ResourceTarget{
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: true,
			},
		},
	}
	a := map[string]module.ResourceTarget{
		mp: {
			Ref:      res,
			ReadOnly: true,
		},
	}
	if err := SetResources(mfRs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[ref]; reflect.DeepEqual(a, ms.Resources) == false {
		t.Errorf("%v != %v", a, ms.Resources)
	}
	// --------------------------------
	mfRs[res] = model.Resource{
		Targets: []model.ResourceTarget{
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: true,
			},
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: true,
			},
		},
	}
	if err := SetResources(mfRs, mSs); err != nil {
		t.Error("err != nil")
	} else if ms := mSs[ref]; reflect.DeepEqual(a, ms.Resources) == false {
		t.Errorf("%v != %v", a, ms.Resources)
	}
	// --------------------------------
	mfRs[res] = model.Resource{
		Targets: []model.ResourceTarget{
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: false,
			},
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: true,
			},
		},
	}
	if err := SetResources(mfRs, mSs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mfRs[res] = model.Resource{
		Targets: []model.ResourceTarget{
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: false,
			},
		},
	}
	mfRs["test"] = model.Resource{
		Targets: []model.ResourceTarget{
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{ref},
				},
				ReadOnly: false,
			},
		},
	}
	if err := SetResources(mfRs, mSs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mfRs[res] = model.Resource{
		Targets: []model.ResourceTarget{
			{
				ResourceTargetBase: model.ResourceTargetBase{
					MountPoint: mp,
					Services:   []string{"b"},
				},
				ReadOnly: true,
			},
		},
	}
	if err := SetResources(mfRs, mSs); err == nil {
		t.Error("err == nil")
	}
}

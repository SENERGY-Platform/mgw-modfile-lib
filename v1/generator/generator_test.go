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

package generator

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGenerator(t *testing.T) {
	if _, err := generator(model.ModFile{}); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	sA := "a"
	sB := "b"
	ig := "ig"
	mf := model.ModFile{
		ID:             "id",
		Name:           "nme",
		Description:    "dsc",
		Tags:           []string{"tag"},
		License:        "lcs",
		Author:         "ath",
		Version:        "ver",
		Type:           "typ",
		DeploymentType: "dtp",
		Architectures:  []string{"arch"},
		Services: map[string]model.Service{
			sA: {},
			sB: {},
		},
		ServiceReferences: map[string][]model.DependencyTarget{
			sB: {
				{
					RefVar:   "rVar1",
					Services: []string{sA},
				},
			},
		},
		Volumes: map[string][]model.VolumeTarget{
			"vol": {
				{
					MountPoint: "mnt1",
					Services:   []string{sA},
				},
			},
		},
		Dependencies: map[string]model.ModuleDependency{
			"ext": {
				Version: "ver",
				RequiredServices: map[string][]model.DependencyTarget{
					"c": {
						{
							RefVar:   "rVar2",
							Services: []string{sA},
						},
					},
				},
			},
		},
		Resources: map[string]model.Resource{
			"res": {
				ResourceBase: model.ResourceBase{
					Tags: []string{"tag"},
					UserInput: &model.UserInput{
						Group: &ig,
					},
				},
				Targets: []model.ResourceTarget{
					{
						ResourceTargetBase: model.ResourceTargetBase{
							MountPoint: "mnt2",
							Services:   []string{sA},
						},
					},
				},
			},
		},
		Secrets: map[string]model.Secret{
			"sec": {
				ResourceBase: model.ResourceBase{
					Tags: []string{"tag"},
					UserInput: &model.UserInput{
						Group: &ig,
					},
				},
				Targets: []model.ResourceTargetBase{
					{
						MountPoint: "mnt3",
						Services:   []string{sA},
					},
				},
			},
		},
		Configs: map[string]model.ConfigValue{
			"cfg": {
				DataType: module.StringType,
				UserInput: &model.UserInput{
					Group: &ig,
				},
				Targets: []model.ConfigTarget{
					{
						RefVar:   "rVar3",
						Services: []string{sA},
					},
				},
			},
		},
		InputGroups: map[string]model.InputGroup{
			ig: {},
		},
	}
	mc := make(module.Configs)
	mc.SetString("cfg", nil, nil, false, "", nil)
	a := module.Module{
		ID:             "id",
		Name:           "nme",
		Description:    "dsc",
		Tags:           map[string]struct{}{"tag": {}},
		License:        "lcs",
		Author:         "ath",
		Version:        "ver",
		Type:           "typ",
		DeploymentType: "dtp",
		Architectures:  map[module.CPUArch]struct{}{"arch": {}},
		Services: map[string]*module.Service{
			sA: {
				BindMounts:    map[string]module.BindMount{},
				Tmpfs:         map[string]module.TmpfsMount{},
				HttpEndpoints: map[string]module.HttpEndpoint{},
				RequiredSrv:   map[string]struct{}{},
			},
			sB: {
				BindMounts:    map[string]module.BindMount{},
				Tmpfs:         map[string]module.TmpfsMount{},
				HttpEndpoints: map[string]module.HttpEndpoint{},
				RequiredSrv:   map[string]struct{}{},
			},
		},
		Volumes: map[string]struct{}{
			"vol": {},
		},
		Dependencies: map[string]string{
			"ext": "ver",
		},
		Resources: map[string]map[string]struct{}{
			"res": {
				"tag": {},
			},
		},
		Secrets: map[string]module.Secret{
			"sec": {
				Tags: map[string]struct{}{
					"tag": {},
				},
			},
		},
		Configs: mc,
		Inputs: module.Inputs{
			Resources: map[string]module.Input{
				"res": {
					Group: &ig,
				},
			},
			Secrets: map[string]module.Input{
				"sec": {
					Group: &ig,
				},
			},
			Configs: map[string]module.Input{
				"cfg": {
					Group: &ig,
				},
			},
			Groups: map[string]module.InputGroup{
				ig: {},
			},
		},
	}
	if b, err := generator(&mf); err != nil {
		t.Error("err != nil")
	} else if cmp.Equal(a, *b) == false {
		t.Errorf("%+v != %+v", a, *b)
	}
}

func TestGetGenerator(t *testing.T) {
	if s, _ := GetGenerator(); s != model.Version {
		t.Errorf("%s != %s", s, model.Version)
	}
}

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

package v1gen

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"reflect"
	"testing"
	"time"
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
				RunConfig: module.RunConfig{
					MaxRetries:  3,
					RunOnce:     false,
					StopTimeout: 5 * time.Second,
					StopSignal:  nil,
					PseudoTTY:   false,
				},
				BindMounts: map[string]module.BindMount{},
				Tmpfs:      map[string]module.TmpfsMount{},
				Volumes: map[string]string{
					"mnt1": "vol",
				},
				Resources: map[string]module.ResourceTarget{
					"mnt2": {
						Ref:      "res",
						ReadOnly: false,
					},
				},
				Secrets:       map[string]string{"mnt3": "sec"},
				Configs:       map[string]string{"rVar3": "cfg"},
				SrvReferences: map[string]string{"rVar1": sB},
				HttpEndpoints: map[string]module.HttpEndpoint{},
				RequiredSrv:   map[string]struct{}{},
				RequiredBySrv: nil,
				ExtDependencies: map[string]module.ExtDependencyTarget{
					"rVar2": {
						ID:      "ext",
						Service: "c",
					},
				},
				Ports: nil,
			},
			sB: {
				RunConfig: module.RunConfig{
					MaxRetries:  3,
					RunOnce:     false,
					StopTimeout: 5 * time.Second,
					StopSignal:  nil,
					PseudoTTY:   false,
				},
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
	} else if reflect.DeepEqual(a, *b) == false {
		t.Errorf("%+v != %+v", a, *b)
	}
	//--------------------------------
	mf = model.ModFile{
		Configs: map[string]model.ConfigValue{
			"cfg": {},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		Services: map[string]model.Service{
			"": {
				Ports: []model.SrvPort{
					{
						Name:     nil,
						Port:     "",
						HostPort: nil,
						Protocol: nil,
					},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		ServiceReferences: map[string][]model.DependencyTarget{
			"": {
				{
					Services: []string{""},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		Volumes: map[string][]model.VolumeTarget{
			"": {
				{
					Services: []string{""},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		Dependencies: map[string]model.ModuleDependency{
			"": {
				RequiredServices: map[string][]model.DependencyTarget{
					"": {
						{
							Services: []string{""},
						},
					},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		Resources: map[string]model.Resource{
			"": {
				Targets: []model.ResourceTarget{
					{
						ResourceTargetBase: model.ResourceTargetBase{
							Services: []string{""},
						},
					},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		Secrets: map[string]model.Secret{
			"": {
				Targets: []model.ResourceTargetBase{
					{
						Services: []string{""},
					},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mf = model.ModFile{
		Configs: map[string]model.ConfigValue{
			"": {
				DataType: module.StringType,
				Targets: []model.ConfigTarget{
					{
						Services: []string{""},
					},
				},
			},
		},
	}
	if _, err := generator(&mf); err == nil {
		t.Error("err == nil")
	}
}

func TestGetGenerator(t *testing.T) {
	if s, _ := GetGenerator(); s != model.Version {
		t.Errorf("%s != %s", s, model.Version)
	}
}

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
	module_lib "github.com/SENERGY-Platform/mgw-module-lib/model"
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
	aA := "a"
	aB := "b"
	ig := "ig"
	sMnt := "mnt3"
	strType := module_lib.StringType
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
		AuxServices: map[string]model.AuxService{
			aA: {},
			aB: {},
		},
		AuxImageSources: []string{"test"},
		ServiceReferences: map[string][]model.DependencyTarget{
			sB: {
				{
					RefVar:      "rVar1",
					Services:    []string{sA},
					AuxServices: []string{aA},
				},
			},
		},
		Volumes: map[string][]model.VolumeTarget{
			"vol": {
				{
					MountPoint:  "mnt1",
					Services:    []string{sA},
					AuxServices: []string{aA},
				},
			},
		},
		Dependencies: map[string]model.ModuleDependency{
			"ext": {
				Version: "ver",
				RequiredServices: map[string][]model.DependencyTarget{
					"c": {
						{
							RefVar:      "rVar2",
							Services:    []string{sA},
							AuxServices: []string{aA},
						},
					},
				},
			},
		},
		HostResources: map[string]model.HostResource{
			"res": {
				Resource: model.Resource{
					Tags: []string{"tag"},
					UserInput: &model.UserInput{
						Group: &ig,
					},
					Optional: false,
				},
				Targets: []model.HostResourceTarget{
					{
						MountPoint: "mnt2",
						Services:   []string{sA},
					},
				},
			},
		},
		Secrets: map[string]model.Secret{
			"sec": {
				Resource: model.Resource{
					Tags: []string{"tag"},
					UserInput: &model.UserInput{
						Group: &ig,
					},
					Optional: false,
				},
				Targets: []model.SecretTarget{
					{
						MountPoint: &sMnt,
						Services:   []string{sA},
					},
				},
			},
		},
		Configs: map[string]model.ConfigValue{
			"cfg": {
				DataType: &strType,
				UserInput: &model.ConfigUserInput{
					UserInput: model.UserInput{
						Group: &ig,
					},
				},
				Optional: false,
				Targets: []model.ConfigTarget{
					{
						RefVar:      "rVar3",
						Services:    []string{sA},
						AuxServices: []string{aA},
					},
				},
			},
		},
		InputGroups: map[string]model.InputGroup{
			ig: {},
		},
	}
	mc := make(module_lib.Configs)
	mc.SetString("cfg", nil, nil, false, "", nil, true)
	a := module_lib.Module{
		ID:             "id",
		Name:           "nme",
		Description:    "dsc",
		Tags:           map[string]struct{}{"tag": {}},
		License:        "lcs",
		Author:         "ath",
		Version:        "ver",
		Type:           "typ",
		DeploymentType: "dtp",
		Architectures:  map[module_lib.CPUArch]struct{}{"arch": {}},
		Services: map[string]*module_lib.Service{
			sA: {
				RunConfig: module_lib.RunConfig{
					MaxRetries:  5,
					RunOnce:     false,
					StopTimeout: 5 * time.Second,
					StopSignal:  nil,
					PseudoTTY:   false,
				},
				BindMounts: map[string]module_lib.BindMount{},
				Tmpfs:      map[string]module_lib.TmpfsMount{},
				Volumes: map[string]string{
					"mnt1": "vol",
				},
				HostResources: map[string]module_lib.HostResTarget{
					"mnt2": {
						Ref:      "res",
						ReadOnly: false,
					},
				},
				SecretMounts:  map[string]module_lib.SecretTarget{"mnt3": {Ref: "sec"}},
				Configs:       map[string]string{"rVar3": "cfg"},
				SrvReferences: map[string]module_lib.SrvRefTarget{"rVar1": {Ref: sB}},
				HttpEndpoints: map[string]module_lib.HttpEndpoint{},
				RequiredSrv:   map[string]struct{}{},
				RequiredBySrv: nil,
				ExtDependencies: map[string]module_lib.ExtDependencyTarget{
					"rVar2": {
						ID:      "ext",
						Service: "c",
					},
				},
				Ports: nil,
			},
			sB: {
				RunConfig: module_lib.RunConfig{
					MaxRetries:  5,
					RunOnce:     false,
					StopTimeout: 5 * time.Second,
					StopSignal:  nil,
					PseudoTTY:   false,
				},
				BindMounts:    map[string]module_lib.BindMount{},
				Tmpfs:         map[string]module_lib.TmpfsMount{},
				HttpEndpoints: map[string]module_lib.HttpEndpoint{},
				RequiredSrv:   map[string]struct{}{},
			},
		},
		AuxServices: map[string]*module_lib.AuxService{
			aA: {
				RunConfig: module_lib.RunConfig{
					MaxRetries:  5,
					RunOnce:     false,
					StopTimeout: 5 * time.Second,
					StopSignal:  nil,
					PseudoTTY:   false,
				},
				BindMounts: map[string]module_lib.BindMount{},
				Tmpfs:      map[string]module_lib.TmpfsMount{},
				Volumes: map[string]string{
					"mnt1": "vol",
				},
				Configs:       map[string]string{"rVar3": "cfg"},
				SrvReferences: map[string]module_lib.SrvRefTarget{"rVar1": {Ref: sB}},
				ExtDependencies: map[string]module_lib.ExtDependencyTarget{
					"rVar2": {
						ID:      "ext",
						Service: "c",
					},
				},
			},
			aB: {
				RunConfig: module_lib.RunConfig{
					MaxRetries:  5,
					RunOnce:     false,
					StopTimeout: 5 * time.Second,
					StopSignal:  nil,
					PseudoTTY:   false,
				},
				BindMounts:      map[string]module_lib.BindMount{},
				Tmpfs:           map[string]module_lib.TmpfsMount{},
				Volumes:         nil,
				Configs:         nil,
				SrvReferences:   nil,
				ExtDependencies: nil,
			},
		},
		AuxImgSrc: map[string]struct{}{"test": {}},
		Volumes: map[string]struct{}{
			"vol": {},
		},
		Dependencies: map[string]string{
			"ext": "ver",
		},
		HostResources: map[string]module_lib.HostResource{
			"res": {
				Resource: module_lib.Resource{
					Tags: map[string]struct{}{
						"tag": {},
					},
					Required: true,
				},
			},
		},
		Secrets: map[string]module_lib.Secret{
			"sec": {
				Resource: module_lib.Resource{
					Tags: map[string]struct{}{
						"tag": {},
					},
					Required: true,
				},
			},
		},
		Configs: mc,
		Inputs: module_lib.Inputs{
			Resources: map[string]module_lib.Input{
				"res": {
					Group: &ig,
				},
			},
			Secrets: map[string]module_lib.Input{
				"sec": {
					Group: &ig,
				},
			},
			Configs: map[string]module_lib.Input{
				"cfg": {
					Group: &ig,
				},
			},
			Groups: map[string]module_lib.InputGroup{
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
	if _, err := generator(&mf); err != nil {
		t.Error("err != nil")
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
		HostResources: map[string]model.HostResource{
			"": {
				Targets: []model.HostResourceTarget{
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
		Secrets: map[string]model.Secret{
			"": {
				Targets: []model.SecretTarget{
					{
						MountPoint: &sMnt,
						Services:   []string{""},
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
				DataType: &strType,
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

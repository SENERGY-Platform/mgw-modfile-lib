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
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator/configs"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator/generic"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator/inputs"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator/mounts"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator/services"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	module_lib "github.com/SENERGY-Platform/mgw-module-lib/model"
	"gopkg.in/yaml.v3"
)

func GetModule(yn *yaml.Node) (module_lib.Module, error) {
	var mf model.ModFile
	err := yn.Decode(&mf)
	if err != nil {
		return module_lib.Module{}, err
	}
	return generateModule(mf)
}

func generateModule(mf model.ModFile) (module_lib.Module, error) {
	mCs, err := configs.GenConfigs(mf.Configs)
	if err != nil {
		return module_lib.Module{}, err
	}
	mSs, err := services.GenServices(mf.Services)
	if err != nil {
		return module_lib.Module{}, err
	}
	mAs, err := services.GenAuxServices(mf.AuxServices)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetSrvReferences(mf.ServiceReferences, mSs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetAuxSrvReferences(mf.ServiceReferences, mAs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetVolumes(mf.Volumes, mSs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetAuxVolumes(mf.Volumes, mAs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetExtDependencies(mf.Dependencies, mSs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetAuxExtDependencies(mf.Dependencies, mAs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetHostResources(mf.HostResources, mSs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetSecrets(mf.Secrets, mSs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetConfigs(mf.Configs, mSs)
	if err != nil {
		return module_lib.Module{}, err
	}
	err = services.SetAuxConfigs(mf.Configs, mAs)
	if err != nil {
		return module_lib.Module{}, err
	}
	return module_lib.Module{
		ID:             mf.ID,
		Name:           mf.Name,
		Description:    mf.Description,
		Tags:           generic.GenStringSet(mf.Tags),
		License:        mf.License,
		Author:         mf.Author,
		Version:        mf.Version,
		Type:           mf.Type,
		DeploymentType: mf.DeploymentType,
		Architectures:  generic.GenStringSet(mf.Architectures),
		Services:       mSs,
		AuxServices:    mAs,
		AuxImgSrc:      generic.GenStringSet(mf.AuxImageSources),
		Volumes:        mounts.GenVolumes(mf.Volumes),
		Dependencies:   mounts.GenDependencies(mf.Dependencies),
		HostResources:  mounts.GenHostResources(mf.HostResources),
		Secrets:        mounts.GenSecrets(mf.Secrets),
		Configs:        mCs,
		Inputs: module_lib.Inputs{
			Resources: inputs.GenInputs(mf.HostResources),
			Secrets:   inputs.GenInputs(mf.Secrets),
			Configs:   inputs.GenInputs(mf.Configs),
			Groups:    inputs.GenInputGroups(mf.InputGroups),
		},
	}, nil
}

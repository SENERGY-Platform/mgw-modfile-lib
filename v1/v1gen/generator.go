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
	"errors"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/v1gen/configs"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/v1gen/generic"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/v1gen/inputs"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/v1gen/mounts"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/v1gen/services"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
)

func generator(f any) (*module.Module, error) {
	mf, ok := f.(*model.ModFile)
	if !ok {
		return nil, errors.New("invalid type")
	}
	mCs, err := configs.GenConfigs(mf.Configs)
	if err != nil {
		return nil, err
	}
	mSs, err := services.GenServices(mf.Services)
	if err != nil {
		return nil, err
	}
	err = services.SetSrvReferences(mf.ServiceReferences, mSs)
	if err != nil {
		return nil, err
	}
	err = services.SetVolumes(mf.Volumes, mSs)
	if err != nil {
		return nil, err
	}
	err = services.SetExtDependencies(mf.Dependencies, mSs)
	if err != nil {
		return nil, err
	}
	err = services.SetHostResources(mf.Resources, mSs)
	if err != nil {
		return nil, err
	}
	err = services.SetSecrets(mf.Secrets, mSs)
	if err != nil {
		return nil, err
	}
	err = services.SetConfigs(mf.Configs, mSs)
	if err != nil {
		return nil, err
	}
	return &module.Module{
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
		Volumes:        mounts.GenVolumes(mf.Volumes),
		Dependencies:   mounts.GenDependencies(mf.Dependencies),
		HostResources:  mounts.GenHostResources(mf.Resources),
		Secrets:        mounts.GenSecrets(mf.Secrets),
		Configs:        mCs,
		Inputs: module.Inputs{
			Resources: inputs.GenInputs(mf.Resources),
			Secrets:   inputs.GenInputs(mf.Secrets),
			Configs:   inputs.GenInputs(mf.Configs),
			Groups:    inputs.GenInputGroups(mf.InputGroups),
		},
	}, nil
}

func GetGenerator() (string, func(any) (*module.Module, error)) {
	return model.Version, generator
}

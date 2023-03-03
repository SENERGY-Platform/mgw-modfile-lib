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

package v1

import "github.com/SENERGY-Platform/mgw-module-lib/module"

func (mf *ModFile) GenModule() (*module.Module, error) {
	mCs, err := GenConfigs(mf.Configs)
	if err != nil {
		return nil, err
	}
	mSs, err := GenServices(mf.Services)
	if err != nil {
		return nil, err
	}
	err = SetSrvReferences(mf.ServiceReferences, mSs)
	if err != nil {
		return nil, err
	}
	err = SetVolumes(mf.Volumes, mSs)
	if err != nil {
		return nil, err
	}
	err = SetExtDependencies(mf.Dependencies, mSs)
	if err != nil {
		return nil, err
	}
	err = SetResources(mf.Resources, mSs)
	if err != nil {
		return nil, err
	}
	err = SetSecrets(mf.Secrets, mSs)
	if err != nil {
		return nil, err
	}
	err = SetConfigs(mf.Configs, mSs)
	if err != nil {
		return nil, err
	}
	return &module.Module{
		ID:             mf.ID,
		Name:           mf.Name,
		Description:    mf.Description,
		Tags:           GenStringSet(mf.Tags),
		License:        mf.License,
		Author:         mf.Author,
		Version:        mf.Version,
		Type:           mf.Type,
		DeploymentType: mf.DeploymentType,
		Architectures:  GenStringSet(mf.Architectures),
		Services:       mSs,
		Volumes:        GenVolumes(mf.Volumes),
		Dependencies:   GenDependencies(mf.Dependencies),
		Resources:      GenResources(mf.Resources),
		Secrets:        GenSecrets(mf.Secrets),
		Configs:        mCs,
		Inputs: module.Inputs{
			Resources: GenInputs(mf.Resources),
			Secrets:   GenInputs(mf.Secrets),
			Configs:   GenInputs(mf.Configs),
			Groups:    GenInputGroups(mf.InputGroups),
		},
	}, nil
}

func GenStringSet(sl []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, s := range sl {
		set[s] = struct{}{}
	}
	return set
}

func GenVolumes(mfVs map[string][]VolumeTarget) map[string]struct{} {
	set := make(map[string]struct{})
	for mfV := range mfVs {
		set[mfV] = struct{}{}
	}
	return set
}

func GenDependencies(mfMDs map[string]ModuleDependency) map[string]string {
	mDs := make(map[string]string)
	for id, mfMD := range mfMDs {
		mDs[id] = mfMD.Version
	}
	return mDs
}

func GenResources(mfRs map[string]Resource) map[string]map[string]struct{} {
	mRs := make(map[string]map[string]struct{})
	for ref, mfR := range mfRs {
		mRs[ref] = GenStringSet(mfR.Tags)
	}
	return mRs
}

func GenSecrets(mfSs map[string]Secret) map[string]module.Secret {
	mSs := make(map[string]module.Secret)
	for ref, mfS := range mfSs {
		mSs[ref] = module.Secret{
			Type: mfS.Type,
			Tags: GenStringSet(mfS.Tags),
		}
	}
	return mSs
}

func GenInputs[T configurable](mfCs map[string]T) map[string]module.Input {
	mIs := make(map[string]module.Input)
	for ref, mfC := range mfCs {
		mfUI := mfC.GetUserInput()
		if mfUI != nil {
			mIs[ref] = module.Input(*mfUI)
		}
	}
	return mIs
}

func GenInputGroups(mfIGs map[string]InputGroup) map[string]module.InputGroup {
	mIGs := make(map[string]module.InputGroup)
	for ref, mfIG := range mfIGs {
		mIGs[ref] = module.InputGroup{
			Name:        mfIG.Name,
			Description: mfIG.Description,
			Group:       mfIG.Group,
		}
	}
	return mIGs
}

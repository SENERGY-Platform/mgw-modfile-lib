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
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/v1gen/generic"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
)

func GenVolumes(mfVs map[string][]model.VolumeTarget) map[string]struct{} {
	set := make(map[string]struct{})
	for mfV := range mfVs {
		set[mfV] = struct{}{}
	}
	return set
}

func GenDependencies(mfMDs map[string]model.ModuleDependency) map[string]string {
	mDs := make(map[string]string)
	for id, mfMD := range mfMDs {
		mDs[id] = mfMD.Version
	}
	return mDs
}

func GenResources(mfRs map[string]model.Resource) map[string]map[string]struct{} {
	mRs := make(map[string]map[string]struct{})
	for ref, mfR := range mfRs {
		mRs[ref] = generic.GenStringSet(mfR.Tags)
	}
	return mRs
}

func GenSecrets(mfSs map[string]model.Secret) map[string]module.Secret {
	mSs := make(map[string]module.Secret)
	for ref, mfS := range mfSs {
		mSs[ref] = module.Secret{
			Type: mfS.Type,
			Tags: generic.GenStringSet(mfS.Tags),
		}
	}
	return mSs
}

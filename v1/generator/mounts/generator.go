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
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/generator/generic"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	module_lib "github.com/SENERGY-Platform/mgw-module-lib/model"
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

func GenHostResources(mfRs map[string]model.HostResource) map[string]module_lib.HostResource {
	mRs := make(map[string]module_lib.HostResource)
	for ref, mfR := range mfRs {
		mRs[ref] = module_lib.HostResource{
			Resource: module_lib.Resource{
				Tags:     generic.GenStringSet(mfR.Tags),
				Required: !mfR.Optional,
			},
		}
	}
	return mRs
}

func GenSecrets(mfSs map[string]model.Secret) map[string]module_lib.Secret {
	mSs := make(map[string]module_lib.Secret)
	for ref, mfS := range mfSs {
		mSs[ref] = module_lib.Secret{
			Resource: module_lib.Resource{
				Tags:     generic.GenStringSet(mfS.Tags),
				Required: !mfS.Optional,
			},
			Type: mfS.Type,
		}
	}
	return mSs
}

func GenFiles(mfFiles map[string]model.File) map[string]module_lib.File {
	mFiles := make(map[string]module_lib.File)
	for ref, file := range mfFiles {
		mFiles[ref] = module_lib.File{
			Source:   file.Source,
			Type:     file.UserInput.Type,
			Required: !file.Optional,
		}
	}
	return mFiles
}

func GenFileGroups(mfFileGroups map[string]model.FileGroup) map[string]struct{} {
	mFileGroups := make(map[string]struct{})
	for ref := range mfFileGroups {
		mFileGroups[ref] = struct{}{}
	}
	return mFileGroups
}

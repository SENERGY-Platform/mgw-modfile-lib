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
	"fmt"

	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	module_lib "github.com/SENERGY-Platform/mgw-module-lib/model"
)

func SetSrvReferences(mfSRs map[string][]model.DependencyTarget, mSs map[string]module_lib.Service) error {
	for ref, mfDTs := range mfSRs {
		for _, mfDT := range mfDTs {
			for _, tRef := range mfDT.Services {
				mS, ok := mSs[tRef]
				if !ok {
					return fmt.Errorf("invalid service reference: service '%s' not defined", tRef)
				}
				if mS.SrvReferences == nil {
					mS.SrvReferences = make(map[string]module_lib.SrvRefTarget)
				}
				if r, k := mS.SrvReferences[mfDT.RefVar]; k {
					if r.Ref == ref {
						continue
					}
					return fmt.Errorf("service '%s' invalid service reference: duplicate '%s'", tRef, mfDT.RefVar)
				}
				mS.SrvReferences[mfDT.RefVar] = module_lib.SrvRefTarget{
					Ref:      ref,
					Template: mfDT.Template,
				}
				mSs[tRef] = mS
			}
		}
	}
	return nil
}

func SetAuxSrvReferences(mfSRs map[string][]model.DependencyTarget, mAs map[string]module_lib.AuxService) error {
	for ref, mfDTs := range mfSRs {
		for _, mfDT := range mfDTs {
			for _, tRef := range mfDT.AuxServices {
				mA, ok := mAs[tRef]
				if !ok {
					return fmt.Errorf("invalid service reference: aux service '%s' not defined", tRef)
				}
				if mA.SrvReferences == nil {
					mA.SrvReferences = make(map[string]module_lib.SrvRefTarget)
				}
				if r, k := mA.SrvReferences[mfDT.RefVar]; k {
					if r.Ref == ref {
						continue
					}
					return fmt.Errorf("aux service '%s' invalid service reference: duplicate '%s'", tRef, mfDT.RefVar)
				}
				mA.SrvReferences[mfDT.RefVar] = module_lib.SrvRefTarget{
					Ref:      ref,
					Template: mfDT.Template,
				}
				mAs[tRef] = mA
			}
		}
	}
	return nil
}

func SetVolumes(mfVs map[string][]model.VolumeTarget, mSs map[string]module_lib.Service) error {
	for mfV, mfVTs := range mfVs {
		for _, mfVT := range mfVTs {
			for _, ref := range mfVT.Services {
				mS, ok := mSs[ref]
				if !ok {
					return fmt.Errorf("invalid volume: service '%s' not defined", ref)
				}
				if mS.Volumes == nil {
					mS.Volumes = make(map[string]string)
				}
				if v, k := mS.Volumes[mfVT.MountPoint]; k {
					if v == mfV {
						continue
					}
					return fmt.Errorf("service '%s' invalid volume: duplicate '%s'", ref, mfVT.MountPoint)
				}
				mS.Volumes[mfVT.MountPoint] = mfV
				mSs[ref] = mS
			}
		}
	}
	return nil
}

func SetAuxVolumes(mfVs map[string][]model.VolumeTarget, mAs map[string]module_lib.AuxService) error {
	for mfV, mfVTs := range mfVs {
		for _, mfVT := range mfVTs {
			for _, ref := range mfVT.AuxServices {
				mA, ok := mAs[ref]
				if !ok {
					return fmt.Errorf("invalid volume: aux service '%s' not defined", ref)
				}
				if mA.Volumes == nil {
					mA.Volumes = make(map[string]string)
				}
				if v, k := mA.Volumes[mfVT.MountPoint]; k {
					if v == mfV {
						continue
					}
					return fmt.Errorf("aux service '%s' invalid volume: duplicate '%s'", ref, mfVT.MountPoint)
				}
				mA.Volumes[mfVT.MountPoint] = mfV
				mAs[ref] = mA
			}
		}
	}
	return nil
}

func SetExtDependencies(mfMDs map[string]model.ModuleDependency, mSs map[string]module_lib.Service) error {
	for extId, mfMD := range mfMDs {
		for extRef, mfDTs := range mfMD.RequiredServices {
			for _, mfDT := range mfDTs {
				for _, ref := range mfDT.Services {
					mS, ok := mSs[ref]
					if !ok {
						return fmt.Errorf("invalid module dependency: service '%s' not defined", ref)
					}
					if mS.ExtDependencies == nil {
						mS.ExtDependencies = make(map[string]module_lib.ExtDependencyTarget)
					}
					if etd, k := mS.ExtDependencies[mfDT.RefVar]; k {
						if etd.ID == extId && etd.Service == extRef {
							continue
						}
						return fmt.Errorf("service '%s' invalid module dependency: duplicate '%s'", ref, mfDT.RefVar)
					}
					mS.ExtDependencies[mfDT.RefVar] = module_lib.ExtDependencyTarget{
						ID:       extId,
						Service:  extRef,
						Template: mfDT.Template,
					}
					mSs[ref] = mS
				}
			}
		}
	}
	return nil
}

func SetAuxExtDependencies(mfMDs map[string]model.ModuleDependency, mAs map[string]module_lib.AuxService) error {
	for extId, mfMD := range mfMDs {
		for extRef, mfDTs := range mfMD.RequiredServices {
			for _, mfDT := range mfDTs {
				for _, ref := range mfDT.AuxServices {
					mA, ok := mAs[ref]
					if !ok {
						return fmt.Errorf("invalid module dependency: aux service '%s' not defined", ref)
					}
					if mA.ExtDependencies == nil {
						mA.ExtDependencies = make(map[string]module_lib.ExtDependencyTarget)
					}
					if etd, k := mA.ExtDependencies[mfDT.RefVar]; k {
						if etd.ID == extId && etd.Service == extRef {
							continue
						}
						return fmt.Errorf("aux service '%s' invalid module dependency: duplicate '%s'", ref, mfDT.RefVar)
					}
					mA.ExtDependencies[mfDT.RefVar] = module_lib.ExtDependencyTarget{
						ID:       extId,
						Service:  extRef,
						Template: mfDT.Template,
					}
					mAs[ref] = mA
				}
			}
		}
	}
	return nil
}

func SetHostResources(mfRs map[string]model.HostResource, mSs map[string]module_lib.Service) error {
	for rRef, mfR := range mfRs {
		for _, mfRT := range mfR.Targets {
			for _, sRef := range mfRT.Services {
				mS, ok := mSs[sRef]
				if !ok {
					return fmt.Errorf("invalid resource: service '%s' not defined", sRef)
				}
				if mS.HostResources == nil {
					mS.HostResources = make(map[string]module_lib.HostResTarget)
				}
				if mRT, k := mS.HostResources[mfRT.MountPoint]; k {
					if mRT.Ref == rRef && mRT.ReadOnly == mfRT.ReadOnly {
						continue
					}
					return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", mRT.Ref, rRef, sRef, mfRT.MountPoint)
				}
				mS.HostResources[mfRT.MountPoint] = module_lib.HostResTarget{
					Ref:      rRef,
					ReadOnly: mfRT.ReadOnly,
				}
				mSs[sRef] = mS
			}
		}
	}
	return nil
}

func SetFiles(mfFiles map[string]model.File, mSs map[string]module_lib.Service) error {
	for fRef, file := range mfFiles {
		for _, target := range file.Targets {
			for _, sRef := range target.Services {
				mS, ok := mSs[sRef]
				if !ok {
					return fmt.Errorf("invalid file: service '%s' not defined", sRef)
				}
				if mS.Files == nil {
					mS.Files = make(map[string]module_lib.FileTarget)
				}
				if mFT, k := mS.Files[target.MountPoint]; k {
					if mFT.Ref == fRef && mFT.ReadOnly == target.ReadOnly {
						continue
					}
					return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", mFT.Ref, fRef, sRef, target.MountPoint)
				}
				mS.Files[target.MountPoint] = module_lib.FileTarget{
					Ref:      fRef,
					ReadOnly: target.ReadOnly,
				}
				mSs[sRef] = mS
			}
		}
	}
	return nil
}

func SetFileGroups(mfFileGroups map[string]model.FileGroup, mSs map[string]module_lib.Service) error {
	for gRef, fileGroup := range mfFileGroups {
		for _, target := range fileGroup.Targets {
			for _, sRef := range target.Services {
				mS, ok := mSs[sRef]
				if !ok {
					return fmt.Errorf("invalid file group: service '%s' not defined", sRef)
				}
				if mS.FileGroups == nil {
					mS.FileGroups = make(map[string]string)
				}
				if r, k := mS.FileGroups[target.BasePath]; k {
					if r == gRef {
						continue
					}
					return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", r, gRef, sRef, target.BasePath)
				}
				mS.FileGroups[target.BasePath] = gRef
				mSs[sRef] = mS
			}
		}
	}
	return nil
}

func SetSecrets(mfSecrets map[string]model.Secret, mServices map[string]module_lib.Service) error {
	for secRef, mfSecret := range mfSecrets {
		for _, mfSecretTarget := range mfSecret.Targets {
			if mfSecretTarget.MountPoint != "" {
				for _, mfSrvRef := range mfSecretTarget.Services {
					mService, ok := mServices[mfSrvRef]
					if !ok {
						return fmt.Errorf("invalid secret: service '%s' not defined", mfSrvRef)
					}
					if mService.SecretMounts == nil {
						mService.SecretMounts = make(map[string]module_lib.SecretTarget)
					}
					if mSecretTarget, k := mService.SecretMounts[mfSecretTarget.MountPoint]; k {
						if mSecretTarget.Ref == secRef {
							continue
						}
						return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", mSecretTarget.Ref, secRef, mfSrvRef, mfSecretTarget.MountPoint)
					}
					mService.SecretMounts[mfSecretTarget.MountPoint] = module_lib.SecretTarget{
						Ref:  secRef,
						Item: mfSecretTarget.Item,
					}
					mServices[mfSrvRef] = mService
				}
			}
			if mfSecretTarget.RefVar != "" {
				for _, mfSrvRef := range mfSecretTarget.Services {
					mService, ok := mServices[mfSrvRef]
					if !ok {
						return fmt.Errorf("invalid secret: service '%s' not defined", mfSrvRef)
					}
					if mService.SecretVars == nil {
						mService.SecretVars = make(map[string]module_lib.SecretTarget)
					}
					if mSecretTarget, k := mService.SecretVars[mfSecretTarget.RefVar]; k {
						if mSecretTarget.Ref == secRef {
							continue
						}
						return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", mSecretTarget.Ref, secRef, mfSrvRef, mfSecretTarget.RefVar)
					}
					mService.SecretVars[mfSecretTarget.RefVar] = module_lib.SecretTarget{
						Ref:  secRef,
						Item: mfSecretTarget.Item,
					}
					mServices[mfSrvRef] = mService
				}
			}
		}
	}
	return nil
}

func SetConfigs(mfCVs map[string]model.ConfigValue, mSs map[string]module_lib.Service) error {
	for cRef, mfCV := range mfCVs {
		for _, mfCT := range mfCV.Targets {
			for _, sRef := range mfCT.Services {
				mS, ok := mSs[sRef]
				if !ok {
					return fmt.Errorf("invalid config: service '%s' not defined", sRef)
				}
				if mS.Configs == nil {
					mS.Configs = make(map[string]string)
				}
				if r, k := mS.Configs[mfCT.RefVar]; k {
					if r == cRef {
						continue
					}
					return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", r, cRef, sRef, mfCT.RefVar)
				}
				mS.Configs[mfCT.RefVar] = cRef
				mSs[sRef] = mS
			}
		}
	}
	return nil
}

func SetAuxConfigs(mfCVs map[string]model.ConfigValue, mAs map[string]module_lib.AuxService) error {
	for cRef, mfCV := range mfCVs {
		for _, mfCT := range mfCV.Targets {
			for _, sRef := range mfCT.AuxServices {
				mA, ok := mAs[sRef]
				if !ok {
					return fmt.Errorf("invalid config: aux service '%s' not defined", sRef)
				}
				if mA.Configs == nil {
					mA.Configs = make(map[string]string)
				}
				if r, k := mA.Configs[mfCT.RefVar]; k {
					if r == cRef {
						continue
					}
					return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", r, cRef, sRef, mfCT.RefVar)
				}
				mA.Configs[mfCT.RefVar] = cRef
				mAs[sRef] = mA
			}
		}
	}
	return nil
}

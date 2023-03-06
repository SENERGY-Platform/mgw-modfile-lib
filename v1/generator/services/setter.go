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
	"github.com/SENERGY-Platform/mgw-module-lib/module"
)

func SetSrvReferences(mfSRs map[string][]model.DependencyTarget, mSs map[string]*module.Service) error {
	for ref, mfDTs := range mfSRs {
		for _, mfDT := range mfDTs {
			for _, tRef := range mfDT.Services {
				if mS, ok := mSs[tRef]; ok {
					if mS.SrvReferences == nil {
						mS.SrvReferences = make(map[string]string)
					}
					if r, k := mS.SrvReferences[mfDT.RefVar]; k {
						if r == ref {
							continue
						}
						return fmt.Errorf("service '%s' invalid service reference: duplicate '%s'", tRef, mfDT.RefVar)
					}
					mS.SrvReferences[mfDT.RefVar] = ref
				} else {
					return fmt.Errorf("invalid service reference: service '%s' not defined", tRef)
				}
			}
		}
	}
	return nil
}

func SetVolumes(mfVs map[string][]model.VolumeTarget, mSs map[string]*module.Service) error {
	for mfV, mfVTs := range mfVs {
		for _, mfVT := range mfVTs {
			for _, ref := range mfVT.Services {
				if mS, ok := mSs[ref]; ok {
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
				} else {
					return fmt.Errorf("invalid volume: service '%s' not defined", ref)
				}
			}
		}
	}
	return nil
}

func SetExtDependencies(mfMDs map[string]model.ModuleDependency, mSs map[string]*module.Service) error {
	for extId, mfMD := range mfMDs {
		for extRef, mfDTs := range mfMD.RequiredServices {
			for _, mfDT := range mfDTs {
				for _, ref := range mfDT.Services {
					if mS, ok := mSs[ref]; ok {
						if mS.ExtDependencies == nil {
							mS.ExtDependencies = make(map[string]module.ExtDependencyTarget)
						}
						if etd, k := mS.ExtDependencies[mfDT.RefVar]; k {
							if etd.ID == extId && etd.Service == extRef {
								continue
							}
							return fmt.Errorf("service '%s' invalid module dependency: duplicate '%s'", ref, mfDT.RefVar)
						}
						mS.ExtDependencies[mfDT.RefVar] = module.ExtDependencyTarget{
							ID:      extId,
							Service: extRef,
						}
					} else {
						return fmt.Errorf("invalid module dependency: service '%s' not defined", ref)
					}
				}
			}
		}
	}
	return nil
}

func SetResources(mfRs map[string]model.Resource, mSs map[string]*module.Service) error {
	for rRef, mfR := range mfRs {
		for _, mfRT := range mfR.Targets {
			for _, sRef := range mfRT.Services {
				if mS, ok := mSs[sRef]; ok {
					if mS.Resources == nil {
						mS.Resources = make(map[string]module.ResourceTarget)
					}
					if mRT, k := mS.Resources[mfRT.MountPoint]; k {
						if mRT.Ref == rRef && mRT.ReadOnly == mfRT.ReadOnly {
							continue
						}
						return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", mRT.Ref, rRef, sRef, mfRT.MountPoint)
					}
					mS.Resources[mfRT.MountPoint] = module.ResourceTarget{
						Ref:      rRef,
						ReadOnly: mfRT.ReadOnly,
					}
				} else {
					return fmt.Errorf("invalid resource: service '%s' not defined", sRef)
				}
			}
		}
	}
	return nil
}

func SetSecrets(mfSCTs map[string]model.Secret, mSs map[string]*module.Service) error {
	for sctRef, mfSCT := range mfSCTs {
		for _, mfRTB := range mfSCT.Targets {
			for _, sRef := range mfRTB.Services {
				if mS, ok := mSs[sRef]; ok {
					if mS.Secrets == nil {
						mS.Secrets = make(map[string]string)
					}
					if r, k := mS.Secrets[mfRTB.MountPoint]; k {
						if r == sctRef {
							continue
						}
						return fmt.Errorf("'%s' & '%s' -> '%s' -> '%s'", r, sctRef, sRef, mfRTB.MountPoint)
					}
					mS.Secrets[mfRTB.MountPoint] = sctRef
				} else {
					return fmt.Errorf("invalid secret: service '%s' not defined", sRef)
				}
			}
		}
	}
	return nil
}

func SetConfigs(mfCVs map[string]model.ConfigValue, mSs map[string]*module.Service) error {
	for cRef, mfCV := range mfCVs {
		for _, mfCT := range mfCV.Targets {
			for _, sRef := range mfCT.Services {
				if mS, ok := mSs[sRef]; ok {
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
				} else {
					return fmt.Errorf("invalid config: service '%s' not defined", sRef)
				}
			}
		}
	}
	return nil
}

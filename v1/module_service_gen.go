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

import (
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"io/fs"
	"time"
)

func GenServices(mfSs map[string]Service) (map[string]*module.Service, error) {
	mSs := make(map[string]*module.Service)
	for ref, mfS := range mfSs {
		mBMs, err := GenBindMounts(mfS.Include)
		if err != nil {
			return nil, fmt.Errorf("service '%s' invalid bind mount: %s", ref, err)
		}
		mTMs, err := GenTmpfsMounts(mfS.Tmpfs)
		if err != nil {
			return nil, fmt.Errorf("service '%s' invalid tmpfsMount: %s", ref, err)
		}
		mHEs, err := GenHttpEndpoints(mfS.HttpEndpoints)
		if err != nil {
			return nil, fmt.Errorf("service '%s' invalid http endpoint: %s", ref, err)
		}
		mPs, err := GenPorts(mfS.Ports)
		if err != nil {
			return nil, fmt.Errorf("service '%s' invalid port mapping: %s", ref, err)
		}
		mSs[ref] = &module.Service{
			Name:          mfS.Name,
			Image:         mfS.Image,
			RunConfig:     GenRunConfig(mfS.RunConfig),
			BindMounts:    mBMs,
			Tmpfs:         mTMs,
			HttpEndpoints: mHEs,
			RequiredSrv:   GenStringSet(mfS.RequiredServices),
			Ports:         mPs,
		}
	}
	return mSs, nil
}

func GenRunConfig(mfRC RunConfig) module.RunConfig {
	mRC := module.RunConfig{
		MaxRetries: mfRC.MaxRetries,
		RunOnce:    mfRC.RunOnce,
		StopSignal: mfRC.StopSignal,
		PseudoTTY:  mfRC.PseudoTTY,
	}
	if mfRC.StopTimeout != nil {
		mRC.StopTimeout = (*time.Duration)(mfRC.StopTimeout)
	}
	return mRC
}

func GenBindMounts(mfBMs []BindMount) (map[string]module.BindMount, error) {
	mBMs := make(map[string]module.BindMount)
	for _, mfBM := range mfBMs {
		if v, ok := mBMs[mfBM.MountPoint]; ok {
			if v.Source == mfBM.Source && v.ReadOnly == mfBM.ReadOnly {
				continue
			}
			return nil, fmt.Errorf("duplicate '%s'", mfBM.MountPoint)
		}
		mBMs[mfBM.MountPoint] = module.BindMount{
			Source:   mfBM.Source,
			ReadOnly: mfBM.ReadOnly,
		}
	}
	return mBMs, nil
}

func GenTmpfsMounts(mfTMs []TmpfsMount) (map[string]module.TmpfsMount, error) {
	mTMs := make(map[string]module.TmpfsMount)
	for _, mfTM := range mfTMs {
		if _, ok := mTMs[mfTM.MountPoint]; ok {
			return nil, fmt.Errorf("duplicate '%s'", mfTM.MountPoint)
		}
		mTM := module.TmpfsMount{Size: uint64(mfTM.Size)}
		if mfTM.Mode != nil {
			mTM.Mode = (*fs.FileMode)(mfTM.Mode)
		}
		mTMs[mfTM.MountPoint] = mTM
	}
	return mTMs, nil
}

func GenHttpEndpoints(mfHEs []HttpEndpoint) (map[string]module.HttpEndpoint, error) {
	mHEs := make(map[string]module.HttpEndpoint)
	for _, mfHE := range mfHEs {
		p := mfHE.Path
		if mfHE.ExtPath != nil {
			p = *mfHE.ExtPath
		}
		if v, ok := mHEs[p]; ok {
			if v.Name == mfHE.Name && v.Port == mfHE.Port && v.Path == mfHE.Path {
				continue
			}
			return nil, fmt.Errorf("duplicate '%s'", mfHE.Path)
		}
		mHEs[p] = module.HttpEndpoint{
			Name: mfHE.Name,
			Port: mfHE.Port,
			Path: mfHE.Path,
		}
	}
	return mHEs, nil
}

func GenPorts(mfSPs []SrvPort) ([]module.Port, error) {
	var mPs []module.Port
	for _, mfSP := range mfSPs {
		proto := module.TcpPort
		if mfSP.Protocol != nil {
			proto = *mfSP.Protocol
		}
		ep, err := mfSP.Port.Parse()
		if err != nil {
			return nil, err
		}
		var hp []uint
		if mfSP.HostPort != nil {
			hp, err = mfSP.HostPort.Parse()
			if err != nil {
				return nil, err
			}
		}
		lep := len(ep)
		lhp := len(hp)
		if lep > lhp {
			return nil, errors.New("range mismatch: ports > host ports")
		}
		if lep > 1 && lep < lhp {
			return nil, errors.New("range mismatch: ports < host ports")
		}
		if lep == 1 {
			mPs = append(mPs, module.Port{
				Name:     mfSP.Name,
				Number:   ep[0],
				Protocol: proto,
				Bindings: hp,
			})
		} else {
			for i, n := range ep {
				mPs = append(mPs, module.Port{
					Name:     mfSP.Name,
					Number:   n,
					Protocol: proto,
					Bindings: []uint{hp[i]},
				})
			}
		}
	}
	return mPs, nil
}
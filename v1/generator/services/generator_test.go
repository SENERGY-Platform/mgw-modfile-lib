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
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"reflect"
	"testing"
	"time"
)

func TestGenRunConfig(t *testing.T) {
	a := module.RunConfig{
		MaxRetries:  3,
		RunOnce:     false,
		StopTimeout: 5 * time.Second,
		StopSignal:  nil,
		PseudoTTY:   false,
	}
	if b := GenRunConfig(model.RunConfig{}); reflect.DeepEqual(a, b) == false {
		t.Errorf("%+v != %+v", a, b)
	}
	str := "test"
	i := 1
	d := model.Duration(1 * time.Second)
	c := model.RunConfig{
		MaxRetries:  &i,
		RunOnce:     true,
		StopTimeout: &d,
		StopSignal:  &str,
		PseudoTTY:   true,
	}
	a = module.RunConfig{
		MaxRetries:  1,
		RunOnce:     true,
		StopTimeout: 1 * time.Second,
		StopSignal:  &str,
		PseudoTTY:   true,
	}
	if b := GenRunConfig(c); reflect.DeepEqual(a, b) == false {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestGenBindMounts(t *testing.T) {
	var mfBMs []model.BindMount
	if bm, err := GenBindMounts(mfBMs); err != nil {
		t.Error("err != nil")
	} else if len(bm) != 0 {
		t.Errorf("len(%v) != 0", bm)
	}
	// --------------------------------
	str := "test"
	str2 := "test2"
	mfBMs = append(mfBMs, model.BindMount{
		MountPoint: str,
		Source:     str2,
		ReadOnly:   true,
	})
	a := module.BindMount{
		Source:   str2,
		ReadOnly: true,
	}
	if bm, err := GenBindMounts(mfBMs); err != nil {
		t.Error("err != nil")
	} else if len(bm) != 1 {
		t.Errorf("len(%v) != 1", bm)
	} else if b, ok := bm[str]; !ok {
		t.Errorf("b, ok := bm[%s]; !ok", str)
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%+v != %+v", a, b)
	}
	// --------------------------------
	mfBMs = append(mfBMs, model.BindMount{
		MountPoint: str,
		Source:     str2,
		ReadOnly:   true,
	})
	if bm, err := GenBindMounts(mfBMs); err != nil {
		t.Error("err != nil")
	} else if len(bm) != 1 {
		t.Errorf("len(%v) != 1", bm)
	}
	mfBMs = mfBMs[:len(mfBMs)-1]
	// --------------------------------
	mfBMs = append(mfBMs, model.BindMount{
		MountPoint: str,
		Source:     "test3",
		ReadOnly:   true,
	})
	if _, err := GenBindMounts(mfBMs); err == nil {
		t.Error("err == nil")
	}
	mfBMs = mfBMs[:len(mfBMs)-1]
	// --------------------------------
	mfBMs = append(mfBMs, model.BindMount{
		MountPoint: str,
		Source:     str2,
		ReadOnly:   false,
	})
	if _, err := GenBindMounts(mfBMs); err == nil {
		t.Error("err == nil")
	}
}

func TestGenTmpfsMounts(t *testing.T) {
	var mfTMs []model.TmpfsMount
	if tm, err := GenTmpfsMounts(mfTMs); err != nil {
		t.Error("err != nil")
	} else if len(tm) != 0 {
		t.Errorf("len(%v) != 0", tm)
	}
	// --------------------------------
	str := "test"
	mfTMs = append(mfTMs, model.TmpfsMount{
		MountPoint: str,
		Size:       64,
		Mode:       nil,
	})
	a := module.TmpfsMount{
		Size: 64,
		Mode: 504,
	}
	if tm, err := GenTmpfsMounts(mfTMs); err != nil {
		t.Error("err != nil")
	} else if len(tm) != 1 {
		t.Errorf("len(%v) != 1", tm)
	} else if b, ok := tm[str]; !ok {
		t.Errorf("m, ok := tm[%s]; !ok", str)
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%+v != %+v", a, b)
	}
	// --------------------------------
	mfTMs = append(mfTMs, model.TmpfsMount{
		MountPoint: str,
		Size:       64,
		Mode:       nil,
	})
	if tm, err := GenTmpfsMounts(mfTMs); err != nil {
		t.Error("err != nil")
	} else if len(tm) != 1 {
		t.Errorf("len(%v) != 1", tm)
	}
	mfTMs = mfTMs[:len(mfTMs)-1]
	// --------------------------------
	fm := model.FileMode(504)
	mfTMs = append(mfTMs, model.TmpfsMount{
		MountPoint: str,
		Size:       64,
		Mode:       &fm,
	})
	if tm, err := GenTmpfsMounts(mfTMs); err != nil {
		t.Error("err != nil")
	} else if len(tm) != 1 {
		t.Errorf("len(%v) != 1", tm)
	}
	mfTMs = mfTMs[:len(mfTMs)-1]
	// --------------------------------
	mfTMs = append(mfTMs, model.TmpfsMount{
		MountPoint: str,
		Size:       32,
		Mode:       nil,
	})
	if _, err := GenTmpfsMounts(mfTMs); err == nil {
		t.Error("err == nil")
	}
	mfTMs = mfTMs[:len(mfTMs)-1]
	// --------------------------------
	fm = model.FileMode(511)
	mfTMs = append(mfTMs, model.TmpfsMount{
		MountPoint: str,
		Size:       64,
		Mode:       &fm,
	})
	if _, err := GenTmpfsMounts(mfTMs); err == nil {
		t.Error("err == nil")
	}
	mfTMs = mfTMs[:len(mfTMs)-1]
	// --------------------------------
	str2 := "test2"
	fm2 := model.FileMode(511)
	mfTMs = append(mfTMs, model.TmpfsMount{
		MountPoint: str2,
		Size:       64,
		Mode:       &fm2,
	})
	a = module.TmpfsMount{
		Size: 64,
		Mode: 511,
	}
	if tm, err := GenTmpfsMounts(mfTMs); err != nil {
		t.Error("err != nil")
	} else if len(tm) != 2 {
		t.Errorf("len(%v) != 2", tm)
	} else if b, ok := tm[str2]; !ok {
		t.Errorf("m, ok := tm[%s]; !ok", str2)
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestGenHttpEndpoints(t *testing.T) {
	var mfHEs []model.HttpEndpoint
	if ep, err := GenHttpEndpoints(mfHEs); err != nil {
		t.Error("err != nil")
	} else if len(ep) != 0 {
		t.Errorf("len(%v) != 0", ep)
	}
	// --------------------------------
	str := "test"
	str2 := "test2"
	mfHEs = append(mfHEs, model.HttpEndpoint{
		Name:    str,
		Path:    str2,
		Port:    nil,
		ExtPath: nil,
	})
	a := module.HttpEndpoint{
		Name: str,
		Port: nil,
		Path: str2,
	}
	if ep, err := GenHttpEndpoints(mfHEs); err != nil {
		t.Error("err != nil")
	} else if len(ep) != 1 {
		t.Errorf("len(%v) != 1", ep)
	} else if b, ok := ep[str2]; !ok {
		t.Errorf("b, ok := ep[%s]; !ok", str2)
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// --------------------------------
	mfHEs = append(mfHEs, model.HttpEndpoint{
		Name:    str,
		Path:    str2,
		Port:    nil,
		ExtPath: nil,
	})
	if ep, err := GenHttpEndpoints(mfHEs); err != nil {
		t.Error("err != nil")
	} else if len(ep) != 1 {
		t.Errorf("len(%v) != 1", ep)
	}
	mfHEs = mfHEs[:len(mfHEs)-1]
	// --------------------------------
	str3 := "test3"
	i := 80
	mfHEs = append(mfHEs, model.HttpEndpoint{
		Name:    str,
		Path:    str2,
		Port:    &i,
		ExtPath: &str3,
	})
	a = module.HttpEndpoint{
		Name: str,
		Port: &i,
		Path: str2,
	}
	if ep, err := GenHttpEndpoints(mfHEs); err != nil {
		t.Error("err != nil")
	} else if len(ep) != 2 {
		t.Errorf("len(%v) != 2", ep)
	} else if b, ok := ep[str3]; !ok {
		t.Errorf("b, ok := ep[%s]; !ok", str3)
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	mfHEs = mfHEs[:len(mfHEs)-1]
	// --------------------------------
	mfHEs = append(mfHEs, model.HttpEndpoint{
		Name:    str,
		Path:    str2,
		Port:    &i,
		ExtPath: nil,
	})
	if _, err := GenHttpEndpoints(mfHEs); err == nil {
		t.Error("err == nil")
	}
}

func TestGenPorts(t *testing.T) {
	var mfSPs []model.SrvPort
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 0 {
		t.Errorf("len(%v) != 0", p)
	}
	// --------------------------------
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80",
			HostPort: nil,
			Protocol: nil,
		},
	}
	a := module.Port{
		Name:     nil,
		Number:   80,
		Protocol: module.TcpPort,
		Bindings: nil,
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 1 {
		t.Errorf("len(%v) != 1", p)
	} else if reflect.DeepEqual(a, p[0]) == false {
		t.Errorf("%v != %v", a, p[0])
	}
	// --------------------------------
	str := "test"
	mfSPs = []model.SrvPort{
		{
			Name:     &str,
			Port:     "80",
			HostPort: nil,
			Protocol: &str,
		},
	}
	a = module.Port{
		Name:     &str,
		Number:   80,
		Protocol: str,
		Bindings: nil,
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 1 {
		t.Errorf("len(%v) != 1", p)
	} else if reflect.DeepEqual(a, p[0]) == false {
		t.Errorf("%v != %v", a, p[0])
	}
	// --------------------------------
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80-81",
			HostPort: nil,
			Protocol: nil,
		},
	}
	a2 := []module.Port{
		{
			Name:     nil,
			Number:   80,
			Protocol: module.TcpPort,
			Bindings: nil,
		},
		{
			Name:     nil,
			Number:   81,
			Protocol: module.TcpPort,
			Bindings: nil,
		},
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 2 {
		t.Errorf("len(%v) != 2", p)
	} else if reflect.DeepEqual(a2, p) == false {
		t.Errorf("%v != %v", a2, p)
	}
	// --------------------------------
	mfSPs = []model.SrvPort{
		{
			Name:     &str,
			Port:     "80-81",
			HostPort: nil,
			Protocol: &str,
		},
	}
	a2 = []module.Port{
		{
			Name:     &str,
			Number:   80,
			Protocol: str,
			Bindings: nil,
		},
		{
			Name:     &str,
			Number:   81,
			Protocol: str,
			Bindings: nil,
		},
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 2 {
		t.Errorf("len(%v) != 2", p)
	} else if reflect.DeepEqual(a2, p) == false {
		t.Errorf("%v != %v", a2, p)
	}
	// --------------------------------
	hp := model.Port("80")
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80",
			HostPort: &hp,
			Protocol: nil,
		},
	}
	a = module.Port{
		Name:     nil,
		Number:   80,
		Protocol: module.TcpPort,
		Bindings: []uint{80},
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 1 {
		t.Errorf("len(%v) != 1", p)
	} else if reflect.DeepEqual(a, p[0]) == false {
		t.Errorf("%v != %v", a, p[0])
	}
	// --------------------------------
	hp2 := model.Port("80-81")
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80",
			HostPort: &hp2,
			Protocol: nil,
		},
	}
	a = module.Port{
		Name:     nil,
		Number:   80,
		Protocol: module.TcpPort,
		Bindings: []uint{80, 81},
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 1 {
		t.Errorf("len(%v) != 1", p)
	} else if reflect.DeepEqual(a, p[0]) == false {
		t.Errorf("%v != %v", a, p[0])
	}
	// --------------------------------
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80-81",
			HostPort: &hp2,
			Protocol: nil,
		},
	}
	a2 = []module.Port{
		{
			Name:     nil,
			Number:   80,
			Protocol: module.TcpPort,
			Bindings: []uint{80},
		},
		{
			Name:     nil,
			Number:   81,
			Protocol: module.TcpPort,
			Bindings: []uint{81},
		},
	}
	if p, err := GenPorts(mfSPs); err != nil {
		t.Error("err != nil")
	} else if len(p) != 2 {
		t.Errorf("len(%v) != 2", p)
	} else if reflect.DeepEqual(a2, p) == false {
		t.Errorf("%v != %v", a2, p)
	}
	// --------------------------------
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "test",
			HostPort: nil,
			Protocol: nil,
		},
	}
	if _, err := GenPorts(mfSPs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	hp3 := model.Port("test")
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80",
			HostPort: &hp3,
			Protocol: nil,
		},
	}
	if _, err := GenPorts(mfSPs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	hp4 := model.Port("80")
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80-81",
			HostPort: &hp4,
			Protocol: nil,
		},
	}
	if _, err := GenPorts(mfSPs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	hp5 := model.Port("80-82")
	mfSPs = []model.SrvPort{
		{
			Name:     nil,
			Port:     "80-81",
			HostPort: &hp5,
			Protocol: nil,
		},
	}
	if _, err := GenPorts(mfSPs); err == nil {
		t.Error("err == nil")
	}
}
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

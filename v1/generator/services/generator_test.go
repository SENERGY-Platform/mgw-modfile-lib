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
	// --------------------------------
	mfBMs = append(mfBMs, model.BindMount{
		MountPoint: str,
		Source:     "test3",
		ReadOnly:   true,
	})
	if _, err := GenBindMounts(mfBMs); err == nil {
		t.Error("err == nil")
	}
	// --------------------------------
	mfBMs = mfBMs[:len(mfBMs)-1]
	mfBMs = append(mfBMs, model.BindMount{
		MountPoint: str,
		Source:     str2,
		ReadOnly:   false,
	})
	if _, err := GenBindMounts(mfBMs); err == nil {
		t.Error("err == nil")
	}
}

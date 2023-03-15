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

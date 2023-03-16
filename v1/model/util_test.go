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

package model

import (
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

func TestByteFmt_UnmarshalYAML(t *testing.T) {
	a := ByteFmt(67108864)
	var b ByteFmt
	if err := yaml.Unmarshal([]byte("64Mb"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%d != %d", a, b)
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("67108864"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%d != %d", a, b)
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("test"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("1.1"), &b); err == nil {
		t.Error("err == nil")
	}
}

func TestDuration_UnmarshalYAML(t *testing.T) {
	a := Duration(time.Second)
	var b Duration
	if err := yaml.Unmarshal([]byte("1s"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%d != %d", a, b)
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("1"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("test"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("test:"), &b); err == nil {
		t.Error("err == nil")
	}
}

func TestFileMode_UnmarshalYAML(t *testing.T) {
	a := FileMode(504)
	var b FileMode
	if err := yaml.Unmarshal([]byte("770"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%d != %d", a, b)
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("0770"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%d != %d", a, b)
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("test"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("test:"), &b); err == nil {
		t.Error("err == nil")
	}
}

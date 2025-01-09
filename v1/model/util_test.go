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
	"reflect"
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

func TestPort_UnmarshalYAML(t *testing.T) {
	a := Port("80")
	var b Port
	if err := yaml.Unmarshal([]byte("80"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%s != %s", a, b)
	}
	// ---------------------------
	a = "80-81"
	if err := yaml.Unmarshal([]byte("80-81"), &b); err != nil {
		t.Error("err != nil")
	} else if a != b {
		t.Errorf("%s != %s", a, b)
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("80-81-"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("-1"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("1.1"), &b); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if err := yaml.Unmarshal([]byte("test"), &b); err == nil {
		t.Error("err == nil")
	}
}

func TestPort_Parse(t *testing.T) {
	a := []uint{80}
	if b, err := Port("80").Parse(); err != nil {
		t.Error("err != nil")
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ---------------------------
	a = []uint{80, 81}
	if b, err := Port("80-81").Parse(); err != nil {
		t.Error("err != nil")
	} else if reflect.DeepEqual(a, b) == false {
		t.Errorf("%v != %v", a, b)
	}
	// ---------------------------
	if _, err := Port("81-80").Parse(); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if _, err := Port("80-81-").Parse(); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if _, err := Port("80-test").Parse(); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if _, err := Port("test").Parse(); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if _, err := Port("-").Parse(); err == nil {
		t.Error("err == nil")
	}
	// ---------------------------
	if _, err := Port("").Parse(); err == nil {
		t.Error("err == nil")
	}
}

func TestConfigValue_GetUserInput(t *testing.T) {
	str := "test"
	cv := ConfigValue{UserInput: &ConfigUserInput{UserInput: UserInput{Name: str}}}
	if ui := cv.GetUserInput(); ui.Name != str {
		t.Error("wrong user input")
	}
}

func TestResourceBase_GetUserInput(t *testing.T) {
	str := "test"
	rb := Resource{UserInput: &UserInput{Name: str}}
	if ui := rb.GetUserInput(); ui.Name != str {
		t.Error("wrong user input")
	}
}

func TestStrOrSlice_UnmarshalYAML(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		a := StrOrSlice{"test"}
		var b StrOrSlice
		if err := yaml.Unmarshal([]byte("test"), &b); err != nil {
			t.Error("err != nil")
		} else if !reflect.DeepEqual(a, b) {
			t.Errorf("%v != %v", a, b)
		}
	})
	t.Run("slice", func(t *testing.T) {
		a := StrOrSlice{"test", "test"}
		var b StrOrSlice
		if err := yaml.Unmarshal([]byte("[test, test]"), &b); err != nil {
			t.Error("err != nil")
		} else if !reflect.DeepEqual(a, b) {
			t.Errorf("%v != %v", a, b)
		}
	})
}

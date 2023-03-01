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

package modfile

import (
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/mgw-modfile-lib/itf"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"gopkg.in/yaml.v3"
	"reflect"
	"testing"
)

type testFile struct {
	Field string `yaml:"field"`
}

func (f *testFile) Parse() (module.Module, error) {
	return module.Module{Name: f.Field}, nil
}

func testDecode(yn *yaml.Node) (itf.ModFile, error) {
	var f testFile
	if err := yn.Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}

func testErrDecode(_ *yaml.Node) (itf.ModFile, error) {
	return nil, errors.New("test")
}

func genTestYml(ver string, val string) []byte {
	return []byte(fmt.Sprintf("modfileVersion: %s\nfield: %s", ver, val))
}

func TestModFile_UnmarshalYAML(t *testing.T) {
	ver := "vTest"
	val := "test"
	d := make(Decoders)
	d[ver] = testDecode
	mf := NewModFile(d)
	// ------------------
	if err := yaml.Unmarshal(genTestYml(ver, val), &mf); err != nil {
		t.Error("err != nil")
	}
	if mf.Version != ver {
		t.Errorf("\"%s\" != \"%s\"", mf.Version, ver)
	}
	if reflect.TypeOf(mf.modFile).Elem() != reflect.TypeOf(testFile{}) {
		t.Errorf("%s != %s", reflect.TypeOf(mf.modFile).Elem(), reflect.TypeOf(testFile{}))
	}
	if reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String() != val {
		t.Errorf("\"%s\" != \"%s\"", reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String(), val)
	}
	// ------------------
	val2 := "test2"
	if err := yaml.Unmarshal(genTestYml(ver, val2), &mf); err != nil {
		t.Error("err != nil")
	}
	if reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String() != val2 {
		t.Errorf("\"%s\" != \"%s\"", reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String(), val2)
	}
	// ------------------
	mf = NewModFile(d)
	if err := yaml.Unmarshal(genTestYml("vErr", val), &mf); err == nil {
		t.Errorf("yaml.Unmarshal(genTestYml(\"vErr\", \"%s\"), &mf); err == nil", val)
	} else {
		if mf.modFile != nil {
			t.Error("mf.modFile != nil")
		}
	}
	// ------------------
	mf = NewModFile(d)
	testYml := []byte("test: test")
	if err := yaml.Unmarshal(testYml, &mf); err == nil {
		t.Errorf("yaml.Unmarshal(\"%s\", &mf); err == nil", string(testYml))
	} else {
		if mf.modFile != nil {
			t.Error("mf.modFile != nil")
		}
	}
	// ------------------
	mf = NewModFile(d)
	testYml2 := []byte("1")
	if err := yaml.Unmarshal(testYml2, &mf); err == nil {
		t.Errorf("yaml.Unmarshal(\"%s\", &mf); err == nil", string(testYml2))
	} else {
		if mf.modFile != nil {
			t.Error("mf.modFile != nil")
		}
	}
	// ------------------
	ver2 := "vErr"
	d[ver2] = testErrDecode
	mf = NewModFile(d)
	if err := yaml.Unmarshal(genTestYml(ver2, val), &mf); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Error("mf.modFile != nil")
		}
	}
	// ------------------
	var mf2 ModFile
	if err := yaml.Unmarshal(genTestYml(ver, val), &mf2); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Error("mf.modFile != nil")
		}
	}
}

func TestModFile_GetModule(t *testing.T) {
	val := "test"
	mf := NewModFile(nil)
	mf.modFile = &testFile{Field: val}
	m, err := mf.GetModule()
	if err != nil {
		t.Error("err != nil")
	}
	if m.Name != val {
		t.Errorf("\"%s\" != \"%s\"", m.Name, val)
	}
}

func TestDecoders_Add(t *testing.T) {
	d := make(Decoders)
	ver := "vTest"
	d.Add(func() (string, func(*yaml.Node) (itf.ModFile, error)) {
		return ver, testErrDecode
	})
	if dc, ok := d[ver]; !ok {
		t.Errorf("dc, ok := d[\"%s\"]; !ok", ver)
	} else {
		_, err := dc(nil)
		if err == nil {
			t.Error("wrong decoder")
		}
	}
}

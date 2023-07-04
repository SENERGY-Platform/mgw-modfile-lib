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
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"gopkg.in/yaml.v3"
	"reflect"
	"testing"
)

type testFile struct {
	Field string `yaml:"field"`
}

func testGen(mf any) (*module.Module, error) {
	tf, ok := mf.(*testFile)
	if !ok {
		return nil, errors.New("test")
	}
	return &module.Module{Name: tf.Field}, nil
}

func testDecode(yn *yaml.Node) (any, error) {
	var f testFile
	if err := yn.Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}

func testErrDecode(_ *yaml.Node) (any, error) {
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
	mf := New(d, nil)
	// ------------------
	if err := yaml.Unmarshal(genTestYml(ver, val), &mf); err != nil {
		fmt.Println(err)
		t.Error("err != nil")
	}
	if mf.version != ver {
		t.Errorf("%s != %s", mf.version, ver)
	}
	if reflect.TypeOf(mf.modFile).Elem() != reflect.TypeOf(testFile{}) {
		t.Errorf("%s != %s", reflect.TypeOf(mf.modFile).Elem(), reflect.TypeOf(testFile{}))
	}
	if reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String() != val {
		t.Errorf("%s != %s", reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String(), val)
	}
	// ------------------
	val2 := "test2"
	if err := yaml.Unmarshal(genTestYml(ver, val2), &mf); err != nil {
		t.Error("err != nil")
	}
	if reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String() != val2 {
		t.Errorf("%s != %s", reflect.ValueOf(mf.modFile).Elem().FieldByName("Field").String(), val2)
	}
	// ------------------
	mf = New(d, nil)
	if err := yaml.Unmarshal(genTestYml("vErr", val), &mf); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Errorf("%v != nil", mf.modFile)
		}
	}
	// ------------------
	mf = New(d, nil)
	testYml := []byte("test: test")
	if err := yaml.Unmarshal(testYml, &mf); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Errorf("%v != nil", mf.modFile)
		}
	}
	// ------------------
	mf = New(d, nil)
	testYml2 := []byte("1")
	if err := yaml.Unmarshal(testYml2, &mf); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Errorf("%v != nil", mf.modFile)
		}
	}
	// ------------------
	ver2 := "vErr"
	d[ver2] = testErrDecode
	mf = New(d, nil)
	if err := yaml.Unmarshal(genTestYml(ver2, val), &mf); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Errorf("%v != nil", mf.modFile)
		}
	}
	// ------------------
	var mf2 MfWrapper
	if err := yaml.Unmarshal(genTestYml(ver, val), &mf2); err == nil {
		t.Error("err == nil")
	} else {
		if mf.modFile != nil {
			t.Errorf("%v != nil", mf.modFile)
		}
	}
}

func TestModFile_GetModule(t *testing.T) {
	ver := "vTest"
	val := "test"
	g := make(Generators)
	g[ver] = testGen
	mf := New(nil, g)
	mf.version = ver
	mf.modFile = &testFile{Field: val}
	m, err := mf.GetModule()
	if err != nil {
		t.Error("err != nil")
	}
	if m.Name != val {
		t.Errorf("%s != %s", m.Name, val)
	}
	mf = New(nil, g)
	mf.modFile = &testFile{Field: val}
	_, err = mf.GetModule()
	if err == nil {
		t.Error("err == nil")
	}
	mf = New(nil, g)
	mf.version = ver
	_, err = mf.GetModule()
	if err == nil {
		t.Error("err == nil")
	}
	mf = New(nil, nil)
	_, err = mf.GetModule()
	if err == nil {
		t.Error("err == nil")
	}
}

func TestDecoders_Add(t *testing.T) {
	d := make(Decoders)
	ver := "vTest"
	d.Add(func() (string, func(*yaml.Node) (any, error)) {
		return ver, testErrDecode
	})
	if dc, ok := d[ver]; !ok {
		t.Errorf("dc, ok := d[%s]; !ok", ver)
	} else {
		_, err := dc(nil)
		if err == nil {
			t.Error("wrong decoder")
		}
	}
}

func TestGenerators_Add(t *testing.T) {
	g := make(Generators)
	ver := "vTest"
	g.Add(func() (string, func(any) (*module.Module, error)) {
		return ver, testGen
	})
	if gn, ok := g[ver]; !ok {
		t.Errorf("gn, ok := g[%s]; !ok", ver)
	} else {
		_, err := gn(nil)
		if err == nil {
			t.Error("wrong generator")
		}
	}
}

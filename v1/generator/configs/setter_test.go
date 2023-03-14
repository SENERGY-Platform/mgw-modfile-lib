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

package configs

import (
	"errors"
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/SENERGY-Platform/mgw-module-lib/module"
	"reflect"
	"testing"
)

func TestParseConfigOptions(t *testing.T) {
	var opt []any
	o, err := parseConfigOptions(opt, func(a any) (int, error) {
		return 0, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(o) != 0 {
		t.Errorf("len(%v) != 0", o)
	}
	opt = append(opt, 1)
	o, err = parseConfigOptions(opt, func(a any) (int, error) {
		return a.(int) + 1, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(o) != 1 {
		t.Errorf("len(%v) != 1", o)
	} else if o[0] != 2 {
		t.Errorf("%d != 2", o[0])
	}
	o, err = parseConfigOptions(opt, func(a any) (int, error) {
		return 0, errors.New("test")
	})
	if err == nil {
		t.Error("err == nil")
	}
}

func TestParseConfigValueString(t *testing.T) {
	str := "test"
	if v, err := parseConfigValueString(str); err != nil {
		t.Errorf("v, err := parseConfigValueString(\"%s\"); err != nil", str)
	} else if v != str {
		t.Errorf("\"%s\" != \"%s\"", v, str)
	}
	if _, err := parseConfigValueString(1); err == nil {
		t.Error("v, err := parseConfigValueString(1); err == nil")
	}
}

func TestParseConfigValueBool(t *testing.T) {
	b := true
	if v, err := parseConfigValueBool(b); err != nil {
		t.Errorf("v, err := parseConfigValueString(%v); err != nil", b)
	} else if v != b {
		t.Errorf("%v != %v", v, b)
	}
	if _, err := parseConfigValueBool(1); err == nil {
		t.Error("v, err := parseConfigValueBool(1); err == nil")
	}
}

func TestParseConfigValueInt64(t *testing.T) {
	i := int64(1)
	if v, err := parseConfigValueInt64(int(i)); err != nil {
		t.Errorf("v, err := parseConfigValueInt64(int(%d)); err != nil", i)
	} else if v != i {
		t.Errorf("%d != %d", v, i)
	}
	if v, err := parseConfigValueInt64(int8(i)); err != nil {
		t.Errorf("v, err := parseConfigValueInt64(int8(%d)); err != nil", i)
	} else if v != i {
		t.Errorf("%d != %d", v, i)
	}
	if v, err := parseConfigValueInt64(int16(i)); err != nil {
		t.Errorf("v, err := parseConfigValueInt64(int16(%d)); err != nil", i)
	} else if v != i {
		t.Errorf("%d != %d", v, i)
	}
	if v, err := parseConfigValueInt64(int32(i)); err != nil {
		t.Errorf("v, err := parseConfigValueInt64(int32(%d)); err != nil", i)
	} else if v != i {
		t.Errorf("%d != %d", v, i)
	}
	if v, err := parseConfigValueInt64(i); err != nil {
		t.Errorf("v, err := parseConfigValueInt64(%d); err != nil", i)
	} else if v != i {
		t.Errorf("%d != %d", v, i)
	}
	if _, err := parseConfigValueInt64(""); err == nil {
		t.Error("v, err := parseConfigValueInt64(\"\"); err == nil")
	}
}

func TestParseConfigValueFloat64(t *testing.T) {
	f := 1.0
	if v, err := parseConfigValueFloat64(float32(f)); err != nil {
		t.Errorf("v, err := parseConfigValueFloat64(float32(%f)); err != nil", f)
	} else if v != f {
		t.Errorf("%f != %f", v, f)
	}
	if v, err := parseConfigValueFloat64(f); err != nil {
		t.Errorf("v, err := parseConfigValueFloat64(%f); err != nil", f)
	} else if v != f {
		t.Errorf("%f != %f", v, f)
	}
	if _, err := parseConfigValueFloat64(""); err == nil {
		t.Error("v, err := parseConfigValueFloat64(\"\"); err == nil")
	}
}

func TestParseConfigTypeOptions(t *testing.T) {
	var opt map[string]any
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if len(cto) != 0 {
		t.Errorf("len(%v) != 0", cto)
	}
	// ---------------------------------------
	opt = make(map[string]any)
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if len(cto) != 0 {
		t.Errorf("len(%v) != 0", cto)
	}
	// ---------------------------------------
	str := "test"
	opt[str] = str
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.StringType {
		t.Error("o.DataType != module.StringType")
	} else if reflect.DeepEqual(o.Value, str) == false {
		t.Errorf("reflect.DeepEqual(%v, \"%s\") == false", o.Value, str)
	}
	// ---------------------------------------
	i := int64(1)
	opt[str] = int(i)
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Int64Type {
		t.Error("o.DataType != module.Int64Type")
	} else if reflect.DeepEqual(o.Value, i) == false {
		t.Errorf("reflect.DeepEqual(%v, %d) == false", o.Value, i)
	}
	// ---------------------------------------
	opt[str] = int8(i)
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Int64Type {
		t.Error("o.DataType != module.Int64Type")
	} else if reflect.DeepEqual(o.Value, i) == false {
		t.Errorf("reflect.DeepEqual(%v, %d) == false", o.Value, i)
	}
	// ---------------------------------------
	opt[str] = int16(i)
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Int64Type {
		t.Error("o.DataType != module.Int64Type")
	} else if reflect.DeepEqual(o.Value, i) == false {
		t.Errorf("reflect.DeepEqual(%v, %d) == false", o.Value, i)
	}
	// ---------------------------------------
	opt[str] = int32(i)
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Int64Type {
		t.Error("o.DataType != module.Int64Type")
	} else if reflect.DeepEqual(o.Value, i) == false {
		t.Errorf("reflect.DeepEqual(%v, %d) == false", o.Value, i)
	}
	// ---------------------------------------
	opt[str] = i
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Int64Type {
		t.Error("o.DataType != module.Int64Type")
	} else if reflect.DeepEqual(o.Value, i) == false {
		t.Errorf("reflect.DeepEqual(%v, %d) == false", o.Value, i)
	}
	// ---------------------------------------
	f := 1.0
	opt[str] = float32(f)
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Float64Type {
		t.Error("o.DataType != module.Float64Type")
	} else if reflect.DeepEqual(o.Value, f) == false {
		t.Errorf("reflect.DeepEqual(%v, %f) == false", o.Value, f)
	}
	// ---------------------------------------
	opt[str] = f
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.Float64Type {
		t.Error("o.DataType != module.Float64Type")
	} else if reflect.DeepEqual(o.Value, f) == false {
		t.Errorf("reflect.DeepEqual(%v, %f) == false", o.Value, f)
	}
	// ---------------------------------------
	b := true
	opt[str] = b
	if cto, err := parseConfigTypeOptions(opt); err != nil {
		t.Errorf("cto, err := parseConfigTypeOptions(%v); err != nil", opt)
	} else if o, ok := cto[str]; !ok {
		t.Errorf("o, ok := cto[\"%s\"]; !ok", str)
	} else if o.DataType != module.BoolType {
		t.Error("o.DataType != module.BoolType")
	} else if reflect.DeepEqual(o.Value, b) == false {
		t.Errorf("reflect.DeepEqual(%v, %v) == false", o.Value, b)
	}
	// ---------------------------------------
	opt[str] = uint(1)
	if _, err := parseConfigTypeOptions(opt); err == nil {
		t.Errorf("_, err := parseConfigTypeOptions(%v); err == nil", opt)
	}
}

func TestParseConfig(t *testing.T) {
	var opt []any
	var ctOpt map[string]any
	p, o, to, err := parseConfig(nil, opt, ctOpt, func(a any) (int, error) {
		return 0, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if p != nil {
		t.Error("p != nil")
	} else if len(o) != 0 {
		t.Errorf("len(%v) != 0", o)
	} else if len(to) != 0 {
		t.Errorf("len(%v) != 0", to)
	}
	// ---------------------------------------
	p, o, to, err = parseConfig(1, opt, ctOpt, func(a any) (int, error) {
		return a.(int) + 1, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if p == nil {
		t.Error("p != nil")
	} else if *p != 2 {
		t.Errorf("%d != 2", *p)
	}
	// ---------------------------------------
	_, _, _, err = parseConfig(1, opt, ctOpt, func(a any) (int, error) {
		return 0, errors.New("test")
	})
	if err == nil {
		t.Error("err == nil")
	}
	// ---------------------------------------
	ctOpt = make(map[string]any)
	ctOpt["test"] = uint(1)
	_, _, _, err = parseConfig(nil, opt, ctOpt, func(a any) (int, error) {
		return 0, nil
	})
	if err == nil {
		t.Error("err == nil")
	}
	ctOpt = nil
	// ---------------------------------------
	opt = append(opt, 1)
	p, o, to, err = parseConfig(nil, opt, ctOpt, func(a any) (int, error) {
		return a.(int) + 1, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(o) != 1 {
		t.Errorf("len(%v) != 1", o)
	} else if o[0] != 2 {
		t.Errorf("%d != 2", o[0])
	}
	// ---------------------------------------
	_, _, _, err = parseConfig(nil, opt, ctOpt, func(a any) (int, error) {
		return 0, errors.New("test")
	})
	if err == nil {
		t.Error("err == nil")
	}
}

func TestParseConfigSlice(t *testing.T) {
	var opt []any
	var ctOpt map[string]any
	sl, o, to, err := parseConfigSlice(nil, opt, ctOpt, func(a any) (int, error) {
		return 0, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(sl) != 0 {
		t.Errorf("len(%v) != 0", sl)
	} else if len(o) != 0 {
		t.Errorf("len(%v) != 0", o)
	} else if len(to) != 0 {
		t.Errorf("len(%v) != 0", to)
	}
	// ---------------------------------------
	_, _, _, err = parseConfigSlice("", opt, ctOpt, func(a any) (int, error) {
		return 0, nil
	})
	if err == nil {
		t.Error("err == nil")
	}
	// ---------------------------------------
	var val []any
	sl, o, to, err = parseConfigSlice(val, opt, ctOpt, func(a any) (int, error) {
		return 0, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(sl) != 0 {
		t.Errorf("len(%v) != 0", sl)
	}
	// ---------------------------------------
	ctOpt = make(map[string]any)
	ctOpt["test"] = uint(1)
	_, _, _, err = parseConfigSlice(nil, opt, ctOpt, func(a any) (int, error) {
		return 0, nil
	})
	if err == nil {
		t.Error("err == nil")
	}
	ctOpt = nil
	// ---------------------------------------
	opt = append(opt, 1)
	sl, o, to, err = parseConfigSlice(nil, opt, ctOpt, func(a any) (int, error) {
		return a.(int) + 1, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(o) != 1 {
		t.Errorf("len(%v) != 1", o)
	} else if o[0] != 2 {
		t.Errorf("%d != 2", o[0])
	}
	// ---------------------------------------
	_, _, _, err = parseConfigSlice(nil, opt, ctOpt, func(a any) (int, error) {
		return 0, errors.New("test")
	})
	if err == nil {
		t.Error("err == nil")
	}
	// ---------------------------------------
	opt = []any{}
	val = append(val, 1)
	sl, o, to, err = parseConfigSlice(val, opt, ctOpt, func(a any) (int, error) {
		return a.(int) + 1, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(sl) != 1 {
		t.Errorf("len(%v) != 1", sl)
	} else if sl[0] != 2 {
		t.Errorf("%d != 2", sl[0])
	}
	// ---------------------------------------
	sl, o, to, err = parseConfigSlice(val, opt, ctOpt, func(a any) (int, error) {
		return 0, errors.New("test")
	})
	if err == nil {
		t.Error("err == nil")
	}
}

func TestSetValue(t *testing.T) {
	mCs := make(module.Configs)
	if err := SetValue("", model.ConfigValue{}, mCs); err == nil {
		t.Error("err == nil")
	} else if len(mCs) != 0 {
		t.Errorf("len(%v) != 0", mCs)
	}
}

func TestSetValueStr(t *testing.T) {
	str := "test"
	testSetValue[string](t, str, []any{str}, module.StringType, 1)
}

func TestSetValueBool(t *testing.T) {
	b := true
	testSetValue[bool](t, b, []any{b}, module.BoolType, "")
}

func TestSetValueInt64(t *testing.T) {
	i := int64(1)
	testSetValue[int64](t, i, []any{i}, module.Int64Type, "")
}

func TestSetValueFloat64(t *testing.T) {
	f := 1.0
	testSetValue[float64](t, f, []any{f}, module.Float64Type, "")
}

func testSetValue[T comparable](t *testing.T, value any, options []any, dataType string, errVal any) {
	mCs := make(module.Configs)
	if err := SetValue("", model.ConfigValue{Value: errVal, DataType: dataType}, mCs); err == nil {
		t.Error("err == nil")
	} else if len(mCs) != 0 {
		t.Errorf("len(%v) != 0", mCs)
	}
	str := "test"
	cv := model.ConfigValue{
		Value:       value,
		Options:     options,
		OptionsExt:  true,
		Type:        str,
		TypeOptions: map[string]any{str: str},
		DataType:    dataType,
	}
	if err := SetValue(str, cv, mCs); err != nil {
		t.Error("err != nil")
	} else if c, ok := mCs[str]; !ok {
		t.Errorf("c, ok := mCs[\"%s\"]; !ok", str)
	} else if cv.DataType != c.DataType {
		t.Errorf("%v != %v", cv.DataType, c.DataType)
	} else if reflect.DeepEqual(cv.Value, c.Default) == false {
		t.Errorf("reflect.DeepEqual(%v, %v) == false", cv.Value, c.Default)
	} else if len(c.Options.([]T)) == 0 {
		t.Errorf("len(%v) == 0", c.Options)
	} else if reflect.DeepEqual(cv.Options[0], c.Options.([]T)[0]) == false {
		t.Errorf("reflect.DeepEqual(%v, %v) == false", cv.Options[0], c.Options.([]T)[0])
	} else if cv.OptionsExt != c.OptExt {
		t.Errorf("%v != %v", cv.OptionsExt, c.OptExt)
	} else if cv.Type != c.Type {
		t.Errorf("%v != %v", cv.Type, c.Type)
	} else if to, k := c.TypeOpt[str]; !k {
		t.Errorf("to, k := c.TypeOpt[\"%s\"]; !k", str)
	} else if reflect.DeepEqual(cv.TypeOptions[str], to.Value) == false {
		t.Errorf("reflect.DeepEqual(%v, %v) == false", cv.TypeOptions[str], to.Value)
	} else if cv.IsList != c.IsSlice {
		t.Errorf("%v != %v", cv.IsList, c.IsSlice)
	} else if c.Delimiter != "" {
		t.Errorf("%v != \"\"", c.Delimiter)
	}
}

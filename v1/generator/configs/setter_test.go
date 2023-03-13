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
	"testing"
)

func TestParseConfigOptions(t *testing.T) {
	o, err := parseConfigOptions([]any{1}, func(a any) (int, error) {
		return a.(int) + 1, nil
	})
	if err != nil {
		t.Error("err != nil")
	} else if len(o) != 1 {
		t.Errorf("len(%v) != 1", o)
	} else if o[0] != 2 {
		t.Errorf("%d != 2", o[0])
	}
	o, err = parseConfigOptions([]any{1}, func(a any) (int, error) {
		return 0, errors.New("test")
	})
	if err == nil {
		t.Error("err == nil")
	} else if len(o) != 0 {
		t.Errorf("len(%v) != 0", o)
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
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

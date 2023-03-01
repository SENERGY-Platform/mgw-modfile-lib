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

package v1

import (
	"gopkg.in/yaml.v3"
	"reflect"
	"testing"
)

func TestGetDecoder(t *testing.T) {
	v, d := GetDecoder()
	if v != ver {
		t.Errorf("\"%s\" != \"%s\"", v, ver)
	}
	f, err := d(&yaml.Node{})
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(f).Elem() != reflect.TypeOf(ModFile{}) {
		t.Errorf("%s != %s", reflect.TypeOf(f).Elem(), reflect.TypeOf(ModFile{}))
	}
}

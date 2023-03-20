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

package inputs

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"testing"
)

type test struct {
	uInput *model.UserInput
}

func (t test) GetUserInput() *model.UserInput {
	return t.uInput
}

func TestGenInputs(t *testing.T) {
	var mfCs map[string]test
	if im := GenInputs(mfCs); len(im) != 0 {
		t.Errorf("len(%v) != 0", im)
	}
	mfCs = make(map[string]test)
	mfCs["a"] = test{}
	if im := GenInputs(mfCs); len(im) != 0 {
		t.Errorf("len(%v) != 0", im)
	}
	str := "test"
	mfCs[str] = test{uInput: &model.UserInput{Name: str}}
	if im := GenInputs(mfCs); len(im) != 1 {
		t.Errorf("len(%v) != 1", im)
	} else if ui, ok := im[str]; !ok {
		t.Errorf("ui, ok := im[%s]; !ok", str)
	} else if ui.Name != str {
		t.Errorf("%s != %s", ui.Name, str)
	}
}

func TestGenInputGroups(t *testing.T) {
	var mfIGs map[string]model.InputGroup
	if igm := GenInputGroups(mfIGs); len(igm) != 0 {
		t.Errorf("len(%v) != 0", igm)
	}
	mfIGs = make(map[string]model.InputGroup)
	str := "test"
	mfIGs[str] = model.InputGroup{Name: str}
	if igm := GenInputGroups(mfIGs); len(igm) != 1 {
		t.Errorf("len(%v) != 1", igm)
	} else if ig, ok := igm[str]; !ok {
		t.Errorf("ig, ok := igm[%s]; !ok", str)
	} else if ig.Name != str {
		t.Errorf("%s != %s", ig.Name, str)
	}
}

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

package generic

import "testing"

func TestGenStringSet(t *testing.T) {
	var sl []string
	if set := GenStringSet(sl); len(set) != 0 {
		t.Errorf("len(%v) != 0", set)
	}
	str := "test"
	str2 := "test2"
	sl = append(sl, str, str2, str)
	if set := GenStringSet(sl); len(set) != 2 {
		t.Errorf("len(%v) != 2", set)
	} else if _, ok := set[str]; !ok {
		t.Errorf("_, ok := set[%s]; !ok", str)
	} else if _, ok := set[str2]; !ok {
		t.Errorf("_, ok := set[%s]; !ok", str2)
	}
}

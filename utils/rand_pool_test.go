/*
   Copyright 2020 LittleBear(1018589158@qq.com)

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package utils

import (
	"testing"
)

func TestMyRand_NextInt64(t *testing.T) {
	my := NewMyRand()
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: 0,
		},
		{
			name: "test1",
			want: 0,
		},
		{
			name: "test1",
			want: 0,
		},
		{
			name: "test1",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(my.NextInt64())
		})
	}
}

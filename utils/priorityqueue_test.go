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
	"reflect"
	"testing"
)

func TestPriorityQueue_Insert(t *testing.T) {

	//type fields struct {
	//	itemHeap *itemHeap
	//	lookup   map[interface{}]*item
	//}

	prq := NewPriorityQueue()
	type args struct {
		v        interface{}
		priority int64
	}
	tests := []struct {
		name   string
		fields PriorityQueue
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "test1",
			fields: prq,
			args: args{
				"1", 1,
			},
		},
		{
			name:   "test2",
			fields: prq,
			args: args{
				"2", 2,
			},
		},
		{
			name:   "test3",
			fields: prq,
			args: args{
				"3", 3,
			},
		},
		{
			name:   "test4",
			fields: prq,
			args: args{
				"5", 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Insert(tt.args.v, tt.args.priority)
			for _, v := range *tt.fields.itemHeap {
				println(v.value.(string))
			}
		})
	}
}

func TestPriorityQueue_Len(t *testing.T) {

	prq := NewPriorityQueue()
	type args struct {
		v        interface{}
		priority int64
	}
	tests := []struct {
		name   string
		fields PriorityQueue
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "test1",
			fields: prq,
			args: args{
				"1", 1,
			}, want: 1,
		},
		{
			name:   "test2",
			fields: prq,
			args: args{
				"2", 2,
			}, want: 2,
		},
		{
			name:   "test3",
			fields: prq,
			args: args{
				"3", 3,
			}, want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Insert(tt.args.v, tt.args.priority)
			if got := prq.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	prq := NewPriorityQueue()

	prq.Insert("1", 1)
	prq.Insert("2", 2)
	prq.Insert("3", 3)

	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			want:    "1",
			wantErr: false,
		},
		{
			name:    "test2",
			want:    "2",
			wantErr: false,
		},
		{
			name:    "test3",
			want:    "3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prq.Pop()
			println(got.(string))
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Remove(t *testing.T) {

	prq := NewPriorityQueue()
	prq.Insert("1", 1)
	prq.Insert("2", 2)
	prq.Insert("3", 3)

	type args struct {
		v interface{}
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				"2",
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "test1",
			args: args{
				"3",
			},
			want:    "3",
			wantErr: false,
		},
		{
			name: "test1",
			args: args{
				"1",
			},
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, it := range *prq.itemHeap {
				println(it.value.(string), it.priority, it.index)
			}

			got, err := prq.Remove(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		ih   itemHeap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ih.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		ih   itemHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ih.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		ih   itemHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ih.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		ih   itemHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_itemHeap_Remove(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		ih   itemHeap
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ih.Remove(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		ih   itemHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	prq := NewPriorityQueue()
	prq.Insert("2", 2)
	prq.Insert("3", 3)
	prq.Insert("1", 1)

	tests := []struct {
		name    string
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			want:    "1",
			wantErr: false,
		},
		{
			name:    "test1",
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prq.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

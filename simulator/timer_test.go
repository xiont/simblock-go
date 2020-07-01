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
package simulator

import (
	"reflect"
	"testing"
)

type TestTask struct {
	interval int64
}

func (tt *TestTask) Run() {
	println(tt, "run..", tt.interval)
}
func (tt *TestTask) GetInterval() int64 {
	return tt.interval
}

func NewTestTask(interval int64) *TestTask {
	return &TestTask{
		interval: interval,
	}
}

func TestPutTask(t *testing.T) {

	task1 := NewTestTask(1)
	task2 := NewTestTask(2)

	type args struct {
		task ITask
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				task: task1,
			},
		},
		{
			name: "test1",
			args: args{
				task: task2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PutTask(tt.args.task)
			b := GetTask()
			b.(*TestTask).Run()
			RemoveTask(tt.args.task)
		})
	}
}

func TestGetTask(t *testing.T) {
	task1 := NewTestTask(1)
	task2 := NewTestTask(2)
	task3 := NewTestTask(3)
	PutTask(task2)
	PutTask(task1)
	PutTask(task3)

	tests := []struct {
		name string
		want ITask
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: task1,
		},
		{
			name: "test2",
			want: task2,
		},
		{
			name: "test3",
			want: task3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTask(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() = %v, want %v", got, tt.want)
			}
			RemoveTask(tt.want)
		})
	}
}

func TestGetCurrentTime(t *testing.T) {
	task1 := NewTestTask(1)
	task2 := NewTestTask(2)
	task3 := NewTestTask(3)
	PutTask(task2)
	PutTask(task1)
	PutTask(task3)
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: 1,
		},
		{
			name: "test2",
			want: 2,
		},
		{
			name: "test3",
			want: 3,
		},
	}
	for _, tt := range tests {
		RunTask()
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentTime(); got != tt.want {
				t.Errorf("GetCurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentTime2(t *testing.T) {
	task1 := NewTestTask(1)
	task2 := NewTestTask(2)
	task3 := NewTestTask(3)
	task4 := NewTestTask(3)
	PutTask(task2)
	PutTask(task1)
	PutTask(task3)
	RunTask()
	PutTask(task4)
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			want: 2,
		},
		{
			name: "test2",
			want: 3,
		},
		{
			name: "test3",
			want: 4,
		},
	}
	for _, tt := range tests {
		RunTask()
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentTime(); got != tt.want {
				t.Errorf("GetCurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveTask(t *testing.T) {
	task1 := NewTestTask(1)
	task2 := NewTestTask(2)
	task3 := NewTestTask(3)
	PutTask(task2)
	PutTask(task1)
	PutTask(task3)

	type args struct {
		task ITask
	}
	tests := []struct {
		name string
		args args
		want ITask
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{task: task1},
			want: task2,
		},
		{
			name: "test2",
			args: args{task: task2},
			want: task3,
		},
		{
			name: "test3",
			args: args{task: task3},
			want: nil,
		},
	}
	for _, tt := range tests {
		RemoveTask(tt.args.task)
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTask(); got != tt.want {
				t.Errorf("GetCurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

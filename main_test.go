package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_makeRandomList(t *testing.T) {
	type args struct {
		distribution []float64
		facum        bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				distribution: []float64{
					0.2, 0.4, 0.3, 0.1,
				},
				facum: false,
			},
		},
		{
			name: "test1",
			args: args{
				distribution: []float64{
					0.2, 0.6, 0.9, 1.0,
				},
				facum: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRandomList(tt.args.distribution, tt.args.facum); !reflect.DeepEqual(got, tt.want) {
				t.Logf("makeRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}

type MyStruct struct {
	i int
	s string
}

func foo0() int {
	fmt.Println("running foo0: ")
	return 100
}

func foo1(a int) (string, string) {
	fmt.Println("running foo1: ", a)
	return "aaaa", "bbb"
}

func foo2(a, b int, c string) MyStruct {
	fmt.Println("running foo2: ", a, b, c)
	return MyStruct{10, "ccc"}
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func Test_reflectFunc(t *testing.T) {
	funcs := map[string]interface{}{
		"foo0": foo0,
		"foo1": foo1,
		"foo2": foo2,
	}

	// call foo0
	if result, err := Call(funcs, "foo0"); err == nil {
		for _, r := range result {
			fmt.Printf("  return: type=%v, value=[%d]\n", r.Type(), r.Int())
		}
	}

	// call foo1
	if result, err := Call(funcs, "foo1", 1); err == nil {
		for _, r := range result {
			fmt.Printf("  return: type=%v, value=[%s]\n", r.Type(), r.String())
		}
	}

	// call foo2
	if result, err := Call(funcs, "foo2", 1, 2, "aa"); err == nil {
		for _, r := range result {
			fmt.Printf("  return: type=%v, value=[%+v]\n", r.Type(), r.Interface().(MyStruct))
		}
	}
}

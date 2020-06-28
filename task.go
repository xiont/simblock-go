package main

type ITask interface {
	Run()
	GetInterval() int64
}

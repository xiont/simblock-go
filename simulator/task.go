package simulator

type ITask interface {
	Run()
	GetInterval() int64
}

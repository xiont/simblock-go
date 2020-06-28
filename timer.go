package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"simblock-go/utils"
	"strconv"
)

type IScheduledTask interface {
	getScheduledTime() int64
	compareTo(newST *ScheduledTask) int
	getTask() ITask
}

type ScheduledTask struct {
	task          ITask
	scheduledTime int64
}

func NewScheduledTask(task ITask, scheduledTime int64) *ScheduledTask {
	scheduledTask := ScheduledTask{
		task:          task,
		scheduledTime: scheduledTime,
	}
	return &scheduledTask
}

/**
 * Gets the task.
 *
 * @return the {@link Task} instance
 */
func (st ScheduledTask) getTask() ITask {
	return st.task
}

/**
 * Gets the scheduled time at which the task is to be executed.
 *
 * @return the scheduled time
 */
func (st ScheduledTask) getScheduledTime() int64 {
	return st.scheduledTime
}

/**
 * Compares the two scheduled tasks.
 *
 * @param o other task
 * @return 1 if self is executed later, 0 if concurrent and -1 if self is to be executed before.
 */
func (st ScheduledTask) compareTo(newST *ScheduledTask) int {
	if st == *newST {
		return 0
	}
	order := st.scheduledTime - newST.scheduledTime

	if order != 0 {
		if order > 0 {
			return 1
		} else {
			return -1
		}
	} else {
		stNum, _ := strconv.ParseInt(fmt.Sprintf("%p", &st), 16, 64)
		newSTNum, _ := strconv.ParseInt(fmt.Sprintf("%p", newST), 16, 64)
		if stNum > newSTNum {
			return 1
		} else {
			return 0
		}
	}
}

type timer struct {
	taskQueue   utils.PriorityQueue
	taskMap     *hashmap.Map
	currentTime int64
}

var Timer = timer{
	taskQueue:   utils.NewPriorityQueue(),
	taskMap:     hashmap.New(),
	currentTime: 0,
}

/**
 * Runs a {@link ScheduledTask}.
 */
func RunTask() {
	// If there are any tasks
	if Timer.taskQueue.Len() > 0 {
		// Get the next ScheduledTask
		// 或取并且删除权值最低的元素（Timer.ScheduleTask）
		icurrentScheduledTask, _ := Timer.taskQueue.Pop()
		currentScheduledTask := icurrentScheduledTask.(ScheduledTask)
		// 获取任务(Timer.ScheduleTask.Task)
		icurrentTask := currentScheduledTask.getTask()
		// 获取调度任务的调度时间，同时修改currentTime
		Timer.currentTime = currentScheduledTask.getScheduledTime()
		// Remove the task from the mapping of all tasks
		// 从hash表中删除该任务 <Timer.ScheduleTask.Task,Timer.ScheduleTask >
		Timer.taskMap.Remove(icurrentTask)
		// Execute
		// 获取到的Task 开始运行
		icurrentTask.Run()
	}
}

/**
 * Remove task from the mapping of all tasks and from the execution queue.
 *
 * @param task the task to be removed
 */
func RemoveTask(task ITask) {
	ScheduledTask, ok := Timer.taskMap.Get(task)
	if ok {
		Timer.taskQueue.Remove(ScheduledTask)
		Timer.taskMap.Remove(task)
	}
}

/**
 * Get the {@link Task} from the execution queue to be executed next.
 *
 * @return the task from the queue or null if task queue is empty.
 */
func GetTask() ITask {
	if Timer.taskQueue.Len() > 0 {
		// 返回权值最小的元素，但是不删除元素，时间复杂度是O(1)
		currentTask, _ := Timer.taskQueue.Peek()
		return currentTask.(ScheduledTask).getTask()
	} else {
		return nil
	}
}

/**
 * Schedule task to be executed at the current time incremented by the task duration.
 *
 * @param task the task
 */
func PutTask(task ITask) {
	scheduledTask := ScheduledTask{
		task:          task,
		scheduledTime: Timer.currentTime + task.GetInterval(),
	}
	Timer.taskMap.Put(task, scheduledTask)
	Timer.taskQueue.Insert(scheduledTask, scheduledTask.scheduledTime)
}

/**
 * Schedule task to be executed at the provided absolute timestamp.
 *
 * @param task the task
 * @param time the time in milliseconds
 */
func PutTaskAbsoluteTime(task ITask, time int64) {
	scheduledTask := NewScheduledTask(task, time)
	Timer.taskMap.Put(task, scheduledTask)
	Timer.taskQueue.Insert(scheduledTask, scheduledTask.scheduledTime)
}

/**
 * Get current time in milliseconds.
 *
 * @return the time
 */
func GetCurrentTime() int64 {
	return Timer.currentTime
}

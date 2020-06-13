package db

import (
	"Beq/dispurser/model"
	"container/list"
	"errors"
	"sync"
)

//JobQueue is hold current jobs
type JobQueue struct {
	List *list.List
}

var instance JobQueue

var totalJobCount int

var once sync.Once

//GetRequestQueue Initiating list database
func GetRequestQueue() *JobQueue {
	once.Do(func() {
		instance.List = list.New()
		totalJobCount = 0
	})
	return &instance
}

//AddJob  is used for add JOBs
func (obj *JobQueue) AddJob(Job model.Job) error {
	if instance.List != nil {
		instance.List.PushBack(Job)
		totalJobCount++
		return nil
	}
	return errors.New("No Data Base Initiate")

}

//GetJob  is used for get a JOB
// pre condition check jpbqueue should be non empty before proceeding this
// function
func (obj *JobQueue) GetJob(Job model.Job) (interface{}, error) {
	if instance.List != nil {
		if instance.List.Back() != nil {
			job := instance.List.Front().Value
			instance.List.Remove(instance.List.Front())
			return job, nil
		}
	}
	return nil, errors.New("No Data Base Initiate")
}

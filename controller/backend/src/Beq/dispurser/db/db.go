package db

import (
	"Beq/dispurser/model"
	"Beq/dispurser/utils"
	"container/list"
	"errors"
	"sync"
	"unsafe"
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

//Dispurse  is used for Dispurse JOBs
func (obj *JobQueue) Dispurse() error {
	if instance.List != nil {
		if instance.List.Back() != nil {
			currentJob := instance.List.Front().Value
			cJob, valid := currentJob.(model.Job)
			instance.List.Remove(instance.List.Front())
			if !valid {
				return errors.New("Invlide JOB")
			}
			utils.Dispurse(&cJob)
		}
		return nil
	}
	return errors.New("No Data Base Initiate")

}

//GetJob  is used for get a JOB
// pre condition check jpbqueue should be non empty before proceeding this
// function
func (obj *JobQueue) GetJob(Job model.Job) (*model.Job, error) {
	if instance.List != nil {
		if instance.List.Back() != nil {
			job := instance.List.Front().Value
			instance.List.Remove(instance.List.Front())
			return (*model.Job)(unsafe.Pointer(&job)), nil
		}
	}
	return nil, errors.New("No Data Base Initiate")
}

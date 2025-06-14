package task

import (
	"errors"
	"slices"
	"sync"
	"time"
)

type Repository interface {
	Save(task Task) (Task, error)
	GetById(id uint) (Task, error)
	GetUserTasks(userID uint) ([]Task, error)
	Update(task Task) (Task, error)
	Delete(id uint) error
	ChagneStatus(id uint) error
}

type inMemoryDatabase struct {
	tasks  []Task
	nextID uint
	mutex  *sync.Mutex
}

func NewRepository() Repository {
	return &inMemoryDatabase{
		tasks:  []Task{},
		nextID: 1,
		mutex:  &sync.Mutex{},
	}
}

func (db *inMemoryDatabase) Save(task Task) (Task, error) {
	task.ID = db.nextID
	task.UpdatedAt = time.Now()
	task.CreatedAt = time.Now()
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.tasks = append(db.tasks, task)
	db.nextID++
	return task, nil
}
func (db *inMemoryDatabase) GetById(id uint) (Task, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for _, task := range db.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, errors.New("not found")
}
func (db *inMemoryDatabase) GetUserTasks(userID uint) ([]Task, error) {
	var tasks []Task
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for _, task := range db.tasks {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func (db *inMemoryDatabase) Update(task Task) (Task, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for i := 0; i < len(db.tasks); i++ {
		if db.tasks[i].ID == task.ID {
			task.UpdatedAt = time.Now()
			db.tasks[i] = task
			return task, nil
		}
	}
	return Task{}, errors.New("task not found")
}

func (db *inMemoryDatabase) Delete(id uint) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for i := range db.tasks {
		if db.tasks[i].ID == id {
			db.tasks = slices.Delete(db.tasks, i, i+1)
			return nil
		}
	}
	return nil
}
func (db *inMemoryDatabase) ChagneStatus(id uint) error {
	return nil
}

package fake

import (
	"sync"

	m "github.com/devbv/go-server-boilerplate/internal/model"
)

type FakeTaskRepository struct {
	list  []*m.Task
	mutex *sync.Mutex
}

func NewFakeTaskRepository() *FakeTaskRepository {
	return &FakeTaskRepository{
		list:  make([]*m.Task, 0),
		mutex: &sync.Mutex{},
	}
}

func (repo *FakeTaskRepository) Add(task *m.Task) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.list = append(repo.list, task)

	return nil
}

func (repo *FakeTaskRepository) List() ([]*m.Task, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.list[:], nil
}

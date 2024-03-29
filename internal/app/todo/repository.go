package todo

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

// ErrTaskNotFound вказує на те, що задачу не знайдено.
var ErrTaskNotFound = errors.New("task not found")

// Repository визначає інтерфейс для сховища задач.
type Repository interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Update(task *Task) error
	Delete(id string) error
	Check(id string) error
	Uncheck(id string) error
}

// InMemoryRepository є in-memory реалізацією сховища задач.
type InMemoryRepository struct {
	tasks map[string]*Task
	mu    sync.RWMutex
}

// NewInMemoryRepository створює новий екземпляр InMemoryRepository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		tasks: make(map[string]*Task),
	}
}

func (r *InMemoryRepository) Create(task *Task) error {
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
	return nil
}

// GetAll повертає усі задачі з сховища.
func (r *InMemoryRepository) GetAll() ([]Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, *task)
	}
	return tasks, nil
}

// GetByID знаходить задачу за ідентифікатором.
func (r *InMemoryRepository) GetByID(id string) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (r *InMemoryRepository) Update(task *Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	updatedAt := time.Now()
	task.UpdatedAt = &updatedAt
	r.tasks[task.ID] = task
	return nil
}

// Delete видаляє задачу з сховища.
func (r *InMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}

// Check відмічає задачу як виконану.
func (r *InMemoryRepository) Check(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return ErrTaskNotFound
	}
	task.Completed = true
	updatedAt := time.Now()
	task.UpdatedAt = &updatedAt
	return nil
}

// Uncheck відмічає задачу як не виконану.
func (r *InMemoryRepository) Uncheck(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return ErrTaskNotFound
	}
	updatedAt := time.Now()
	task.UpdatedAt = &updatedAt
	task.Completed = false
	return nil
}

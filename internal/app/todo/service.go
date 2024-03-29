package todo

// TaskService надає методи для роботи з задачами.
type TaskService struct {
	repo Repository
}

// NewTaskService створює новий екземпляр TaskService.
func NewTaskService(repo Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// CreateTask створює нову задачу.
func (s *TaskService) CreateTask(task *Task) error {
	return s.repo.Create(task)
}

// GetAllTasks повертає усі задачі.
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAll()
}

// GetTaskByID знаходить задачу за ідентифікатором.
func (s *TaskService) GetTaskByID(id string) (*Task, error) {
	return s.repo.GetByID(id)
}

// UpdateTask оновлює задачу.
func (s *TaskService) UpdateTask(task *Task) error {
	return s.repo.Update(task)
}

// DeleteTask видаляє задачу.
func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}

// Додай check і uncheck методи

// CheckTask відмічає задачу як виконану.
func (s *TaskService) CheckTask(id string) error {
	return s.repo.Check(id)
}

// UncheckTask відмічаы задачу як не виконану.
func (s *TaskService) UncheckTask(id string) error {
	return s.repo.Uncheck(id)
}

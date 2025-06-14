package task

type Service interface {
	Save(createTaskRequest CreateTaskRequest, userID uint) (Task, error)
	GetById(id uint) (Task, error)
	GetUserTasks(userID uint) ([]Task, error)
	Delete(id uint) error
	ChagneStatus(updateTaskRequest UpdateTaskRequest, id uint) (Task, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Save(createTaskRequest CreateTaskRequest, userID uint) (Task, error) {
	var task = Task{
		Title:  createTaskRequest.Title,
		UserID: userID,
	}
	task, err := s.repo.Save(task)
	return task, err
}
func (s *service) GetById(id uint) (Task, error) {
	task, err := s.repo.GetById(id)
	return task, err
}
func (s *service) GetUserTasks(userID uint) ([]Task, error) {
	tasks, err := s.repo.GetUserTasks(userID)
	return tasks, err
}
func (s *service) Delete(id uint) error {
	return nil
}
func (s *service) ChagneStatus(updateTaskRquest UpdateTaskRequest, id uint) (Task, error) {
	task, err := s.repo.GetById(id)
	if err != nil {
		return Task{}, err
	}
	task.IsDone = updateTaskRquest.IsDone
	task, err = s.repo.Update(task)
	return task, err
}

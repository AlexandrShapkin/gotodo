package main

type TaskManager interface {
	Cteate(task *Task) (*Task, error)
	Read(pattern *Task) ([]*Task, error)
	Update(updated *Task) (*Task, error)
	Delete(pattern *Task) ([]*Task, error)
}

type JsonTaskManager struct {
	storage Storage
}

func NewTaskManager(storage Storage) TaskManager {
	return &JsonTaskManager{
		storage: storage,
	}
}

func (tm *JsonTaskManager) Cteate(task *Task) (*Task, error) {
	tasks, err := tm.storage.Read()
	if err != nil {
		return nil, err
	}

	id := 0
	for _, t := range tasks {
		if t.ID > id {
			id = t.ID
		}
	}
	id++

	task.ID = id
	tasks = append(tasks, task)
	err = tm.storage.Save(tasks)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (tm *JsonTaskManager) Read(pattern *Task) ([]*Task, error) {
	panic("not implemented") // TODO: Implement
}

func (tm *JsonTaskManager) Update(updated *Task) (*Task, error) {
	panic("not implemented") // TODO: Implement
}

func (tm *JsonTaskManager) Delete(pattern *Task) ([]*Task, error) {
	panic("not implemented") // TODO: Implement
}



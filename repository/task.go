package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	task := []entity.Task{}
	err := r.db.Where("user_id = ?", id).Find(&task).Error
	if err != nil {
		return []entity.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	errtask := r.db.Create(&task).Error
	if err != nil {
		return 0, errtask
	}
	return task.ID, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	err := r.db.Where("id = ?", id).Find(&task).Error
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	task := []entity.Task{}
	err := r.db.Where("category_id = ?", catId).Find(&task).Error
	if err != nil {
		return []entity.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	errtask := r.db.Model(&task).Where("id = ?", task.ID).Updates(task).Error
	if errtask != nil {
		return errtask
	}
	return nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	errtask := r.db.Where("id = ?", id).Delete(&entity.Task{}).Error
	if errtask != nil {
		return errtask
	}
	return nil
}

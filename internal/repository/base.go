package repository

import "gorm.io/gorm"

type Model interface {
	TableName() string
}

type Repository[T Model] interface {
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
	FindById(id uint) (*T, error)
	FindAll() ([]T, error)
	FindOne(condition *T) (*T, error)
}

type BaseRepository[T Model] struct {
	db *gorm.DB
}

func NewBaseRepository[T Model](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) FindOne(condition *T) (*T, error) {
	var entity T
	err := r.db.Where(condition).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"wallet_engine/internals/core/ports"
	"wallet_engine/pkg/utils"
)

type Repository[T ports.RequestDTO] struct {
	db *gorm.DB
}

func NewRepository[T ports.RequestDTO](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (r *Repository[T]) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var payload []T
	if err := r.db.Scopes(utils.Paginate(payload, pagination, r.db)).Find(&payload).Error; err != nil {
		fmt.Println(err.Error(), "repo")
		return nil, err
	}
	pagination.Rows = payload
	return pagination, nil
}

func (r *Repository[T]) GetByID(id string) (*T, error) {
	var payload T
	if err := r.db.Where("id = ?", id).First(&payload).Error; err != nil {
		return nil, err
	}
	return &payload, nil
}

func (r *Repository[T]) GetByIDForUpdate(id string) (*T, error) {
	var payload T
	if err := r.db.Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).Where("id = ?", id).First(&payload).Error; err != nil {
		return nil, err
	}
	return &payload, nil
}

func (r *Repository[T]) Persist(payload *T) error {
	if err := r.db.Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) Update(payload *T) error {
	if err := r.db.Save(payload).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) Delete(id string, entity interface{}) error {
	if err := r.db.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) DeleteAll(entity string) error {
	if err := r.db.Exec(fmt.Sprintf("DELETE FROM %v", entity)).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository[T]) WithTx(tx *gorm.DB) *Repository[T] {
	return NewRepository[T](tx)
}

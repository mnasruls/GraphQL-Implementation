package repositories

import "gorm.io/gorm"

type Repositories struct{
	modelDb *gorm.DB
}

func NewRepositories(db *gorm.DB)*Repositories  {
	return &Repositories{
		modelDb: db,
	}
}

func (db *Repositories)StorePosition()  {
	
}
package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Areas struct {
	gorm.Model
	ParentId    uint
	Name        string
	Description string
	Coord       string
}

type AreasManager struct {
	db *DB
}

func NewAreasManager(db *DB) (*AreasManager, error) {
	db.AutoMigrate(&Areas{})
	manager := AreasManager{}
	manager.db = db
	return &manager, nil
}

func (state *AreasManager) AllAreas() []Areas {
	var entities []Areas
	state.db.Find(&entities)
	return entities
}

func (state *AreasManager) AreaByID(id string) *Areas {
	entity := Areas{}
	state.db.First(&entity, id)
	return &entity
}

func (state *AreasManager) CreateArea(parentId uint, name string, description string, coord string) *Areas {
	entity := Areas{
		ParentId:    parentId,
		Name:        name,
		Description: description,
		Coord:       coord,
	}
	state.db.Create(&entity)
	return &entity
}

func (state *AreasManager) UpdateArea(id string, parentId uint, name string, description string, coord string) *Areas {
	entity := state.AreaByID(id)
	entity.ParentId = parentId
	entity.Name = name
	entity.Description = description
	entity.Coord = coord
	state.db.Save(&entity)
	return entity
}

func (state *AreasManager) DeleteArea(id string) *Areas {
	entity := Areas{}
	state.db.Delete(&entity, id)
	return &entity
}

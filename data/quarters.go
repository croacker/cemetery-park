package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Quarters struct {
	gorm.Model
	Name        string
	Description string
	Coord       string
}

type QuartersManager struct {
	db *DB
}

func NewQuartersManager(db *DB) (*QuartersManager, error) {
	db.AutoMigrate(&Quarters{})
	manager := QuartersManager{}
	manager.db = db
	return &manager, nil
}

func (state *QuartersManager) AllQuarters() []Quarters {
	var entities []Quarters
	state.db.Find(&entities)
	return entities
}

func (state *QuartersManager) QuarterByID(id string) *Quarters {
	entity := Quarters{}
	state.db.First(&entity, id)
	return &entity
}

func (state *QuartersManager) CreateQuarter(name string, description string, coord string) *Quarters {
	entity := Quarters{
		Name:        name,
		Description: description,
		Coord:       coord,
	}
	state.db.Create(&entity)
	return &entity
}

func (state *QuartersManager) UpdateQuarter(id string, name string, description string, coord string) *Quarters {
	entity := state.QuarterByID(id)
	entity.Name = name
	entity.Description = description
	entity.Coord = coord
	state.db.Save(&entity)
	return entity
}

func (state *QuartersManager) DeleteQuarter(id string) *Quarters {
	entity := Quarters{}
	state.db.Delete(&entity, id)
	return &entity
}

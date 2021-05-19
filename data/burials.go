package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Burials struct {
	gorm.Model
	ParentId    uint
	Name        string
	Description string
	Coord       string
}

type BurialsManager struct {
	db *DB
}

func NewBurialsManager(db *DB) (*BurialsManager, error) {
	db.AutoMigrate(&Burials{})
	manager := BurialsManager{}
	manager.db = db
	return &manager, nil
}

func (state *BurialsManager) AllBurials() []Burials {
	var entities []Burials
	state.db.Find(&entities)
	return entities
}

func (state *BurialsManager) BurialByID(id string) *Burials {
	entity := Burials{}
	state.db.First(&entity, id)
	return &entity
}

func (state *BurialsManager) CreateBurial(parentId uint, name string, description string, coord string) *Burials {
	entity := Burials{
		ParentId:    parentId,
		Name:        name,
		Description: description,
		Coord:       coord,
	}
	state.db.Create(&entity)
	return &entity
}

func (state *BurialsManager) UpdateBurial(id string, parentId uint, name string, description string, coord string) *Burials {
	entity := state.BurialByID(id)
	entity.ParentId = parentId
	entity.Name = name
	entity.Description = description
	entity.Coord = coord
	state.db.Save(&entity)
	return entity
}

func (state *BurialsManager) DeleteBurial(id string) *Burials {
	entity := Burials{}
	state.db.Delete(&entity, id)
	return &entity
}

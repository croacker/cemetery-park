package data

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Users struct {
	gorm.Model
	Name     string
	Password string
}

type UsersManager struct {
	db *DB
}

func NewUsersManager(db *DB) (*UsersManager, error) {
	db.AutoMigrate(&Users{})
	manager := UsersManager{}
	manager.db = db
	manager.predefined()
	return &manager, nil
}

func (state *UsersManager) AllUsers() []Users {
	var entities []Users
	state.db.Find(&entities)
	return entities
}

func (state *UsersManager) UserByID(id string) *Users {
	entity := Users{}
	state.db.First(&entity, id)
	return &entity
}

func (state *UsersManager) UserByName(name string) *Users {
	entity := Users{}
	err := state.db.Where("name=?", name).Find(&entity).Error
	handleError(err)
	return &entity
}

func (state *UsersManager) CreateUser(name string, password string) *Users {
	entity := Users{
		Name:     name,
		Password: password,
	}
	state.db.Create(&entity)
	return &entity
}

func (state *UsersManager) UpdateUser(id string, name string, password string) *Users {
	entity := state.UserByID(id)
	entity.Name = name
	entity.Password = password
	state.db.Save(&entity)
	return entity
}

func (state *UsersManager) DeleteUser(id string) *Users {
	entity := Users{}
	state.db.Delete(&entity, id)
	return &entity
}

func (state *UsersManager) predefined() {
	entity := state.UserByName("user")
	if entity.ID == 0 {
		log.Printf("User not exists '%s'", "user")
		state.CreateUser("user", "password")
	} else {
		log.Printf("User exists '%s'", entity.ID)
	}
}

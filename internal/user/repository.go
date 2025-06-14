package user

import (
	"errors"
	"sync"
)

type Repository interface {
	GetByEmail(email string) (User, error)
	GetByID(id uint) (User, error)
	Save(user User) (User, error)
}

type inMemoryDatabase struct {
	users  []User
	nextID uint
	mutex  sync.Mutex
}

func NewRepository() Repository {
	return &inMemoryDatabase{
		users:  []User{},
		nextID: 1,
		mutex:  sync.Mutex{},
	}
}

func (db *inMemoryDatabase) GetByEmail(email string) (User, error) {
	for _, user := range db.users {
		if email == user.Email {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (db *inMemoryDatabase) GetByID(id uint) (User, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for _, user := range db.users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("not found")
}

func (db *inMemoryDatabase) Save(user User) (User, error) {
	user.ID = db.nextID
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.users = append(db.users, user)
	db.nextID++
	return user, nil
}

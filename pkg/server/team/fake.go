package team

import (
	"sync"

	"github.com/luizalabs/teresa-api/models/storage"
	"github.com/luizalabs/teresa-api/pkg/server/user"
)

type FakeOperations struct {
	mutex   *sync.RWMutex
	Storage map[string]*storage.Team

	UserOps user.Operations
}

func (f *FakeOperations) Create(name, email, url string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if _, found := f.Storage[name]; found {
		return ErrTeamAlreadyExists
	}

	f.Storage[name] = &storage.Team{Name: name, Email: email, URL: url}
	return nil
}

func (f *FakeOperations) AddUser(name, email string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	t, found := f.Storage[name]
	if !found {
		return ErrNotFound
	}

	u, err := f.UserOps.GetUser(email)
	if err != nil {
		return err
	}

	for _, userOfTeam := range t.Users {
		if userOfTeam.Email == email {
			return ErrUserAlreadyInTeam
		}
	}

	t.Users = append(t.Users, *u)
	return nil
}

func NewFakeOperations() Operations {
	return &FakeOperations{
		mutex:   &sync.RWMutex{},
		Storage: make(map[string]*storage.Team),
		UserOps: user.NewFakeOperations()}
}

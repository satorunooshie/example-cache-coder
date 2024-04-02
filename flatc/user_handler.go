package flatc

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/satorunooshie/example-cache-coder/model"
)

type UserHandler struct{}

func (h UserHandler) Make(b *flatbuffers.Builder, v *model.User) ([]byte, error) {
	b.Reset()
	// must create all objects before Start.
	// If you create after Start, it will be panic with following error.
	// `Incorrect creation order: object must not be nested.`
	name := b.CreateString(v.Name)
	email := b.CreateString(v.Email)
	UserStart(b)
	UserAddName(b, name)
	UserAddId(b, int32(v.ID))
	UserAddEmail(b, email)
	pos := UserEnd(b)
	b.Finish(pos)
	return b.Bytes[b.Head():], nil
}

func (h UserHandler) Read(b []byte) (*model.User, error) {
	user := GetRootAsUser(b, 0)
	name := user.Name()
	id := user.Id()
	email := user.Email()
	return &model.User{
		Name:  string(name),
		ID:    int(id),
		Email: string(email),
	}, nil
}

type UsersHandler struct{}

func (h UsersHandler) Make(b *flatbuffers.Builder, list []model.User) ([]byte, error) {
	b.Reset()
	UsersStartUsersVector(b, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		UserAddName(b, b.CreateString(list[i].Name))
		UserAddId(b, int32(list[i].ID))
		UserAddEmail(b, b.CreateString(list[i].Email))
	}
	users := b.EndVector(len(list))
	UsersStart(b)
	UsersAddUsers(b, users)
	pos := UsersEnd(b)
	b.Finish(pos)
	return b.Bytes[b.Head():], nil
}

func (h UsersHandler) Read(b []byte) ([]model.User, error) {
	users := GetRootAsUsers(b, 0)
	usersLen := users.UsersLength()
	result := make([]model.User, usersLen)
	for i := 0; i < usersLen; i++ {
		user := new(User)
		if !users.Users(user, i) {
			return nil, nil
		}
		result[i] = model.User{
			Name:  string(user.Name()),
			ID:    int(user.Id()),
			Email: string(user.Email()),
		}
	}
	return []model.User(result), nil
}

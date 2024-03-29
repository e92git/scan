package store

import (
	"scan/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// FindBySession ...
func (r *UserRepository) FindBySession(session string) (*model.User, error) {
	u := &model.User{}
	res := r.store.db.Where("session = ?", session).First(u)
	return u, res.Error
}

// First ...
func (r *UserRepository) First(id int64) (*model.User, error) {
	u := &model.User{ID: id}
	res := r.store.db.First(u)
	return u, res.Error
}

// // Create ...
// func (r *UserRepository) Create(u *model.User) error {
// 	if err := u.Validate(); err != nil {
// 		return err
// 	}

// 	if err := u.BeforeCreate(); err != nil {
// 		return err
// 	}

// 	return r.store.db.QueryRow(
// 		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
// 		u.Email,
// 		u.EncryptedPassword,
// 	).Scan(&u.ID)
// }

// Find ...
// func (r *UserRepository) Find(id int) (*model.User, error) {
// 	u := &model.User{}
// 	if err := r.store.db.QueryRow(
// 		"SELECT id, email, encrypted_password FROM users WHERE id = ?",
// 		id,
// 	).Scan(
// 		&u.ID,
// 		&u.Email,
// 		&u.EncryptedPassword,
// 	); err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, ErrRecordNotFound
// 		}

// 		return nil, err
// 	}

// 	return u, nil
// }

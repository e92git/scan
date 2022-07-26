package store

import (
	"scan/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// FindByApiKey ...
func (r *UserRepository) FindByApiKey(apiKey string) (*model.User, error) {
	u := &model.User{}
	res := r.store.db.Where("api_key = ?", apiKey).First(u)
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

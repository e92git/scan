package model

import (
)

type User struct {
	ID        int64  `json:"id" example:"234"`
	Name      string `json:"name" example:"ivan_v"`
	Role      string `json:"role" example:"client" enums:"client,show_api,manager,admin"`
	Session   string `json:"session" example:""`
	CreatedAt string `json:"created_at" example:"2022-07-23 11:23:55"`
}

var UserRoles = struct {
	Client  string
	ShowApi string
	Manager string
	Admin   string
}{
	Client:  "client",   // Покупатели в магазине (нет доступа к API)
	ShowApi: "show_api", // client + GET запросы в API
	Manager: "manager",  // show_api + Все API-запросы доступные менеджеру магазина
	Admin:   "admin",    // manager + Все API-запросы
}

func (u *User) HasShowApi() bool {
	return u.Role == UserRoles.ShowApi || u.Role == UserRoles.Manager || u.Role == UserRoles.Admin
}

func (u *User) HasManager() bool {
	return u.Role == UserRoles.Manager || u.Role == UserRoles.Admin
}

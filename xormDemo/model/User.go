package model

import "time"

type CreateName string

func (j CreateName) MarshalJSON() ([]byte, error) {
	return []byte("陈俊"), nil
}

type User struct {
	Id         int64
	Name       string `xorm:"varchar(50) not null unique 'user_name'"`
	Age        int64
	CreateDate time.Time `xorm:"created"`
	Version    int64     `xorm:"version"`
}

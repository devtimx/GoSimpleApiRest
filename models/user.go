package models

import "time"

/*User, the model for the database*/
type User struct {
	ID        uint      `json:"id,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	BirthDate time.Time `json:"birthDate,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Profile   string    `json:"profile,omitempty"`
	Status    bool      `json:"status,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdateAt  time.Time `json:"updatedAt,omitempty"`
}

/*
CREATE TABLE users (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    birth_date DATE NOT NULL,
	email VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
	profile varchar(256) NOT NULL,
	status boolean,
	created_at DATE NOT NULL,
	updated_at DATE NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);
*/

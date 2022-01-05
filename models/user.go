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
}

/*
CREATE TABLE users (
    id serial NOT NULL,
    firstName VARCHAR(150) NOT NULL,
    lastName VARCHAR(150) NOT NULL,
    birthDate DATE NOT NULL,
	email VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);
*/

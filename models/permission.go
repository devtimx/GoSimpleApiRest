package models

import "time"

/*User, the model for the database*/
type Permission struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"Name,omitempty"`
	GuardName string    `json:"guardName,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdateAt  time.Time `json:"updatedAt,omitempty"`
}

/*
CREATE TABLE permissions (
    id serial NOT NULL,
    name VARCHAR(150) NOT NULL,
    guard_name VARCHAR(150) NOT NULL,
    created_at DATE NOT NULL,
	updated_at DATE NOT NULL,
    CONSTRAINT pk_permissions PRIMARY KEY(id)
);
*/

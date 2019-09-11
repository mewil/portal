package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mewil/portal/common/database"
	"github.com/mewil/portal/common/logger"
	"go.uber.org/zap"
)

type AuthRepository interface {
	StoreAuthRecord(string, string, string, bool) error
	EmailExists(string) (bool, error)
	LoadPasswordHash(string) ([]byte, error)
	LoadUserIdAndAdminStatus(string) (string, bool, error)
}

const authSchema string = `create table if not exists auth (
	email varchar(64) collate utf8mb4_bin not null unique,
	user_uuid binary(16) not null unique,
	password_hash binary(64) not null,
	admin tinyint(1) default 0,
	updated_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	primary key (email)
) engine InnoDB default charset utf8mb4 collate utf8mb4_bin;`

type repository struct {
	log logger.Logger
	db  database.DB
}

func NewAuthRepository(log logger.Logger, db database.DB, adminEmail, adminPassword string) (AuthRepository, error) {
	r := repository{
		log: log.(*zap.SugaredLogger).Named("repository"),
		db:  db,
	}
	if _, err := db.Exec(authSchema); err != nil {
		return nil, fmt.Errorf("failed to create auth table %s", err.Error())
	}
	adminUserExists, err := r.EmailExists(adminEmail)
	if err != nil {
		return nil, fmt.Errorf("failed to query auth table %s", err.Error())
	}
	if !adminUserExists {
		adminId, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		if err := r.StoreAuthRecord(adminEmail, adminId.String(), adminPassword, true); err != nil {
			return nil, fmt.Errorf("failed to insert admin auth record %s", err.Error())
		}
	}
	return &r, nil
}

func (r *repository) StoreAuthRecord(email, userId, password string, isAdmin bool) error {
	passwordHash, err := HashPassword(password)
	if err != nil {
		return err
	}
	if _, err := r.db.Exec(
		"insert auth set email=?, user_uuid=UUID_TO_BIN(?), password_hash=?, admin=?",
		email,
		userId,
		passwordHash,
		isAdmin,
	); err != nil {
		return err
	}
	r.log.Info("created new auth record for user " + email)
	return nil
}

func (r *repository) EmailExists(email string) (exists bool, err error) {
	err = r.db.QueryRow("select exists (select 1 from user where email=?)", email).Scan(
		&exists,
	)
	return
}

func (r *repository) LoadPasswordHash(email string) ([]byte, error) {
	passwordHash := []byte{}
	if err := r.db.QueryRow(
		"select password_hash from user where email=? limit 1",
		email,
	).Scan(
		&passwordHash,
	); err != nil {
		return nil, err
	}
	return passwordHash, nil
}

func (r *repository) LoadUserIdAndAdminStatus(email string) (string, bool, error) {
	userId, isAdmin := *new(string), *new(bool)
	if err := r.db.QueryRow(
		"select BIN_TO_UUID(user_uuid), admin from auth where email=? limit 1",
		email,
	).Scan(
		&userId,
		&isAdmin,
	); err != nil {
		return "", false, err
	}
	return userId, isAdmin, nil
}

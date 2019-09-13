package main

import (
	"fmt"

	"github.com/mewil/portal/common/database"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
	"go.uber.org/zap"
)

type UserRepository interface {
	CreateUser(string, string, string) error
	UserIDExists(string) (bool, error)
	UsernameExists(string) (bool, error)
	GetUser(string) (*pb.User, error)
	FollowUser(string, string) error
	UnfollowUser(string, string) error
	GetFollowers(string, uint32) ([]*pb.User, error)
	GetFollowing(string, uint32) ([]*pb.User, error)
}

func NewUserRepository(log logger.Logger, db database.DB) (UserRepository, error) {
	r := repository{
		log: log.(*zap.SugaredLogger).Named("repository"),
		db:  db,
	}
	for _, schema := range []string{userSchema, followingSchema} {
		if _, err := db.Exec(schema); err != nil {
			return nil, fmt.Errorf("failed to create table %s", err.Error())
		}
	}
	return &r, nil
}

const userSchema string = `create table if not exists users (
	user_id binary(16) not null unique,
	username varchar(64) collate utf8mb4_bin not null unique,
	file_id binary(16) not null,
	name varchar(64) collate utf8mb4_bin,
	description varchar(256) collate utf8mb4_bin,
	updated_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	primary key (user_id)
) engine InnoDB default charset utf8mb4 collate utf8mb4_bin;`

const followingSchema string = `create table if not exists following (
	user_id binary(16) not null unique,
	following_id binary(16) not null unique,
	created_at timestamp default current_timestamp,
	key users(user_id),
	key users(following_id)
) engine InnoDB;`

type repository struct {
	log logger.Logger
	db  database.DB
}

const userPageSize = 25

func (r *repository) CreateUser(userID, username, email string) (err error) {
	_, err = r.db.Exec(
		"insert users set user_id=UUID_TO_BIN(?), username=?, email=?",
		userID,
		username,
		email,
	)
	return
}

func (r *repository) UserIDExists(userID string) (exists bool, err error) {
	err = r.db.QueryRow("select exists (select 1 from users where user_id=?)", userID).Scan(
		&exists,
	)
	return
}

func (r *repository) UsernameExists(username string) (exists bool, err error) {
	err = r.db.QueryRow("select exists (select 1 from users where username=?)", username).Scan(
		&exists,
	)
	return
}

func (r *repository) GetUser(userID string) (*pb.User, error) {
	u := pb.User{}
	if err := r.db.QueryRow(
		"select BIN_TO_UUID(user_id), name, BIN_TO_UUID(file_id), description, username from users where user_id=UUID_TO_BIN(?) limit 1",
		userID,
	).Scan(
		&u.UserID,
		&u.Name,
		&u.FileID,
		&u.Description,
		&u.Username,
	); err != nil {
		return nil, err
	}
	err := *new(error)
	u.FollowerCount, err = r.getFollowerCount(userID)
	if err != nil {
		return nil, err
	}
	u.FollowingCount, err = r.getFollowingCount(userID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) getFollowerCount(userID string) (n uint32, err error) {
	err = r.db.QueryRow(
		"select count(*) from following where following_id=UUID_TO_BIN(?)",
		userID,
	).Scan(&n)
	return
}

func (r *repository) getFollowingCount(userID string) (n uint32, err error) {
	err = r.db.QueryRow(
		"select count(*) from following where user_id=UUID_TO_BIN(?)",
		userID,
	).Scan(&n)
	return
}

func (r *repository) FollowUser(userID string, followID string) (err error) {
	_, err = r.db.Exec(
		"insert following set user_id=UUID_TO_BIN(?), following_id=UUID_TO_BIN(?)",
		userID,
		followID,
	)
	return
}

func (r *repository) UnfollowUser(userID string, followID string) (err error) {
	_, err = r.db.Exec(
		"delete from following where user_id=UUID_TO_BIN(?), following_id=UUID_TO_BIN(?) limit 1",
		userID,
		followID,
	)
	return
}

const selectFollowersQuery = `
select BIN_TO_UUID(user.user_id), user.name, BIN_TO_UUID(user.file_id), user.description, user.username
from users, following
where following.following_id=UUID_TO_BIN(?) and user.user_id=following.user_id
order by user.username asc
limit ?,?
`

func (r *repository) GetFollowers(userID string, page uint32) ([]*pb.User, error) {
	return r.getUserPage(selectFollowersQuery, userID, page)
}

const selectFollowingQuery = `
select BIN_TO_UUID(user.user_id), user.name, BIN_TO_UUID(user.file_id), user.description, user.username
from users, following
where following.user_id=UUID_TO_BIN(?) and user.user_id=following.user_id
order by user.username asc
limit ?,?
`

func (r *repository) GetFollowing(userID string, page uint32) ([]*pb.User, error) {
	return r.getUserPage(selectFollowersQuery, userID, page)
}

func (r *repository) getUserPage(query string, userID string, page uint32) ([]*pb.User, error) {
	offset := page * userPageSize
	rows, err := r.db.Query(query, userID, offset, userPageSize)
	if err != nil {
		return nil, err
	}
	results := make([]*pb.User, 0, userPageSize)
	for rows.Next() {
		u := pb.User{}
		if err := rows.Scan(
			&u.UserID,
			&u.Name,
			&u.FileID,
			&u.Description,
			&u.Username,
		); err != nil {
			return nil, err
		}
		results = append(results, &u)
	}
	return results, nil
}

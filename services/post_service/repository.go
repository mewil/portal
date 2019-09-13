package main

import (
	"fmt"

	"github.com/mewil/portal/common/database"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
	"go.uber.org/zap"
)

type PostRepository interface {
	CreatePost(postID, userID, fileID, caption string) error
	CreatePostLike(postID, userID, likeID string) error
	CreateComment(postID, userID, commentID, text string) error
	CreateCommentLike(commentID, userID, likeID string) error
	DeletePost(postID string) error
	DeletePostLike(likeID string) error
	DeleteComment(commentID string) error
	DeleteCommentLike(likeID string) error
	GetPost(postID string) (*pb.Post, error)
	GetProfile(userID string, page uint32) ([]*pb.Post, error)
	GetFeed(userID string, page uint32) ([]*pb.Post, error)
	GetPostLikes(postID string, page uint32) ([]*pb.PostLike, error)
	GetPostComments(postID string, page uint32) ([]*pb.Comment, error)
	GetCommentLikes(commentID string, page uint32) ([]*pb.CommentLike, error)
	GetComment(commentID string) (*pb.Comment, error)
}

func NewPostRepository(log logger.Logger, db database.DB) (PostRepository, error) {
	r := repository{
		log: log.(*zap.SugaredLogger).Named("repository"),
		db:  db,
	}
	for _, schema := range []string{postsSchema, commentsSchema, postLikesSchema, commentLikesSchema} {
		if _, err := db.Exec(schema); err != nil {
			return nil, fmt.Errorf("failed to create table %s", err.Error())
		}
	}
	return &r, nil
}

const postsSchema string = `create table if not exists posts (
	post_id binary(16) not null unique,
	user_id binary(16) not null,
	file_id binary(16) not null,
	caption varchar(1024) collate utf8mb4_bin,
	updated_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	primary key (post_id),
	key users(user_id)
) engine InnoDB default charset utf8mb4 collate utf8mb4_bin;`

const commentsSchema string = `create table if not exists comments (
	comment_id binary(16) not null unique,
	user_id binary(16) not null,
	post_id binary(16) not null,
	text varchar(1024) collate utf8mb4_bin,
	updated_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	primary key (comment_id),
	key users(user_id),
	key posts(post_id)
) engine InnoDB;`

const postLikesSchema string = `create table if not exists post_likes (
	like_id binary(16) not null unique,
	user_id binary(16) not null,
	post_id binary(16) not null,
	created_at timestamp default current_timestamp,
	primary key (like_id),
	key users(user_id),
	key posts(post_id)
) engine InnoDB;`

const commentLikesSchema string = `create table if not exists comment_likes (
	like_id binary(16) not null unique,
	user_id binary(16) not null,
	comment_id binary(16) not null,
	created_at timestamp default current_timestamp,
	primary key (like_id),
	key users(user_id),
	key comments(comment_id)
) engine InnoDB;`

type repository struct {
	log logger.Logger
	db  database.DB
}

const (
	postPageSize    = 25
	likePageSize    = 50
	commentPageSize = 25
)

func (r *repository) CreatePost(postID, userID, fileID, caption string) (err error) {
	_, err = r.db.Exec(
		"insert posts set post_id=UUID_TO_BIN(?), user_id=UUID_TO_BIN(?), file_id=UUID_TO_BIN(?), caption=?",
		postID,
		userID,
		fileID,
		caption,
	)
	return
}

func (r *repository) CreatePostLike(postID, userID, likeID string) (err error) {
	_, err = r.db.Exec(
		"insert post_likes set post_id=UUID_TO_BIN(?), user_id=UUID_TO_BIN(?), like_id=UUID_TO_BIN(?)",
		postID,
		userID,
		likeID,
	)
	return
}

func (r *repository) CreateComment(postID, userID, commentID, text string) (err error) {
	_, err = r.db.Exec(
		"insert comments set post_id=UUID_TO_BIN(?), user_id=UUID_TO_BIN(?), comment_id=UUID_TO_BIN(?), text=?",
		postID,
		userID,
		commentID,
		text,
	)
	return
}

func (r *repository) CreateCommentLike(commentID, userID, likeID string) (err error) {
	_, err = r.db.Exec(
		"insert comment_likes set comment_id=UUID_TO_BIN(?), user_id=UUID_TO_BIN(?), comment_id=UUID_TO_BIN(?)",
		commentID,
		userID,
		likeID,
	)
	return
}

func (r *repository) DeletePost(postID string) (err error) {
	_, err = r.db.Exec(
		"delete from posts where post_id=?",
		postID,
	)
	if err != nil {
		return
	}
	_, err = r.db.Exec(
		"delete from post_likes where post_id=?",
		postID,
	)
	if err != nil {
		return
	}
	_, err = r.db.Exec(
		"delete comment_likes from comment_likes inner join on comments.comment_id = comment_likes.comment_id where comments.post_id=?",
		postID,
	)
	if err != nil {
		return
	}
	_, err = r.db.Exec(
		"delete from comments where post_id=?",
		postID,
	)
	return
}

func (r *repository) DeletePostLike(likeID string) (err error) {
	_, err = r.db.Exec(
		"delete from post_likes where like_id=?",
		likeID,
	)
	return
}

func (r *repository) DeleteComment(commentID string) (err error) {
	_, err = r.db.Exec(
		"delete from comments where comment_id=?",
		commentID,
	)
	if err != nil {
		return
	}
	_, err = r.db.Exec(
		"delete from comment_likes where comment_id=?",
		commentID,
	)
	return
}

func (r *repository) DeleteCommentLike(likeID string) (err error) {
	_, err = r.db.Exec(
		"delete from comment_likes where like_id=?",
		likeID,
	)
	return
}

func (r *repository) GetPost(postID string) (*pb.Post, error) {
	p := pb.Post{}
	if err := r.db.QueryRow(
		"select BIN_TO_UUID(post_id), BIN_TO_UUID(user_id), BIN_TO_UUID(file_id), caption, created_at from posts where post_id=UUID_TO_BIN(?) limit 1",
		postID,
	).Scan(
		&p.PostID,
		&p.UserID,
		&p.FileID,
		&p.Caption,
		&p.CreatedAt,
	); err != nil {
		return nil, err
	}
	err := *new(error)
	p.LikeCount, err = r.getPostLikeCount(postID)
	if err != nil {
		return nil, err
	}
	p.CommentCount, err = r.getPostCommentCount(postID)
	if err != nil {
		return nil, err
	}
	p.TopComments, err = r.getTopPostComments(postID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) getPostLikeCount(postID string) (n uint32, err error) {
	err = r.db.QueryRow(
		"select count(*) from post_likes where post_id=UUID_TO_BIN(?)",
		postID,
	).Scan(&n)
	return
}

func (r *repository) getCommentLikeCount(commentID string) (n uint32, err error) {
	err = r.db.QueryRow(
		"select count(*) from comment_likes where comment_id=UUID_TO_BIN(?)",
		commentID,
	).Scan(&n)
	return
}

func (r *repository) getPostCommentCount(postID string) (n uint32, err error) {
	err = r.db.QueryRow(
		"select count(*) from comments where post_id=UUID_TO_BIN(?)",
		postID,
	).Scan(&n)
	return
}

const topCommentLimit = 5

func (r *repository) getTopPostComments(postID string) ([]*pb.Comment, error) {
	rows, err := r.db.Query(`select BIN_TO_UUID(comments.comment_id), BIN_TO_UUID(comments.user_id), BIN_TO_UUID(comments.post_id), comments.text, comments.created_at COUNT(like_id) AS like_count
from comments
left join comment_likes 
on comments.comment_id = comment_likes.comment_id
group by comments.comment_id
order by like_count
limit ?`, postID, topCommentLimit,
	)
	if err != nil {
		return nil, err
	}
	return r.readCommentRows(rows)
}

func (r *repository) GetProfile(userID string, page uint32) ([]*pb.Post, error) {
	return r.getPostPage(`
select IN_TO_UUID(post_id), BIN_TO_UUID(user_id), BIN_TO_UUID(file_id), caption, created_at
from posts
where user_id=UUID_TO_BIN(?)
order by created at desc
limit ?,?`, userID, page)
}

func (r *repository) GetFeed(userID string, page uint32) ([]*pb.Post, error) {
	return r.getPostPage(`
select BIN_TO_UUID(post_id), BIN_TO_UUID(user_id), BIN_TO_UUID(file_id), caption, created_at
from posts
inner join following
on posts.user_id = following.following_id
where following.user_id=UUID_TO_BIN(?)
order by posts.created_at desc
limit ?,?`, userID, page)
}

func (r *repository) getPostPage(query string, userID string, page uint32) ([]*pb.Post, error) {
	offset := page * postPageSize
	rows, err := r.db.Query(query, userID, offset, postPageSize)
	if err != nil {
		return nil, err
	}
	return r.readPostRows(rows)
}

func (r *repository) GetPostLikes(postID string, page uint32) ([]*pb.PostLike, error) {
	offset := page * likePageSize
	rows, err := r.db.Query(`
select BIN_TO_UUID(like_id), BIN_TO_UUID(user_id), BIN_TO_UUID(post_id), created_at
from post_likes
where post_id=UUID_TO_BIN(?)
order by created at desc
limit ?,?`, postID, offset, likePageSize)
	if err != nil {
		return nil, err
	}
	return r.readPostLikeRows(rows)
}

func (r *repository) GetPostComments(postID string, page uint32) ([]*pb.Comment, error) {
	offset := page * postPageSize
	rows, err := r.db.Query(`
select BIN_TO_UUID(comment_id), BIN_TO_UUID(user_id), BIN_TO_UUID(post_id), text, created_at
from comments
where post_id=UUID_TO_BIN(?)
order by created at desc
limit ?,?`, postID, offset, postPageSize)
	if err != nil {
		return nil, err
	}
	return r.readCommentRows(rows)
}

func (r *repository) GetCommentLikes(commentID string, page uint32) ([]*pb.CommentLike, error) {
	offset := page * postPageSize
	rows, err := r.db.Query(`
select BIN_TO_UUID(like_id), BIN_TO_UUID(user_id), BIN_TO_UUID(comment_id), created_at
from comment_likes
where post_id=UUID_TO_BIN(?)
order by created at desc
limit ?,?`, commentID, offset, postPageSize)
	if err != nil {
		return nil, err
	}
	return r.readCommentLikeRows(rows)
}

func (r *repository) GetComment(commentID string) (*pb.Comment, error) {
	c := pb.Comment{}
	if err := r.db.QueryRow(
		"select BIN_TO_UUID(comment_id), BIN_TO_UUID(user_id), BIN_TO_UUID(post_id), text, created_at from comments where comment_id=UUID_TO_BIN(?) limit 1",
		commentID,
	).Scan(
		&c.CommentID,
		&c.UserID,
		&c.PostID,
		&c.Text,
		&c.CreatedAt,
	); err != nil {
		return nil, err
	}
	err := *new(error)
	c.LikeCount, err = r.getCommentLikeCount(c.PostID)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *repository) readCommentRows(rows database.Rows) ([]*pb.Comment, error) {
	results := make([]*pb.Comment, 0)
	for rows.Next() {
		c := pb.Comment{}
		if err := rows.Scan(
			&c.CommentID,
			&c.UserID,
			&c.PostID,
			&c.Text,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}
		err := *new(error)
		c.LikeCount, err = r.getCommentLikeCount(c.PostID)
		if err != nil {
			return nil, err
		}
		results = append(results, &c)
	}
	return results, nil
}

func (r *repository) readCommentLikeRows(rows database.Rows) ([]*pb.CommentLike, error) {
	results := make([]*pb.CommentLike, 0)
	for rows.Next() {
		l := pb.CommentLike{}
		if err := rows.Scan(
			&l.LikeID,
			&l.UserID,
			&l.CommentID,
			&l.CreatedAt,
		); err != nil {
			return nil, err
		}
		results = append(results, &l)
	}
	return results, nil
}

func (r *repository) readPostLikeRows(rows database.Rows) ([]*pb.PostLike, error) {
	results := make([]*pb.PostLike, 0)
	for rows.Next() {
		l := pb.PostLike{}
		if err := rows.Scan(
			&l.LikeID,
			&l.UserID,
			&l.PostID,
			&l.CreatedAt,
		); err != nil {
			return nil, err
		}
		results = append(results, &l)
	}
	return results, nil
}

func (r *repository) readPostRows(rows database.Rows) ([]*pb.Post, error) {
	results := make([]*pb.Post, 0, postPageSize)
	for rows.Next() {
		p := pb.Post{}
		if err := rows.Scan(
			&p.PostID,
			&p.UserID,
			&p.FileID,
			&p.Caption,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		err := *new(error)
		p.LikeCount, err = r.getPostLikeCount(p.PostID)
		if err != nil {
			return nil, err
		}
		p.CommentCount, err = r.getPostCommentCount(p.PostID)
		if err != nil {
			return nil, err
		}
		p.TopComments, err = r.getTopPostComments(p.PostID)
		if err != nil {
			return nil, err
		}
		results = append(results, &p)
	}
	return results, nil
}

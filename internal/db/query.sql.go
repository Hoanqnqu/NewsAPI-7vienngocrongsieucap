// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteCategory = `-- name: DeleteCategory :exec
UPDATE categories SET deleted_at = NOW() WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const deleteDisLike = `-- name: DeleteDisLike :exec
DELETE from dislikes Where news_id = $1 and user_id = $2
`

type DeleteDisLikeParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) DeleteDisLike(ctx context.Context, arg DeleteDisLikeParams) error {
	_, err := q.db.Exec(ctx, deleteDisLike, arg.NewsID, arg.UserID)
	return err
}

const deleteLike = `-- name: DeleteLike :exec
DELETE from likes Where news_id = $1 and user_id = $2
`

type DeleteLikeParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) DeleteLike(ctx context.Context, arg DeleteLikeParams) error {
	_, err := q.db.Exec(ctx, deleteLike, arg.NewsID, arg.UserID)
	return err
}

const deleteNews = `-- name: DeleteNews :exec
UPDATE news SET deleted_at = NOW() WHERE id = $1
`

func (q *Queries) DeleteNews(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteNews, id)
	return err
}

const deleteSave = `-- name: DeleteSave :exec
DELETE from saves Where news_id = $1 and user_id = $2
`

type DeleteSaveParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) DeleteSave(ctx context.Context, arg DeleteSaveParams) error {
	_, err := q.db.Exec(ctx, deleteSave, arg.NewsID, arg.UserID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users SET deleted_at = NOW() WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getAdmin = `-- name: GetAdmin :many
SELECT id, auth_id, email, password, name, role, image_url, created_at, updated_at, deleted_at
FROM users
WHERE
    users.email = $1
    AND users.password = $2
    AND users.role = 'admin'
LIMIT 1
`

type GetAdminParams struct {
	Email    pgtype.Text
	Password pgtype.Text
}

func (q *Queries) GetAdmin(ctx context.Context, arg GetAdminParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getAdmin, arg.Email, arg.Password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.AuthID,
			&i.Email,
			&i.Password,
			&i.Name,
			&i.Role,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllCategories = `-- name: GetAllCategories :many
SELECT id, name, created_at, updated_at, deleted_at from categories
`

func (q *Queries) GetAllCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, getAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllNews = `-- name: GetAllNews :many
SELECT
    n.id AS id,
    n.author,
    n.title,
    n.description,
    n.content,
    n.url,
    n.image_url,
    n.publish_at,
    n.created_at AS created_at,
    n.updated_at AS updated_at,
    n.deleted_at AS deleted_at,
    json_agg(hc.category_id::uuid) AS category_ids
FROM news n
    Left JOIN has_categories hc ON n.id = hc.news_id
GROUP BY
    n.id,
    n.author,
    n.title,
    n.description,
    n.content,
    n.url,
    n.image_url,
    n.publish_at,
    n.created_at,
    n.updated_at,
    n.deleted_at
`

type GetAllNewsRow struct {
	ID          pgtype.UUID
	Author      pgtype.Text
	Title       pgtype.Text
	Description pgtype.Text
	Content     pgtype.Text
	Url         pgtype.Text
	ImageUrl    pgtype.Text
	PublishAt   pgtype.Timestamp
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	DeletedAt   pgtype.Timestamp
	CategoryIds []byte
}

func (q *Queries) GetAllNews(ctx context.Context) ([]GetAllNewsRow, error) {
	rows, err := q.db.Query(ctx, getAllNews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllNewsRow
	for rows.Next() {
		var i GetAllNewsRow
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.Url,
			&i.ImageUrl,
			&i.PublishAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.CategoryIds,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, auth_id, email, password, name, role, image_url, created_at, updated_at, deleted_at from users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.AuthID,
			&i.Email,
			&i.Password,
			&i.Name,
			&i.Role,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDislike = `-- name: GetDislike :one
SELECT news_id, user_id from dislikes Where news_id = $1 and user_id = $2
`

type GetDislikeParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) GetDislike(ctx context.Context, arg GetDislikeParams) (Dislike, error) {
	row := q.db.QueryRow(ctx, getDislike, arg.NewsID, arg.UserID)
	var i Dislike
	err := row.Scan(&i.NewsID, &i.UserID)
	return i, err
}

const getLike = `-- name: GetLike :one
SELECT news_id, user_id from likes Where news_id = $1 and user_id = $2
`

type GetLikeParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) GetLike(ctx context.Context, arg GetLikeParams) (Like, error) {
	row := q.db.QueryRow(ctx, getLike, arg.NewsID, arg.UserID)
	var i Like
	err := row.Scan(&i.NewsID, &i.UserID)
	return i, err
}

const getNews = `-- name: GetNews :one
SELECT
    n.id AS id,
    n.author,
    n.title,
    n.description,
    n.content,
    n.url,
    n.image_url,
    n.publish_at,
    n.created_at AS created_at,
    n.updated_at AS updated_at,
    n.deleted_at AS deleted_at,
    json_agg(hc.category_id::uuid) AS category_ids
FROM news n
    Left JOIN has_categories hc ON n.id = hc.news_id
where
    id = $1
GROUP BY
    n.id,
    n.author,
    n.title,
    n.description,
    n.content,
    n.url,
    n.image_url,
    n.publish_at,
    n.created_at,
    n.updated_at,
    n.deleted_at
`

type GetNewsRow struct {
	ID          pgtype.UUID
	Author      pgtype.Text
	Title       pgtype.Text
	Description pgtype.Text
	Content     pgtype.Text
	Url         pgtype.Text
	ImageUrl    pgtype.Text
	PublishAt   pgtype.Timestamp
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	DeletedAt   pgtype.Timestamp
	CategoryIds []byte
}

func (q *Queries) GetNews(ctx context.Context, id pgtype.UUID) (GetNewsRow, error) {
	row := q.db.QueryRow(ctx, getNews, id)
	var i GetNewsRow
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Url,
		&i.ImageUrl,
		&i.PublishAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CategoryIds,
	)
	return i, err
}

const getSaves = `-- name: GetSaves :many
SELECT news_id from saves Where user_id = $1
`

func (q *Queries) GetSaves(ctx context.Context, userID pgtype.UUID) ([]pgtype.UUID, error) {
	rows, err := q.db.Query(ctx, getSaves, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.UUID
	for rows.Next() {
		var news_id pgtype.UUID
		if err := rows.Scan(&news_id); err != nil {
			return nil, err
		}
		items = append(items, news_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByAuthID = `-- name: GetUserByAuthID :one
SELECT id, auth_id, email, password, name, role, image_url, created_at, updated_at, deleted_at FROM users WHERE users.auth_id = $1 LIMIT 1
`

func (q *Queries) GetUserByAuthID(ctx context.Context, authID string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByAuthID, authID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AuthID,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.Role,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const insertCategory = `-- name: InsertCategory :exec
INSERT INTO categories (id, name, created_at) VALUES ($1, $2, NOW())
`

type InsertCategoryParams struct {
	ID   pgtype.UUID
	Name pgtype.Text
}

func (q *Queries) InsertCategory(ctx context.Context, arg InsertCategoryParams) error {
	_, err := q.db.Exec(ctx, insertCategory, arg.ID, arg.Name)
	return err
}

const insertDisLike = `-- name: InsertDisLike :exec
INSERT INTO dislikes (news_id, user_id) VALUES ($1, $2)
`

type InsertDisLikeParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) InsertDisLike(ctx context.Context, arg InsertDisLikeParams) error {
	_, err := q.db.Exec(ctx, insertDisLike, arg.NewsID, arg.UserID)
	return err
}

const insertHasCategory = `-- name: InsertHasCategory :exec
INSERT INTO has_categories (news_id, category_id) VALUES ($1, $2)
`

type InsertHasCategoryParams struct {
	NewsID     pgtype.UUID
	CategoryID pgtype.UUID
}

func (q *Queries) InsertHasCategory(ctx context.Context, arg InsertHasCategoryParams) error {
	_, err := q.db.Exec(ctx, insertHasCategory, arg.NewsID, arg.CategoryID)
	return err
}

const insertLike = `-- name: InsertLike :exec
INSERT INTO likes (news_id, user_id) VALUES ($1, $2)
`

type InsertLikeParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) InsertLike(ctx context.Context, arg InsertLikeParams) error {
	_, err := q.db.Exec(ctx, insertLike, arg.NewsID, arg.UserID)
	return err
}

const insertNews = `-- name: InsertNews :exec
INSERT INTO
    news (
        id,
        author,
        title,
        description,
        content,
        url,
        image_url,
        publish_at,
        created_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        NOW()
    )
`

type InsertNewsParams struct {
	ID          pgtype.UUID
	Author      pgtype.Text
	Title       pgtype.Text
	Description pgtype.Text
	Content     pgtype.Text
	Url         pgtype.Text
	ImageUrl    pgtype.Text
	PublishAt   pgtype.Timestamp
}

func (q *Queries) InsertNews(ctx context.Context, arg InsertNewsParams) error {
	_, err := q.db.Exec(ctx, insertNews,
		arg.ID,
		arg.Author,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.Url,
		arg.ImageUrl,
		arg.PublishAt,
	)
	return err
}

const insertSave = `-- name: InsertSave :exec
Insert into saves (news_id, user_id) values ($1, $2)
`

type InsertSaveParams struct {
	NewsID pgtype.UUID
	UserID pgtype.UUID
}

func (q *Queries) InsertSave(ctx context.Context, arg InsertSaveParams) error {
	_, err := q.db.Exec(ctx, insertSave, arg.NewsID, arg.UserID)
	return err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO
    users (
        id,
        auth_id,
        email,
        name,
        role,
        image_url,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6, NOW())
`

type InsertUserParams struct {
	ID       pgtype.UUID
	AuthID   string
	Email    pgtype.Text
	Name     pgtype.Text
	Role     pgtype.Text
	ImageUrl pgtype.Text
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
		arg.ID,
		arg.AuthID,
		arg.Email,
		arg.Name,
		arg.Role,
		arg.ImageUrl,
	)
	return err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET
    name = $1,
    updated_at = NOW()
WHERE
    id = $2
`

type UpdateCategoryParams struct {
	Name pgtype.Text
	ID   pgtype.UUID
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.Exec(ctx, updateCategory, arg.Name, arg.ID)
	return err
}

const updateNews = `-- name: UpdateNews :exec
UPDATE news
SET
    title = $1,
    description = $2,
    content = $3,
    author = $4,
    url = $5,
    image_url = $6,
    publish_at = $7,
    updated_at = NOW()
WHERE
    id = $8
`

type UpdateNewsParams struct {
	Title       pgtype.Text
	Description pgtype.Text
	Content     pgtype.Text
	Author      pgtype.Text
	Url         pgtype.Text
	ImageUrl    pgtype.Text
	PublishAt   pgtype.Timestamp
	ID          pgtype.UUID
}

func (q *Queries) UpdateNews(ctx context.Context, arg UpdateNewsParams) error {
	_, err := q.db.Exec(ctx, updateNews,
		arg.Title,
		arg.Description,
		arg.Content,
		arg.Author,
		arg.Url,
		arg.ImageUrl,
		arg.PublishAt,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
    name = $1,
    image_url = $2,
    updated_at = NOW()
WHERE
    id = $3
`

type UpdateUserParams struct {
	Name     pgtype.Text
	ImageUrl pgtype.Text
	ID       pgtype.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.Name, arg.ImageUrl, arg.ID)
	return err
}

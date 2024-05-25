-- name: InsertUser :exec
INSERT INTO users (id,
                   auth_id,
                   email,
                   name,
                   role,
                   image_url,
                   created_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW());

-- name: GetAllUsers :many
SELECT *
from users;

-- name: GetUserByAuthID :one
SELECT *
FROM users
WHERE users.auth_id = $1 LIMIT 1;
-- name: UpdateUser :exec
UPDATE users
SET name       = $1,
    image_url  = $2,
    updated_at = NOW()
WHERE id = $3;
-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;

-- name: InsertCategory :exec
INSERT INTO categories (id, name, created_at)
VALUES ($1, $2, NOW());
-- name: UpdateCategory :exec
UPDATE categories
SET name       = $1,
    updated_at = NOW()
WHERE id = $2;
-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1;

-- name: GetAllCategories :many
SELECT *
from categories;

-- name: InsertNews :exec
INSERT INTO news (id,
                  author,
                  title,
                  description,
                  content,
                  url,
                  image_url,
                  publish_at,
                  created_at)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        NOW());

-- name: UpdateNews :exec
UPDATE news
SET title       = $1,
    description = $2,
    content     = $3,
    author      = $4,
    url         = $5,
    image_url   = $6,
    publish_at  = $7,
    updated_at  = NOW()
WHERE id = $8;
-- name: DeleteNews :exec
UPDATE news
SET deleted_at = NOW()
WHERE id = $1;

-- name: GetAllNews :many
SELECT *
from news;


-- name: InsertLike :exec
INSERT INTO likes(news_id,
                  user_id)
VALUES ($1,
        $2);

-- name: DeleteLike :exec
DELETE
from likes
Where news_id = $1
  and user_id = $2;

-- name: InsertDisLike :exec
INSERT INTO likes(news_id,
                  user_id)
VALUES ($1,
        $2);

-- name: DeleteDisLike :exec
DELETE
from likes
Where news_id = $1
  and user_id = $2;
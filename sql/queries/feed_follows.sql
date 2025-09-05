-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING
    feed_follows.*,
    (SELECT name FROM users WHERE id = feed_follows.user_id) AS user_name,
    (SELECT name FROM feeds WHERE id = feed_follows.feed_id) AS feed_name;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1
LIMIT 1;

-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
JOIN feeds ON feed_follows.feed_id = feeds.id
JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1
ORDER BY feed_follows.created_at DESC;

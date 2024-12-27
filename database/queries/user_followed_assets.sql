-- name: CreateUserFollowedAsset :one
INSERT INTO "user_followed_assets" (
    user_id, asset_id, followed_at 
) VALUES (
    $1, $2, now()
) RETURNING *;

-- name: DeleteUserFollowedAsset :exec
DELETE FROM "user_followed_assets" WHERE user_id = $1 AND asset_id = $2;
-- name: CreateAsset :one
INSERT INTO "assets" (
    cmc_id, name, slug, price, 
    percent_change_1h, percent_change_24h, 
    percent_change_7d, market_cap, 
    volume_24h, circulating_supply,
    all_time_high, all_time_low,
    turnover, total_supply, 
    max_supply, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8,
    $9, $10, $11, $12, $13, $14, $15, now(), now()
) RETURNING *;

-- name: GetAsset :one
SELECT * FROM "assets" 
WHERE cmc_id = $1 LIMIT 1;

-- name: UpdateAsset :one
UPDATE "assets"
SET
    price = coalesce(sqlc.narg('price'), price),
    percent_change_1h = coalesce(sqlc.narg('percent_change_1h'), percent_change_1h),
    percent_change_24h = coalesce(sqlc.narg('percent_change_24h'), percent_change_24h),
    percent_change_7d = coalesce(sqlc.narg('percent_change_7d'), percent_change_7d),
    market_cap = coalesce(sqlc.narg('market_cap'), market_cap),
    volume_24h = coalesce(sqlc.narg('volume_24h'), volume_24h),
    circulating_supply = coalesce(sqlc.narg('circulating_supply'), circulating_supply),
    all_time_high = coalesce(sqlc.narg('all_time_high'), all_time_high),
    all_time_low = coalesce(sqlc.narg('all_time_low'), all_time_low),
    turnover = coalesce(sqlc.narg('turnover'), turnover),
    total_supply = coalesce(sqlc.narg('total_supply'), total_supply),
    max_supply = coalesce(sqlc.narg('max_supply'), max_supply),
    updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
    WHERE cmc_id = sqlc.arg('cmc_id')
RETURNING *;

-- name: ListAssets :many
SELECT * FROM "assets"
ORDER BY cmc_id
LIMIT $1
OFFSET $2;
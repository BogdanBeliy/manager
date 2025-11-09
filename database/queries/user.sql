-- name: FetchAll :many
SELECT *
FROM hronos.public.users;

-- name: FetchOne :one
select *
from hronos.public.users
where id = $1
limit 1;

-- name: Create :one
insert into hronos.public.users (email, password, owner)
values ($1, $2, $3)
returning email, first_name, last_name;

-- name: Delete :one



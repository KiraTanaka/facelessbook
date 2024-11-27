SELECT id,
	created_time,
	author_id,
	text
FROM posts 
WHERE id = $1
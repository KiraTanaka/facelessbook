SELECT count(*)
FROM post_likes 
WHERE post_id = $1 and is_like
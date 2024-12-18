INSERT INTO posts
			(author_id,
			text)
VALUES     ($1,
			$2)
RETURNING id
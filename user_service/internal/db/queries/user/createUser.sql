INSERT INTO users
			(phone,
			pass_hash)
VALUES     ($1,
			$2)
RETURNING id
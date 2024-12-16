SELECT id,
	phone,
	pass_hash
FROM users 
WHERE phone = $1
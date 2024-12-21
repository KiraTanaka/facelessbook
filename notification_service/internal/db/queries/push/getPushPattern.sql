SELECT subject,
		push_message,
		full_message
FROM push_patterns
WHERE name = $1
INSERT INTO push_notification
			(user_id,
			subject,
			push_message,
			full_message)
VALUES     ($1,
			$2,
			$3,
			$4)
RETURNING id
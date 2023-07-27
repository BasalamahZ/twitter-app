package postgresql

const queryCreateTweet = `
	INSERT INTO 
		tweets 
	(
		title,
		description,
		create_time
	) VALUES (
		:title,
		:description,
		:create_time
	) RETURNING
		id
`

const queryGetTweet = `
	SELECT
		t.id,
		t.title,
		t.description
	FROM
		tweets t
	%s
`

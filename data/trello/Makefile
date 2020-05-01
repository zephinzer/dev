api_samples:
	curl "https://api.trello.com/1/members/me?key=${TRELLO_KEY}&token=${TRELLO_TOKEN}" | jq '.' | tee ./members-me.json
	curl "https://api.trello.com/1/members/me/notifications?key=${TRELLO_KEY}&token=${TRELLO_TOKEN}&read_filter=unread" | jq '.' | tee ./members-me-notifications.json

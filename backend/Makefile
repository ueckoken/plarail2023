up:
	docker compose down
	docker compose up --build -d

pull:
	git pull

test:
	docker compose -f compose.debug.yaml up --build -d

test-status:
	docker compose -f compose.debug.yaml ps

test-logs:
	docker compose -f compose.debug.yaml logs -f


build:
	docker compose build

run: build
	docker compose up

test-build:
	docker compose -f docker-compose-test.yml build

test-run:
	docker compose -f docker-compose-test.yml run --rm testapp

test-cleanup:
	docker compose -f docker-compose-test.yml rm --stop --force

test: test-build test-run test-cleanup

.PHONY: docker-up
docker-up:
	docker-compose -f build/dev/docker-compose.yaml up --build
 
.PHONY: docker-down
docker-down:
	docker-compose -f build/dev/docker-compose.yaml down
	docker system prune --volumes --force
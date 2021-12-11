default:
	echo 'nothing to do as default'
start-kafka:
	docker-compose up -d
stop-kafka:
	docker-compose down
run:
	go run .
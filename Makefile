default: mongo-start

mongo-start: mongo-stop
	cd test && docker-compose up -d

mongo-stop:
	cd test && docker-compose down

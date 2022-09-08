docker-image-build:
	@ docker build -t nat/meuapp:1.0 .

run: docker-image-build
	@ docker-compose up 

stop:
	@ docker-compose down -v
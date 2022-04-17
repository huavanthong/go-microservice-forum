swagger: 
	swagger.bat generate spec -o ./swagger.yaml --scan-models

docker-build:
	docker build -t product-api .

docker-run:
	docker run -p 8080:8080 -it product-api
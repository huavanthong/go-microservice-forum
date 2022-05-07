swagger: 
	swagger.bat generate spec -o ./swagger.yaml --scan-models

docker-build:
	docker build -t product-api .

docker-run:
	docker run -p 9090:9090 -it product-api 

generate_client:
	cd sdk && swagger.bat generate client -f ../swagger.yaml -A product-api
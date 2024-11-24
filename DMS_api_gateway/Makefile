build-image:
	docker build . -t dms_api_gateway:1.0.0

start-container:
	docker run --name dms_api_gateway_container -p 80:8080  -it  dms_api_gateway:1.0.0
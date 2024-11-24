build-image:
	docker build . -t authorization_api:1.0.0

start-container:
	docker run --name auth_container -p 4040:4040  -it  authorization_api:1.0.0
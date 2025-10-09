gotenberg:
	docker run \
		-p 3000:3000 \
		--name gotenberg \
		--add-host="host.docker.internal:host-gateway" \
		gotenberg/gotenberg:8

gotenberg-start:
	docker start gotenberg

gotenberg-stop:
	docker stop gotenberg
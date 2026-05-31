gotenberg-run:
	docker run \
		-p 3000:3000 \
		--name gotenberg \
		--add-host="host.docker.internal:host-gateway" \
		gotenberg/gotenberg:8

gotenberg:
	docker start gotenberg
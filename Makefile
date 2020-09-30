SERVICE_NAME=live-reload

image:
	docker build -t docker.io/goforbroke1006/live-reload:latest ./
	docker push docker.io/goforbroke1006/live-reload:latest

GOOSes = solaris linux windows darwin freebsd
GOARCHes = amd64 386

release:
	rm -f ./build/*
	$(foreach goos,$(GOOSes), \
		$(foreach goarch,$(GOARCHes), \
			GOOS=${goos} GOARCH=${goarch} go build -o ./build/${SERVICE_NAME}_${goos}_${goarch} ./ || true ; \
		) \
	)
	#zip -mr ./build/release-`date +"%y-%m-%d-%H-%M-%S"`.zip ./build/*

ll: fmt build

build:
	gb build

fmt:
	@gofmt -w ./

clean:
	rm -fr target bin pkg

docker: fmt
	docker run  \
		-e "TARGETS=linux/amd64" \
		-v `pwd`:/build quay.io/opsee/build-go \
		&& docker build -t dis/dis .

run: docker
	docker run \
		-e SLACK_URL=$(SLACK_TOKEN) \
		-p 9009:9009 \
		--rm \
		dis/dis

.PHONY: docker run clean all

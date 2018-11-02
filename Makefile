USERNAME	:=	95af6355-a497-4702-93cb-23c88418bb68
TEMPLATE	:=	.ild.1\d5
repository	:=	bitbucket.org/242617/blockchain.com-bruteforce
image		:=	blockchain.com-bruteforce
container	:=	242617/${image}

windows: windows-setup build
macos: macos-setup build
linux: linux-setup build

test:
	# go test ${repository}

windows-setup:
	$(eval goos=windows)
	$(eval goarch=amd64)
	$(eval output=bruteforce.exe)

macos-setup:
	$(eval goos=darwin)
	$(eval goarch=amd64)
	$(eval output=bruteforce)

linux-setup:
	$(eval goos=linux)
	$(eval goarch=amd64)
	$(eval output=bruteforce)

build: test
	GOOS=${goos} GOARCH=${goarch} \
	go build \
		-o build/${output} \
		-ldflags " \
			-X main.username=${USERNAME} \
			-X main.template=${TEMPLATE} \
		" \
		${repository}

docker:
	docker build -t ${container} .
	docker run \
		-it --rm \
		--name ${image} \
		-v `pwd`:/build \
		${container}
	docker push ${container}
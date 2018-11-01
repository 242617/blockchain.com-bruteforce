USERNAME	:=	login
repository	:=	bitbucket.org/242617/blockchain.com-bruteforce
image		:=	blockchain.com-bruteforce
container	:=	242617/${image}

all: windows macos linux
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
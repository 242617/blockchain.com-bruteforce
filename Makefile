USERNAME	:=	ogin
TEMPLATE	:=	templat[a-z]
repository	:=	bitbucket.org/242617/blockchain.com-bruteforce

windows: windows-setup build
mac: mac-setup build
linux: linux-setup build

test:
	# go test ${repository}

windows-setup:
	$(eval goos=windows)
	$(eval goarch=amd64)
	$(eval output=bruteforce.exe)

mac-setup:
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
		-o ${output} \
		-ldflags " \
			-X main.username=${USERNAME} \
			-X main.template=${TEMPLATE} \
		" \
		${repository}

repository	:=	github.com/242617/blockchain.com-bruteforce

windows: windows-setup build
macos: macos-setup build
linux: linux-setup build

list-windows: windows-setup list
list-macos: macos-setup list
list-linux: linux-setup list

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
		${repository}
clean:
	clear

build: clean
	go build \
		-o bruteforce \
		-ldflags " \
			-X main.login=${LOGIN} \
			-X main.template=passwor[a-z] \
		" \
		bitbucket.org/242617/blockchain.com-bruteforce

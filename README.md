# blockchain.com bruteforce

Application for bruteforcing you lost password (assuming you remember some part of your password).

This will open headless Chrome and will make appempts for log you in using provided `username` and `password` template.

All output is written into `bruteforce.log` file.

## Build

Download
```
go get -u github.com/242617/blockchain.com-bruteforce
go build \
    -o bruteforce \
    github.com/242617/blockchain.com-bruteforce
```
You can use Makefile to build for differenct platforms.
```
make macos // for MacOS binary
make linux // for Linux binary
make windows // for Windows binary
```


## Run

* `username` - blockchain.com wallet id
* `password_mask` - password mask (regexp: e. g. `(a|b)`, `\d`, `\S`, `(A|z){2,3}`.)
* `resume_from` - word to resume from

```
./bruteforce \
    -username={username} \
    -password="{password_mask}" \
    -resume={resume_from}
```

### List mode

This mode is to list all combinations for testing your password mask. `username` is unnecessary in this case.
```
./bruteforce \
    -password="{password_mask}" \
    -resume={resume_from}  \
    -list
```
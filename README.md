# blockchain.com bruteforce
You can try to restore you lost password if you remember some of your password letters.

This will open headless Chrome and will make appempts for log you in using provided `username` and `password` template.

## Build
Download
```
go get -u github.com/242617/blockchain.com-bruteforce
```
Use `go build`
```
go build \
    -o bruteforce \
    github.com/242617/blockchain.com-bruteforce
```
Or `make`
```
make macos // for MacOS binary
make linux // for Linux binary
make windows // for Windows binary
```


## Run
* `username` - blockchain.com wallet id
* `password` - password mask (regexp: e. g. `(a|b)`, `\d`, `\S`, `(A|z){2,3}`.)
* `resume` - password to resume from

```
./bruteforce \
    -username={username} \
    -password="{password_mask}" \
    -resume={resume_from}
```

### List mode
This mode is to list all combinations for testing your password mask.
```
./bruteforce \
    -password="{password_mask}" \
    -resume={resume_from}  \
    -list
```
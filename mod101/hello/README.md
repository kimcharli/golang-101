

- https://golang.org/doc/code.html
- https://golang.org/pkg/testing/

```
ckim-mbp:hello ckim$ go mod init example.com/user/hello
go: creating new go.mod: module example.com/user/hello
ckim-mbp:hello ckim$ vi hello.go

ckim-mbp:hello ckim$ 
ckim-mbp:hello ckim$ go install example.com/user/hello
ckim-mbp:hello ckim$ ll ~/go/bin/hello 
-rwxr-xr-x  1 ckim  staff  2178184 Aug  3 11:44 /Users/ckim/go/bin/hello*
ckim-mbp:hello ckim$ 
```

```
ckim-mbp:hello ckim$ ~/go/bin/hello 
Hello, world.
ckim-mbp:hello ckim$ 
```


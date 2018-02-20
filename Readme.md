# gpurl

`gpurl` is small shell utility to extract parts from a URL. It uses go standard library URL parser to parse URLs. `gpurl` is written as a UNIX cli tool, you can pass the URLs as arguments separated by space or piped into stdin

## Installing
Get it with go get

```bash
go get github.com/dbalan/gpurl
```

This will create the gpurl executable under your `$GOPATH/bin` directory.

## Example

```bash
% gpurl -p host http://dbalan.in/hello
dbalan.in
% gpurl -p path http://dbalan.in/hello
/hello

% echo http://dbalan.in/hello | gpurl -p path -e
/hello
```

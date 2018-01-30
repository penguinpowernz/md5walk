# MD5 Walk

This is a small command to walk through a directory recursively and generate MD5 sums for
all files found.  Like the `md5sum` command on linux, but recursive.

## Installation

    go get github.com/penguinpowernz/md5walk
    go install github.com/penguinpowernz/md5walk/cmd/md5walk

## Usage

Once you have the CLI built, run it like this:

    $ md5walk .
    38334fb8161eeef13e59e4f72315ff45  README.md
    8624efb023799899a5f2bdb059d97e47  cmd/md5walk/main.go
    bfdcc4ccb7dc1ce4ede190e4ddc4e734  walk.go

If you want to use it in your code, it works like this:

```go
import "github.com/penguinpowernz/md5walk"

sums, err := md5walk.Walk("/path/to/dir")
if err != nil {
    panic(err)
}

for filename, hash := range sum {
    fmt.Println(filename, hash)
}
```

Easy!
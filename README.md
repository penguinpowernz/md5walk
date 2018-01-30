# MD5 Walk

This is a small command to walk through a directory recursively and generate MD5 sums for
all files found.  Like the `md5sum` command on linux, but recursive.

## Installation

    go get github.com/penguinpowernz/md5walk
    go install github.com/penguinpowernz/md5walk/cmd/md5walk

## Usage

Once you have the CLI built, run it like this:

    $ md5walk .
    4cf2d64e44205fe628ddd534e1151b58  .git/HEAD
    5b603c2c0801a9ded3b79159fc38b404  .git/config
    a0a7c3fff21f2aea3cfa1d0316dd816c  .git/description
    ce562e08d8098926a3862fc6e7905199  .git/hooks/applypatch-msg.sample
    579a3c1e12a1e74a98169175fb913012  .git/hooks/commit-msg.sample
    2b7ea5cee3c49ff53d41e00785eb974c  .git/hooks/post-update.sample
    054f9ffb8bfe04a599751cc757226dda  .git/hooks/pre-applypatch.sample
    01b1688f97f94776baae85d77b06048b  .git/hooks/pre-commit.sample
    3c5989301dd4b949dfa1f43738a22819  .git/hooks/pre-push.sample
    3ff6ba9cf6d8e5332978e057559b5562  .git/hooks/pre-rebase.sample
    7dfe15854212a30f346da5255c1d794b  .git/hooks/prepare-commit-msg.sample
    517f14b9239689dff8bda3022ebd9004  .git/hooks/update.sample
    036208b4a1ab4a235d75c181e685e5a3  .git/info/exclude
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
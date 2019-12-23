# jc-project

## Using this library
This code structure assumes that the library is part of a bigger mono-repo style codebase. It is also assumed that the `action`
package is a public library within that repo.

### Get to your GOPATH
- If you have a `GOPATH` defined, `cd $GOPATH`
- If not, `cd $HOME/go`

### Get and build the code
From your `GOPATH` run the following in order to build, test, and run a test harness around the package:
```bash
mkdir -p src/github.com/amazing0x41
cd src/github.com/amazing0x41
git clone https://github.com/aMaZing0x41/jc-project.git
cd jc-project
./build.sh
```

## TODO
- document functions
- implement all TODOs
- write benchmark tests
- make code concurrent
- review requirements
- reivew naming and give suggestions for idomatic go
- how to build/test

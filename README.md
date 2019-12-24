# jc-project

## Using this library
This code structure assumes that the library is part of a bigger mono-repo style codebase. It is also assumed that the `action`
package is a public library within that repo.

### Environment
This package should build and run on OSX (darwin), Linux, and Windows. The lastest version of go (Go 1.13 as of Dec 2109) must be installed and setup on the host.
Follow the instructions [here](https://golang.org/dl/) to get up-and-running.

__Note: These instructions are optomized for a linux or OSX environment. If using windows, further support is needed. Continue at your own risk!__

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

### Test Harness
At `cmd/main.go` there is a test harness that imports the action package and calls the add/get functions concurrently.
The `NUM_ITERS` const defines how many iterations to run through in the main loop. Currently this needs to be updated and the
code rebuilt.

### Benchmark Tests
The 

## Continuous Integration
This repo is set up with a github action that runs on every push. The action runs the `build.sh` script. If this fails, merges will not be allowed. 

_FUTURE: May want to consider code coverage and possibly performance benchmarks as future actions._

### Benchmark Tests
The package tests contain some very simple benchmark tests. There is definitely room for improvement here, but the intial results show what is expected.

Run the following to invoke the benchmark tests (pwd needs to be the action folder):
```
cd pkg/action
go test -run=XXX -bench=.
```

Here is some sample output from one of the test runs:

```
goos: darwin
goarch: amd64
pkg: github.com/amazing0x41/jc-project/pkg/action
BenchmarkAddAction1-4         	  103650	     10915 ns/op
BenchmarkAddAction10-4        	   11061	    107365 ns/op
BenchmarkAddAction100-4       	    1086	   1034081 ns/op
BenchmarkAddAction1000-4      	     108	  10733620 ns/op
BenchmarkAddAction100000-4    	       1	1190839328 ns/op
BenchmarkAddAction1000000-4   	       1	10767939648 ns/op
BenchmarkGetStats1-4          	  100886	     13076 ns/op
BenchmarkGetStats10-4         	   94556	     12063 ns/op
BenchmarkGetStats100-4        	   93639	     12862 ns/op
BenchmarkGetStats1000-4       	   84751	     14230 ns/op
BenchmarkGetStats100000-4     	       1	1519720700 ns/op
BenchmarkGetStats1000000-4    	       1	13268265757 ns/op
PASS
ok  	github.com/amazing0x41/jc-project/pkg/action	39.451s
```
 
 These results are a little tricky to read. The numbers after the names, e.g. 1, 10, 1000, are how many operations are performed on the actions map. For the `GetStats` that number should be roughly the number of itmes in the map. The `GetStats` timings also take into account adding the data to the map. This is why things start to fall off in the two largest cases. Dave Cheney has a great write up on benchmark tests in go [here](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go).

 Although maps don't explicitly provide any guarantess in go, we can see that the `AddAction` benchmarks exhibit approximately O(N) (maybe O(N log N)) time complexity where N is the number of adds to the map and the `GetStats` is roughly constant until the size of the map overtakes the timing (cases 100K and 1M).

 More time could definitely be spent to improve the overall tests and tune the results so that identifying any sort of regression is more apparent.


## Assumptions
- Requirements ask for `addAction` and `getStats` signatures - because of how go exports the names are `AddAction` and `GetStats`
- `GetStats` return data does not need to be sorted in any fashion or "pretty printed". _Note: sorting the final array is not difficult, just saving some CPU cycles._


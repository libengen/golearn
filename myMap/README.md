## benchmark 
```
$go test -bench=.
goos: darwin
goarch: amd64
pkg: howework/lesson3/group8/ben
BenchmarkMyMap_All-4      	 2270334	       528 ns/op
BenchmarkMyMap_All2-4     	 2017182	       587 ns/op
BenchmarkMyMap_Store-4    	 2525980	       462 ns/op
BenchmarkMyMap_Store2-4   	 2051406	       577 ns/op
BenchmarkMyMap_Load-4     	10178641	       116 ns/op
BenchmarkMyMap_Load2-4    	24184500	        48.1 ns/op
PASS
ok  	howework/lesson3/group8/ben	9.904s
```

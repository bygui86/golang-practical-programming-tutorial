
# GoLang tutorials

If you are creating a Go application that will not have subpackages, you can develop the code anywhere on your filesystem. However, since this application (and most other apps you will develop) will use subpackages, your code must live in your GOPATH folder.

## Topics

### Basic

* hello world
* variables
* constants
* functions
* pointers
* struct
* null check
* methods
  * pointer receiver
  * value receiver
* init method
* packages
* multiple-return
* nested-import
* types
* type-defaults
* type-conversion
* float
* math
* if-else
* switch
* loops
* data structures
  * array
  * slice
  * map
  * range
  * list
* print
* parsing
  * number
  * xml (web-server_v4.go, xml-parsing_v1.go, xml-parsing_v2.go, xml-parsing_v3.go)
  * html (web-server_v3.go)
  * json
  * url
* web
  * server
  * client
* sleep
* shift
* defer
* panic
* file
  * reading
  * writing
* environment variable
  * get
  * set
* logs
  * normal
  * level + format
  * to file
* command line
  * arguments
  * flags
* lambda
* closures
* variadic functions
* interface
* error
* time
  * formatting
  * parsing
* epoch
* alias
* strings
  * formatting
  * functions
* regular expressions
* sha1 hashes
* base64 encoding
* sorting
* sorting by functions
* recursion
* exit
* casting
* reflection

### Advanced

* (unix) signals
* timer & ticker
* dynamic-type
* type-assertion - `IN PROGRESS`
* goroutines - `IN PROGRESS`

---

## Open topics

### Basic

:white_check_mark: `DONE`

### Advanced

* goroutines   https://medium.com/@riteeksrivastava/a-complete-journey-with-goroutines-8472630c7f5c
* goroutine pooling
* channels   https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8
* channel buffering
* channel synchronization
* channel directions
* select
* timeouts
* non-blocking channel operations
* closing channels
* range over channels
* worker pools
* rate limiting
* atomic counters
* mutex
* stateful goroutines
* spawning processes
* exec'ing processes

### News

* modules (try `go help modules`)

## Specific tecnologies

* db
  * `[OK]` mysql > see [this repo](https://github.com/bygui86/go-todo-rest-api-example)
  * `[OK]` cassandra > see [this repo](https://github.com/bygui86/go-rest-cassandra)
  * `[TODO]` mongodb
  * `[TODO]` redis
* protocols
  * `[OK]` http/rest > see [this repo](https://github.com/bygui86/go-rest-cassandra) and [this repo](https://github.com/bygui86/go-service)
  * `[OK]` grpc > see [this repo](https://github.com/bygui86/grpc-samples)
  * `[TODO]` websocket
* docker
  * `[OK]` dockerfile > see [this repo](https://github.com/bygui86/go-kafka/blob/master/consumer/Dockerfile)
* others
  * `[OK]` protobuf > see [this repo](https://github.com/bygui86/go-protobuf)
  * `[OK]` kafka > see [this repo](https://github.com/bygui86/go-kafka)
* testing
  * `[TODO]` principles
  * `[TODO]` moking   https://godoc.org/google.golang.org/grpc/examples/helloworld/mock_helloworld
  * `[TODO]` tools

### Tools

* Govendor   https://github.com/kardianos/govendor

---

## Links

### Tutorials
* https://pythonprogramming.net/go/introduction-go-language-programming-tutorial/ - `IN PROGRESS`
	[next: Goroutines - Concurrency in Goprogramming]
* https://gobyexample.com/ - `IN PROGRESS`
	[next: Goroutines]
* https://icyapril.com/go/programming/2017/12/17/object-orientation-in-go.html - `TODO`
* https://www.golang-book.com/books/intro - `TODO`

### Govendor
    https://github.com/kardianos/govendor
    https://blog.gopheracademy.com/advent-2015/vendor-folder/
    https://github.com/kardianos/spider

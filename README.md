
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
  * concatenation
* regular expressions
* sha1 hashes
* base64 encoding
* sorting
* sorting by functions
* recursion
* exit
* casting
* reflection
* iota

### Advanced
* (unix) signals
* timer & ticker
* dynamic-type
* type-assertion
* enum
* goroutines
  * pooling
  * rate limiting
  * atomic counters
  * mutex
  * stateful
* channels
  * directions
  * buffering
  * synchronization
  * select
  * timeouts
  * non-blocking operations
  * closing
  * range over
* spawning processes
* exec'ing processes

---

## Open topics

### Advanced
* wait-groups
  * https://tutorialedge.net/golang/go-waitgroup-tutorial/
* modules (try `go help modules`)
  * https://github.com/golang/go/wiki/Modules
  * https://blog.golang.org/using-go-modules
  * https://roberto.selbach.ca/intro-to-go-modules/

## Specific tecnologies
* db
  - [x] mysql > see [this repo](https://github.com/bygui86/go-todo-rest-api-example)
  - [x] cassandra > see [this repo](https://github.com/bygui86/go-rest-cassandra)
  - [ ] mongodb
  - [ ] redis
  - [ ] solr
* protocols
  - [x] http/rest > see [this repo](https://github.com/bygui86/go-rest-cassandra) and [this repo](https://github.com/bygui86/go-service)
  - [x] grpc > see [this repo](https://github.com/bygui86/grpc-samples)
  - [ ] websocket > see [this repo](https://github.com/bygui86/websocket-samples)
* docker
  - [x] dockerfile > see [this repo](https://github.com/bygui86/go-kafka/blob/master/consumer/Dockerfile)
* kubernetes
  - [ ] prometheus metrics
  - [ ] jaeger tracing
* others
  - [x] protobuf > see [this repo](https://github.com/bygui86/go-protobuf)
  - [x] kafka > see [this repo](https://github.com/bygui86/go-kafka)
* testing
  - [ ] principles
  - [ ] moking
  - [ ] tools

### Tools
- [x] Govendor
- [ ] Go kit
  * https://dev.to/napolux/how-to-write-a-microservice-in-go-with-go-kit-a66
  * https://medium.com/@shijuvar/go-microservices-with-go-kit-introduction-43a757398183
- [ ] Go Echo
  * https://echo.labstack.com/
  * https://github.com/labstack/echo

---

## Links

### Articles
#### Goroutines
* https://medium.com/@riteeksrivastava/a-complete-journey-with-goroutines-8472630c7f5c
#### Channels
* https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/
* https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8
### Environment variables
* https://stackoverflow.com/questions/40326540/how-to-assign-default-value-if-env-var-is-empty
* https://codereview.stackexchange.com/questions/108563/reading-environment-variables-of-various-types

### Tutorials
* https://gobyexample.com/
* https://pythonprogramming.net/go/introduction-go-language-programming-tutorial/ - `IN PROGRESS`
	[next: Goroutines - Concurrency in Goprogramming / https://pythonprogramming.net/go/goroutines-go-language-programming-tutorial/]
* https://tour.golang.org/ - `IN PROGRESS`
	[next: https://tour.golang.org/methods/2]
* https://icyapril.com/go/programming/2017/12/17/object-orientation-in-go.html
* https://www.golang-book.com/books/intro
* https://dlintw.github.io/gobyexample/public/index.html#by-sequence
#### Testing
* https://godoc.org/google.golang.org/grpc/examples/helloworld/mock_helloworld

### Govendor
* https://github.com/kardianos/govendor
* https://blog.gopheracademy.com/advent-2015/vendor-folder/
* https://github.com/kardianos/spider

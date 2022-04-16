
# Chapter 4  
This chapter will help you understand that how can you test something such as Backend API?  

# Introduction
When you try to define what testing is, you will come up with a multitude of answers, and many of us will not understand the 
full benefits of testing until we've been burnt by buggy software or we have tried to change a complex code base which has no tests.
```
"The art of a good night's sleep is knowing you will not get woken by a support call and the piece of mind from being able to
confidently change your software in an always moving market."
=> Nghĩa là:
"Nghệ thuật của một giấc ngủ ngon là biết rằng bạn sẽ không bị đánh thức bởi một cuộc gọi hỗ trợ và phần tâm trí không thể tự tin 
thay đổi phần mềm của bạn trong một thị trường luôn chuyển động." 
```
OK, so I am trying to be funny, but the concept is correct:  
- Nobody enjoys debugging poorly written code, and indeed, 
- Nobody enjoys the stress caused when a system fails. 
- Starting out with a mantra of quality first can alleviate many of these problems.

# Table of contents
* [The testing pyramid](#the-testing-pyramid) 
* [Outside-in development](#Outside-in-development) 
* [Unit test](#Unit-test) 
* [Dependency injection and mocking](#dependency-injection-and-mocking)
* [Code coverage](#Code-coverage)
* [Behavioral Driven Development](#behavirol-driven-development)
* [Testing with Docker compose](#Testing-Docker-compose)
* [Benchmarking and profiling](#benchmarking-and-profiling)
* [How to run this project](#prepare-environment)

# Questions
## About the testing pyramid
* [What is model for testing in Web application? What purpose for each component?](#model-pyramid)
* [What is automated testing?](#Automated-testing)
* [What is problem if we use traditional test for a web-app?](#problem)
## About Outside-in development
* [What is outside-in development process?](#outside-in-development)

## About Unit Test
* [What the laws for you when to implement Unit test?](#law)
* [What ideas do you need to test a web app?](#ideas)
* [What is project structure for implementing Unit Test?](#project-structure)
* [What rules for naming test case for Unit Test?](#naming-for-test-case)

## About Dependency injection and mocking
* [Why we need to manage the dependencies on our handler?](#dependency-injection-and-mocking)
* [What is mocking?](#what-is-mocking)
* [How do you design a mocking method?](#how-to-implement-mocking)
* [How do you implement Unit Test for Data Store?](#unit-test-for-data-store)
* [What is Testify package? Purpose for using this package?](#testify-package)
* [How do you design testify package for testing data store?](#how-to-use-testify-package)

## About Code Coverage
* [How to run test coverage for this project?](#code-coverage)

## About Behavirol Driven Development
* [What is BBD?](#what-is-bbd)
* [How do you run Cucumber on this project?](#how-to-run-cucumber)

## About Testing with Docker compose
* [What is Docker compose?](#docker-compose)
* [Why this project can't see Dockerfile?](#why-it-lack-of-dockerfile)
* [How to run docker compose?](#run-with-docker-compose)
* [Do you understand about work-flow happened if you run command docker-compose up or not?](#run-with-docker-compose)

## About Benchmarking and profiling
* [What is benchmark tool?](#benchmarking-and-profiling)
* [How do you make sure the result measurement is correct?](#how-to-make-sure-that-the-result-of-measurement-is-correct)
* [How do you run benchmark tool with this project?](#benchmarks)

**********************************************************************************************************************************
## The testing pyramid
### Model pyramid
> **************************  
> *********   UI   *********  
> *****     Service    *****  
> ***        Unit        ***  
> **************************  
* Unit test: Implement to detect error on your design.
* Service: where you define a flow service in your application.
* UI: where user will use your applicaiton, and detect error.

### Automated testing
In the early days of automated testing, all the testing was completed at the top of the pyramid. While this did work from a quality perspective, it meant the process of debugging the area at fault would be incredibly complicated and time-consuming.
#### Problem
* If you were lucky, there might be a complete failure which could be tracked down to a stack trace. 
* If you were unlucky, then the problem would be behavioral; and even if you knew the system inside out, it would involve plowing through thousands of lines of code and manually repeating the action to reproduce the failure.
## Outside-in development
```
When writing tests, I like to follow a process called outside-in development. 

With outside-in development, you start by writing your tests almost at the top of the pyramid, determine what the functionality is going to be for the story you are working on, and then write some failing test for this story. 

Then you work on implementing the unit tests and code which starts to get the various steps in the behavioral tests to pass.
```
More details: [here](https://www.meisternote.com/app/note/cH2i2kXrT9kA/outside-in-development)
## Unit-test
Our unit tests go right down to the bottom of the pyramid. 
### Law
* **First law:** You may not write production code until you have written a failing unit test
* **Second law:**  You may not write more of a unit test than is sufficient to fail, and not compiling is failing
* **Third law:**  You may not write more production code than is sufficient to pass the currently failing test

**Note:**
- Production code: Production means anything that you need to work reliably, and consistently.  
Refer: [here](https://stackoverflow.com/questions/490289/what-exactly-defines-production)
### Ideas 
One of the most effective ways to test a microservice in Go is not to fall into the trap of trying to execute all the tests through the HTTP interface.  
**Follow steps:**
**Step 1:** Create a pattern for test program  
    - We need develop **a pattern that avoids** creating a physical web server for testing our handlers, the code to create this kind of test is slow to run and incredibly tedious to write.
**Step 2:** Implement Unit test
    - What need to be doing is to test our handlers and the code within them as **Unit test**. 
    - These tests will run far quicker than testing through the web server.
**Step 3:** Get coverage.
    - And if we think about coverage, we will be able to test the writing of the handlers in the **Cucumber** tests that execute a request to the running server which overall gives us 100% coverage of our code.
### Project Structure
```
chapter4
| - main.go
| --- handlers                          # Split the handlers our into a seperate package.
| --- | - search.go                     # Our handlers
| --- | - search_test.go                # Unit test for our handlers
```

At main.go
```
	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
```

The signature for a test method looks like this:
```
func TestXxx(*testing.T)
```
### Naming for test case
The name of the test must have a particular name beginning with Test and then immediately following this an uppercase
character or number.  
For a example:
- Do not: TestmyHandler
- Should: Test1Handler
- Should: TestMyHandler
- Recommend: Test1MyHandler
- Recommend: TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent

## Dependency injection and mocking
To get the tests that return items from the Search handler to pass, we are going to need a data store. Whether we implement our
data store in a database or a simple in-memory store we do not want to run our tests against the actual data store as we will be
checking both data store and our handler. For this reason, we are going to need to manage the dependencies on our handler so
that we can replace them in our tests. To do this, we are going to use a technique called dependency injection where we will
pass our dependencies into our handler rather than creating them internally.
### What is Mocking?
This method allows us to replace these dependencies with stubs or mocks when we are testing the handler, making it possible to
control the behavior of the dependency and check how the calling code responds to this.

### How to implement mocking
Before we do anything, we need to create our dependency. In our simple example, we are going to create an in-memory data store which has a single method:
```
Search(string) []Kitten
```
To replace the type with a mock, we need to change our handler to depend on an interface which represents our data store. We
can then interchange this with either an actual data store or a mock instance of the store without needing to change the
underlying code:
```
type Store interface {
	Search(name string) []Kitten
}
```
We can now go ahead and create the implementation for this. Since this is a simple example, we are going to hardcode our list of kittens as a slice and the search method will just select from this slice when the criteria given as a parameter matches the name of the kitten.
```
Search {
    Store data.Store
}
```
More details: 
- Create mock for Search => [here](https://github.com/huavanthong/build-microservice-golang/blob/master/01_GettingStarted/book-build-microservice/chapter4/data/datastore.go)
### Unit Test for Data Store
Now, back to our unit tests: we would like to ensure that, when we call the ServeHTTP method with a search string, we are
querying the data store and returning the kittens from it.  
#### Testify package
To do this, we are going to create a mock instance of our data store. We could create the mock ourselves; however, there is an
excellent package by Matt Ryer who incidentally is also a Packt author. 
* Testify (https://github.com/stretchr/testify.git) has a fully featured mocking framework with assertions. 
* It also has an excellent package for testing the equality of objects in
our tests and removes quite a lot of the boilerplate code we have to write.
#### How to use Testify package
In the data package, we are going to create a new file called **mockstore.go**. This structure will be our mock implementation of
the data store:
More details: [here](https://www.meisternote.com/app/note/FkfCXX4bfrLQ/dependency-injection-and-mocking)

## Code-coverage
To run test coverage.
```
go test -cover ./...
```
## Behavirol Driven Development
### What is BBD?
* Behavioral Driven Development (BDD) and is a technique often executed by an application framework called Cucumber.
* It was developed by Dan North and was designed to create a common ground between developers and product owners.

## Testing Docker compose
If you investigate this project, you will know that this project don't has any Dockerfile for this chapter 4.
The question, we need to anwer is why this project don't need it?
### Why it lack of Dockerfile

### Docker compose
docker-compose.yml
```
version: '2'
services:
  mongodb:
    image: mongo
    ports:
      - 27017:27017

```
## Benchmarking and profiling
Benchmarking is a way of measuring the performance of your code by executing it multiple times with a fixed workload. 

### How to make sure that the result of measurement is correct
When running a benchmark, Go needs to run it multiple times to get an accurate reason.  
Therefore, we need to loop your service.
```
  for n := 0; n < b.N; n++
```

**********************************************************************************************************************************
# Prepare environment
If you want to run this project, please prepare tools following steps below:
## Golang
To install Golang
```

```
Check GOPATH is existent on your ennvironment
```
$echo %GOPATH%
C:\Go\bin
```
## Make
Install Make for Windows.
```
winget install gnuwin32.make
```

And remember to set PATH for your environment
```
path to: C:\Program Files (x86)\GnuWin32\bin
```
Refer: [here](https://www.technewstoday.com/install-and-use-make-in-windows/)
## MGO
To install mgo
```
go get gopkg.in/mgo.v2 
```

Add library to go.mod
```
require(
    gopkg.in/mgo.v2
)
```
More details: [here](https://github.com/tobyzxj/mgo)
## Cucumber
### Install Godog
To install Cucumber (godog) package
```
go get github.com/DATA-DOG/godog/cmd/godog
```

Add library to go.mod
```
require(
    gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
    github.com/cucumber/godog v0.12.4
)
```
### Set environment
Export PATH point to GOLANG environment
```
#=============== Linux ===================#
# the executable is here after installation
# $GOPATH/bin/godog
export PATH=$PATH:$GOPATH

#=============== Window ===================#
set PATH=%PATH%;%GOPATH%

```
### Folder contain godog library
To check godog.exe is exist in your environment
```
# godog.exe is in directory C:\Go\bin
C:\Go\bin\godog.exe
```
More details: [here](https://techblog.fexcofts.com/2019/08/09/go-and-test-cucumber/)


### Issue knowledge when download
#### Issue 1: wrong link to download
> go get github.com/DATA-DOG/godog/cmd/godog
```
module declares its path as: github.com/cucumber/godog
                but was required as: github.com/DATA-DOG/godog
```
To fix it.
> go get github.com/cucumber/godog/cmd/godog

More details: [here](https://github.com/cucumber/godog/issues/211)
#### Issue 2: Wrong path to working godog

**Behavior:** When you run command cucumber in Makefile
```
make cucumber
```

**Error:** You will see error below:
```
- Container chapter4-mongodb-1  Running                                                                                                                                                                                                0.0s
cd features && godog ./
failed to compile tested package: C:\01_Work\01_Programming\build-microservice-golang\01_GettingStarted\book-build-microservice\chapter4\features, reason: exit status 1, output: WORK=C:\Users\huava\AppData\Local\Temp\go-build013581706
# docker-compose/features
search_test.go:14:2: import "github.com/cucumber/godog/cmd/godog" is a program, not an importable package
FAIL    docker-compose/features [setup failed]

make: *** [cucumber] Error 1
```
**Solution:**
```
Change command from Makefile.
From:
    cd features && godog ./
To: 
    cd godog ./
```
### Run godog
To run this godog for this project
```
godog ./
```

To run with docker-compose
```
make cucumber
```
## Run with Docker compose
When we run docker-compose up command, we will download the image of MongoDB and start an instance exposing these ports on our local host.  
To run it.
```
make cucumber
```

Output
```
chapter4> make cucumber
docker-compose up -d
[+] Running 1/0
 - Container chapter4-mongodb-1  Running                                                                                                                                                                        0.0s 
cd features && godog ./
Feature: As a user when I call the search endpoint, I would like to receive a list of kittens
Server running with pid: 26688
  Scenario: Invalid query                       # search.feature:4
    Given I have no search criteria             # search_test.go:33 -> iHaveNoSearchCriteria
    When I call the search endpoint             # search_test.go:41 -> iCallTheSearchEndpoint
    Then I should receive a bad request message # search_test.go:52 -> iShouldReceiveABadRequestMessage
Server running with pid: 7868
  Scenario: Valid query                     # search.feature:9
    Given I have a valid search criteria    # search_test.go:60 -> iHaveAValidSearchCriteria
    When I call the search endpoint         # search_test.go:41 -> iCallTheSearchEndpoint
    Then I should receive a list of kittens # search_test.go:66 -> iShouldReceiveAListOfKittens

2 scenarios (2 passed)
6 steps (6 passed)
9.205181s
docker-compose stop
[+] Running 1/1
 - Container chapter4-mongodb-1  Stopped
```
### Start server
To start this server with docker-compose.
```
chapter4>make run
docker-compose up -d
[+] Running 1/0
 - Container chapter4-mongodb-1  Running                                                                                                                                                      0.0s 
go run main.go
```
### How to run


### Issue knowledge
#### Issue 1: Exists port on your machine
**Behavior:** When you run command cucumber in Makefile
```
make run
```

**Error:** You will see error below:
```
chapter4> make run
docker-compose up -d
[+] Running 1/1
 - Container chapter4-mongodb-1  Started                                                                                                                                                      3.3s 
go run main.go
2022/03/07 15:18:07 listen tcp :8323: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
exit status 1
make: *** [run] Error 1
```

**Solution:** Find the exist port and kill it
```
chapter4>netstat -ano | findstr :8323
  TCP    0.0.0.0:8323           0.0.0.0:0              LISTENING       276
  TCP    [::]:8323              [::]:0                 LISTENING       276

chapter4>taskkill /F /PID 276
SUCCESS: The process with PID 276 has been terminated.
```
To see the setting port. Refer: [here](https://github.com/huavanthong/build-microservice-golang/blob/docker-postgre/01_GettingStarted/book-build-microservice/chapter4/main.go)

## Testing
### Unit Test
To run Unit Test
```
make unit
```

Output
```
go test -race -v ./...
testing: warning: no tests to run
PASS
ok      docker-compose  (cached) [no tests to run]
?       docker-compose/benchmark        [no test files]
# docker-compose/benchmark/loadtest
benchmark\loadtest\main.go:29:16: not enough arguments in call to bench.New
        have (number, time.Duration, time.Duration, time.Duration)
        want (bool, int, time.Duration, time.Duration, time.Duration)
note: module requires Go 1.16
=== RUN   TestReturns1KittenWhenSearchGarfield
--- PASS: TestReturns1KittenWhenSearchGarfield (0.00s)
=== RUN   TestReturns0KittenWhenSearchTom
--- PASS: TestReturns0KittenWhenSearchTom (0.00s)
PASS
ok      docker-compose/data     (cached)

No scenarios
No steps
0s
ok      docker-compose/features 1.578s
=== RUN   TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent
--- PASS: TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent (0.00s)
=== RUN   TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent
--- PASS: TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent (0.00s)
=== RUN   TestSearchHandlerCallsDataStoreWithValidQuery
    search_test.go:42: PASS:    Search(string)
--- PASS: TestSearchHandlerCallsDataStoreWithValidQuery (0.00s)
=== RUN   TestSearchHandlerReturnsKittensWithValidQuery
--- PASS: TestSearchHandlerReturnsKittensWithValidQuery (0.00s)
PASS
ok      docker-compose/handlers (cached)
make: *** [unit] Error 2
```

### Cucumber Test
To run Cucumber testing
```
make test
```

Output
```
go test -race -v ./...
# docker-compose/benchmark/loadtest
benchmark\loadtest\main.go:29:16: not enough arguments in call to bench.New
        have (number, time.Duration, time.Duration, time.Duration)
        want (bool, int, time.Duration, time.Duration, time.Duration)
note: module requires Go 1.16
testing: warning: no tests to run
PASS
ok      docker-compose  1.636s [no tests to run]
?       docker-compose/benchmark        [no test files]
=== RUN   TestReturns1KittenWhenSearchGarfield
--- PASS: TestReturns1KittenWhenSearchGarfield (0.00s)
=== RUN   TestReturns0KittenWhenSearchTom
--- PASS: TestReturns0KittenWhenSearchTom (0.00s)
PASS
ok      docker-compose/data     1.247s

No scenarios
No steps
1.0002ms
ok      docker-compose/features 1.890s
=== RUN   TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent
--- PASS: TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent (0.00s)
=== RUN   TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent
--- PASS: TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent (0.00s)
=== RUN   TestSearchHandlerCallsDataStoreWithValidQuery
    search_test.go:42: PASS:    Search(string)
--- PASS: TestSearchHandlerCallsDataStoreWithValidQuery (0.00s)
=== RUN   TestSearchHandlerReturnsKittensWithValidQuery
--- PASS: TestSearchHandlerReturnsKittensWithValidQuery (0.00s)
PASS
ok      docker-compose/handlers 1.430s
make: *** [unit] Error 2

```

### Issue knowledge when compiler
#### Issue 1: Wrong MINGW compiler architecture
**Behavior:** When you run test
```
make unit 
or 
make test
```

**Error:** You will see error below:
```
chapter4> make unit
go test -race -v ./...
# runtime/cgo
cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
FAIL    docker-compose [build failed]
FAIL    docker-compose/data [build failed]
FAIL    docker-compose/features [build failed]
FAIL    docker-compose/handlers [build failed]
make: *** [unit] Error 2
```

**Solution:** Please install correct MINGW architecture x64 bit
```
Step 1: Download x86_64-win32-seh at link
> https://sourceforge.net/projects/mingw-w64/files/

Step 2: Add to folder C:\MINGIW

Step 3: Add path to C:\MinGW\mingw64\bin
```
## Benchmarks
To run benchmark
```
> go test -bench=. -benchmem
PASS
ok      docker-compose  0.598s
```

One of the other nice features of benchmark tests is that we can run them and it outputs profiles which can be used with pprof:
```
chapter4> go test -bench=. -cpuprofile=cpu.prof -blockprofile=block.prof -memprofile=mem.prof
```



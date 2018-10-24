01/hello.go
-----------

* package, import, function
* string is UTF-8

02/slice.go
-----------

* nil slice is empty slice
* make creates zeroed array with specific size
* declare variable with "var name type" vs := operator
* discard operator
* array is recreated when there is not enough capacity
* array is represented as a slice because of pointers - it is possible to do [lowerBound:upperBound] but it returns pointers to the same array
* for loops variants (classic, foreach, while, forever)

__switch__ and __if__ are almost the same like usual

03/map.go
-----------

* struct - two fields with same type can be on one line or separated
* nil map is not ready to use
* create with literal
* create empty (not nil) with make function
* when entry not present - returned zero value and with false

04/pointers.go
-------------

* like C but without pointer artithmetic

05/defer.go
------------

* values are calculated and saved, but function is executed at the end
* like using in C#
* deferred calls are saved on stack
* can catch panic state

06/package.go
---------------

* lower case private
* upper case public
* embedded structure (sort of inheritance)

07/functions.go
---------------

* multiple output values
* higher order functions
* variadic functions
* methods
    * classic function - quack the duck
    * function receiver - duck quacks

08/interfaces.go
-----------------

* interfaces are implemented implicitly
* empty interface like object in C#
* type assertion
* mention reader and writer

09/errors.go
------------

* error is just any object with Error() method returning string (see implementation of errors.New(""))
* panic/recover is like try/catch
* panic is always fatal to your program, never assume caller can solve the problem, impossible to continue

10/goroutine.go
---------------

* sync package
* when parameter is passed by value, the race condition is gone
* go run -race 03\goroutine.go

11/channel-pingpong.go
-----------------------

* "go func(ch chan string)" both direction
* "<-chan" to sender
* "chan<-" receiver
* when close is missing, there is deadlock because wait for channel will never end

12/channel-parallel.go
-----------------------

* make(chan string) - unbuffered channel
* make(chan string, 5) - buffered channel

13/select.go
------------

* The select statement lets a goroutine wait on multiple communication operations.
* A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
* The default case in a select is run if no other case is ready.

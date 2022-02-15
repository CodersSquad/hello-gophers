---
marp: true
theme: gaia
_class: lead
paginate: true
backgroundColor: #fff
backgroundImage: url('https://marp.app/assets/hero-background.svg')
---

![bg left:40% 80%](https://marp.app/assets/marp.svg)

# **Hello Gophers**
## This is just the beginning

https://talks.obedmr.com/

---
# Agenda
- *YaPL?* - Yet another Programming Language?
- The origin of The *Gophermania*
- *Who* in the world is using Go?
- *Why* should I care?
- *People* want to `code`
- Simple Web Server in Go
- What else?
- Handy *stuff* (links, books, videos, etc)

---

## YaPL? - Yet another Programming Language?
- Initially for large scale systems at Google
- Compiled languages advantages
- Takes the best of the most successful programming languages
- _"Complexity_ _is_ _multiplicative"_ by Rob Pike
- Looking for a distributed by design alternative

---

## The origin of The Gophermania
- At Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson
- Influenced by Niklaus Wirth,  Pascal and Modula-2 inspired the package concept
- Sometimes called "C for the 21st century"
- Came from an explosion of complexity

---

## Who in the world is using Go?
- A better question can be, who's not using it?

---

## Why should I care?
- Highly used and required in most of Cloud Service Providers
- Distributed by design
- Simplicity
- The Golang Tooling makes your code pretty and easy to read and contribute to

---

## From Golang Proverbs
- Concurrency is not parallelism - _from_ *Go* *Proverbs*
- Don't communicate by sharing memory, share memory by communicating - _from_ *Go* *Proverbs*

---

# People want to code ...

---

```golang
package main

import (
	"fmt"
)

func main() {
	topics := map[string]string{
		"easy":     "Program Structure",
		"inspired": "Basic and Composite types",
		"what?":    "Functions vs Methods",
		"cool":     "Goroutines",
		"crazy":    "Channels",
	}

	fmt.Printf("Let's have fun with: \n")
	for key, topic := range topics {
		fmt.Printf(" - \"%s\" which is: '%v' \n", topic, key)
	}
}

```

---
```golang
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

echo1 source code [link](https://github.com/adonovan/gopl.io/blob/master/ch1/echo1/main.go) 

---

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
```

---

```golang
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```
dup3 source code [link](https://github.com/adonovan/gopl.io/blob/master/ch1/dup3/main.go)

---

## Basic and Composite types (1/3)
- Integers
- Float
- Complex numbers
- Booleans
- Strings
- Constants
- Arrays
- Slices
- Maps
- Structs
- JSON
- Text and HTML templates

<style scoped>
ul { columns: 2; }
</style>
---

## Basic and Composite types (2/3)

```golang
const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
```
Basic types (surface) source code [link](https://github.com/adonovan/gopl.io/blob/master/ch3/surface/main.go)

---

## Basic and Composite types (3/3)
[Composite types (slices)](https://github.com/adonovan/gopl.io/blob/master/ch4/append/main.go)

```golang
 letters := []string{"a", "b", "c", "d"}

 func make([]T, len, cap) []T

 var s []byte
 s = make([]byte, 5, 5)
 // s == []byte{0, 0, 0, 0, 0}
```
[Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals)

--- 

## Let's code: Slices and Maps in Go

https://classify.obedmr.com/get-lab/ap-labs/go-slices-maps


---

## Functions vs Methods (1/2)
[Geometry example](https://github.com/adonovan/gopl.io/blob/master/ch6/geometry/geometry.go)

**Struct**
```golang
type Point struct{ X, Y float64 }
```

**Function**
```golang
// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```
--- 

## Functions vs Methods (2/2)

**Method**

```golang
// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

---

## Let's code: Functions and Methods in Go

https://classify.obedmr.com/get-lab/ap-labs/go-functions-methods

---

## Goroutines
In Go, each concurrently executing activity is called a _goroutine_.

When a program starts, its only goroutine is the one that calls the main function. It's called the _main_goroutine_.

```go
  f()       // call f(); wait for it to return
  go f()    // create a new goroutine that calls f(); don't wait
```

---

## Channels
If goroutines are the activities of a concurrent Go program, *channels* are the connections between them.

- A *channel* is a communication mechanism that lets one goroutine send values to another goroutine.

- A *channel* is a _reference_ to the data structure created by _make_

- A *channel* has 2 operations, _send_ and _receive_, also known as _communications_.

---

```golang
  // Channels Examples

  ch := make(chan int) // ch has type 'chan int'

  ch <- x  // a send statement
  
  x = <-ch // a receive expression in an assignment statement

  <-ch     // a receive statement; result is discarded
  
  close(ch) // To close a channel
```
---

### Simple Web Server in Go

```golang
...

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echo the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

Web Server source [link](src/webServer.go)

---

## What else?
- *Interfaces* (as contracts, values, type assertions, etc)
- [*go* *test*](https://golang.org/cmd/go/#hdr-Testing_flags)
- *Reflection?* yes, but be careful
- *Low* *level* *programming?* yes, it's *unsafe*
- You can also call *C* *from* *Golang* with *cgo*
- and ... many other fancy `use{ful|less}` things

---

## Handy Stuff
- [The Golang Tour](https://tour.golang.org/welcome/1)
- [Golang Blog](https://blog.golang.org/)
- [Golang Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html) guide
- Youtube Channel [_Just for func_](https://www.youtube.com/channel/UC_BzFbxG2za3bp5NRRRXJSw)
- [The Go Programming Language Book](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) *by* Donovan and Kernighan
- [Go Proverbs](https://go-proverbs.github.io/)
- [GopherCon](https://www.gophercon.com/) conference
- [GolangNews](https://golangnews.com/) (news, jobs and more)

---

# Thanks

**Obed N Muñoz Reynoso**
Cloud Software Engineer
obed.n.munoz@``gmail | tec | intel``.com

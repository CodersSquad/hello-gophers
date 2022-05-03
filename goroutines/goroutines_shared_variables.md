---
marp: true
theme: gaia
_class: lead
paginate: true
backgroundColor: #fff
backgroundImage: url('https://marp.app/assets/hero-background.svg')
---

![bg left:40% 80%](https://miro.medium.com/max/1000/0*HzrvQ9UCByp4-LcE.jpg)

# **Shared variables in Go**
## The fun continues

https://talks.obedmr.com/


---

## Based on:

[The Go Programming Language](http://techbus.safaribooksonline.com/book/programming/go/9780134190570) book

**Presentation code and examples:**

https://github.com/CodersSquad/hello-gophers


---

## Race Conditions
A **race condition** is a situation in which the program does not give the correct result for some interleavings of the operations of multiple goroutines.


```
 var balance int

 func Deposit(amount int) {
     balance = balance + amount
 }

 func Balance() int {
     return balance
 }
```


---

## Execution example

```
 // Alice
 go func() {
     bank.Deposit(200)                // A1
     fmt.Println("=", bank.Balance()) // A2
 }()

 // Bob
 go bank.Deposit(100)                 // B
```


```
 Alice first              Bob first        Alice/Bob/Alice
           0                      0                      0
   A1    200              B     100              A1    200
   A2 "= 200"             A1    300              B     300
   B     300              A2 "= 300"             A2 "= 300"

```


---

## Goroutine Monitor

_Do not communicate by sharing memory; instead, share memory by communicating._


---

```
 package bank // Package bank provides a concurrency-safe bank with one account.

 var deposits = make(chan int) // send amount to deposit
 var balances = make(chan int) // receive balance
 func Deposit(amount int) { deposits <- amount }
 func Balance() int       { return <-balances }

 func teller() {
     var balance int // balance is confined to teller goroutine
     for {
         select {
         case amount := <-deposits:
             balance += amount
         case balances <- balance:
         }
     }
 }

 func init() {
     go teller() // start the monitor goroutine
 }
```

---

## Mutual Exclusion

Remember the [concurrent web crawler example](./goroutines_channels_part2.md#13)? we may use a buffered channel as a **counting semaphore** to ensure that no more than 20 goroutines made simultaneous HTTP requests.

A semaphore that counts only to 1 is called a **binary semaphore**.


---

```
 var (
     sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
     balance int
 )

 func Deposit(amount int) {
     sema <- struct{}{} // acquire token
     balance = balance + amount
     <-sema // release token
 }

 func Balance() int {
     sema <- struct{}{} // acquire token
     b := balance
     <-sema // release token
     return b
 }
```


---

## Mutual Exclusion with `Mutexes`

```
 import "sync"

 var (
     mu      sync.Mutex // guards balance
     balance int
 )

```

---

```
 func Deposit(amount int) {
     mu.Lock()
     balance = balance + amount
     mu.Unlock()
 }

 func Balance() int {
     mu.Lock()
     defer mu.Unlock()
     return balance
 }
```

- Why we used `defer`?


---

## Read/Write Mutexes

Imagine the case of thousands of `Balance` requests. What will happen?

```
 var mu sync.RWMutex
 var balance int

 func Balance() int {
     mu.RLock() // readers lock
     defer mu.RUnlock()
     return balance
 }
```

---

The `Balance` function now calls the `RLock` and `RUnlock` methods to acquire and releases multiple readers or shared lock.

The `Deposit` function, which is unchanged, calls the `mu.Lock` and `mu.Unlock` methods to acquire and release a writer or exclusive lock.


---

##  Memory Synchronization


```
 var x, y int
 go func() {
     x = 1                   // A1
     fmt.Print("y:", y, " ") // A2
 }()
 go func() {
     y = 1                   // B1
     fmt.Print("x:", x, " ") // B2
 }()
```

Possible outputs:

```
 y:0 x:1    x:1 y:1    x:0 y:0 // This is a weird one
 x:0 y:1    y:1 x:1    y:0 x:0 // This is another weird one
```


---

## That's why `Synchronization` patterns are required on this type of concurrent operations.


---

## Let's code: Goroutines vs OS Threads

Follow instructions from:

https://github.com/CodersSquad/go-goroutines-vs-threads


---

## Golang Race Detector

Take a look on:

https://blog.golang.org/race-detector


---

# Thanks

- **Obed N Muñoz Reynoso**
	- Cloud Software Engineer
	- obed.n.munoz@``gmail | tec | intel``.com

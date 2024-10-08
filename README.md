# tourofgo with gopl

Summaries and notes on this page were compiled while learning GoLang through [The Go Programming Language](http://www.gopl.io/) and [Tour of Go](https://go.dev/tour/welcome/1).

 <details>
 
 <summary>  BASICS </summary>

<details>
<summary>1.packages</summary>

- ## packages
  - every go program is madeup of packages, program always starts from the main packages

src - [packages.go](0.basics/0.packages.go)

</details>
<details>
<summary>2.imports</summary>

- ## imports

  - we can use the codes from other packages but it need to be imported on the used package and it can be done in various ways

    1. import "package1"

       import "package2"

    2. factored

src - [imports.go](0.basics/1.imports.go)

</details>
<details>
<summary>3.exported name </summary>

- ## exportednames
  - in go if anything starts with capital its exported,
    - ex:
      1. type Exportingname int | where Exportingname can be Exported,
      2. func Kunfuguys(attack Exportingname) Exportingname {} | where the Kunfuguys function can be Exported
    - Thats why the famous fmt.Println() function starts with Capital P from the package fmt.

src - [exportednames.go](0.basics/2.exportednames.go)

</details>
<details> 
<summary>4.functions</summary>

- ## functions
  - in go function can be passed with 0 or more arguments
  - in the arguments type always comes after the variable name
  - named return used to make naked return

src - [functions.go](0.basics/3.functions.go)

</details>

<details>

<summary>5.variables</summary>

- ## variables
  - its stores the value of type we defined
  - in go , some time we can initialize a variable without the keyword var and type
    - ex : var num int = 10 -> num := 10
    - but this only applies inside the function only
    - so all the global variables are start with keyword var

src - [variables.go](0.basics/4.variables.go)

</details>

<details>

<summary>6.basic types</summary>

- ## basic types

  1. the basic data types contains

     - bool - (true or false) only has two values , default = false.
     - string - its a read-only slice of bytes , default = ""
     - int - we have more variation of this , like
       - signed - int8 , int16, int32, int64
       - unsigned - uint8, uint16, uint32, uint64 and uintptr
       - the only difference btw these twos are the values it contains.
         - int8 and uint8 total value is 2^8 or 2pow(8), which is 256, so int8(signed) can store from -128 to 127.
         - but for the uint8 we can store from 0 to 255.
     - float64 / float32 - used to store the decimal point numbers

  2. type conversion : it is nothing but converting a value of different
     type to a another type

     - a := t(v) = here v of type1 is converted to a type2

  3. tupple assignment : its another way of assigning value to
     the variables

     - ex : a , b = b , a
     - where right side of the value gets initialized before updating to it, like if a = 10 and b = 20,
     - the a, b = b ,a -> a , b = 20 , 10
       then a = 20 , b = 10.

  4. constants : from the name itself its deriving that its a constant, so it cant be changed,
     - ex : const newOne int = 1 or const py = 3.143 //thats i love you

src - [basictypes](0.basics/5.basictypes.go)

</details>

</details>

<details>

<summary> CONTROL FLOW </summary>

<details>

<summary> 1. for loops</summary>

- ## for loops

  1. go doesnt have much loop other than for
     - but it is very easy to initialize it,
     - it only has 3 components
       1. first one is for variable declaration;
       2. next one is conditional expression;
       3. final one is the iteration;
     - ex for i := 0 ; i < 10 ; i++ {
       contents gets looped
       until condittional became false
       }
  2. we can use the for like while in c also
  3. if there is no variable declaration and conditional expression,
     then its a forever loop

     - for {

       }

src - [for loops](1.controlflow/0.forloops.go)

</details>

<details>

<summary> 2. if conditions </summary>

- ## if condition
  - in go, like forloop we can assign shortstatements which is only accesible inside the scope of if contd.

src - [if condition](1.controlflow/1.ifcontd.go)

</details>

<details>

<summary> 3. switch </summary>

- ## switch
  - it executes its first cases which matches the condittion rather than all possibillities.
  - if a switch is without a condittion , it acts as switch true {}

src - [switches.go](/1.controlflow/2.switch.go), [typeswitch](/2.moretypes/1.structs.go)

</details>

<details>

<summary> 4. Defer </summary>

- # DEFER

  - defer defers the function which means its get executed all after function gets executed including a defer function which initiated after this

  - its in the order of last in first out due to its stack based structure

src - [defers.go](/1.controlflow/3.defers.go)

</details>

</details>

<details>

<summary> MORE TYPES</summary>

<details>

<summary> 1. pointers</summary>

- # POINTERS
  - pointers holds a memory address of a variable
  - & holds - address of a variable where \* holds the value of the variable.
  - ex : a := 10; i := &a; \*i += 2 ; //now a = 12

src - [pointers.go](/2.moretypes/0.pointers.go)

</details>

<details>

<summary> 2. structs </summary>

- # STRUCTS
  - a struct is a collection of fields and these fields are accessed by dots.
  - struct field can be accesed through a struct pointer
    - ex : p = &v ; p.x = 10 | v = st{x : 1 , y : 2}

src - [structs.go](/2.moretypes/1.structs.go)

</details>

<details>

<summary> 3. arrays  </summary>

- # ARRAYS
  - its a linear data structure where all elements are arranged sequently
  - var a [10]int declares a variable array of ten integers

src - [arrays.go](/2.moretypes/2.arrays.go)

</details>

<details>

<summary> 4. slices </summary>

- # SLICES
  - nothing but a dynamically sized array
    - ex : a := []int{2,3,4,5,6,7}
    - a[2:4] -> a[low : high] -> index starts from 0
  - we slice a slice by low, high params | [starting from index :ending before index] = [2 : 6] -> we get 2 to 5 index values
  - starting from index has the ability to drop the values and ending before index just shrink the values.
  - len() / cap() two functions, len() : length describes no. of elements , cap() : describes total no. of underlying array.
  - we can create slice using make()
    - ex : slic := make([]int, 0, 5) -> len = 0 , cap = 5
    - but when we append values to the slice which exceeds the capacity, its just add another 5 cap or the capacity we mentioned to it.
    - slic = append(slic, 10, 12, 14, 16, 18, 20) | len : 6 , cap : 10

src - [slices.go](/2.moretypes/3.slices.go)

</details>

<details>

<summary> 5. range </summary>

- # RANGE
  - It iterates over all values of maps and slice
  - It can be used with string too, because in go string is read-only slice of bytes.

src - [ranges.go](/2.moretypes/4.ranges.go)

</details>

<details>

<summary> 6. maps </summary>

- # MAP
  - its a key and value data structure, in go its only initialized by make only
    - ex : samplemap := make(map[string]float64) | map[key] = value -> version["go"] = 1.22.4 where version is a map name , go is a key and 1.22.4 is the value

src - [maps.go](/2.moretypes/5.maps.go)

</details>

<details>

<summary> 7. function values </summary>

- # FUNCTION VALUES
  - in go functions are also values, so it can be passed as arguments.
  - a function maybe a closure(in go it allows a function to capture variables from its surrounding)

src - [functionvalues.go](/2.moretypes/6.functionvalues.go)

</details>

</details>

<details>

<summary> METHODS AND INTERFACE </summary>

<details>

<summary> 1. methods </summary>

- # METHODS
  - go method is a function with special receiver argument
  - we can create a method func for non struct type too
  - pointer indirection is supported
    - ex : a.abs(); b := &a; b.abs() -> (\*b).abs()

src - [methods.go](/3.methodsandinterfaces/0.methods.go)

</details>

<details>

<summary> 2. interfaces </summary>

- # INTERFACES
  - its a type in go, which interfaces to many types with same methods
  - a interface can initialized with nil type to call a method but a nil interface cant call a method (gives runtime error , because it doesnt know which type of method to call )
  - a empty interface is differ from nil interface, it holds all type or any type we never imagined
  - we can assert an interface by its type and the same thing happens in type switch too
  - There is a interface called Stringer which accesible by fmt package so all the functions in the fmt package can read our method too, so we can customize how the string in a struct or simply a string should appear.
  - if any packages have some methods which accepts some Interface then we can create our own implementation on that method function
  - we can also assert a interface by its method type

src - [interfaces.go](/3.methodsandinterfaces/1.interfaces.go)

src - [typeassertion.go](/3.methodsandinterfaces/5.typeassertion.go)

</details>

</details>

<details>

<summary> GENERICS </summary>

<details>

<summary> 1. generics </summary>
    
*   # Generics     
    *   inorder create a new method , structure and function for different concrete types, we can use generics in go
        1. type parrameter -> where we use generics as parrameters
            *   ex :  func ind[T interface{int | float64}] (a,b T) T {}
        2. Generic types -> where a type itself a generic
            *   ex :   type dog[T interface{string | []byte}] struct {
                name T
            }


src - [genericallinkedlist](4.generics/2.llgen.go)

</details>

</details>

<details>

<summary> CONCURRENCY </summary>

<details>

<summary> 1. goroutines </summary>

- # goroutines
  - to execute functions independently or simultaneously, ex: `go run()`

src - [goroutine.go](/5.concurrency/0.goroutines.go)

</details>

<details>

<summary> 2. Channels </summary>

- # Channels

  - used to communicate between go routines and to synchronize the execution
  - 1.  `ch1 <- 32 ` // sending 32 to channel ch1
    2.  `val := <-ch1` // receiving 32 via channel
  - like map and slice , it should be created before use, to create `ch1 := make(chan int)`
  - we can create buffered channels - Generally the channels are unbuffered which means they dont have storage for to hold values ` ex: each send operation has to wait for the corresponding receive operation to complete`
    but in buffererd it doesnt block until its full `ch1 := make(chan int, 20)`
  - close(chanel) - we can close a channel to say a go routine there are no more values `ex : val , ok := <-chan` where ok is false if there are no more values to receive
    , only sender should close | sending a closed a channel will cause panic

  - Range - we can range a channel too or iterate a channel

  - select - used to select multiple channel operation send and receive

src - [channels.go](/5.concurrency/1.channels.go)

</details>

<details>

<summary> 3. Mutex </summary>

- # MUTEX

  - we already know that go keyword used to create go routine which will run concurrently to main routine and for the communication between them we use channel but what if we create a x number go routines which all access a same a variable, it leads to a race condition thats why we use a mutex for that from the package sync.

  - this datastructure help us lock on different sections of data so that only one goroutine at a time can access it.

src - [mutex.go](/5.concurrency/3.sync.go)

</details>

</details>

<details>

<summary> MISC </summary>

<details>

<summary> 1. panic and recover </summary>

- # PANIC and RECOVER

src - [panicandrecover.go](/misc/panic/panicandrecover.go)

</details>

<! -- <details>

<summary> 2. </summary>

src - []()

</details>

<details>

<summary> 3. </summary>

src - []()

</details>

<details>

<summary> 4. </summary>

src - []()

</details> -->

</details>

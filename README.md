#  tourofgo with gopl

Summaries and notes on this page were compiled while learning GoLang through [The Go Programming Language](http://www.gopl.io/) and [Tour of Go](https://go.dev/tour/welcome/1).

 <details>
 
 <summary>  BASICS </summary>

<details>
<summary>1.packages</summary>

* ## packages
    * every go program is madeup of packages, program always starts from the main packages

src - [packages.go](0.basics/0.packages.go)
</details>
<details>
<summary>2.imports</summary>

* ## imports
    * we can use the codes from other packages but it need to be imported on the used package and it can be done in various ways

        1. import "package1"
        
           import "package2"

        2. factored 

src - [imports.go](0.basics/1.imports.go)      
</details>
<details>
<summary>3.exported name </summary>

* ## exportednames
    * in go if anything starts with capital its exported,
        * ex: 
            1. type Exportingname int  | where Exportingname can be Exported,
            2. func Kunfuguys(attack Exportingname) Exportingname {} | where the Kunfuguys function can be Exported
        * Thats why the famous fmt.Println() function starts with Capital P from the package fmt.

src - [exportednames.go](0.basics/2.exportednames.go)
</details>
<details> 
<summary>4.functions</summary>

* ## functions
    * in go function can be passed with 0 or more arguments
    * in the arguments type always comes after the variable name
    * named return used to make naked return

src - [functions.go](0.basics/3.functions.go)
</details>

<details>

<summary>5.variables</summary>

*  ## variables
    * its stores the value of type we defined
    * in go , some time we can initialize a variable without the keyword var and type 
         *   ex : var num int = 10 -> num := 10
         *   but this only applies inside the function only
         *   so all the global variables are start with keyword var

src - [variables.go](0.basics/4.variables.go)

</details>

<details>

<summary>6.basic types</summary>

*  ## basic types
    1. the basic data types contains 
        * bool - (true or false) only has two values , default = false.
        * string - its a read-only slice of bytes , default = ""
        * int - we have more variation of this , like
            * signed - int8 , int16, int32, int64
            * unsigned - uint8, uint16, uint32, uint64 and uintptr
            * the only difference btw these twos are the values it contains.
                * int8 and uint8 total value is 2^8 or 2pow(8), which is 256, so int8(signed) can store from -128 to 127.
                * but for the uint8 we can store from 0 to 255.
        * float64 / float32 - used to store the decimal point numbers

    2. type conversion : it is nothing but converting a value of different
                         type to a another type

        * a := t(v) = here v of type1 is converted to a type2 

    3. tupple assignment : its another way of assigning value to 
        the variables
        * ex : a , b = b , a
        * where right side of the value gets initialized before updating to it, like if a = 10 and b = 20,
        * the a, b = b ,a -> a , b = 20 , 10
        then a = 20 , b = 10.
    
    4. constants : from the name itself its deriving that its a constant, so it cant be changed, 
        * ex : const newOne int = 1 or const py = 3.143 //thats i love you
        
src - [basictypes](0.basics/5.basictypes.go)

</details>

</details>

<details>

<summary> CONTROL FLOW </summary>

<details>

<summary> 1. for loops</summary>

*   ##  for loops

    1. go doesnt have much loop other than for
        * but it is very easy to initialize it,
        * it only has 3 components 
            1. first one is for variable declaration;
            2. next one is conditional expression;
            3. final one is the iteration;
        * ex for i := 0 ; i < 10 ; i++ {
            contents gets looped
            until condittional became false
        }
    2. we can use the for like while in c also
    3. if there is no variable declaration and conditional expression,
        then its a forever loop 
        * for {

          }

src - [for loops](1.controlflow/0.forloops.go)

</details>

<details>

<summary> 2. if conditions </summary>

*   ## if condition
    * in go, like forloop we can assign shortstatements which is only accesible inside the scope of if contd.

src - [if condition](1.controlflow/1.ifcontd.go)

</details>

<details>

<summary> 3. switch </summary>

*   ## switch
    * it executes its first cases which matches the condittion rather than all possibillities.
    * if a switch is without a condittion , it acts as switch true {}

src - [switches.go](/1.controlflow/2.switch.go)

</details>

<details>

<summary> 4. Defer </summary>

*   # DEFER
    * defer defers the function which means its get executed all after function gets executed including a defer function which initiated after this 

    * its in the order of last in first out due to its stack based structure

src - [defers.go](/1.controlflow/3.defers.go)

</details>


</details>


<details>

<summary> MORE TYPES</summary>

<details>

<summary> 1. pointers</summary>

src - []()

</details>

<details>

<summary> 2. structs </summary>

src - []()

</details>


<details>

<summary>  </summary>

src - []()

</details>


</details>

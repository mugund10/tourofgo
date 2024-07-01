#  tourofgo

<details>
<summary>packages</summary>

* ## packages
    * every go program is madeup of packages, program always starts from the main packages

src - [packages.go](0.basics/0.packages.go)
</details>
<details>
<summary>imporsts</summary>

* ## imports
    * we can use the codes from other packages but it need to be imported on the used package and it can be done in various ways

        1. import "package1"
        
           import "package2"

        2. factored 

src - [imports.go](0.basics/1.imports.go)      
</details>
<details>
<summary>exported name </summary>

* ## exportednames
    * in go if anything starts with capital its exported,
        * ex: 
            1. type Exportingname int  | where Exportingname can be Exported,
            2. func Kunfuguys(attack Exportingname) Exportingname {} | where the Kunfuguys function can be Exported
        * Thats why the famous fmt.Println() function starts with Capital P from the package fmt.

src - [exportednames.go](0.basics/2.exportednames.go)
</details>
<details> 
<summary>functions</summary>

* ## functions
    * in go function can be passed with 0 or more arguments
    * in the arguments type always comes after the variable name
    * named return used to make naked return

src - [functions.go](0.basics/3.functions.go)
</details>

<details>

<summary>variables</summary>

*  ## variables
    * its stores the value of type we defined
    * in go , some time we can initialize a variable without the keyword var and type 
         *   ex : var num int = 10 -> num := 10
         *   but this only applies inside the function only
         *   so all the global variables are start with keyword var

-src: [variables.go](0.basics/4.variables.go)

</details>
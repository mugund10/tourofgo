#  tourofgo

* ## packages
    * every go program is madeup of packages, program always starts from the main packages

src - [packages.go](0.basics/0.packages.go)

* ## imports
    * we can use the codes from other packages but it need to be imported on the used package and it can be done in various ways

        1. import "package1"
        
           import "package2"

        2. factored 
src - [imports.go](0.basics/1.imports.go)      

* ## exported name
    * in go if anything starts with capital its exported,
        * ex: 
            1. type Exportingname int  | where Exportingname can be Exported,
            2. func Kunfuguys(attack Exportingname) Exportingname {} | where the Kunfuguys function can be Exported
        * Thats why the famous fmt.Println() function starts with Capital P from the package fmt.
src - [exportedname.go](0.basics/2.exportedname.go)
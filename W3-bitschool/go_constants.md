## Go Constants
    Variables with a fixed and unchangeable value are called constants. And are declared using the "const" keyword, name and value as the type can always be inferred by the compiler from the value.
    The value must be assigned as it is being delcared.

    e.g

    package main
    import ("fmt")

    const PI = 3.14

    func main() {
    fmt.Println(PI)
}

# Naming constants
    The rules for naming constants are the same as those of a variable only most times, they are written in blocks to easily differentiate them from constants.
    Constants can be declared inside and outside of functions.

# Types of constant
    There are two types of constants:
** Typed constant
** Untyped Constant

    The only difference is that typed constants are declared without a data type meanwhile, untyped ones are declared without a data type and are inferred by the compiler from the value.

    Multiple constants can be declared by grouping them in a block for readability just as when making multiple imports.
    
     
    NOTE: When constants are declared, it is impossible to change their values after declaration.
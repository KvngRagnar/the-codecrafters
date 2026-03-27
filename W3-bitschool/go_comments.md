## Go Comments

    Comments in go are represented by "//" at the beginning of a sentence for single line comments and (*/) at the beginning and end of a sentence for multi lined comments and are igored upon execution of a code.
    
    They makes codes more readable and to explain codes. They can also be used to prevent a code from running when testing an alternative one.

    // This is a comment
    package main
    import ("fmt")

    func main() {
    
    fmt.Println("Hello World!") // This is a comment
    }


## Using Comments To Prevent Code Execution

    Comments as stated earlier, can be used to prevent a particular code from executing as shown below.

    The commented code can be saved for later reference and troubleshooting.


    package main
    import ("fmt")

    func main() {
    fmt.Println("Hello World!")
    // fmt.Println("This line does not execute")
    }
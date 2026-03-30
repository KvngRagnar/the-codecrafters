## Vaiable

# Variable Types in Go

    In go there are various variable types;
    int; which stores integers of both positive and negative values
    string; which stores texts such as "hello"
    float32; this stores floating point numbers with decimals, both positive and negative.
    bool; This stores bolean values, true or false.

# Declaring variables in Go

    In go there are two methods of declaring variables. 

** using the "var" key word or by utilising the ":=" sign.

    When the "var" key word is used it isn't compulsory to give the value of the variable right after declaring it and that can be done in and out of functions in Go. When using the "var" keyword, the vaue or type of variables has to be declared. either one or both but never none.

    Mean while, when using ":=", the value has to be declared immediately and can only be used inside of functions and never outside as that prevents codes from running. Here, the variable tye is inferred by the compiler from the value.

    In GO, variables can be declared with and without initial values. when declared with initial values, those values are reflected in the out put but when the values are undeclared and since Go is initialised, the default values of the data types are printed, mainly from their origin point.
    Values can be assigned after declaration asnot all values may be initially known.
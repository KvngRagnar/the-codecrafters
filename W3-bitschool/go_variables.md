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

    Mean while, when using ":=", the value has to be declared immediately and can only be used inside of functions and never outside as that prevents codes from running. Here, the variable tye is inferred by the compiler from the value. Value and type declaration must be done on the same line.


    In GO, variables can be declared with and without initial values. when declared with initial values, those values are reflected in the out put but when the values are undeclared and since Go is initialised, the default values of the data types are printed, mainly from their origin point.
    Values can be assigned after declaration asnot all values may be initially known.

# Multiple variable declaration

    In Go it is possible to declare multiple variables of different data types on one line but htis is only possible when the "type" keyword isn't used as that would limit the data types to be declared on that line.

# Naming Variables
    In Go, a set of rules govern variable naming and they include:

        A variable name must start with a letter or an underscore character (_)
        A variable name cannot start with a digit
        A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
        Variable names are case-sensitive (age, Age and AGE are three different variables)
        There is no limit on the length of the variable name
        A variable name cannot contain spaces
        The variable name cannot be any Go keywords


# Multi-Word Variable Names
    Variables with more than one word can be difficult to read so there are acceptable ways of writing them.
        1. Camel Case
            Here, each word except the first  begins with a capital letter, e.g. goVariableDeclaration.

        2. Paschal Case
            Here, all the first letters of the words used in naming the variable are capitalised. e.g. GoVariableDeclaration

        3. Snake Case
            In this method, underscore separates ecah word used in naming. e.g. go_variable_declaration
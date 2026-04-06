# GO FUNCTIONS

    Functions in Go are blocks of code designed to perform specific tasks and can be reused multiple times within a program. They are created using the func keyword, followed by a function name, parentheses, and a block of code inside curly braces.

    A function does not run automatically; it only executes when it is called by its name. This allows programmers to organize code efficiently and avoid repetition by reusing the same function whenever needed.

    Functions can be called multiple times, and each call will execute the same set of instructions. Proper naming rules must be followed, such as starting with a letter, using only alphanumeric characters or underscores, and avoiding spaces.

    Overall, functions help make programs more organized, readable, and reusable.

## PARAMETERS AND ARGUMENTS

    In Go, parameters are used to pass information into functions, allowing them to work with different values. Parameters are defined inside the function’s parentheses along with their data types and act as variables within the function.

    When a function is called, the actual values passed to it are called arguments. These arguments replace the parameters and are used when the function executes.

    Functions can have one or multiple parameters, separated by commas. When calling such functions, the number of arguments must match the number of parameters, and they must be provided in the correct order.

    Overall, parameters and arguments make functions more flexible and reusable by allowing them to handle different inputs.

## RETURN VALUES

    In Go, functions can return values by specifying a return type and using the return keyword. This allows a function to send results back to where it was called, such as returning the sum of two numbers.

    Go also supports named return values, where the return variables are defined in the function signature. These values can be returned directly using a “naked return” without explicitly naming them again.

    The returned value from a function can be stored in a variable for later use. Additionally, Go functions can return multiple values at once, which can be assigned to multiple variables when the function is called.

    If some returned values are not needed, they can be ignored using an underscore (_). Overall, return values make functions more powerful by allowing them to produce and share results.
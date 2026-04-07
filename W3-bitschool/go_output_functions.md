# GO OUTPUT FUNCTIONS

    Go has three functions to output text:

        Print()
        Println()
        Printf()


## The Print() Function

    The Print() function in Go is used to display output in its default format. It prints values exactly as they are without automatically adding spaces or new lines between string arguments. To create a new line, the escape character \n must be used manually. Multiple values can be printed in a single statement, and spaces can be added explicitly using " ". However, when printing non-string values, Print() automatically inserts a space between them. 
    Overall, the Print() function provides basic control over how output is displayed in Go programs.

## The Println() Function

    The Println() function in Go is similar to Print(), but it automatically adds a space between arguments and inserts a new line at the end of the output. This makes it more convenient for displaying readable output without manually adding spaces or newline characters. It is commonly used when printing multiple values on separate lines.
    Overall, Println() simplifies output formatting for basic use cases.

## The Printf() Function

    The Printf() function in Go is used for formatted output, allowing developers to control how values are displayed. It uses formatting verbs such as %v to print values and %T to display the data type of variables. This function is especially useful when you need precise control over the format of the output. Unlike Print() and Println(), it does not automatically add spaces or new lines unless specified. 
    Overall, Printf() provides flexibility and customization in output formatting.

# General Formatting Verbs

    General formatting verbs in Go are used with the Printf() function to display values in different formats regardless of their data type. The %v verb prints values in their default format, while %#v shows them in Go-syntax format. The %T verb displays the data type of a value, and %% is used to print a literal percent sign. These verbs provide flexibility for general-purpose output formatting.

## Integer Formatting Verbs

    Integer formatting verbs are used to display numbers in different number systems and formats. For example, %b prints binary, %d prints decimal, %o and %O print octal, while %x and %X print hexadecimal values. Additional options allow formatting with signs, padding, and alignment, such as %4d or %04d. These verbs are useful for controlling how integer values are presented.

## String Formatting Verbs

    String formatting verbs control how text is displayed. The %s verb prints a plain string, while %q prints it with double quotes. Formatting options like %8s and %-8s adjust alignment and spacing, and %x can display the string as hexadecimal byte values. These options make it easier to format and manipulate text output.

## Boolean Formatting Verbs

    Boolean formatting uses the %t verb to display values as true or false. It works similarly to %v but is specifically designed for boolean values. This ensures clear and readable output when working with logical conditions.
 
## Float Formatting Verbs

    Float formatting verbs are used to display decimal numbers in different styles. The %f verb shows standard decimal format, while %e uses scientific notation. Precision can be controlled using formats like %.2f, and %g displays only necessary digits. These verbs help control the accuracy and appearance of floating-point numbers.
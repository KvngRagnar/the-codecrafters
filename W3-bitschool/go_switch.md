# The switch Statement

    The switch statement in Go is used to select and execute one block of code from multiple options. It evaluates an expression once and compares its value with different cases. When a match is found, the corresponding code block is executed. Unlike some other languages, Go automatically stops after a match, so a break statement is not required. This makes the switch statement cleaner and easier to use.

## Single-Case switch Syntax

    The switch syntax consists of an expression followed by multiple case statements and an optional default block. Each case represents a possible value that the expression can match. If a match occurs, the associated code runs. The default case is executed when none of the cases match. This structure helps simplify multiple conditional checks.

## Single-Case switch Example

    In practice, a switch statement can be used to map values to outputs, such as converting a number into a weekday name. The expression is compared against each case, and when a match is found, the corresponding output is displayed. This approach makes the code more readable than using many if-else statements.

## The default Keyword

    The default keyword is used to define a fallback block of code when no case matches the expression. It ensures that the program still produces an output even when none of the specified conditions are met. Additionally, all case values must match the type of the switch expression, otherwise the program will produce a compilation error.

## The Multi-case switch Statement

    The multi-case switch statement in Go allows multiple values to be grouped under a single case. This means that if the expression matches any of the listed values, the same block of code will be executed. It helps reduce repetition and makes the code more concise and readable. The default case is still optional and runs if none of the values match. Overall, multi-case switch statements are useful for handling multiple related conditions efficiently.
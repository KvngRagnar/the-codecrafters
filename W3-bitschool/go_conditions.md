# GO CONDITIONS

    Conditional statements in Go are used to execute different actions based on whether a condition is true or false. These conditions are created using comparison operators such as <, >, <=, >=, ==, and !=, as well as logical operators like && (AND), || (OR), and ! (NOT).

    Go provides several conditional statements to control program flow. The if statement executes a block of code when a condition is true, while the else statement runs when the condition is false. The else if statement allows additional conditions to be tested if the previous ones are false.

    The switch statement is used when there are multiple possible conditions, making the code cleaner and easier to read compared to many if-else statements.

    Overall, conditional statements help programs make decisions and respond dynamically to different situations.

## THE IF STATEMENT

    The if statement in Go is used to execute a block of code only when a specified condition is true. It is written in lowercase (if), as using uppercase will result in an error.

    A condition inside an if statement is evaluated, and if it returns true, the code inside the block is executed. This condition can be a direct comparison (e.g., 20 > 18) or involve variables.

    For example, variables like x and y can be compared using operators such as >. If the condition x > y is true, the program executes the corresponding code, such as printing a message.

    Overall, the if statement allows programs to make decisions and perform actions based on specific conditions.

## THE ELSE STATEMENT

    The else statement in Go is used together with an if statement to execute a block of code when the condition is false. If the if condition is not satisfied, the program automatically runs the code inside the else block.

    This allows programs to handle alternative outcomes. For example, if a condition like time < 18 is false, the program can display a different message such as “Good evening” instead of “Good day.”

    The if and else statements must be written together properly, with the else placed directly after the closing bracket of the if block (} else {). Writing them on separate lines incorrectly can cause errors.

    Overall, the else statement ensures that a program can respond appropriately when a condition is not met.

## THE ELSE IF STATEMENT

   he else if statement in Go is used to check additional conditions when the first if condition is false. It allows multiple conditions to be tested in sequence, giving the program more decision-making options.

    If the first condition is false, the program evaluates the else if condition. If that is also false, it can continue to other else if statements or execute the else block if all conditions fail.

    Only one block of code is executed in an if–else if–else structure. Once a true condition is found, the program runs that block and skips the remaining conditions.

    Overall, the else if statement helps handle multiple possible outcomes efficiently, making programs more flexible and easier to control.

## THE MESTED IF

    A nested if statement in Go is an if statement placed inside another if statement. It is used to check multiple conditions in a structured way, where the inner condition is only evaluated if the outer condition is true.

    If the first condition is satisfied, the program executes its block and then checks the inner if condition. The inner block runs only when both conditions are true.

    If the outer condition is false, the nested if is skipped entirely, and the program may execute an else block instead.

    Overall, nested if statements allow for more detailed decision-making by combining multiple related conditions within a program.
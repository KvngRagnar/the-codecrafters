# Go for Loop

    The for loop is the only looping construct in Go and is used to repeatedly execute a block of code. It is useful when the same code needs to run multiple times with different values, and each repetition is called an iteration. The loop can include three components: initialization, condition, and increment, which control how the loop starts, runs, and ends. Although these components are optional, they must exist in some form within the loop logic. Overall, the for loop provides a flexible and powerful way to handle repetition in Go programs.

## for Loop Examples

    Examples of for loops show how they can be used to count numbers or perform repeated tasks. A loop can start from a value, check a condition, and increment until the condition becomes false. Different increments, such as increasing by 1 or by 10, allow for customized iteration patterns. These examples demonstrate how loops can be adapted for various counting and repetitive operations.

## The continue Statement

    The continue statement is used to skip the current iteration of a loop and move to the next one. It is typically used with conditions to ignore specific values during iteration. This helps control the flow of the loop without completely stopping it. Overall, continue allows selective execution within loops.

## The break Statement

    The break statement is used to immediately terminate a loop when a certain condition is met. Once break is executed, the loop stops running, even if the condition for continuation is still true. It is useful when a specific stopping point is required. This makes it easier to control loop execution and avoid unnecessary iterations.

## Nested Loops

    Nested loops occur when one loop is placed inside another loop. The inner loop runs completely for each iteration of the outer loop. This structure is useful when working with combinations of data, such as pairing elements from two arrays. Nested loops allow handling more complex repetitive tasks.

## The Range Keyword

    The range keyword is used to simplify iteration over arrays, slices, and maps. It returns both the index and the value of each element during iteration. Developers can choose to use both values or ignore one using an underscore (_). This makes looping cleaner and more efficient when working with collections of data.
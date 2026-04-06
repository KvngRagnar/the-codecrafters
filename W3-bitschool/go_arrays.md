# GO ARRAYS

    Arrays in Go are used to store multiple values of the same data type in a single variable. They have a fixed length, which can either be defined explicitly or automatically inferred by the compiler based on the number of values provided.

    Arrays can be declared using either the var keyword or the shorthand := syntax. Each element in an array is accessed using an index, starting from 0, meaning the first element is at index 0, the second at index 1, and so on.

    You can modify array elements by assigning a new value to a specific index. If an array is not fully initialized, its elements are automatically assigned default values, such as 0 for integers and an empty string for strings.

    It is also possible to initialize only specific elements in an array by specifying their index positions. Additionally, the len() function is used to determine the total number of elements in an array.

    Overall, arrays in Go provide a simple way to manage collections of fixed-size data efficiently.
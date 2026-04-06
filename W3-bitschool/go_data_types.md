# GO DATA TYPES

    Data types in Go define the kind of data a variable can store and determine its size and behavior. Go is a statically typed language, meaning once a variable’s type is declared, it cannot store values of a different type.

    Go has three main basic data types. The bool type represents boolean values, which can only be true or false. The numeric type includes integers, floating-point numbers, and complex numbers used for mathematical operations.

    The string type is used to store text values. Overall, data types help ensure that data is used correctly and efficiently within a program.

## BOOLEAN DATA TYPE

    The boolean data type in Go is declared using the bool keyword and can only hold two values: true or false. It is mainly used for decision-making and conditional statements in programs.

    Boolean variables can be declared with or without an initial value. When a boolean is declared without being assigned a value, it automatically takes the default value of false.

    Go also allows both typed declarations (explicitly specifying bool) and untyped declarations using shorthand syntax. Regardless of how they are declared, boolean variables always store either true or false.

    Overall, the boolean data type is essential for controlling program flow through conditions and logical operations.

## INTEGER DATA TYPE

    Integer data types in Go are used to store whole numbers without decimal points, such as positive and negative values. They are divided into two categories: signed integers, which can store both positive and negative numbers, and unsigned integers, which can store only non-negative values.

    Signed integers include types like int, int8, int16, int32, and int64, each with different sizes and value ranges. The default integer type is int, and its size depends on the system architecture (32-bit or 64-bit).

    Unsigned integers include types like uint, uint8, uint16, uint32, and uint64, and they only store values greater than or equal to zero. Like signed integers, their sizes determine the range of values they can hold.

    Choosing the correct integer type depends on the size of the value you need to store, as using a type with insufficient range can cause errors such as overflow.

## FLOAT DATA TYPE

    Float data types in Go are used to store numbers with decimal points, including both positive and negative values. Go provides two float types: float32 and float64, which differ in size and the range of values they can store.

    The float32 type uses 32 bits and can store smaller decimal values, while float64 uses 64 bits and can store much larger and more precise numbers. By default, Go assigns the float64 type to floating-point numbers if no type is specified.

    Choosing the appropriate float type depends on the size and precision of the value needed. Using a type with a smaller range, such as float32, for very large numbers can cause overflow errors.

    Overall, float data types are essential for handling decimal and scientific calculations in Go programs.

## STRING DATA TYPE

    The string data type in Go is used to store text, which is a sequence of characters enclosed in double quotes. It allows programs to handle words, sentences, and other textual data.

    String variables can be declared with or without an initial value. If a string is declared without being assigned a value, it defaults to an empty string ("").

    Go supports both typed declarations (explicitly using string) and shorthand declarations using :=. Regardless of how they are declared, all string variables store text data.

    Overall, the string data type is essential for working with and displaying text in Go programs.


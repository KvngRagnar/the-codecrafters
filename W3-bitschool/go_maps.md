# GO MAPS

    Maps in Go are used to store data in key-value pairs, where each key is unique and maps to a specific value. They are unordered, meaning the elements are not stored in a fixed order, and they do not allow duplicate keys.

    Maps can be created using the var keyword, shorthand :=, or the make() function. The make() function is the preferred way to create empty maps, as using other methods without initialization can lead to runtime errors.

    You can access, add, and update elements in a map using keys, and remove elements using the delete() function. The len() function is used to determine the number of elements in a map.

    Maps also allow checking if a key exists using a special syntax that returns both the value and a boolean indicator. Additionally, maps are reference types, meaning changes made to one variable can affect another if they point to the same map.

    Overall, maps are powerful for storing and retrieving data efficiently using unique keys.
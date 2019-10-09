Any type that can have a method associated with it can implement an integer
An empty interface(type Example interface{}) is an interface that has no behaviour/method in it

BEST PRACTICES
-Use many, small interfaces
    -Single method interfaces are some of the most powerful and flexible

-Don't export interfaces for types that will be consumed
-Do export interfaces for types that will be used by package
-Design functions and methods to receive interfaces whenever possible
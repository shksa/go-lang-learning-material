## **Go's philosophy**
***Go aims to provide everything that programmers need, including many powerful and convenient features while keeping the language as small, consistent, and fast (to build and run) as possible.***

## **WORKSPACE**: 
- It is the project directory. 
- The location of workspace should be present as one of the colon delimited path string in the GOPATH environment variable value.


## **GOPATH**: 
- GOPATH is an environment variable.
- This variable is used to **resolve import statements in go source files**.
- If you created a project directory without adding it's location to the GOPATH env variable and a package source inside that project has an import statement containing an `import path` to a package in the same project, Go will not be able to resolve that import path and will raise a ***cannot find package error***.
- This is because Go looks for all the import paths in the GOPATH env variable. 
- If your project location is not present in the GOPATH, Go obviously cannot look for any package paths inside that project location, therefore resulting in a ***cannot find package error***.
- You can however use the go build coomand-line tool to build and run a package present at any location provided it only imports packages from Go's standard library.
- ***It is a must to include the path of your project directory in the GOPATH env variable so that Go can resolve import paths present in the sources of that project***.
-  ```
    edit .zshrc file to contain the below line and run source .zshrc
    export GOPATH=/workspace1Path/:/workspace2Path/:/workspace3Path
    ```

## **`go`** **TOOL**:
- The `go` command line tool is used to build, fetch and install Go packages and commands.

## **`go env GOPATH`**: 
- This command returns the value of GOPATH env variable. If you change GOPATH env value, `go env GOPATH` will reflect that change.

## **WORKSPACE STRUCTURE**:
- src - contains Go packages.
- pkg - contains package objects
- bin - contains command executables

## **`src`** **DIRECTORY**: 
- This directory is a must in a Go project. 
- All the packages in a Go project should be rooted at this directory.
- Go's package resolution algorithm resolves all the import paths by searching for those paths inside the `src` directory of projects included in GOPATH env variable.

## **`go build`** vs **`go install`**: 
- `go build [package import path]`
  - on single **`main`** package:- It compiles the package and creates an executable file.
  - on single **`non-main`** package:- It compiles the package but does not produce any output file. **Can be used to check if the package has any compilation errors**.
- **`go install [package import path]`** 
  - command compiles and produces an output file. This output file can be a command executable or a package object depending on the source.

## **PACKAGE:**
- Every piece of Go code exists inside a package.
- Go operates in terms of packages rather than files.
- A package is what Go calls it's code **modules**. 
- A package can be split across many files as we like, from Go's perspective if all the files have same package declaration they are part of the same package.
- **A package in Go is a directory containing one or more Go source files**.
- **A package can be of 2 types**
  - **main** package
  - **non-main** package


## **PACKAGE PATH or IMPORT PATH:**
- A package path is simply the path of a package directory that contains Go source files.
- **This is also the import path of a package.**
- The package path must include a **base path** which usually is a remote repository path like **github.com/shksa**, so then the package path will look like **github.com/shksa/packageName**
- All Go source files inside this package directory must define the same package name which by Go's convention is the last element of the package's import path (this is nothing but the name of the package directory itself).

## **PACKAGE NAMES: The first statement in a Go source file must be**
```
package name
```
- This statement defines a name for all the code in the file to be packaged under that name.
- All source files under a package must use the same package name.
- Go's convention is that the package name is the last element of the import path. package imported as **`crypto/rot13`** should be named as **`rot13`**.
- Packages for building executables should always use the name 'main'.

## **`go get`:**
- `go get` is used to fetch remote packages. 
- The import path or package path contains the path of remote repository is its base, so using that `go get` command will fetch, build and install that package.

## **EXPORTED NAMES:**
- In Go, a name is exported only if it begins with a capital letter. 
- So when importing a package you can only refer to it's exported names, "unexported names" are not accessible from outside the package. If you try to access a unexported name, you will get an error.
  
## **Variable declaration, initialization, zero-values**
- In Go variables can be declared with the **`var`** keyword.
- Initialization of variables can be done along with the declaration statement itsef or with a seperate assignment statement.
- **Go has the feature of automatically initializing variables with the zero-value of the variable's type when they are'nt explicitly initialized**.
- *This feature ensures that Go programs dont suffer from the problem of initialized garbage values that affect other languages*
- **Each type has a zero value**.
- zero value of int is 0, bool is false, string is "".
- zero value of a **array** is defined by the zero value of it's elements.
- zero value of a **slice** is `nil`. Although zero value is nil, it behaves as a length 0 slice which can be appended by values to make the slice grow.

- **Go supports type inference and duck typing**.
- ***Go also supports 'short variable declaration' statements. Such a statement both declares and initializes a variable at the same time without needing to specify the variable's type, Go can deduce that from the initializing value.***
- **Go also supports multiple variable declaration in a single statement**
-  ```go
    // single variable declaration
	var x1 int = 9 // 1st alternative

	var x2 int     // 2nd alternative
	x2 = 9

	var x3 = 9     // 3rd alternative

	x4 := 9        // 4th alternative

	// Multi-variable declaration
	a, b, c := 1, 2, 3 // with short variable declaration syntax

	var x, y, z int    // with 'var' declaration and no explicit initialization

	var m, l, k = 1, 2, 3 // with 'var' declaraation and explicit initialization
    ```
## **Shoadowing**
- Variables are shadowed by variables of the same name in different scope.
- ```go
    a, b, c := 1, 2, 3
    for a := 1 ; a < 7 ; a++ { // a is shadowing the outer a
      b := 11 // b is shadowing the outer b
      c := 12 // c is shadowing the outer c
    }
    fmt.Printf("a = %v, b = %v, c = %v", a, b, c) 
    // a = 1, b = 2, c = 3
    ```

## **Arrays:**
- A Go array is a **fixed-length sequence of items of the same type**.
- The length of array is part of the type.
- Arrays cannot be resized.
- Array can be created using these statement
  - `var array [length]Type`
  - `var array = [N]Type{value1, value2, value3, value4, ..., valueN}`
  - `var array = [...]Type{value1, value2, value3, value4, ..., valueN}`
- `Go guarentees that all array elements are initialized to their zero value if they are not explicity initialized`.
- ```go
    var buffer [10]byte
    // buffer is [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    var grid = [4][4]int{
      {1, 2, 3},
      {4, 5, 6},
      {7, 8, 9},
    }
    /* 
    grid is 
    [
      [1, 2, 3],
      [4, 5, 6],
      [7, 8, 9],
      [0, 0, 0]
    ]
    */
    cities := [...]string{"bangalore", "chennai", "delhi", "kolkata"}
    ```

## **Slices**
- **A Slice is a variable-length fixed capacity sequence of items of the same type**.
- In Go slices are flexible, powerful, convinient than arrays.
- Arrays vs Slices
    | Arrays       | slices |
    | ------------ | ------------ |
    | Arrays are passed by value(i.e copied) | Slices are references, hence they are cheap to pass |
    | Arrays are of fixed size | Slices can be resized |
    | Can be sliced | Can be sliced |
- **Shrinking and growing**
  - Slices can shrunk by slicing.
  - Slices can be grown using the `append` built-in function.
- Slices can be created using these statements
  - `var slice = make([]Type, length, capacity)`
  - `var slice = make([]Type, length)`
  - `var slice = []Type{}`
  - `var slice = []Type{value1, value2, value3, ..., valueN}`
  - `var slice []Type`
- Slices have the type `[]T`.
- Arrays have the type `[n]T`, this is not the same as type of slice which is `[]T`.
- **Slice is a seperate data type in Go, unlike python**.

## Composite literals
- ***Primitive literals construct values for primitive types.***
- ***Composite literals construct values for composite types.***
- Both these literals create a new value each time they are evaluated
- Composite types in Go include arrays, slices, maps, structs.
- **Syntax of composite literal**
  - **`literalType{literalValues}`**
- ```go
    type Point3D struct { x, y, z float64 }
    type Line struct { p, q Point3D }

    origin := Point3D{}        // zero value for Point3D
    line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x
    ```

## **Creating slices with `make([]T, length, capacity)`**
- `make()` is a built-in function to create slices, maps, channels which are all reference type.
- `make()` when used to create slices, creates a **hidden zero-value** initialized array and returns a slice reference that refers to the hidden array.
- **The hidden array like all arrays in Go is fixed in length**. The length of this hidden array is defined to be the capacity of the slice.
- **A slice's length is any amount upto it's capacity**.

## **Creating slices with composite literal and nil slices**
- Creating slices with the composite literal is very convinient because it allows us to create slices with some initial values.
- `[]Type{}` is equivalent to `make([]Type, 0)`. Both create an **empty slice** which is a reference to a length 0 array.
- The built-in **append()** can be used to populate empty slices.
- The declaration `var slice []Type` creates a slice that is equal to `nil`, no hidden array is created. Just a slice variable of the given type is created.
- Zero value of a slice is `nil`.
- nil slice has length and capacity as 0 and no underlying array.
- A nil slice behaves as a length 0 slice, it can be appended with values to grow the slice.
- ```go
    var slice []int
    slice == nil
    len(slice) == 0
    cap(slice) == 0
    // This nil slice can be grown using append() function
    slice = append(slice, 3, 4, 5)
    len(slice) == 3
    cap(slice) == 3
    ```
- > **For practical purposes, when we need to create an initial empty slice it's always better to create the empty slice with `make()` giving it a length 0 and a non-zero capacity that is approximately equal to the number of items you expect the slice to end up with**.

## **Slices are refernces to arrays, not the arrays themselves**
- A slice is a reference to an array and slices of slices are also refernces to the same array.
- So when we change the array data through one of it's slice references, the change is visible to all the slice references.
- ```go
    s := []string{"A", "B", "C", "D", "E", "F", "G"}
    t := s[:5]
    u := s[3:len(s)-1]
    // s is ["A", "B", "C", "D", "E", "F", "G"]
    // t is ["A", "B", "C", "D", "E"]
    // u is ["D", "E", "F", "G"]
    // Now, change the hidden array element value through one of the slices
    u[0] = "ζ"
    // The hidden array is modified, this is visible to all of it's slice references.
    // s is ["A", "B", "C", "ζ", "E", "F", "G"]
    // t is ["A", "B", "C", "ζ", "E"]
    // u is ["ζ", "E", "F", "G"]
     ```
- ![alt-text](src/learningGo/sliceAndInternalArray.png)

## **Slice of pointers**
- Slice can contain pointer types.
- ```go
    type Product struct {
      name string
      price float64
    }

    products := []*Product{
      &Product{"milk", 34.5},   // or just {"milk", 34.5}
      &Product{"bisket", 23.5}, // or just {"bisket", 23.5}
      &Product{"balloon", 4.3}, // or just {"balloon", 4.3}
    }

    for _, product := range products {
      product.price += 0.5
    }
    ```

## **Shrinking and growing slices**
- A slice's length can grown or reduced by reslicing it, it can't be grown beyond it's capacity
- Slice's **capacity** can be grown using the **append()** function.
- ```go
    avengers := []string{"ironman", "thor", "hulk"}
    avengers = append(avengers, "blackwidow", "hawkeye")

    justiceLeague := []string{"batman", "superman", "flash"}

    heros := append(avengers, justiceLeague...) // appending a slice to a slice
    ```

## **Closures**
- Go functions can be closures when they are declared inside an outer function **without a name**.
- **A closure is a function that "captures" constants and variables that are present in the same scope where the function was created if the function refers to them in it's body**.
- **In this sense, the function is bound to those variables**.
- The variables closed over the function are like properties of the function, like property of an object.
- **They do not behave like local variables, they rather behave like static variables, in that they exist even after the function has returned and can can be used in the next call with it's value preserved from the previous call.**
- **In Go**, every anaonymous function is a closure.
- **For a function to be a closure it must be defined with no name**.
```go
func fibonacci() func() int {
  a, b = 1, 0
  return func() int { // anonymous function, therefore a closure
    a, b = b, a+b
    return a
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 6; i++ {
    fmt.Println(f())
  }
}
>> 0
>> 1
>> 1
>> 2
>> 3
>> 5
```

## **Type conversions**
- Conversion syntax: **`resultOfDesiredType:= desiredType(expression)`**
- The **`expression`** argument can be a numeric literal, string literal, variable etc.
- A **`string`** can be converted to **`[]byte`**(**underlying utf-8 bytes**) or to a **`rune[]`**(**unicode code points**) and both **`byte[]`** and **`rune[]`** can be converted to a **`string`**
- A custom type of string slice can be converted to a plain string slice.
- ```go
    stringVar := "sleptking"
    // converting string type to byte slice type []byte
    // using the call Type(expression)
    byteSlice := []byte(stringVar)
    fmt.Printf("%v \n", byteSlice) // [115 108 101 112 116 107 105 110 103]

    // converting string type to rune slice type []byte
    // using the call Type(expression)
    runeSlice := []rune(stringVar)
    fmt.Printf("%v \n", runeSlice) // [115 108 101 112 116 107 105 110 103]

    // A custom string slice type
    type StringSlice []string

    fancy := StringSlice{"lithium", "sodium", "potassium"}

    // converting custom string slice to plain string slice type []string
    plainStringSlice := []string(fancy)
    ```


## **Strong typing**
- Go is strongly typed.
- This means when a variable is declared as an `int`, we can only assign `int` type values to it. 
- Assigning values other than `int` will raise a compile time error.
- Simply put, strong typing means `a variable declared with a certain type can only accept values of that type`.

## **Constants**
- Constants can be declared using the `const` keyword
```go
  const pi = 3.14 // constant; type-compatible with any numeric type
  const top uint8 = 42 // constant; type: uint8
```
- There are two kinds of numeric constants
  - `Typed numeric constants` -  can only be used in expressions with numbers of the same type
  - `Untyped numeric constants` - can be used in expressions with numbers of any built-in type

## **Enumerations**
- Go has bare-bones support for enumerations.
- It is acheived by grouping several constant declarations using the const declaration once and the `iota` keyword.
```go
  const (
    Cyan = iota //0, value of iota is 0 now
    Magenta     //1, also iota but now iota's value is 1
    Yellow      //2, also iota but now iota's value is 2
  )
```
- **`iota`** keyword represents successive untyped integer constants.
- **It's value is reset to zero whenever the keyword const appears and increments by one for each constant declaration**.
- ***When to use enums?***
  - **enums are used in a case when you want just want constants to have distinct values and don't really care what those values are.**

## **Variables**
- Variables in Go can hold
  - ***values*** or
  - ***references*** or
  - ***pointers***
- *A variable is the name given to a piece of memory that holds a value*.
- ***We can think of variables as being the value it holds***.
- *Go treats a variable as synonymous with the memory that stores the value assigned to that variable*.
-   ```go
    y := 1.5 // variable: y, value: 1.5, type: float64, memory address: 0xf870000f78
    ```


## **Types in Go**
- Fundamentally types in Go can divided into 3 groups.
  1. **Value** types, ex: `int`, `string`, `bool` types
  2. **Pointer** types, ex: `*int`, `*bool` type
  3. **Reference** types, ex: `[]int`, `[]string`, `map[int]string` types




## **Pointers**
- Pointer is a variable that stores another variable's memory address.
- A variable that is pointed-to by a pointer can be modified through that pointer.
- **Ponited**-to variables persist in memory as long as there exists a pointer that points to them.
- `&` is called the *address*-of operator.
- \* as a unary operator: `*` operator when used with a pointer provides access to value of variable the pointer points to.
```go
x := 2 // variable: x, type: int
px := &x // variable: px, type: *int
*px == 2 // *px and x can be used interchangebly now.
// *px is called dereferencing the pointer px.
```
- `*` is called `contents-of` or `dereference` or `indirection` operator.
- `Dereferencing a pointer`: This means to access value of the variable it points to using the `*` operator.
- `Indirection`: Using the pointer to refer to a value is called `indirection`.
- `pointed-to value` - Value of variable that a pointer points to.
- `pointed-to variable` - Variable that a pointer points to.
```go
*px++ // increments the pointed-to value
*px == 3
x == 3
```
- `* as type modifier`: 
  - `*` operator when placed to the left of a type name changes the meaning of the name from specifying the `type of value` to specifying the `pointer to a value of the given type`.

## **new() and & operator to create pointers**
- `There are 2 ways in Go to create variables and acquire pointers to them at the same time!`.
- `using new(type)`:
  - Using a type we can 
    - `Create values of that type` or 
    - `Pointers to values of that type`.
  - `new(type)` does the second thing
  - ```go
    type composer struct {
      name      string
      birthYear int
    }

    rahman := composer{"A.R Rahman", 1280} // creating a value of type 'composer'
    
    zimmer := new(composer) // creating a pointer to a value of type 'composer'
    
    zimmer.name, zimmer.birthYear = "hans zimmer", 1325 // modyfying the pointed-to value
    ```
  - `new(Type) ≡ &Type{}`
  - Both these syntaxes allocate a new zeroed value of the given `Type` and return a pointer to that value.
  - `If a type cannot be initialized using braces then we can only use the built-in new() function`.
  - Using `&Type{initial values}` for `structs` we can specify initial field values
  - ```go
    beethovan := &composer{"blind beethovan", 4512} // pointer to a 'composer' type value with the specified fields.
    ```

## **Why pass pointer type variables instead of regular variables to functions?**
- ***To modify the data in-place and pointers are a lot cheaper to pass***.
- Anytime you pass value type variables as arguments to a function like for example **`swap(x, y)`**, the function receives a copy of those values.
- Then these received values are set as values of parameters of the function.
- If you want your function to receive an argument as-is instead of it's copy, you need to send a pointer of that variable to the function.
- Once you have a pointer of the argument, you can access it's value by ***dereferencing the pointer***.
- A swap function in Go looks like this.
- ```go
    func swap(px, py *int) {
      // px == 0xf5622d3276... something
      // py == 0xfd6237da67... something
      // px is a variable which serves as a pointer to variable x of type int
      // py is a variable which serves as a pointer to variable y of type int
      // *px accesses the value of variable x
      // *py accesses the value of variable y
      // You can change value of variable x through it's pointer variable px by the statement *px = NEW_VALUE
      *px, *py = *py, *px
    }

    func main() {
      x, y := "foo", "bar"
      swap(&x, &y)
      x == "bar" // true
      y == "foo" // true
    }
    ```

## **Reference types**
- *A variable of a reference type refers to a hidden value in memory that stores the actual data*.
- Variables of reference types are cheap to pass.
- They have the same syntax to work with as variables of value types.
- ***slices, maps, channels, functions, methods*** are reference types.
- **Unlike pointers, there is no special syntax for reference types since they are used just like values**.
- ```go
    grades := []int{23, 67, 45, 12, 86, 56, 32, 87, 57}
    // variable grades is of type slice of ints which is a refernce type.
    // This variable refers to a hidden array value but does not hold any data like variables of value types.
    ```

## **init() and main() function**
- Go reserves two function names for special purposes.
  1. **init()** - in all packages
  2. **main()** - only in *main* packages
- A Go program's startup sequence 
- ![alt-text](src/learningGo/programStartUpSequence.png)
- When a package is imported, 
  1. If it has it's own imports they are performed first.
  2. The the package-level constants and variables are created.
  3. And then the **`init()`** function is called if the package has it.

## Errors
- Errors in Go are expressed with values of the built-in **`error`** type.
- The **`error`** type is a built-in interface just like the fmt.Stringer interface
- ```go
    type error interface {
      Error() string
    }
    ```
- How to use errors in programs
  - Function should return **`error`** values and the calling code should handle the errors by testing whether the error equals **`nil`**.
  - A **nil** **`error`** value denotes **success**, a **non-nil** **`error`** value denotes **failure**.
  - ```go
      i, err := strconv.Atoi("42")
      if err != nil {
          fmt.Printf("couldn't convert number: %v\n", err)
      }
      fmt.Println("Converted integer:", i)
      ```
# Object Orientation in Go
## No inheritence
- Go does not support inheritence like python, c++, java does.
- **Go takes a radically different approach to object orientation**.
- ***Go avoids talking about classes, objects, instances altogether***
- ***Go talks only about "types" and "values" where values of certain types have methods***.

## Aggregation and embedding
- Go supports ***aggregation*** and ***embedding*** in struct types.
- ```go
    type ColoredPoint struct {
      x, y        int // Named fields (aggregation)
      color.Color     // Anonymous field (embedding)
    }

    func main() {
      cp := ColoredPoint{}
      fmt.Printf("%#v", cp)
    }
    ```
- **Anonymous fields** have no variable name and are called ***embedded fields***.
- **Named fields** are called ***named aggregate fields***.

## Type-safe duck typing
- Go has excellent support for duck typing.
- ***Interfaces*** provide you with the flexibity of doing things that duck typing allows you to.
- If a value has the methods that an interface requires then that value can be used anywhere the interface is expected.
- It does'nt matter what the actual type of the value is, it only matters that it fulfills the required interface.
  
## Seperation of ***Interfaces, values, methods***.
- Interfaces are used to specify method signatures.
- Structs are used to specify aggregated and embedded values.
- Methods are used to specify operations on custom types.
- **There is explicit connection b/w a custom type's methods and any particular interface**.
- ***But if the type's methods fulfill one or more interfaces, values of that type can be used wherever values of those interfaces are expected***.
  
## IS-A and HAS-A relationship
- The ***is-a*** relationship are defined by interfaces, i.e purely in terms of ***method signatures***.
- A value that satisfies io.Reader interface ***is'nt a reader*** because of what it ***is***, ***but because of what methods it provides***, in this case ***Read()*** .
- The ***has-a*** relationship is expressed by using structs in which we aggregate or embed values of particular types.

## Built-in types and methods.
- We cannot add methods to built-in types.
- But it is very easy to create custom types based on built-in times and add any methods to it.

## Custom types
- A **custom type** is defined on a **base type**.
- The base type can be any built-in type like ints, bools, slices, channels or **structs** or **interfaces** or **function signature**.
- Go has a **type statement** with the syntax
  - **`type typeName baseTypeSpecification`**
- ***Custom types based on structs and interfaces are much more powerful than other base types because structs and interfaces provide features like embedding and aggregation which play a vital role in Go's object-orientation***.
- ```go
    type Count int
    type StringMap map[string]string
    type IntList []int

    var i Count = 7
    fmt.Printf("i: Type: %T and Value: %v \n", i, i) // i: Type: main.Count and Value: 7

    var ownerOf = StringMap{
      "rocky": "johncena",
    }
    fmt.Printf("ownerOf: Type: %T and Value: %#v \n", ownerOf, ownerOf) // ownerOf: Type: main.StringMap and Value: main.StringMap{"rocky":"johncena"}

    var scores = IntList{
      4, 7, 3, 2,
    }
    fmt.Printf("scores: Type: %T and Value: %v", scores, scores) // scores: Type: main.IntList and Value: [4 7 3 2]
    ```

## Advantages of custom types
- ***The interface and struct custom types have the ability to embed and aggregate other custom types which form the basis of object-orientation programs in Go***
- It improves program readability and provide an abstraction mechanism.
- ***Custom function types***
  - **When working with higher order functions it is convenient to define custom types for the signature of the functions we want to pass.**
  
- ```go
    type StringForString func(string) string 
    // A custom type that specifies a function signature

    // removePunctuation function has the type specified by the "StringForString" custom type
    func removePunctuation(testString string) string {
      newString := ""
      for _, runeCharacter := range testString {
        switch character := string(runeCharacter); character {
        case "a", "e", "i", "o", "u":
        default:
          newString += string(character)
        }
      }
      return newString
    }

    func processStrings(removePunctuation StringForString, strings []string) []string {
      result := []string{}
      for _, testString := range strings {
        result = append(result, removePunctuation(testString))
      }
      return result
    }

    var stringsWithoutVowels = processStrings(removePunctuation, []string{
      "apple",
      "peanut",
      "domestic",
    })

    fmt.Printf("%v \n", stringsWithoutVowels) // [ppl pnt dmstc]
    ```

## **Go methods**
- A method is a special kind of function that can be called on a value of a custom type.
- The value is passed as a pointer or a value depending upon how the method is defined.
- A method has a **receiver** b/w func keyword and the name.
- The **receiver type** is type the method belongs to. 
- So now, the function is associated with a **custom type**.
- We can call all the functions that are defined to be methods of the custom type on values of that type.
- When the method is called, the receiver's variable(if present) is automatically set to the value or the pointer on which the method is called.

## Pointer receivers, Value receivers and indirection
- ```go
      type Vertex struct {x, y int}

      func (pv *Vertex) Scale(f float64) {
        pv.x = pv.x * f
        pv.y = pv.y * f
      }

      func (pv *Vertex) Move(dx, dy float64) {
        pv.x = pv.x + dx
        pv.y = pv.y + dy
      }

      /* Scale and Move are functions that are defined to be received by values of type *Vertex */

      func main() {
        v := Vertex{3, 4}
        v.Scale(2) // Go interprts this as (&v).Scale(2)
        fmt.Printf("%#v \n", v) // main.Vertex{x: 6, y:8}
        v.Move(5, 5) // Go interprets this as (&v).Move(5, 5)
        fmt.Printf("%#v \n", v) // main.Vertex{x: 11, y:13}
      }
    ```
- **Pointer receivers**:- Methods with a pointer receiver operates on pointer to a value and can modify the value on which the method is called through indirection.
- **Value receivers**:- Methods with a value receiver operates on a copy of the original value on which the method is called and hence it cannot modify the original value.
- **Methods often need to modify their receiver, so pointer receivers are much more common than value receivers**.
- **Methods with pointer receivers can be called on either a pointer or a value**.
- **Go provides this as a convinience for method calls, it iterprets `v.Scale(2)` as `(&v).Scale(2)`**
- ```go
    type Vertex struct {x, y int}

    func (v Vertex) Scale(f float64) {
      v.x = v.x * f
      v.y = v.y * f
    }

    func (v Vertex) Move(dx, dy float64) {
      v.x = v.x + dx
      v.y = v.y + dy
    }

    /* Scale and Move are functions that are defined to be received by values of type Vertex */

    func main() {
      v := Vertex{3, 4}
      (&v).Scale(2) // Go interprts this as (*&v).Scale(2)
      fmt.Printf("%#v \n", v) // main.Vertex{x: 3, y:4}
      (&v).Move(5, 5) // Go interprets this as (*&v).Move(5, 5)
      fmt.Printf("%#v \n", *pv) // main.Vertex{x: 3, y:4}
    }
    ```
- **Methods with value receivers can be called on either a value or a pointer**.
- **Go provides this as a convenience also, it interprets `(&v).Scale(2)` as `(*&v).Scale(2)`**

## Method set
- ***Method set of a type*** - A type's method set consists of all the methods that can be called on a value of that type.

## Embedded Fields
- structs in Go can include one or types as ***embedded fields***
- The fields of embedded type can the accessed by embeder through name of the embedded Type if the embeder also has a field with the same name.
- If there is no field in embeder with same name of embedded's field's then you can access the embedded'd fields through the plain dor notation.
- ```go
    type Item struct {
      id string
      price float64
      quantity int
    }

    func (item *Item) Cost() float64 {
      return item.price * float64(item.quantity)
    }
    
    // SpecialItem is a custom type that embeds the Item type
    type SpecialItem struct {
      Item // Anonymous field (embedding)
      catalogID int
    }

    func main() {
      specialItem := SpecialItem{
        Item{"1xd4r", 23.56, 3}, 1234569,
      }
      fmt.Printf("%#v \n", specialItem) 
      // main.SpecialItem{Item:main.Item{id:"1xd4r", price:23.56, quantity:3}, catalogID:1234569}
      fmt.Printf("%v \n", specialItem.Cost()) // 70.67999999999999
      // Can call Item type's Cost() method on SpecialItem type

      fmt.Printf("specialItem.id: %v, specialItem.price: %v, specialItem.quantity: %v, specialItem.catalogID: %v", specialItem.id, specialItem.price, specialItem.quantity, specialItem.catalogID)
      // specialItem.id: 1xd4r, specialItem.price: 23.56, specialItem.quantity: 3, specialItem.catalogID: 12
    }
    ```
- **Key concept of embedding** - If a type B embed's type A, then we can call type A's methods on type B's values.
- **Also, when a method of the embedded type is called on the embeder's type, only the embedded value is passed to the method, not the embeder's value**.

## Overriding methods
- You can define methods for the embeder type that override methods of the embedded type.
- ```go
    type LuxuryItem struct {
      Item // embedded type
      markup float64
    }

    func (luxuryItem *LuxuryItem) Cost() float64 {
      return luxuryItem.Item.Cost() * item.markup // Makes use of embedded type's method.
    }

    func main() {
      luxuryItem := LuxuryItem{
        Item{"1xd4r", 23.56, 3}, 3.5,
      }
      fmt.Printf("%#v \n", luxuryItem) 
      // main.LuxuryItem{Item:main.Item{id:"1xd4r", price:23.56, quantity:3}, markup:3.5}
      fmt.Printf("%v \n", luxuryItem.Cost()) // 247.37999999999997
    }
     ```

## Interfaces
- An ***Interface*** is **custom type** in Go that specifies a set of ***method signatures***.
- A type is said to ***satisfy*** an interface when the type has all the methods which the interface requires.
- Values of such types can be used wherever it's interface type is expected.
- Interfaces on their are of no use, we need ***concrete types*** that implement them to make them useful.
- ***In Go, duck typing is acheived by using interfaces.***
- **Interface naming convention**: In Go, names of interfaces should end with 'er'. ex:- Exchanger, Stringer etc.

## No "implements" statement
- ***Also, there is no need of explicit statements to specify that a particular custom type implements a particular interface using keywords like "inherits", "extends", "implements".*** 
- ***It is sufficient for a custom type to provide the methods specified by an interface for Go to know that the type satisfies the interface***
- ***The above feature of Go makes it very flexible, i.e it becomes very easy to add new interfaces, new types and methods without thinking about any concept of inheritence.***

## Interface Values
- **Under the hood, value of an interface type can be thought of as a tuple of a value and a concrete type - (value, concrete type)**.
- An interface holds a value of a specific underlying concrete type.
- ```go
    // Exchanger is a custom interface type that specifies a single method
    type Exchanger interface {
      Exchange()
    }

    // StringPair is a custom structure type that aggregates 2 string type fields.
    type StringPair struct {
      first, second string
    }

    // Exchange is a method on *StringPair type
    func (pair *StringPair) Exchange() {
      pair.first, pair.second = pair.second, pair.first
    }

    // String is a method on StringPair type
    func (pair StringPair) String() string {
      return fmt.Sprintf("%q+%q", pair.first, pair.second)
    }

    // Point is a custom integer array type
    type Point [2]int

    // Exchange is a method on *Point type
    func (point *Point) Exchange() {
      point[0], point[1] = point[1], point[0]
    }

    // String is a method on Point type
    func (point Point) String() string {
      return fmt.Sprintf("%v, %v", point[0], point[1])
    }

    // ExchangeThese takes Exchanger slice type as asrgument and calls the Exchange() method on elements of the Exchanger slice
    func ExchangeThese(exchangers []Exchanger) {
      for _, exchanger := range exchangers {
        exchanger.Exchange()
      }
    }

    func main() {
      var cavil Exchanger  // variable of type Exchanger interface
	cavil = &StringPair{"henry", "cavil"} // value of the Exchanger interface type
      point := Point{3, 7}
      fmt.Printf("before Exchange call:\n %v %v \n", cavil, point)
      // before Exchange call: "henry"+"cavil" 3, 7

      cavil.Exchange()
      point.Exchange()
      fmt.Printf("after Exchange call:\n %v %v \n", cavil, point)
      // after Exchange call: "cavil"+"henry" 7, 3

      banner := StringPair{"bruce", "banner"}
      fmt.Printf("before exchangeThese call:\n %v %v %v \n", cavil, banner, point)
      // before exchangeThese call: "cavil"+"henry" "bruce"+"banner" 7, 3

      exchangers := []Exchanger{
        // The following pointers are of Exchanger type because they have a method with signature Exchange().
        cavil, &banner, &point,
      }
      ExchangeThese(exchangers)
      fmt.Printf("after exchangeThese call:\n %v %v %v \n", cavil, banner, point)
      // after exchangeThese call: "henry"+"cavil" "banner"+"bruce" 3, 7

      readers := []io.Reader{
        // The following pointers are of io.Reader type because they have a method with signature Read([]byte) (int, error)
        &StringPair{"joe", "rogan"},
        &StringPair{"bobby", "lee"},
      }

      for _, reader := range readers {
        raw, err := ToBytes(reader, 16)
        if err != nil {
          fmt.Println(err)
        }
        fmt.Println(raw)
        fmt.Printf("%q\n", raw)
      }
      // [106 111 101 114 111 103 97 110]
      // "joerogan"
      // [98 111 98 98 121 108 101 101]
      // "bobbylee"
    }
    ```
- Some built-in interfaces:- 
  - **`fmt.Stringer`** - specifies 1 method
  - **`io.Reader`** - specifies 1 method
  - **`io.Writer`** - specifies 1 method
- The empty interface type **`interface{}`** is satisfied or implemented by all types. 
- So every value is of the type **`interface{}`**.

## Type assertions
- **Type assertion provides access to an interface value's underlying concrete type.**
- **Type assertion is a form of type convertion in which you convert a value an **`interface`** type into the value's underlying concrete type**.
- The same type convertion can be acheived using **type switch** and **introspection** using reflect package.
- Syntax for type assertions
  - **`resultOfConcreteType, boolean := expression.(concreteType)`** // Checked
  - **`resultOfConcreteType := expression.(concreteType)`** // Unchecked; panic() on failure.
- **```t := i.(T)```**
- The statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to a variable t.
- ```go
    func main() {
      var i interface{} = 99
      var s interface{} = []string{"bob", "dylan"}
      /* Note that if we printed the original i and s variables (both of type interface{}) they would be printed as an int and a []string. This is because when the fmt package’s print functions are faced with interface{} types, they are sensible enough to print the actual underlying values. */


      if i, ok := i.(int); ok {
        fmt.Printf("%T→%d\n", i, i) // i is a shadow variable of type int
        // int→99
      }

      if s, ok := s.([]string); ok {
        fmt.Printf("%T→%q\n", s, s) // s is a shadow variable of type []string
        // []string→["bob" "dylan"]
      }

      // Checking if a value is of of a specific interface type and only
      // if it is, call the methods specified by that interface on that
      // value
      type IsValider interface {
        IsValid() bool
      }

      // x is a variable defined here

      if thing, ok := x.(IsValider) ; ok {
        thing.IsValid()
        // do stuff
      } else {
        // x is not an IsValider
      }
    }
    ```
- It is quite common when doing type assertions to use the same name for the result value as for the original value i.e to shadow variables.

## Interface embedding
- Interfaces have excellent support for embedding other interfaces.
- The effect is that the embedded interface's method signatures become a part of the embeder's interface.
- This way we can make use of existing interfaces and create new interfaces very easily.
- **A pattern of tiny methods that rely on functions to do the work is very common in Go.**

- ```go
    type StringPair struct {
      first, second string
    }

    type LowerCaser interface {
      LowerCase()
    }

    type UpperCaser interface {
      UpperCase()
    }

    type LowerUpperCaser interface {
      LowerCaser // As if we had written LowerCase()
      UpperCaser // As if we had written UpperCase()
    }

    type FixCaser interface {
      FixCase()
    }

    type ChangeCaser interface {
      LowerUpperCaser // As if we had written LowerCase(); UpperCase()
      FixCaser        // As if we had written FixCase()
    }

    func (pair *StringPair) UpperCase() {
      pair.first = strings.ToUpper(pair.first)
      pair.second = strings.ToUpper(pair.second)
    }

    func (pair *StringPair) LowerCase() {
      pair.first = strings.ToLower(pair.first)
      pair.second = strings.ToLower(pair.second)
    }

    func (pair *StringPair) FixCase() {
      pair.first = fixCase(pair.first)
      pair.second = fixCase(pair.second)
    }

    func fixCase(s string) string {
      var chars []rune
      upper := true
      for _, char := range s {
        if upper {
          char = unicode.ToUpper(char)
        } else {
          char = unicode.ToLower(char)
        }
        chars = append(chars, char)
        upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
      }
      return string(chars)
    }

    func main() {
      lobelia := StringPair{"LOBELIA", "SACKVILLE-BAGGINS"}
      // Values of *StringPair belong to all of these custom interface types
      // 1. LowerCaser
      // 2. UpperCaser
      // 3. LowerUpperCaser
      // 4. FixCaser
      // 5. ChangeCaser
      //
      lobelia.FixCase()
      fmt.Printf("%#v \n", lobelia)
      // main.StringPair{first:"Lobelia", second:"Sackville-Baggins"}
    }
    ```

## Structs
- Structs in Go are used aggregate and embed values together.
- Really useful when values are of different types, slice cannot do this because it needs all of it's values to be of the same type (although we can use type interface{} to store values of any type in a slice).
- ***Go allows us to annotate struct fields with strings (called tags in Go terminology).***
- ***Go allows to create and initialize structs with only those fields we want using the syntax fieldName: fieldValue***.
- ***Anonymous structs*** - These structs can be used like built-in types to declare types and values (with composite literal syntax)
- ```go
    // with 'var' declaration and NO explicit initialization
    var name struct{ first, last string }
    name.first = "andres"
    name.last = "iniesta"
    fmt.Printf("%#v \n", name)


    // SVD with struct type spec and composite literal
    // ex: 1
    person := struct{
      age int 
      name string
      isMarried bool
      ownsABitch bool
    }{
      age: 20,
      name: "lionel messi",
      // the rest of the fields are initialized to their zero value
    }
    fmt.Printf("%#v \n", person)
    // {20 lionel messi false false}

    // ex: 2 // by not including the field names in composite literal
    points := []struct{ x, y int }{
      {4, 6},
      {},
      {-7, 11},
      {15, 17},
      {4, 8},
    }
    for _, point := range points {
      fmt.Printf("(%d, %d) \n", point.x, point.y)
    }
    /*
      (4, 6)
      (0, 0)
      (-7, 11)
      (15, 17)
      (4, 8)
    */
    ```

## Struct embedding.
- Structs can be embedded with other structs just like we do for interfaces, by including the type name of a struct as an anonymous field inside another struct.
- ***The purpose of embedding struct's is to include the fields of those struct's in the ember's struct, and Go makes this very easy***.
- **Embedded field's name can be accessed directly by the .(dot) selector operator without mentioning the type name if there is no collision b/w the names of the embedded struct and the embeder struct**.
- ***Every field name of a struct must be unique, if not Go raises compilation error***.

## Embedding values
- ```go
    // Person is a custom structure type
    type Person struct {
      Name      string
      ForeNames []string
      SurName   string
    }

    // Author is custom structure type that embed's the Person structure type, so that field's of Person type will be available in this type.
    type Author struct {
      Person // anonymous field (embedding)
      BookTitles []string
      YearBorn   int
    }

    func main() {
      author := Author{
        Person{"robert", []string{"louis", "bob", "max"}, "martin"},
        []string{"GOT", "GOT2", "GOT3", "GOT4"},
        1950,
      }

      fmt.Println("author.Person \n", author.Person)
      // author.Person
      // {robert [louis bob max] martin}

      fmt.Println("author.Name, author.ForeNames, author.SurName \n", author.Name, author.ForeNames, author.SurName)
      // author.Name, author.ForeNames, author.SurName
      // robert [louis bob max] martin

      fmt.Println("author.BookTitles, author.YearBorn \n", author.BookTitles, author.YearBorn)
      // author.BookTitles, author.YearBorn
      // [GOT GOT2 GOT3 GOT4] 1950
    }
    ```
## Embedding values that have methods
- By embedding values that have methods, the embeder value will get those methods.
- ***When an embedded field's method is called on a value of the embeder, it is only the embedded field value that gets passed as the method's receiver, not the embeder value***

## Embedding interfaces
- You can also embed interface types as fields in a struct.
- By doing so, you can assign a value which satisfies the interface to that interface type field.
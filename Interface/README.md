# Интерфейсы

## Определение и полиморфизм

Интерфейс - это набор сигнатур методов (список методов без реализации)

Например, есть интерфейс **Shape**, в котором описан метод получения **Периметра**

```
type Shape interface {
	Perimeter() int
}
```

У каждой фигуры своя формула получения периметра. Для того, чтоб удовлетворять интерфейсу **Shape**, нужно реализовать метод **Perimeter() int**

```
// Треугольник
type Triangle struct {
	a, b, c int
}

func (t Triangle) Perimeter() int {
	return t.a + t.b + t.c
}
```

Теперь мы можем написать общий метод, который будет принимать на вход структуру, которая удовлетворяет интерфейсу **Shape**. Любая структура, где реализован метод **Perimeter() int** будет здесь работать

```
func PrintSomeInfo(shape Shape) int {
	fmt.Printf("Фигура: %T\n", shape)
	fmt.Printf("Параметры фигуры: %v\n", shape)
	fmt.Printf("Периметр фигуры: %d\n", shape.Perimeter())
	return 0
}
```

Также можно вызывать метод **Perimeter() int** непосредственно через **Shape**

```
Shape.Perimeter(Rectangle{4, 8})
Shape.Perimeter(Triangle{3, 4, 5})
```


## Встраивание интерфейса (interface embedding)

## Пустой интерфейс

Интерфейс называют пустым, если у него нет методов

`interface{}`

Пустые интерфейсы используют, если тип значения заранее неизвестен. Например, функция `fmt.Println`:

`func Println(a ...interface{}) (n int, err error)`

`any` - псевдоним для `interface{}`. Разницы между ними нет

## Интерфейсы и nil

Переменная типа interface представляется как пара `(type, value)`. Пока интерфейсной переменной не присвоено значение, у нее и `type`, и `value` равны `nil`, поэтому сама переменная считается равной `nil`

Если присвоить интерфейсу тип с nil-value, то интерфейс перестает быть равен `nil`.

```
type greeter interface {
    greet()
}

type english struct {
    name string
}

// e может быть nil!
func (e *english) greet() {
    if e == nil {
        fmt.Println("Value is nil")
        return
    }
    fmt.Println("Hello", e.name)
}

var ivar greeter
ivar = (*english)(nil)
ivar.greet()
// Value is nil
```

## Приведение типа. Переключатель типа

*Приведение типа* позволяет извлекать конкретное значение из переменной интерфейсного типа:

```
var value any = "hello"
str := value.(string)
fmt.Println(str)
// hello
```

Если тип конкретного значения отличается от указанного, то появится ошибка

```
fl := value.(float64)
// ошибка
```

Проверить тип можно следующим образом:

```
str, ok := value.(string)
fmt.Println(str, ok)  // hello true

fl, ok := value.(float64)
fmt.Println(fl, ok)  // 0 false
```

Приведение типа можно использовать вместе со `switch`. Такая конструкция называется *переключателем типа*:

```
var value any = "hello"

switch v := value.(type) {
case string:
    fmt.Printf("%#v is a string\n", v)
case float64:
    fmt.Printf("%#v is a float\n", v)
default:
    fmt.Printf("%#v is a mystery\n", v)
}
// "hello" is a string
```



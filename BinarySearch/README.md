# Алгоритм бинарного поиска

Рекурсивный алгоритм поиска в отсортированном массиве со сложностью O(log n)

Учитываем, что на вход может прийти пустой срез:

```
if len(sortedSlice) == 0 {
    return -1, fmt.Errorf("элемент %d не найден", x)
}
```

Учитываем ситуацию, когда в срезе только один элемент:
```
if len(sortedSlice) == 1 {
    if sortedSlice[0] == x {
        return x, nil
    } else {
        return -1, fmt.Errorf("элемент %d не найден", x)
    }
}
```

Вычисляем середину среза:

```
middleIndex := len(sortedSlice)/2 + len(sortedSlice)%2
```

Сравниваем элемент в середине массива с числом, которое ищем:

```
if sortedSlice[middleIndex] == x {
    return x, nil
}
```

Если элемент в середине больше числа, то ищем в правой части, иначе в левой:
```
if sortedSlice[middleIndex] > x {
    return BinarySearch(sortedSlice[:middleIndex], x)
} else {
    return BinarySearch(sortedSlice[middleIndex:], x)
}
```

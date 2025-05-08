# Тесты

Тесты пишут в отдельном файле с суффиксом `_test`
Тест - это функция с префиксом `Test`

```
func TestFirst(t *testing.T) {
	// Исходные данные
	got := GetSum(5, 6)
	want := 11

	// Проверка
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
```

Выполнить тест: `go test -v`

Чтобы оценить тестовое покрытие, нужно выполнить команду: `go test -cover`

Также можно собрать детальную статистику (profile) и посмотреть отчет в HTML:

```
go test -coverprofile="cover.prof"
go tool cover -html="cover.prof"
```

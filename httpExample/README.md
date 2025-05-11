# HTTP в Go

Пакет net/http в Go отвечает одновременно и за HTTP-клиенты, и за HTTP-серверы.

## Запросы

За HTTP-запросы отвечает структура `http.Client`:

`client := http.Client{Timeout: 3 * time.Second}`

Таймаут - максимальное время, которое клиент готов ждать ответ от сервера, прежде чем вернет ошибку. По умолчанию таймаут не задан, поэтому ждать клиент будет до бесконечности. Лучше всегда явно указывать его.

Клиента достаточно создать один раз и дальше использовать для всех запросов

## Запрос с параметром

Параметры в ссылке  
`GET https://httpbingo.org/get?id=42`

Первым делом создаем клиента и запрос:

```
// Создаем клиента
client := http.Client{Timeout: 3 * time.Second}

// Создаем запрос
const uri = "https://httpbingo.org/get"
req, err := http.NewRequest(http.MethodGet, uri, nil)
if err != nil {
    panic(err)
}
```

Затем нужно наполнить запрос параметрами и выполнить его:

```
// Наполняем запрос
params := url.Values{}
params.Add("id", "42")
req.URL.RawQuery = params.Encode()

// Выполняем запрос
resp, err := client.Do(req)
if err != nil {
    panic(err)
}
```

## Заголовки запроса

Редактируются внутри запроса:

```
// Редактируем заголовки
req.Header.Add("Accept", "application/json")
req.Header.Add("X-Request-Id", "42")
```

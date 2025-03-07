Калькулятор выражений через API
Этот проект предоставляет API для вычисления математических выражений. Вы можете отправить POST-запрос с выражением, и сервер вернет результат вычисления.

Требования
Убедитесь, что у вас установлен curl для Windows. Если нет, скачайте его с официального сайта или используйте Git for Windows, который включает curl.

Использование API
Отправка запроса
Чтобы вычислить выражение, отправьте POST-запрос на эндпоинт /api/v1/calculate с JSON-телом, содержащим выражение.

Пример команды для Windows Command Prompt (cmd):

```
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"
```
Параметры запроса
URL: http://localhost:8080/api/v1/calculate

Метод: POST

Заголовок: Content-Type: application/json

Тело запроса: JSON-объект с ключом expression, содержащим строку с математическим выражением.

Пример тела запроса:

```
{
  "expression": "2+2*2"
}
```
Пример ответа
Если сервер работает корректно, вы получите ответ в формате JSON:

```
{
  "result": 6
}
```
Примеры команд
Вычисление простого выражения:

```
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"3+5*2\"}"
```
Ожидаемый ответ:
```
{
  "result": 13
}
```
Вычисление выражения с дробными числами:
```
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"10.5 / 2 + 3\"}"\
```
Ожидаемый ответ:
```
{
  "result": 8.25
}
```
Вычисление выражения со скобками:
```
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{\"expression\": \"(2+3)*4\"}"
```
Ожидаемый ответ:
```
{
  "result": 20
}
```
Возможные ошибки
Сервер не запущен:
Если сервер не запущен на localhost:8080, вы получите ошибку подключения:
```
curl: (7) Failed to connect to localhost port 8080: Connection refused
```
Неправильный формат JSON:
Если JSON-данные неверно сформированы, сервер может вернуть ошибку:
```
{
  "error": "Invalid JSON format"
}
```
Неподдерживаемое выражение:
Если выражение содержит недопустимые символы или операции, сервер может вернуть ошибку:
```
{
  "error": "Invalid expression"
}
```
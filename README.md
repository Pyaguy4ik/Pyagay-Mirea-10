# Практическое задание № 10

# JWT-аутентификация: создание и проверка токенов. Middleware для авторизации

Студент группы *ЭФМО-02-25 Пягай Даниил Игоревич*

# Описание

**Цели:**

    • Понять устройство JWT и где его уместно применять в REST API. 
    • Сгенерировать и проверить JWT в Go (HS256), передавать его в Authorization: Bearer …. 
    • Реализовать middleware-аутентификацию (достаёт токен, валидирует, кладёт клеймы в context). 
    • Добавить middleware-авторизацию (RBAC/права на эндпоинты). 
    • Встроить это в уже знакомую архитектуру HTTP-сервиса/роутера.


## Инициализация проекта

```bash
mkdir -p ~/pz10-auth
cd ~/pz10-auth
go mod init pz10-auth
go get github.com/go-chi/chi/v5
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

## Создаём структуру файлов и запускаем сервер

![structure](img/structure.jpeg)

# Проверка Health
![login](img/health.jpeg)

# Создаём заметку
![login](img/create_note.jpeg)

# Запрашиваем все заметки
![login](img/get_all_notes.jpeg)

# Обновляем заметку
![login](img/update.jpeg)

# Удаляем заметку
![login](img/delete.jpeg)

# Пробуем найти заметку по id после удаления
![login](img/after_delete.jpeg)

# Создание заметки с невалидными данными
![login](img/post_error_novalid.jpeg)

# Пытемся обновить заметку с несуществующим id
![login](img/update_not_found.jpeg)

## Список проведённых запросов
![post](img/all_requests.jpeg)

# Ответы сервера на проведённые запросы
![login](img/answer.jpeg)

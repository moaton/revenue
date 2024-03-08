# Revenue managment Go

### Структура
```bash
└───backend
    ├───cmd
    ├───docs
    ├───internal
    │   ├───app
    │   ├───dto
    │   ├───entity
    │   ├───interfaces
    │   ├───repository
    │   │   ├───report
    │   │   ├───revenue
    │   │   └───users
    │   ├───transport
    │   │   └───http
    │   │       └───handlers
    │   │           ├───middleware
    │   │           ├───revenue
    │   │           ├───users
    │   │           └───utils
    │   ├───types
    │   ├───usecase
    │   └───utils
    └───pkg
        └───gorm
            └───postgres
                └───migrations
```

### Unit тесты
>   Для запуска тестов:

```bash
make test
```

>   Для проверки покрытия тестов:

```bash
make view_test
```

>   Для проверки покрытия
```bash
go test -v ./ -coverprofile=filename
go tool cover -html=filename
```
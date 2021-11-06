### Создаем миграцию
```bash
    migrate create -ext sql -dir migrations UsersCreationMigration #UsersCreationMigration - имя миграции
```

### Выполняем миграции
```bash
    migrate -path migrations -database postgres://postgres:postgres@localhost:5432/restapi up #up - создаст новую схему, down - откатит миграции
```
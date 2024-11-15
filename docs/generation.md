## Code generating

### Library:
https://github.com/deepmap/oapi-codegen

### Install:
```
$ go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
$ oapi-codegen -version
```

### Запуск генерации
1. Надо добавить в swagger.yaml описание метода
2. Сгенерировать серверную часть

Есть файлы dto.cfg.yml и server.cfg.yml и в них закомментирована команда, которую надо запустить, после изменения swagger файла
# TemplateRestAPIGolang

Небольшой шаблон API на языке Golang

## Структура проекта

* **.infra** - инфраструктурные вещи
* **cmd** - наши сервисы
* **docs** - общая документация
* **internal** - вся внутрянка бекенда
  * **api** - слой контроллеров
  * **common** - файлы конфигурации
  * **domain** - слой моделей
  * **server** - кастомный HTTP-сервер
  * **telemetry** - логи и метрики
* **pkg** - публичные библиотечки


Сваггер доступен по пути http://localhost:8080/api/swagger/index.html

## Сборка проекта


### Бинарники

```bash
 # сборка
 $ cd ./cmd
 $ go mod tidy
 $ go build -o ../../out .
 $ cd ../..
 # запуск (на dev-окружении)
 $ ./out dev.yml ./cmd/config
```

### Swagger

```bash
 $ swag init --output cmd/docs/ --parseInternal \
    -d cmd,internal/domain/base,internal/domain/entity,internal/domain/enum,internal/api
```

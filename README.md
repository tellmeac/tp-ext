# Подготовка к разработке

## Установка telepresence

```bash
sudo curl -fL https://app.getambassador.io/download/tel2/linux/amd64/latest/telepresence -o /usr/local/bin/telepresence
sudo chmod a+x /usr/local/bin/telepresence
```

## Подготовка Helm Chart для тестирвоания

2 контейнера для общения между собой через Nats сервер
Один из них может обрабатывать http запросы
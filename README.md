# Projektname

Für Webprogrammierung eine kleine Wikisoftware.
Von Johannes Foltin

## Voraussetzungen

Stellen Sie sicher, dass die folgenden Voraussetzungen erfüllt sind, bevor Sie das Programm starten:

- Docker installiert und ausgeführt
- Docker Compose installiert und ausgeführt

## Installation

1. Klone das Repository:

    ```shell
    git clone https://github.com/JohannesFoltin/wikijedia.git
    ```

2. Baue das Programm mit Docker Compose:

    ```shell
    docker-compose build
    ```

3. Starte die Applikation

    ```shell
    docker-compose up -d
    ```

    Dadurch werden die erforderlichen Container gestartet und das Programm wird im Hintergrund ausgeführt.

Es ist erreichbar unter Port 80. (http://localhost:80)

## Dev

### Im Frontend:
1. 
```shell
npm i
```
2. 

```shell
npm run dev
```

### Im Backend:

```shell
go run *.go
```
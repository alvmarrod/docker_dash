# Docker Dashboard Project

## How to run it

Please notice that this is a development procedure. If you want to keep it production ready you should build the binary file for backend and compile the ReactJS frontend.

1. Run the backend with

```
go run main.go restapi.go external.go log.go docker.go aux.go
```

2. Run the frontend with

```
npm start
```
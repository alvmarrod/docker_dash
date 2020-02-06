# Docker Dashboard Project

This projects aims to provide an easy-to-use non-CLI interface to manage Docker locally. Its user target is those people who could benefit from using Docker but are not used to its commands or not even technical.

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

## Dependencies

### GO Dependencies

* [Gorilla/Mux](https://github.com/gorilla/mux)

### React Dependencies

* [React-Switch](https://www.npmjs.com/package/react-switch)
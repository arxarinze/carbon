# Paster Binner

<!-- GETTING STARTED -->
## Getting Started Development

To run this as a development enviroment you must download and setup golang.
<p align="right">(<a href="https://go.dev/doc/install">Golang Setup</a>)</p>

### Prerequisites

***change .env.example to .env and use***

***After setting up golang open project root in vscode and use terminal to run these commands:***

```
  go mod tidy
  go run main.go
```



## Just Want To Run The Project
***A binary has been setup for you in the repo under build***

***change .env.example to .env and add into the folder for your os e.g osx/x64/.env***

### Linux/Mac
``` ./carbon ```

### Windows
```Run the .exe file```


## Postman Collection
```Postman collection is included in the build folder just import and use```

## Run Test
***From the root of the project***
``` 
  cd internal/services
  go test 
```
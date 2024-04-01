# go-webapp-sample

[![license](https://img.shields.io/github/license/ybkuroki/go-webapp-sample?style=for-the-badge)](https://github.com/ybkuroki/go-webapp-sample/blob/master/LICENSE)
[![report](https://goreportcard.com/badge/github.com/ybkuroki/go-webapp-sample?style=for-the-badge)](https://goreportcard.com/report/github.com/ybkuroki/go-webapp-sample)
[![workflow](https://img.shields.io/github/actions/workflow/status/ybkuroki/go-webapp-sample/check.yml?label=check&logo=github&style=for-the-badge)](https://github.com/ybkuroki/go-webapp-sample/actions?query=workflow%3Acheck)
[![release](https://img.shields.io/github/release/ybkuroki/go-webapp-sample?style=for-the-badge&logo=github)](https://github.com/ybkuroki/go-webapp-sample/releases)

## Preface
This repository is the sample of web application using golang.
This sample uses [Echo](https://echo.labstack.com/) as web application framework, [Gorm](https://gorm.io/) as OR mapper and [Zap logger](https://pkg.go.dev/go.uber.org/zap) as logger.
This sample application provides only several functions as Web APIs.
Please refer to the 'Service' section about the detail of those functions.

Also, this application contains the static contents such as html file, css file and javascript file which built [vuejs-webapp-sample](https://github.com/ybkuroki/vuejs-webapp-sample) project to easily check the behavior of those functions.
So, you can check this application without starting a web server for front end.
Please refer to the 'Starting Server' section about checking the behavior of this application.

If you would like to develop a web application using golang, please feel free to use this sample.

## Install
Perform the following steps:
1. Download and install [Visual Studio Code(VS Code)](https://code.visualstudio.com/).
1. Download and install [Golang](https://golang.org/).
1. Get the source code of this repository by the following command.
    ```bash
    go install github.com/ybkuroki/go-webapp-sample@latest
    ```

## Starting Server
There are 2 methods for starting server.

### Without Web Server
1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
1. When startup is complete, the console shows the following message:
    ```
    http server started on [::]:8080
    ```
1. Access [http://localhost:8080](http://localhost:8080) in your browser.
1. Login with the following username and password.
    - username : ``test``
    - password : ``test``

### With Web Server
#### Starting Application Server
1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
1. When startup is complete, the console shows the following message:
    ```
    http server started on [::]:8080
    ```
1. Access [http://localhost:8080/api/health](http://localhost:8080/api/health) in your browser and confirm that this application has started.
    ```
    healthy
    ```
#### Starting Web Server
1. Clone [vuejs-webapp-sample](https://github.com/ybkuroki/vuejs-webapp-sample) project and install some tools.
1. Start by the following command.
    ```bash
    npm run dev
    ```
1. When startup is complete, the console shows the following message:
    ```
    > vuejs-webapp-sample@*.*.* dev
    > vite --mode development
    
    
    VITE v*.*.*  ready in 1362 ms
    
    ➜  Local:   http://localhost:3000/
    ➜  press h to show help
    ```
1. Access [http://localhost:3000](http://localhost:3000) in your browser.
1. Login with the following username and password.
    - username : ``test``
    - password : ``test``

## Using Swagger
In this sample, Swagger is enabled only when executed this application on the development environment.
Swagger isn't enabled on the another environments in default.

### Accessing to Swagger
1. Start this application according to the 'Starting Application Server' section.
2. Access [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.

### Updating the existing Swagger document
1. Update some comments of some controllers.
2. Download Swag library. (Only first time)
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
3. Update ``docs/docs.go``.
    ```bash
    swag init
    ```

## Build executable file
Build this source code by the following command.
```bash
go build main.go
```

## Project Map
The following figure is the map of this sample project.

```
- go-webapp-sample
  + config                  … Define configurations of this system.
  + logger                  … Provide loggers.
  + middleware              … Define custom middleware.
  + migration               … Provide database migration service for development.
  + router                  … Define routing.
  + controller              … Define controllers.
  + model                   … Define models.
  + repository              … Provide a service of database access.
  + service                 … Provide a service of book management.
  + session                 … Provide session management.
  + test                    … for unit test
  - main.go                 … Entry Point.
```

## Services
This sample provides 3 services: book management, account management, and master management.
Regarding the detail of the API specification, please refer to the 'Using Swagger' section.

### Book Management
There are the following services in the book management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Get Service|GET|``/api/books/[BOOK_ID]``|Book ID|Get a book data.|
|List/Search Service|GET|``/api/books?query=[KEYWORD]&page=[PAGE_NUMBER]&size=[PAGE_SIZE]``|Page, Keyword(Optional)|Get a list of books.|
|Regist Service|POST|``/api/books``|Book|Regist a book data.|
|Edit Service|PUT|``/api/books``|Book|Edit a book data.|
|Delete Service|DELETE|``/api/books``|Book|Delete a book data.|

### Account Management
There are the following services in the Account management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Login Service|POST|``/api/auth/login``|Session ID, User Name, Password|Session authentication with username and password.|
|Logout Service|POST|``/api/auth/logout``|Session ID|Logout a user.|
|Login Status Check Service|GET|``/api/auth/loginStatus``|Session ID|Check if the user is logged in.|
|Login Username Service|GET|``/api/auth/loginAccount``|Session ID|Get the login user's username.|

### Master Management
There are the following services in the Master management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Category List Service|GET|``/api/categories``|Nothing|Get a list of categories.|
|Format List Service|GET|``/api/formats``|Nothing|Get a list of formats.|

## Tests
Create the unit tests only for the packages such as controller, service, model/dto and util. The test cases is included the regular cases and irregular cases. Please refer to the source code in each packages for more detail.

The command for testing is the following:
```bash
go test ./... -v
```

## Libraries
This sample uses the following libraries.

|Library Name|Version|
|:---|:---:|
|echo|4.11.4|
|gorm|1.25.9|
|go-playground/validator.v9|9.31.0|
|zap|1.26.0|

## Contribution
Please read [CONTRIBUTING.md](https://github.com/ybkuroki/go-webapp-sample/blob/master/CONTRIBUTING.md) for proposing new functions, reporting bugs and submitting pull requests before contributing to this repository.

## License
The License of this sample is *MIT License*.

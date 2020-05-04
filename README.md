# go-webapp-sample

## Preface
This sample project uses [Echo](https://echo.labstack.com/) and [Gorm](https://gorm.io/) written by [Golang](https://golang.org/). It provides only Web API. So, I recommend using a [vuejs-webapp-sample](https://github.com/ybkuroki/vuejs-webapp-sample) project as Web UI.

## Install
Perform the following steps:
1. Download and install [Visual Studio Code(VS Code)](https://code.visualstudio.com/).
1. Download and install [Golang](https://golang.org/).
1. Get the source code of this repository by the following command.
    ```bash
    go get -u github.com/ybkuroki/go-webapp-sample
    ```

## Starting Server
Perform the following steps:
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

## Build executable file
Build this source code by the following command.
```bash
go build main.go
```

## Project Map
The follwing figure is the map of this sample project.

```
- go-webapp-sample
  + common                  … Provide a common service of this system.
  + controller              … Define controllers.
  + model                   … Define models.
  + repository              … Provide a service of database access.
  + service                 … Provide a service of book management.
  - main.go                 … Entry Point.
```

## Services
This sample provides 3 services: book management, account management, and master management.

### Book Management
There are the following services in the book management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|List Service|GET|``/api/book/list``|Page|Get a list of books.|
|Regist Service|POST|``/api/book/new``|Book|Regist a book data.|
|Edit Service|POST|``/api/book/edit``|Book|Edit a book data.|
|Delete Service|POST|``/api/book/delete``|Book|Delete a book data.|
|Search Title Service|GET|``/api/book/search``|Keyword, Page|Search a title with  the specified keyword.|

### Account Management
There are the following services in the Account management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Login Status Check Service|GET|``/api/account/loginStatus``|Session ID|Check if the user is logged in.|
|Login Username Service|GET|``/api/account/loginAccount``|Session ID|Get the login user's username.|

### Master Management
There are the following services in the Master management.

|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Category List Service|GET|``/api/master/category``|Nothing|Get a list of categories.|
|Format List Service|GET|``/api/master/format``|Nothing|Get a list of formats.|

## Libraries
This sample uses the following libraries.

|Library Name|Version|
|:---|:---:|
|Echo|4.1.16|
|Gorm|1.9.12|
|go-playground/validator.v9|9.31.0|

## License
The License of this sample is *MIT License*.

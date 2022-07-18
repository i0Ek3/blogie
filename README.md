# blogie

A blog backend program developed with Gin which integrates many useful features.

## Features

- fully detailed notes and standard coding

- interface validation and API access control

- integrate middlewares like link tracing, scalable

- support swagger API 

- support application configuration

- graceful shutdown and boot

- support i18n

- cross-platform

- find yourself

## Getting Started

### Build & Run

```shell
$ go build . ; ./blogie -h
```

### Usage/Example

TODO

## Architecture

In this part, we'll show you how this project work.

TODO: picture of blogie.

### Structure

```console
├── README.md
├── configs
├── docs
│   └── sql
├── global
├── go.mod
├── go.sum
├── images
│   └── gin-demo.jpg
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── main.go
├── pkg
├── scripts
│   └── install-mysql.sh
├── storage
└── third_party
```

### Database

Install MySQL first, and then create a database which named `blogie` and use this databse:

```shell
# run command under the root folder of project
$ ./scripts/install-mysql.sh # for macOS users

$ mysql -uroot -p
mysql> CREATE DATABASE blogie;
mysql> USE blogie;
mysql> SOURCE ./doc/sql/blog.sql; # import blog.sql

# after import auth.sql(contained in blog.sql), insert following sql
mysql> INSERT INTO `blogie`.`blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES (1, 'i0Ek3', 'blogie', 0, 'i0Ek3', 0, '', 0, 0);
```

After import blog.sql, it will create following four tables:

```console
mysql> show tables;
+------------------+
| Tables_in_blogie |
+------------------+
| blog_article     |
| blog_article_tag |
| blog_auth        |
| blog_tag         |
+------------------+
4 rows in set (0.00 sec)
```

### RESTful API

We use RESTful style to design our API.

#### Tag

| Function           | HTTP Method | Path      |
| ------------------ | ----------- | --------- |
| Add new tag        | POST        | /tags     |
| Delete specify tag | DELETE      | /tags/:id |
| Update specify tag | PUT         | /tags/:id |
| Get tag list       | GET         | /tags     |

#### Article

| Function               | HTTP Method | Path          |
| ---------------------- | ----------- | ------------- |
| Add new article        | POST        | /articles     |
| Delete specify article | DELETE      | /articles/:id |
| Update specify article | PUT         | /articles/:id |
| Get specify article    | GET         | /articles/:id |
| Get article list       | GET         | /articles     |

### Error Code

Resonable error code is very easy to locate problems and locate lines of code. So, we use pure numbers represent that different parts represent different services and different modules.

| Service | Module | Details                   |
| ------- | ------ | ------------------------- |
| 10      | 00     | [Common] Basic error      |
| 10      | 01     | [Common] Database error   |
| 10      | 02     | [Common] Authorized error |
| 10      | 03     | [Common] Other error      |
| 20      | 01     | Tag error                 |
| 20      | 02     | Article error             |
| 20      | 03     | Upload error              |
|         |        |                           |

For example, the error code 100001 means a basic error(success), and the error code 100202 means token error. 

### Logger

In our project, we support seven log levels: Trace, Debug, Info, Warn, Error, Fatal and Panic. But for Trace level, only for learning purpose, so decide it according your own purpose.

| Level | Details |
| ----- | ------- |
| Trace |         |
| Debug |         |
| Info  |         |
| Warn  |         |
| Error |         |
| Panic |         |
| Fatal |         |

Also our logger support output the log into file and os.Stdout.

Add you should know something here, we add lumberjack into our project, so our project support log rotation, if you don't like that, you can only use Logrotate program. In the other hand, add lumberjack will increases the complexity of the log package.

### Common Component

To ensure the standardization of the application, we will abstract the basic functions to the public component of the project.

- error code standardization

- configuration management

- database connection

- log writting

- response processing

### Interface Generation

We use swagger to generate our interface, just write comment for API and then swagger can read it and generate correspoding interface documents.

| Comment  | Details                                                                                              |
| -------- | ---------------------------------------------------------------------------------------------------- |
| @Summary | Summary to describe what function of this API                                                        |
| @Produce | Type of MIME, option: json, xml, html etc.                                                           |
| @Param   | Format of parameters, need to following rules: para_name, para_type, data_type, is_required, comment |
| @Success | Success to response, need to following rules: status_code, para_type, data_type, comment             |
| @Failure | Failed to response, need to following rules: status_code, para_type, data_type, comment              |
| @Router  | Router, need to following rules: router_addr, HTTP_method                                            |

After write comment for APIs, run command `swag init` under the root folder, that would be generated three files: docs.go, swagger.yaml, swagger.json, and then run the server `go run main.go` and open the link `http://127.0.0.1:8080/swagger/index.html` to see what happened.

### Interface Verification

Verification rules used to validate the validity of struct fields, following below:

| Tag      | Meaning                    |
| -------- | -------------------------- |
| required | Must fill something        |
| gt       | >                          |
| gte      | >=                         |
| lt       | <                          |
| lte      | <=                         |
| min      | Mininum                    |
| max      | Maximum                    |
| oneof    | One of set                 |
| len      | Required length equals len |

### Access Control

After we developed of finished some features we want to the other people to see it what it looks like, but we don't want all unreleavant people see that, so we should consider defense-in-depth and access control to API interfaces. 

There are two common API access control schemes on the market today, namely OAuth 2.0 and JWT(JSON Web Token). In our project, we choose JWT to provide access control for API interfaces.

JWT contains three parts:

- Header

- Payload

- Signature

After you finished the access control, you can generate token by run following command:

```shell
curlie -X POST \
  'http://127.0.0.1:8080/auth' \
  -H 'app_key: i0Ek3' \
  -H 'app_secret: blogie'

{"token":"eyJhbG...pXVCJ9.eyJhcH...dpZSJ9.9X4SFy...pxMcs8"}
```

## Issues

### Mod Issues

- Database driver installation
  
  - import `_ "github.com/go-sql-driver/mysql"`
  
  - import `_ "github.com/jinzhu/gorm/dialects/mysql"` for model/model.go

- Swagger files
  
  - import `swaggerFiles "github.com/swaggo/files"` for internal/routers/router.go

### Tag Issues

- After tag deleted by curl, and then you add a new tag with curl but tag id is increase instead of start as 0. This situation may not an error, just a feature in database

### Module Issues

- Upload service: curl: (26) Failed to open/read local data from file/application
  
  - use command `curlie -X POST http://127.0.0.1:8080/upload/file -F file=@./demo.jpg -F type=1` to solve it

## Reference

- Project [iam](https://github.com/marmotedu/iam)

- Project demon

- Project [go-programming-tour-book](https://github.com/go-programming-tour-book)

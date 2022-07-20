# blogie

A blog backend program developed with Gin which integrates many useful features.

## Features

- Fully detailed notes and standard coding

- Interface validation and API access control

- Integrate middlewares like link tracing, scalable

- Support swagger API 

- Support application configuration

- Graceful shutdown

- Support i18n

- Cross-platform

- Find yourself

## Getting Started

### How-To(Build & Run)

```shell
# required
$ export GO111MODULE=on ; export GOPROXY=https://goproxy.cn

# build & run
$ go mod tidy ; make ; ./blogie -h

Usage of ./blogie:
  -config string
        specify config path (default "configs/")
  -mode string
        run in which mode
  -port string
        run in which port
```

## Architecture

In this part, we'll show you how this project work.

TODO

### Database

You should install MySQL on your system first, and then create a database which named `blogie` and use this databse. If you have same system with mine(Unix), you can run the script directly to install and import blog.sql.

```shell
# for Unix/Linux users
# run command under the root folder of project
# if you are macOS user, please make sure you have homebrew installed
$ ./scripts/setup.sh

# or you want to setup it manually
$ mysql -uroot -p
mysql> CREATE DATABASE blogie;
mysql> USE blogie;
mysql> SOURCE ./doc/sql/blog.sql; # import blog.sql
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

- Error code standardization

- Configuration management

- Database connection

- Log writting

- Response processing

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

JWT contains Header, Payload, Signature three parts:

```json
Header {
    "alg": "HS256", # HMAC SHA256
    "typ": "JWT"
}

Payload { # mainly stored in the actual data transmitted in JWT
    "sub": "Topic",
    "name": "i0Ek3",
    "admin": true
}

Signature # Signature of the agreed algorithm and rules for the first two parts (Header+Payload)
   
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

After you finished the access control, you can generate token by run following command:

```shell
$ curlie -X POST \
  'http://127.0.0.1:8080/auth' \
  -H 'app_key: i0Ek3' \
  -H 'app_secret: blogie'

{"token":"eyJhbG...pXVCJ9.eyJhcH...dpZSJ9.9X4SFy...pxMcs8"}
```

### Middleware

#### Access Log

Access log basically records the request method of each request, the start time of the method call, the end time of the method call, the method response result, and the status code of the method response result. and other additional attributes to achieve the effect of log link tracking.

#### Recovery

It is very important for abnormal capture and timely alarm notification, so we need to customize the recovery middleware for our project's own conditions or ecosystem to ensure that abnormalities are being captured normally, and it is necessary to be identified and processed in time.

#### Service Information Storage

Usually we often need to set some internal information in the process, such as the basic information such as the application name and the application version number, or the information storage of business attributes. At this time, there is a unified place to deal with.

#### Interface Limiter

During the operation of the application, new clients will be accessed constantly, and sometimes there will be a peak of traffic (such as marketing activities). It is very likely to cause accidents, so we often have a variety of means to restrict peaks, and the rate-limiting control of the application interface is one of the methods for the application interface.

#### Timeout Control

The mutual influence of upstream and downstream applications leads to a serial response, and eventually makes a certain scale unavailable in the entire cluster application. Therefore, we need to perform the most basic timeout control in all requests in the application.

### Link Tracking

In this part, we use Jaeger to implement link tracking which support OpenTracing specfication.

First, we use docker to install Jaeger with following command:

```shell
$ docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 68
31:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-o
ne:latest
```

And then, you can open `http://localhost:16686/` to check Jaeger's Web UI. 

### Application Configuration

We use fsnotify package to solve config reading issue.

### Complie Program

You can complie program directly by using Go for different platforms. 

```shell
# for Linux platform
$ CGO_ENABLED=0 GOOS=linux go build -a -o blogie .

# for Windows platform
$ CGO_ENABLED=0 GOOS=windows go build -a -o blogie.exe .

# for macOS platform
$ CGO_ENABLED=0 GOOS=darwin go build -a -o blogie .
```

If you want to shrink the size of binary of your program, you can remove debug and flags informations by run following command `go build -ldflags="-w -s"`. But, this stuff will cause callstack have no detailed informations while your program appears panic, also cannot use gdb debug the program.

In our project, we choose `ldflags` to set compile informations for our program, you can run following command to set it:

```shell
$ go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"

build_time: 2022-07-19,14:53:56
build_version: 1.0.0
git_commit_id: xxxxxx
```

### Graceful Shutdown

In this project, we use signal to implement graceful shutdown. On Unix/Linux platform, you can run command `kill -l` to check the signals of your system support. In our project, we accept two signals:

- syscall.SIGINT(2), also you can type Ctrl+C to interrupt the program

- syscall.SIGTERM(15)

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

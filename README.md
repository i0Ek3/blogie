# blogie

A blog backend program developed with Gin which integrates many useful features.

## Features

- Compatible 1.18+

- Fully detailed notes and standard coding

- Support RESTful APIs and APIs documentary generation

- Interface validation and APIs access control

- Integrate middlewares, scalable

- Support static resource file service

- Support application configuration

- Support link tracing with OpenTelemetry specification

- Graceful shutdown and restart

- Support i18n and cross-platform

- Find yourself

## Getting Started

### How-To(Build & Run)

```shell
# Please make sure docker is running
$ make # build blogie
$ ./blogie -h

Usage of ./blogie:
  -config string
        specify config path (default "configs/")
  -mode string
        run in which mode
  -port string
        run in which port
```

If you want to use gdb to debug your program, you can build binary with following command:

```go
$ GOFLAGS="-ldflags=-compressdwarf=false" go build
$ gdb blogie

GNU gdb (GDB) 12.1
Copyright (C) 2022 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
Type "show copying" and "show warranty" for details.
This GDB was configured as "x86_64-apple-darwin21.3.0".
Type "show configuration" for configuration details.
For bug reporting instructions, please see:
<https://www.gnu.org/software/gdb/bugs/>.
Find the GDB manual and other documentation resources online at:
    <http://www.gnu.org/software/gdb/documentation/>.

For help, type "help".
Type "apropos word" to search for commands related to "word"...
Reading symbols from blogie...
Loading Go Runtime support.
(gdb)
```

> ps: macOS users may encounter such a prompt: please check gdb is codesigned - see taskgated(8).

## Architecture

In this part, we'll show you how this project work.

![](https://github.com/i0Ek3/blogie/blob/main/images/blogie.jpg)

```shell
├── README.md           // project instruction
├── Makefile            // used to build program
├── docker-compose.yml  // used to boot jaeger and mysql
├── go.mod
├── go.sum
├── main.go             // program entry
├── configs             // global config file
├── docs                // swagger files and docs
├── example             // some validation examples
├── global              // global variables
├── internal            // internal modules
│   ├── dao             // DAO, data related operations
│   ├── middleware      // HTTP middleware
│   ├── model           // Model, database related operations
│   ├── routers         // route processing logic
│   └── service         // project core business logic
├── pkg                 // project related module packages
├── scripts             // some scripts and sql file
└── storage             // store logs and other files
```

### Database

You should install MySQL on your system first, and then create a database which named `blogie` and use this databse. If you have same system with mine(Unix), you can run the script directly to install and import blog.sql.

```shell
# for Unix/Linux users, run command under the root folder of project
# if you are macOS user, please make sure you have homebrew installed
$ ./scripts/setup.sh

# or you want to setup it manually
$ mysql -uroot -p
mysql> CREATE DATABASE blogie;
mysql> USE blogie;
mysql> SOURCE ./scripts/sql/blogie.sql; # import blogie.sql
```

After import blog.sql, it will create following four tables:

```console
mysql> show tables;
+--------------------+
|  Tables_in_blogie  |
+--------------------+
| blogie_article     |
| blogie_article_tag |
| blogie_auth        |
| blogie_tag         |
+--------------------+
4 rows in set (0.00 sec)
```

At the end of the project, you can export all files in your own database, and use following command to query the specific path:

```shell
mysql> show global variables like "%datadir%";

+---------------+-----------------------+
| Variable_name | Value                 |
+---------------+-----------------------+
| datadir       | /usr/local/var/mysql/ |
+---------------+-----------------------+
1 row in set (0.00 sec)
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

After that, you can use `curl` to test your APIs. If you don't know how to test you APIs with `curl`, you can check this [post](http://www.ruanyifeng.com/blog/2019/09/curl-reference.html) or just run script `./scripts/test.sh`.

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

| Level | Details                                                                                                                                                |
| ----- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Trace | Trace have ability to print more fine-grained log information than debug level.                                                                        |
| Debug | It is used to provide some debuf information, which is convenient for us to locate the problem and observe whether the program meets the expectations. |
| Info  | Default logger level, provide some necessary log information to facilitate troubleshooting.                                                            |
| Warn  | More important than Info level.                                                                                                                        |
| Error | It indicates an error in program execution.                                                                                                            |
| Panic | It indicates that there is a serious error in the program, usually the stack information is printed out, and can also be caught.                       |
| Fatal | It indicates that the program encountered a fatal error and needs to exit.                                                                             |

Also our logger support output the log into file and os.Stdout.

And you should know something here, we add lumberjack into our project, so our project support log rotation, if you don't like that, you can only use Logrotate program. In the other hand, add lumberjack will increases the complexity of the log package.

### Common Component

To ensure the standardization of the application, we will abstract the basic functions to the public component of the project.

- Error code standardization

- Configuration management

- Database connection

- Log writting

- Response processing

### Interface Generation

We use swagger to generate our interface documents, just write comment for APIs and then swagger can read it and generate correspoding interface documents.

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

```Console
Header { # json object
    "alg": "HS256", # HMAC SHA256
    "typ": "JWT"
}

Payload { # json object, mainly stored in the actual data transmitted in JWT
    "sub": "Topic",
    "name": "i0Ek3",
    "admin": true
}

Signature # Signature of the agreed algorithm and rules for the first two parts (Header+Payload)

HMACSHA256(
  base64UrlEncode(Header) + "." +
  base64UrlEncode(Payload),
  secret)
```

After you finished the access control, you can generate token by run following command:

```shell
$ curlie -X POST \
  'http://127.0.0.1:8080/auth' \
  -H 'app_key: i0Ek3' \
  -H 'app_secret: blogie'

{"token":"eyJhbG...pXVCJ9.eyJhcH...dpZSJ9.9X4SFy...pxMcs8"}
         |     Part1     |     Part2     |      Part3    |

Part1 = base64UrlEncode(Header)
Part1 = base64UrlEncode(Payload)
Part3 = HMACSHA256(Part1 + "." + Part2, secret)
```

More details about JWT please check [here](https://jwt.io/introduction/).


### Middleware

Some common application middleware can solve most problems in the project. Therefore, in our project, we implemented some basic application middleware, such as access log, recovery, service information storage, rate limiter, redis cache, circuit breaker, cron, timeout control, etc.

#### Access Log

Access log basically records the request method of each request, the start time of the method call, the end time of the method call, the method response result, and the status code of the method response result. and other additional attributes to achieve the effect of log link tracing.

#### Recovery

It is very important for abnormal capture and timely alarm notification, so we need to customize the recovery middleware for our project's own conditions or ecosystem to ensure that abnormalities are being captured normally, and it is necessary to be identified and processed in time.

#### Service Information Storage

Usually we often need to set some internal information in the process, such as the basic information such as the application name and the application version number, or the information storage of business attributes. At this time, there is a unified place to deal with.

#### Rate Limiter

During the operation of the application, new clients will be accessed constantly, and sometimes there will be a peak of traffic (such as marketing activities). It is very likely to cause accidents, so we often have a variety of means to restrict peaks, and the rate-limiting control of the application interface is one of the methods for the application interface.

#### Cron

In business scenarios, we usually need to delete some invalid data at a fixed point in time, so as to achieve the purpose of scheduled task scheduling management. But if it is hardcoded, it is obviously not elegant. Therefore, we implemented the timer middleware with the help of the cron library in the project to complete the requirement of regularly deleting invalid data. This library implements the cron spec parser and task runner, making it easier to use and integrate in our projects.

#### Circuit Breaker

Circuit breakers are the fuses in the underlying service, and its core is fail fast. When the downstream service is unable to provide the service due to overload or failure, we need to let the upstream service know in time, and temporarily fuse the call chain between the caller and the provider, so as to avoid the chain reaction of small failures causing the entire system to be paralyzed or even damaged. Therefore, we provide circuit breaker middleware in the project to disconnect the call chain of abnormal services, so as not to affect the normal use of other services.

#### Timeout Control

The mutual influence of upstream and downstream applications leads to a serial response, and eventually makes a certain scale unavailable in the entire cluster application. Therefore, we need to perform the most basic timeout control in all requests in the application.

### Link Tracing

In this part, we use Jaeger to implement link tracing which support OpenTracing specfication. 

Usually, when multiple distributed interfaces call each other and the response is particularly slow, we need to locate and solve the problem in time. Therefore, we need to do link tracing.

> Please install docker and boot it first.

First, we use docker to install Jaeger with following command:

```shell
$ docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 68
31:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-o
ne:latest
```

And then, after you run the blogie service, you can open `http://localhost:16686/` to check Jaeger's Web UI to tracing interface calls.

Also you can run the command `docker-compose up -d` under the root folder to boot Jaeger service.

### Application Configuration

#### Config Reading

In fact, we cannot run the Go program directly in other directories, because it will prompt that the configuration file cannot be found. So, if you want to pack config.yaml into execute binary, you can run following command:

```Go
$ go get -u github.com/go-bindata/go-bindata/...
$ go-bindata -o configs/config.go -pkg=configs configs/config.yaml
```

Or, just use commond-line argument to pass the parameters.

#### Configure Hot Update

We use fsnotify package to solve config hot update issue.

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

If you want to shrink the size of binary of your program, you can remove debug and flags information by run following command `go build -ldflags="-w -s"`. But, this stuff will cause callstack have no detailed information while your program appears panic, also cannot use gdb debug the program.

In our project, we choose `ldflags` to set compile informations for our program, you can run following command to set it:

```shell
$ go build -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"

build_time: 2022-07-19,14:53:56
build_version: 1.0.0
git_commit_id: xxxxxx
```

### Graceful Shutdown/Restart

In this project, we use signal to implement graceful shutdown and restart. On Unix/Linux platform, you can run command `kill -l` to check the signals of your system support. In our project, we accept two signals:

- syscall.SIGINT(2), also you can type Ctrl+C to interrupt the program

- syscall.SIGTERM(15)

## Issues

### Mod Issues

- Database driver installation

  - Add following line `_ "github.com/go-sql-driver/mysql"` for example/main.go

  - Add following line `_ "github.com/jinzhu/gorm/dialects/mysql"` for model/model.go

- Swagger files

  - Add following line `swaggerFiles "github.com/swaggo/files"` for internal/routers/router.go

### Tag Issues

- After tag deleted by curl, and then you add a new tag with curl but tag id is increase instead of start as 0. This situation may not an error, just a feature in database

### Module Issues

- Upload service: curl: (26) Failed to open/read local data from file/application

  - Use command `curlie -X POST http://127.0.0.1:8080/upload/file -F file=@./demo.jpg -F type=1` to solve it

## Credit

[marmotedu](https://github.com/marmotedu) | [eddycjy](https://github.com/eddycjy) | [minibear2333](https://github.com/minibear2333)

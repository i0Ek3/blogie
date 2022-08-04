#!/bin/bash

test_api() {
    #
    #   There is some methods to test your APIs:
    #     1. use curl program in terminal to test, also you can use curlie
    #     2. use postman or related API tools to test
    #     3. after run the blogie program, you use swagger to test
    #

    # create new tag without any message
    curl -X POST http://127.0.0.1:8080/api/v1/tags

    # create new article without any message
    curl -X POST http://127.0.0.1:8080/api/v1/articles

    # delete a tag
    curl -X DELETE http://127.0.0.1:8080/api/v1/tags/1

    # delete a article
    curl -X DELETE http://127.0.0.1:8080/api/v1/articles/1

    # get tag list
    curl -X GET http://127.0.0.1:8080/api/v1/tags/1

    # get article list
    curl -X GET http://127.0.0.1:8080/api/v1/articles/1
}

ab_test() {
    ab -c 1000 -n 1000 http://localhost:8080/swagger
    ab -c 1000 -n 1000 http://localhost:8080/upload/file
    ab -c 1000 -n 1000 http://localhost:8080/auth
    ab -c 1000 -n 1000 http://localhost:8080/debug/vars
    ab -c 1000 -n 1000 http://localhost:8080/api/v1/tags
    ab -c 1000 -n 1000 http://localhost:8080/api/v1/articles
}

main() {
    #test_api
    ab_test
}

main

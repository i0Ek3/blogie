#!/bin/bash

# add new article
curl -X 'POST' \
  'http://127.0.0.1:8080/api/v1/articles?tag_id=Gopher&title=How%20to%20test%20API%20with%20curl&cover_image_url=./images/demo.jpg&content=this%20is%20a%20test&created_by=i0Ek3' \
  -H 'accept: application/json'

# add another article
curl -X 'POST' \
  'http://127.0.0.1:8080/api/v1/articles?tag_id=Gopher&title=How%20to%20test%20API%20with%20curl1&cover_image_url=./images/demo.jpg&content=this%20is%20a%20test1&created_by=i0Ek3' \
  -H 'accept: application/json'

# get an article
curl -X 'GET' \
  'http://127.0.0.1:8080/api/v1/articles/1' \
  -H 'accept: application/json'

# get article list
curl -X 'GET' \
  'http://127.0.0.1:8080/api/v1/articles?name=Gopher&tag_id=1&state=0' \
  -H 'accept: application/json'

# update an article
curl -X 'PUT' \
  'http://127.0.0.1:8080/api/v1/articles/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d 'i0Ek3'

# delete an article
curl -X 'DELETE' \
  'http://127.0.0.1:8080/api/v1/articles/1' \
  -H 'accept: application/json'

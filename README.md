# todolist-goapi

Simple Todo list go api to handle add/removal of todolist

Run this go command to build a Go binary ```go build -o bin/rest_api```

The @dockerfile will create the binary within the docker image and run it.

Build docker image ``` docker build -t todo-api:v1 . ```

Run docker image ``` docker run -it -p 3000:8080 todo-api:v1 ```

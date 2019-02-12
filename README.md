#Natural language processor

Natural language processor written in Go.

##Usage

```
$ curl --data-binary @body_of_text.txt localhost:8080/learn
$ curl localhost:8080/generate
```

##Run

Install dependencies, run tests and build with:

```
$ make dep
$ make test
$ make build

```

Run the app on port 8080 with:

```
$ ./web
2019/02/12 00:56:56 listen on port 8080
```
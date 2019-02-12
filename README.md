## Natural language processor

Natural language processor written in Go.

### Usage

```
$ curl --data-binary @body_of_text.txt localhost:8080/learn
$ curl localhost:8080/generate
```

### Run

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

### TODOs

- Switch out pseudo-random number generator for real random number generator
to handle trigram node selection and prefix seeding
- Make the limit of trigrams to return from store configurable (currently
hard-coded to 1000 in main func)

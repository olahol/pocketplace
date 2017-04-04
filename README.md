# pocketplace

> :fireworks: Draw pixels on a canvas with friends.

![Demo](https://cdn.rawgit.com/olahol/pocketplace/master/demo.gif "Demo")

* [x] Completely in-memory, no need for a database.
* [x] Statically linked, everything you need in one binary, including the frontend.

# Install

```bash
  $ go get github.com/olahol/pocketplace
```

# Example

```bash
  $ pocketplace -port 8080 -size 200 -cooldown 0
    Canvas pixel size 200x200
    Drawing cooldown 0s
    Listening on port 8080
```

# Development

To build the frontend into the binary I use [`file2const`](https://github.com/bouk/file2const)
and `go:generate` directives. So if you are modifying the frontend don't forget to:

```bash
  $ go generate
  $ go build
```

to build the binary correctly with the updated frontend.

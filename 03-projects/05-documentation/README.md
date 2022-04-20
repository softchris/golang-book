# Documenting your code

TODO

##  Pre-Lecture Quiz

TODO

## Introduction

This chapter will cover:

- Setting up godoc tool.
- Document your code.
- Generate docs.

## Setup `godoc`

`godoc` generates documentation for your package.

1. Setup `GOPATH`:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

1. Install `godoc`

    ```bash
    go get golang.org/x/tools/cmd/godoc
    ```

1. Register `godoc` path:

   ```bash
   export PATH="$GOPATH/bin:$PATH"
   ```

1. Run `godoc`:

   ```bash
   godoc -http=:6060
   ```

   on `http://localhost:6060` there are now docs. Find your package and you will see it has docs.

1. Add comment as normal commented out comment, near the code

## Assignment

TODO

## Solution

TODO

## Challenge

TODO

##  Post-Lecture Quiz

TODO

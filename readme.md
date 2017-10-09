# Fizz Buzz Go Server

This application is a simple web server that exposes a REST endpoint providing [fizz buzz](https://en.wikipedia.org/wiki/Fizz_buzz) replies as JSON.

## Running the server

    go build
    ./fizzbuzz

The server will listen on port 1337 and respond to GET http requests on the root path.

## Using the service

You can use curl :

    curl 'http://localhost:1337/'

The following parameters can be used :

- string1 (defaults to **fizz**)
- string2 (defaults to **buzz**)
- int1 (defaults to **3**)
- int2 (defaults to **5**)
- limit (defaults to **100**)

Here is a sample with customized parameters :

    curl 'http://localhost:1337/?string1=hey&string2=ho&int1=2&int2=6&limit=200'

# dsb-middleware

## What?

The **mobile-dsb** API doesn't send a `CORS` header.
This API sends a request with the `URI` you enter to the **mobile-dsb** API,
adds the `CORS` header and sends you the body

## Setting up

```sh
$ git clone https://github.com/Tch1b0/dsb-middleware
```

Now you can edit the port the Service will run on in the **.env file**

```sh
$ docker-compose build
```

```sh
$ docker-compose up -d
```

The Service will now run either on port `5010` or the port you declared in the **.env file**

## References

-   [Docker](https://www.docker.com/) _required_
-   [Golang](https://golang.org)
-   [Gin](https://github.com/gin-gonic/gin)

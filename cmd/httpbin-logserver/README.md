# httpbin-logserver

The intent of this is to expose a httpbin API as implemented by
mccutchen/go-httpbin/v2, only adding the http server around it,
along with logging the request details to stdout as JSON.

The request details are used to plot a histogram / graph of the
requests as they hit the service. This is particularly applicable
to testing rate limiters and seeing the request count / rate.

Logging graceful shutdown and program exit is done with `log` package,
which logs the output into the standard error, so it doesn't interfere
with the JSON output in standard output.

## task: Available tasks for this project:

* default:       Run default (install)
* build:         Build from source
* docker:        Build docker image
* install:       Install from source

## Command line arguments

```
Usage of ./httpbin-logserver:
  -addr string
    	address to listen to (default ":8085")
  -output string
    	output, 'stdout' or file (default "stdout")
```

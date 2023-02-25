## whoami 

Tiny Go webserver that prints HTTP requests to output.

The main purpose of the server is to use it as an end service in testing network performance.

## Usage

```shell
docker run -p8888:8081 pprishchepa/whoami
curl http://localhost:8888/foo
``` 

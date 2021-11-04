# Simple todo API in go

This uses the aws SAM stack. The two available endpoints are

`/gettodo` and `/addtodo`

Each endpoint maps to a single lambda function handler. Each is in its own folder and built separately.

Build steps can be found in `build.sh` but probably should not be run as a shell script.

```shell
curl -X POST -d '{"desc":"more todo"}' "https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/todos"
curl "https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/todos"
```

```shell
$ curl -X POST -d '{"desc":"testing easier id"}' "https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/todos"
-> {"id":"8086","desc":"testing easier id"}
$ curl "https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/todos/8086"                                     
-> {"id":"8086","desc":"testing easier id"}
```
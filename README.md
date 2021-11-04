# Simple todo API in go

This uses the aws SAM stack. The two available endpoints are 

`/gettodo` and `/addtodo`

Each endpoint maps to a single lambda function handler. Each is in its own folder and built separately.

Build steps can be found in `build.sh` but probably should not be run as a shell script.


# curl -X POST https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/addtodo/
# curl https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/gettodo/




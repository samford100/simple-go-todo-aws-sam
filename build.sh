# ensure this is fun from the root directory
cd getTodo/
go build
cd ../addTodo
go build
cd ..
sam build
sam deploy --guided


# curl -X POST https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/addtodo/
# curl https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/gettodo/
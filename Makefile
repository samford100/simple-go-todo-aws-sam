build:
	cd addTodo && go build && cd ..
	cd getTodo && go build && cd ..
	cd getTodos && go build && cd ..
	cd updateTodo && go build && cd ..
	cd deleteTodo && go build && cd ..
	sam build

deploy:
	sam deploy --guided

clean:
	rm -rf .aws-sam

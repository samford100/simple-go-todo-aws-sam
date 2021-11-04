<script>
  let todos = [];
  fetch("https://cjo5jmgvm3.execute-api.us-east-1.amazonaws.com/Prod/todos/")
            .then((resp) => {
							return resp.json();
						})
						.then((data) => {
							todos = data;
						})
</script>

<h1>Todo List</h1>
<ul>
  {#each todos as todo}
    <li>
      <span>{todo.desc}</span>
    </li>
  {/each}
</ul>

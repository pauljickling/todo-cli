# todo-cli
A command line tool for managing tasks. Currently in development, and not yet ready for general use.

## config

- where to store todo items
- where to store completed items
- where to store in progress items.
- the time increment to include for recent updates (default is weekly)

## usage

`todo -n "my new task"` add the todo item "my new task". also accepts `--new` 
`todo -l` lists all todos. also accepts `--list`
`todo -r "my new task"` removes the todo item "my new task". also accepts `--remove`
`todo -c "my new task"` marks the todo item "my new task" as complete and moves it to the completed tasks. also accepts `--complete`
`todo -p "my new task"` marks the todo item "my new task" as in-progress. also accepts `--progress`

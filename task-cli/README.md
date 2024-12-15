# [Task Tracker Project](https://roadmap.sh/projects/task-tracker)

Task tracker is a CLI project used to track and manage tasks.

## Build and Install

```sh
go build -o <bin directory> ./
```

## The list of commands and their usage is given below:

### Adding a new task
```sh
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Updating and deleting tasks
```sh
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1
```

### Marking a task as in progress or done
```sh
task-cli mark-in-progress 1
task-cli mark-done 1
```

### Listing all tasks
```sh
task-cli list
```

### Listing tasks by status
```sh
task-cli list done
task-cli list todo
task-cli list in-progress
```


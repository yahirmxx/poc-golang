package main

var tasks = []Task{
	{ID: 1, Title: "Task 1", Description: "Do something"},
	{ID: 2, Title: "Task 2", Description: "Do something else"},
}

func getAllTasksRepository() []Task {
	return tasks
}

func getTaskByIdRepository(taskID int) Task {
	for _, task := range tasks {
		if task.ID == taskID {
			return task
		}
	}
	return Task{}
}

func createTaskRepository(newTask Task) Task {
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	return newTask
}

func updateTaskRepository(taskID int, updatedTask Task) Task {
	for index, task := range tasks {
		if task.ID == taskID {
			tasks[index] = updatedTask

			return updatedTask
		}
	}
	return Task{}
}

func deleteTaskRepository(taskID int) bool {
	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return true
		}
	}
	return false
}

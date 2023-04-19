package tasks

import "github.com/stretchr/testify/assert"

func NewTaskTest() *Task {
	task := NewTask()
	data := TaskData{Title: "title", Description: "description"}
	task.AddData(data)
	return task
}

func (rs *RepositorySuite) TestAddTask() {
	task := NewTaskTest()
	
	_, err := rs.repository.AddTask(task)
	assert.NoError(rs.T(), err)

	recoveredTask := &Task{}
	rs.dbHandler.First(recoveredTask, "id = ?", task.ID)

	assert.Equal(rs.T(), recoveredTask.ID, task.ID)
}

func (rs *RepositorySuite) TestGetTask() {
	task := NewTaskTest()

	rs.dbHandler.Create(task)

	recoveredTask, err := rs.repository.GetTask(task.ID.String())

	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), recoveredTask.ID, task.ID)
}

func (rs *RepositorySuite) TestGetTaskNotFound() {
	_, err := rs.repository.GetTask("id")
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "Task not found", err.Error())
}

func (rs *RepositorySuite) TestGetAllTasks() {
	var numberOfTasks int = 3

	for i := 0; i < numberOfTasks; i++ {
		task := NewTaskTest()
		rs.dbHandler.Create(task)
	}

	tasks, err := rs.repository.GetAllTasks()
	
	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), numberOfTasks, len(*tasks))
}

func (rs *RepositorySuite) TestDeleteTask() {
	task := NewTaskTest()

	rs.dbHandler.Create(task)

	removedTask, err := rs.repository.DeleteTask(task.ID.String())

	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), removedTask.ID, task.ID)
}

func (rs *RepositorySuite) TestDeleteTaskNotFound() {
	_, err := rs.repository.DeleteTask("id")
	assert.Error(rs.T(), err)
	assert.Equal(rs.T(), "Task not found", err.Error())
}

func (rs *RepositorySuite) TestRepositoryUpdateTask() {
	task := NewTaskTest()
	rs.dbHandler.Create(task)

	data := &TaskData{Title: "newTitle", Description: "newDescription"}

	updatedTask, err := rs.repository.UpdateTask(task.ID.String(), data)

	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), updatedTask.ID, task.ID)
	assert.Equal(rs.T(), updatedTask.Title, data.Title)
	assert.Equal(rs.T(), updatedTask.Description, data.Description)
	assert.NotEqual(rs.T(), task.Title, updatedTask.Title)
	assert.NotEqual(rs.T(), task.Description, updatedTask.Description)
}

package services

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"task_management/models"
	"task_management/utils"
)

// Task Service
type TaskService struct {
	tasks    map[string]*models.Task
	comments map[string]*models.Comment
	mu       sync.RWMutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:    make(map[string]*models.Task),
		comments: make(map[string]*models.Comment),
	}
}

// Task Management
func (ts *TaskService) CreateTask(title, description string, project *models.Project, createdBy *models.User) (*models.Task, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	// Check if user has permission to create tasks
	if !createdBy.HasPermission("create_project") && !project.CanUserAccess(createdBy.ID) {
		return nil, fmt.Errorf("user does not have permission to create tasks in this project")
	}
	
	task := models.NewTask(title, description, project, createdBy)
	ts.tasks[task.ID] = task
	return task, nil
}

func (ts *TaskService) GetTask(taskID string) *models.Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.tasks[taskID]
}

func (ts *TaskService) UpdateTask(taskID string, title, description string, user *models.User) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	if task == nil {
		return fmt.Errorf("task not found")
	}
	
	// Check if user has permission to update task
	if !ts.canUserModifyTask(task, user) {
		return fmt.Errorf("user does not have permission to update this task")
	}
	
	task.UpdateDetails(title, description)
	return nil
}

func (ts *TaskService) AssignTask(taskID string, assignee *models.User, assignedBy *models.User) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	if task == nil {
		return fmt.Errorf("task not found")
	}
	
	// Check if assignedBy has permission to assign tasks
	if !assignedBy.HasPermission("assign_tasks") && task.CreatedBy.ID != assignedBy.ID {
		return fmt.Errorf("user does not have permission to assign this task")
	}
	
	// Use the provided assignee user
	// Ensure the assignee has proper fields set
	if assignee.FirstName == "" {
		assignee.FirstName = "Assignee"
	}
	if assignee.LastName == "" {
		assignee.LastName = "User"
	}
	
	if !task.CanBeAssignedTo(assignee) {
		return fmt.Errorf("user cannot be assigned to this task")
	}
	
	task.AssignTo(assignee)
	
	// Add assignment comment
	comment := models.NewComment(
		fmt.Sprintf("Task assigned to %s", assignee.GetFullName()),
		assignedBy,
		task,
		models.AssignmentComment,
	)
	ts.comments[comment.ID] = comment
	
	return nil
}

func (ts *TaskService) ChangeTaskStatus(taskID string, status models.TaskStatus, user *models.User) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	if task == nil {
		return fmt.Errorf("task not found")
	}
	
	// Check if user has permission to change status
	if !ts.canUserModifyTask(task, user) {
		return fmt.Errorf("user does not have permission to change status of this task")
	}
	
	oldStatus := task.Status
	err := task.UpdateStatus(status)
	if err != nil {
		return err
	}
	
	// Add status change comment
	comment := models.NewComment(
		fmt.Sprintf("Status changed from %s to %s", oldStatus.String(), status.String()),
		user,
		task,
		models.StatusChangeComment,
	)
	ts.comments[comment.ID] = comment
	
	return nil
}

func (ts *TaskService) SetTaskPriority(taskID string, priority models.Priority, user *models.User) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	if task == nil {
		return fmt.Errorf("task not found")
	}
	
	// Check if user has permission to modify task
	if !ts.canUserModifyTask(task, user) {
		return fmt.Errorf("user does not have permission to modify this task")
	}
	
	task.SetPriority(priority)
	return nil
}

func (ts *TaskService) SetTaskDueDate(taskID string, dueDate *time.Time, user *models.User) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	if task == nil {
		return fmt.Errorf("task not found")
	}
	
	// Check if user has permission to modify task
	if !ts.canUserModifyTask(task, user) {
		return fmt.Errorf("user does not have permission to modify this task")
	}
	
	task.SetDueDate(dueDate)
	return nil
}

func (ts *TaskService) AddTaskDependency(taskID, dependencyID string, user *models.User) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	dependency := ts.tasks[dependencyID]
	
	if task == nil || dependency == nil {
		return fmt.Errorf("task or dependency not found")
	}
	
	// Check if user has permission to modify task
	if !ts.canUserModifyTask(task, user) {
		return fmt.Errorf("user does not have permission to modify this task")
	}
	
	err := task.AddDependency(dependency)
	if err != nil {
		return err
	}
	
	// Add dependency comment
	comment := models.NewComment(
		fmt.Sprintf("Added dependency: %s", dependency.Title),
		user,
		task,
		models.GeneralComment,
	)
	ts.comments[comment.ID] = comment
	
	return nil
}

func (ts *TaskService) AddSubtask(parentTaskID, title, description string, user *models.User) (*models.Task, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	parentTask := ts.tasks[parentTaskID]
	if parentTask == nil {
		return nil, fmt.Errorf("parent task not found")
	}
	
	// Check if user has permission to create subtasks
	if !ts.canUserModifyTask(parentTask, user) {
		return nil, fmt.Errorf("user does not have permission to create subtasks for this task")
	}
	
	subtask := models.NewTask(title, description, parentTask.Project, user)
	ts.tasks[subtask.ID] = subtask
	parentTask.AddSubtask(subtask)
	
	// Add subtask comment
	comment := models.NewComment(
		fmt.Sprintf("Added subtask: %s", subtask.Title),
		user,
		parentTask,
		models.GeneralComment,
	)
	ts.comments[comment.ID] = comment
	
	return subtask, nil
}

// Comment Management
func (ts *TaskService) AddComment(taskID, content string, author *models.User) (*models.Comment, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	
	task := ts.tasks[taskID]
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}
	
	// Check if user has access to the project
	if !task.Project.CanUserAccess(author.ID) {
		return nil, fmt.Errorf("user does not have access to this project")
	}
	
	comment := models.NewComment(content, author, task, models.GeneralComment)
	ts.comments[comment.ID] = comment
	return comment, nil
}

func (ts *TaskService) GetTaskComments(taskID string) []*models.Comment {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	
	var comments []*models.Comment
	for _, comment := range ts.comments {
		if comment.GetTask().ID == taskID {
			comments = append(comments, comment)
		}
	}
	
	// Sort by creation time
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].GetCreatedAt().Before(comments[j].GetCreatedAt())
	})
	
	return comments
}

// Search and Filtering
func (ts *TaskService) SearchTasks(query string, projectID string) []*models.Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	
	var results []*models.Task
	for _, task := range ts.tasks {
		if projectID != "" && task.Project.ID != projectID {
			continue
		}
		
		// Simple text search in title and description
		if utils.ContainsIgnoreCase(task.Title, query) || utils.ContainsIgnoreCase(task.Description, query) {
			results = append(results, task)
		}
	}
	
	// Sort by priority and creation date
	sort.Slice(results, func(i, j int) bool {
		if results[i].Priority != results[j].Priority {
			return results[i].Priority < results[j].Priority
		}
		return results[i].CreatedAt.After(results[j].CreatedAt)
	})
	
	return results
}

func (ts *TaskService) GetTasksByProject(projectID string) []*models.Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	
	var tasks []*models.Task
	for _, task := range ts.tasks {
		if task.Project.ID == projectID {
			tasks = append(tasks, task)
		}
	}
	
	// Sort by priority and creation date
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Priority != tasks[j].Priority {
			return tasks[i].Priority < tasks[j].Priority
		}
		return tasks[i].CreatedAt.After(tasks[j].CreatedAt)
	})
	
	return tasks
}

func (ts *TaskService) GetTasksByUser(userID string) []*models.Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	
	var tasks []*models.Task
	for _, task := range ts.tasks {
		if task.Assignee != nil && task.Assignee.ID == userID {
			tasks = append(tasks, task)
		}
	}
	
	// Sort by priority and due date
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Priority != tasks[j].Priority {
			return tasks[i].Priority < tasks[j].Priority
		}
		
		// Handle tasks without due dates
		if tasks[i].DueDate == nil && tasks[j].DueDate == nil {
			return tasks[i].CreatedAt.After(tasks[j].CreatedAt)
		}
		if tasks[i].DueDate == nil {
			return false
		}
		if tasks[j].DueDate == nil {
			return true
		}
		
		return tasks[i].DueDate.Before(*tasks[j].DueDate)
	})
	
	return tasks
}

func (ts *TaskService) GetOverdueTasks() []*models.Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	
	var overdueTasks []*models.Task
	for _, task := range ts.tasks {
		if task.IsOverdue() {
			overdueTasks = append(overdueTasks, task)
		}
	}
	
	// Sort by due date
	sort.Slice(overdueTasks, func(i, j int) bool {
		if overdueTasks[i].DueDate == nil || overdueTasks[j].DueDate == nil {
			return false
		}
		return overdueTasks[i].DueDate.Before(*overdueTasks[j].DueDate)
	})
	
	return overdueTasks
}

// Helper Methods
func (ts *TaskService) canUserModifyTask(task *models.Task, user *models.User) bool {
	// Admin can modify any task
	if user.Role == models.Admin {
		return true
	}
	
	// Project manager can modify tasks in their projects
	if user.Role == models.ProjectManager && task.Project.Owner.ID == user.ID {
		return true
	}
	
	// Task creator can modify their own tasks
	if task.CreatedBy.ID == user.ID {
		return true
	}
	
	// Task assignee can modify their assigned tasks
	if task.Assignee != nil && task.Assignee.ID == user.ID {
		return true
	}
	
	return false
}


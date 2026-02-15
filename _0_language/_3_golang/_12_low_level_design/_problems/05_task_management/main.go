package main

import (
	"fmt"
	"time"

	"task_management/models"
	"task_management/services"
)

// Task Management System
type TaskManagementSystem struct {
	ProjectService *services.ProjectService
	TaskService    *services.TaskService
	TimeService    *services.TimeService
}

func NewTaskManagementSystem() *TaskManagementSystem {
	return &TaskManagementSystem{
		ProjectService: services.NewProjectService(),
		TaskService:    services.NewTaskService(),
		TimeService:    services.NewTimeService(),
	}
}

func main() {
	fmt.Println("=== TASK MANAGEMENT SYSTEM DEMONSTRATION ===\n")

	// Create task management system
	tms := NewTaskManagementSystem()

	// Register users
	user1, _ := tms.ProjectService.RegisterUser("alice", "alice@example.com", "hash1", "Alice", "Johnson")
	user2, _ := tms.ProjectService.RegisterUser("bob", "bob@example.com", "hash2", "Bob", "Smith")
	user3, _ := tms.ProjectService.RegisterUser("charlie", "charlie@example.com", "hash3", "Charlie", "Brown")

	// Set user roles
	user1.SetRole(models.ProjectManager)
	user2.SetRole(models.Developer)
	user3.SetRole(models.Developer)

	fmt.Println("1. USER REGISTRATION:")
	fmt.Printf("User 1: %s (%s) - %s\n", user1.GetFullName(), user1.Username, user1.Role.String())
	fmt.Printf("User 2: %s (%s) - %s\n", user2.GetFullName(), user2.Username, user2.Role.String())
	fmt.Printf("User 3: %s (%s) - %s\n", user3.GetFullName(), user3.Username, user3.Role.String())

	// Create teams
	team1, _ := tms.ProjectService.CreateTeam("Frontend Team", "Responsible for frontend development", user1)
	team2, _ := tms.ProjectService.CreateTeam("Backend Team", "Responsible for backend development", user2)

	// Add members to teams
	tms.ProjectService.AddTeamMember(team1.ID, user2.ID)
	tms.ProjectService.AddTeamMember(team2.ID, user3.ID)

	fmt.Println()
	fmt.Println("2. TEAM CREATION:")
	fmt.Printf("Team 1: %s (%d members)\n", team1.Name, team1.GetMemberCount())
	fmt.Printf("Team 2: %s (%d members)\n", team2.Name, team2.GetMemberCount())

	// Create projects
	project1, err := tms.ProjectService.CreateProject("E-commerce Website", "Build a modern e-commerce platform", user1)
	if err != nil {
		fmt.Printf("Error creating project1: %v\n", err)
		return
	}
	project2, err := tms.ProjectService.CreateProject("Mobile App", "Develop a mobile application", user1)
	if err != nil {
		fmt.Printf("Error creating project2: %v\n", err)
		return
	}

	// Add members to projects
	err = tms.ProjectService.AddProjectMember(project1.ID, user2.ID)
	if err != nil {
		fmt.Printf("Error adding user2 to project1: %v\n", err)
	}
	err = tms.ProjectService.AddProjectMember(project1.ID, user3.ID)
	if err != nil {
		fmt.Printf("Error adding user3 to project1: %v\n", err)
	}
	err = tms.ProjectService.AddProjectMember(project2.ID, user3.ID)
	if err != nil {
		fmt.Printf("Error adding user3 to project2: %v\n", err)
	}

	fmt.Println()
	fmt.Println("3. PROJECT CREATION:")
	fmt.Printf("Project 1: %s (%d members)\n", project1.Name, project1.GetMemberCount())
	fmt.Printf("Project 2: %s (%d members)\n", project2.Name, project2.GetMemberCount())

	// Create tasks
	task1, err := tms.TaskService.CreateTask("Design Homepage", "Create the main homepage design", project1, user1)
	if err != nil {
		fmt.Printf("Error creating task1: %v\n", err)
		return
	}
	task2, err := tms.TaskService.CreateTask("Implement User Authentication", "Build user login and registration", project1, user1)
	if err != nil {
		fmt.Printf("Error creating task2: %v\n", err)
		return
	}
	task3, err := tms.TaskService.CreateTask("Create Product Catalog", "Build product listing and search", project1, user1)
	if err != nil {
		fmt.Printf("Error creating task3: %v\n", err)
		return
	}
	task4, err := tms.TaskService.CreateTask("Mobile App UI", "Design mobile app user interface", project2, user1)
	if err != nil {
		fmt.Printf("Error creating task4: %v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("4. TASK CREATION:")
	fmt.Printf("Task 1: %s (Project: %s)\n", task1.Title, task1.Project.Name)
	fmt.Printf("Task 2: %s (Project: %s)\n", task2.Title, task2.Project.Name)
	fmt.Printf("Task 3: %s (Project: %s)\n", task3.Title, task3.Project.Name)
	fmt.Printf("Task 4: %s (Project: %s)\n", task4.Title, task4.Project.Name)

	// Assign tasks
	tms.TaskService.AssignTask(task1.ID, user2, user1)
	tms.TaskService.AssignTask(task2.ID, user2, user1)
	tms.TaskService.AssignTask(task3.ID, user3, user1)
	tms.TaskService.AssignTask(task4.ID, user3, user2)

	fmt.Println()
	fmt.Println("5. TASK ASSIGNMENT:")
	if task1.Assignee != nil {
		fmt.Printf("Task 1 assigned to: %s\n", task1.Assignee.GetFullName())
	} else {
		fmt.Printf("Task 1 not assigned\n")
	}
	if task2.Assignee != nil {
		fmt.Printf("Task 2 assigned to: %s\n", task2.Assignee.GetFullName())
	} else {
		fmt.Printf("Task 2 not assigned\n")
	}
	if task3.Assignee != nil {
		fmt.Printf("Task 3 assigned to: %s\n", task3.Assignee.GetFullName())
	} else {
		fmt.Printf("Task 3 not assigned\n")
	}
	if task4.Assignee != nil {
		fmt.Printf("Task 4 assigned to: %s\n", task4.Assignee.GetFullName())
	} else {
		fmt.Printf("Task 4 not assigned\n")
	}

	// Set task priorities and due dates
	tms.TaskService.SetTaskPriority(task1.ID, models.High, user1)
	tms.TaskService.SetTaskPriority(task2.ID, models.Critical, user1)
	tms.TaskService.SetTaskPriority(task3.ID, models.Medium, user1)
	tms.TaskService.SetTaskPriority(task4.ID, models.Low, user2)

	dueDate1 := time.Now().AddDate(0, 0, 7)  // 1 week from now
	dueDate2 := time.Now().AddDate(0, 0, 3)  // 3 days from now
	dueDate3 := time.Now().AddDate(0, 0, 14) // 2 weeks from now
	dueDate4 := time.Now().AddDate(0, 0, 21) // 3 weeks from now

	tms.TaskService.SetTaskDueDate(task1.ID, &dueDate1, user1)
	tms.TaskService.SetTaskDueDate(task2.ID, &dueDate2, user1)
	tms.TaskService.SetTaskDueDate(task3.ID, &dueDate3, user1)
	tms.TaskService.SetTaskDueDate(task4.ID, &dueDate4, user2)

	fmt.Println()
	fmt.Println("6. TASK CONFIGURATION:")
	fmt.Printf("Task 1: Priority %s, Due %s\n", task1.Priority.String(), dueDate1.Format("2006-01-02"))
	fmt.Printf("Task 2: Priority %s, Due %s\n", task2.Priority.String(), dueDate2.Format("2006-01-02"))
	fmt.Printf("Task 3: Priority %s, Due %s\n", task3.Priority.String(), dueDate3.Format("2006-01-02"))
	fmt.Printf("Task 4: Priority %s, Due %s\n", task4.Priority.String(), dueDate4.Format("2006-01-02"))

	// Add task dependencies
	tms.TaskService.AddTaskDependency(task2.ID, task1.ID, user1) // Authentication depends on Homepage
	tms.TaskService.AddTaskDependency(task3.ID, task2.ID, user1) // Product Catalog depends on Authentication

	fmt.Println()
	fmt.Println("7. TASK DEPENDENCIES:")
	if len(task2.GetDependencies()) > 0 {
		fmt.Printf("Task 2 depends on: %s\n", task2.GetDependencies()[0].Title)
	}
	if len(task3.GetDependencies()) > 0 {
		fmt.Printf("Task 3 depends on: %s\n", task3.GetDependencies()[0].Title)
	}

	// Add comments
	tms.TaskService.AddComment(task1.ID, "Starting work on the homepage design. Will focus on responsive layout.", user2)
	tms.TaskService.AddComment(task2.ID, "Need to research best practices for authentication security.", user2)
	tms.TaskService.AddComment(task3.ID, "Will implement search functionality with filters.", user3)

	fmt.Println()
	fmt.Println("8. TASK COMMENTS:")
	comments1 := tms.TaskService.GetTaskComments(task1.ID)
	fmt.Printf("Task 1 comments: %d\n", len(comments1))
	for _, comment := range comments1 {
		fmt.Printf("  - %s: %s\n", comment.GetAuthor().GetFullName(), comment.GetContent())
	}

	// Change task status
	tms.TaskService.ChangeTaskStatus(task1.ID, models.InProgress, user2)
	tms.TaskService.ChangeTaskStatus(task2.ID, models.Todo, user2)
	tms.TaskService.ChangeTaskStatus(task4.ID, models.InProgress, user3)

	fmt.Println()
	fmt.Println("9. TASK STATUS UPDATES:")
	fmt.Printf("Task 1 status: %s\n", task1.Status.String())
	fmt.Printf("Task 2 status: %s\n", task2.Status.String())
	fmt.Printf("Task 4 status: %s\n", task4.Status.String())

	// Log time entries
	tms.TimeService.LogTime(task1.ID, user2, 2*time.Hour, time.Now(), "Worked on homepage layout")
	tms.TimeService.LogTime(task1.ID, user2, 1*time.Hour, time.Now().Add(-24*time.Hour), "Research and planning")
	tms.TimeService.LogTime(task4.ID, user3, 3*time.Hour, time.Now(), "Mobile UI design and prototyping")

	fmt.Println()
	fmt.Println("10. TIME TRACKING:")
	timeEntries1 := tms.TimeService.GetTimeEntries(task1.ID)
	fmt.Printf("Task 1 time entries: %d\n", len(timeEntries1))
	for _, entry := range timeEntries1 {
		fmt.Printf("  - %s: %s (%.1f hours)\n", 
			entry.GetDate().Format("2006-01-02"), 
			entry.GetDescription(), 
			entry.GetDurationInHours())
	}

	// Generate time reports
	startDate := time.Now().AddDate(0, 0, -7) // 1 week ago
	endDate := time.Now()
	
	user2Report := tms.TimeService.GenerateUserTimeReport(user2.ID, startDate, endDate)
	project1Report := tms.TimeService.GenerateProjectTimeReport(project1.ID, startDate, endDate)

	fmt.Println()
	fmt.Println("11. TIME REPORTS:")
	fmt.Printf("User 2 total time: %.1f hours\n", user2Report.GetTotalHours())
	fmt.Printf("Project 1 total time: %.1f hours\n", project1Report.GetTotalHours())

	// Search functionality
	searchResults := tms.TaskService.SearchTasks("design", project1.ID)
	fmt.Println()
	fmt.Println("12. SEARCH FUNCTIONALITY:")
	fmt.Printf("Search results for 'design' in Project 1: %d tasks\n", len(searchResults))
	for _, task := range searchResults {
		fmt.Printf("  - %s (Priority: %s)\n", task.Title, task.Priority.String())
	}

	// Get user tasks
	user2Tasks := tms.TaskService.GetTasksByUser(user2.ID)
	user3Tasks := tms.TaskService.GetTasksByUser(user3.ID)

	fmt.Println()
	fmt.Println("13. USER TASK ASSIGNMENTS:")
	fmt.Printf("User 2 tasks: %d\n", len(user2Tasks))
	for _, task := range user2Tasks {
		fmt.Printf("  - %s (Status: %s)\n", task.Title, task.Status.String())
	}
	fmt.Printf("User 3 tasks: %d\n", len(user3Tasks))
	for _, task := range user3Tasks {
		fmt.Printf("  - %s (Status: %s)\n", task.Title, task.Status.String())
	}

	// Get project tasks
	project1Tasks := tms.TaskService.GetTasksByProject(project1.ID)
	project2Tasks := tms.TaskService.GetTasksByProject(project2.ID)

	fmt.Println()
	fmt.Println("14. PROJECT TASK OVERVIEW:")
	fmt.Printf("Project 1 tasks: %d\n", len(project1Tasks))
	for _, task := range project1Tasks {
		fmt.Printf("  - %s (Priority: %s, Status: %s)\n", 
			task.Title, task.Priority.String(), task.Status.String())
	}
	fmt.Printf("Project 2 tasks: %d\n", len(project2Tasks))
	for _, task := range project2Tasks {
		fmt.Printf("  - %s (Priority: %s, Status: %s)\n", 
			task.Title, task.Priority.String(), task.Status.String())
	}

	// Create subtasks
	_, _ = tms.TaskService.AddSubtask(task1.ID, "Create Header Component", "Design and implement the website header", user2)
	_, _ = tms.TaskService.AddSubtask(task1.ID, "Create Footer Component", "Design and implement the website footer", user2)

	fmt.Println()
	fmt.Println("15. SUBTASKS:")
	fmt.Printf("Task 1 subtasks: %d\n", len(task1.GetSubtasks()))
	for _, subtask := range task1.GetSubtasks() {
		fmt.Printf("  - %s\n", subtask.Title)
	}

	// Check overdue tasks
	overdueTasks := tms.TaskService.GetOverdueTasks()
	fmt.Println()
	fmt.Println("16. OVERDUE TASKS:")
	fmt.Printf("Overdue tasks: %d\n", len(overdueTasks))
	for _, task := range overdueTasks {
		fmt.Printf("  - %s (Due: %s)\n", task.Title, task.DueDate.Format("2006-01-02"))
	}

	// Project statistics
	allProjects := tms.ProjectService.GetAllProjects()
	activeProjects := tms.ProjectService.GetProjectsByStatus(models.ProjectActive)
	planningProjects := tms.ProjectService.GetProjectsByStatus(models.ProjectPlanning)

	fmt.Println()
	fmt.Println("17. PROJECT STATISTICS:")
	fmt.Printf("Total projects: %d\n", len(allProjects))
	fmt.Printf("Active projects: %d\n", len(activeProjects))
	fmt.Printf("Planning projects: %d\n", len(planningProjects))

	// User statistics
	allUsers := tms.ProjectService.GetAllUsers()
	developers := 0
	projectManagers := 0
	
	for _, user := range allUsers {
		if user.Role == models.Developer {
			developers++
		} else if user.Role == models.ProjectManager {
			projectManagers++
		}
	}

	fmt.Println()
	fmt.Println("18. USER STATISTICS:")
	fmt.Printf("Total users: %d\n", len(allUsers))
	fmt.Printf("Developers: %d\n", developers)
	fmt.Printf("Project Managers: %d\n", projectManagers)

	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}

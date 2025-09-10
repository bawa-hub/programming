package cli

import (
	"cli-todo-app/internal/todo"
	"cli-todo-app/pkg/models"
	"cli-todo-app/pkg/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// CLI represents the command-line interface
type CLI struct {
	service *todo.Service
	scanner *bufio.Scanner
}

// New creates a new CLI instance
func New(service *todo.Service) *CLI {
	return &CLI{
		service: service,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the CLI application
func (cli *CLI) Run() {
	cli.printWelcome()
	cli.printHelp()
	
	for {
		fmt.Print(utils.Bold(utils.Cyan("todo> ")))
		if !cli.scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(cli.scanner.Text())
		if input == "" {
			continue
		}
		
		// Parse and execute command
		if err := cli.executeCommand(input); err != nil {
			fmt.Printf("%s %s\n", utils.Red("Error:"), err.Error())
		}
	}
}

// executeCommand parses and executes a command
func (cli *CLI) executeCommand(input string) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}
	
	command := strings.ToLower(parts[0])
	args := parts[1:]
	
	switch command {
	case "help", "h":
		cli.printHelp()
	case "add", "a":
		return cli.handleAdd(args)
	case "list", "l":
		return cli.handleList(args)
	case "show", "s":
		return cli.handleShow(args)
	case "edit", "e":
		return cli.handleEdit(args)
	case "delete", "d":
		return cli.handleDelete(args)
	case "complete", "c":
		return cli.handleComplete(args)
	case "start":
		return cli.handleStart(args)
	case "priority", "p":
		return cli.handlePriority(args)
	case "due":
		return cli.handleDue(args)
	case "tag":
		return cli.handleTag(args)
	case "category", "cat":
		return cli.handleCategory(args)
	case "search":
		return cli.handleSearch(args)
	case "filter":
		return cli.handleFilter(args)
	case "sort":
		return cli.handleSort(args)
	case "stats":
		cli.handleStats()
	case "export":
		return cli.handleExport(args)
	case "import":
		return cli.handleImport(args)
	case "clear":
		return cli.handleClear(args)
	case "archive":
		return cli.handleArchive()
	case "quit", "q", "exit":
		fmt.Println(utils.Green("Goodbye!"))
		os.Exit(0)
	default:
		return fmt.Errorf("unknown command: %s. Type 'help' for available commands", command)
	}
	
	return nil
}

// Command handlers

func (cli *CLI) handleAdd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: add <title> [description]")
	}
	
	title := args[0]
	description := ""
	if len(args) > 1 {
		description = strings.Join(args[1:], " ")
	}
	
	todo, err := cli.service.AddTodo(title, description)
	if err != nil {
		return err
	}
	
	fmt.Printf("%s Created todo #%d: %s\n", 
		utils.Green("✓"), todo.ID, utils.Bold(todo.Title))
	return nil
}

func (cli *CLI) handleList(args []string) error {
	var todos []*models.Todo
	
	// Parse filter arguments
	if len(args) > 0 {
		switch strings.ToLower(args[0]) {
		case "pending":
			todos = cli.service.GetTodosByStatus(models.Pending)
		case "in-progress", "progress":
			todos = cli.service.GetTodosByStatus(models.InProgress)
		case "completed":
			todos = cli.service.GetTodosByStatus(models.Completed)
		case "overdue":
			todos = cli.service.GetOverdueTodos()
		case "today":
			todos = cli.service.GetTodosDueToday()
		case "all":
			todos = cli.service.GetAllTodos()
		default:
			return fmt.Errorf("unknown filter: %s. Use: pending, in-progress, completed, overdue, today, all", args[0])
		}
	} else {
		todos = cli.service.GetAllTodos()
	}
	
	if len(todos) == 0 {
		fmt.Println(utils.Yellow("No todos found."))
		return nil
	}
	
	// Sort todos by ID by default
	todos = cli.service.SortTodos(todos, "id", true)
	
	cli.printTodoList(todos)
	return nil
}

func (cli *CLI) handleShow(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: show <id>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	todo := cli.service.GetTodo(id)
	if todo == nil {
		return fmt.Errorf("todo #%d not found", id)
	}
	
	cli.printTodoDetail(todo)
	return nil
}

func (cli *CLI) handleEdit(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: edit <id> <title> [description]")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	title := args[1]
	description := ""
	if len(args) > 2 {
		description = strings.Join(args[2:], " ")
	}
	
	if err := cli.service.UpdateTodo(id, title, description); err != nil {
		return err
	}
	
	fmt.Printf("%s Updated todo #%d\n", utils.Green("✓"), id)
	return nil
}

func (cli *CLI) handleDelete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: delete <id>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	if err := cli.service.DeleteTodo(id); err != nil {
		return err
	}
	
	fmt.Printf("%s Deleted todo #%d\n", utils.Green("✓"), id)
	return nil
}

func (cli *CLI) handleComplete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: complete <id>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	if err := cli.service.CompleteTodo(id); err != nil {
		return err
	}
	
	fmt.Printf("%s Completed todo #%d\n", utils.Green("✓"), id)
	return nil
}

func (cli *CLI) handleStart(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: start <id>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	if err := cli.service.StartTodo(id); err != nil {
		return err
	}
	
	fmt.Printf("%s Started todo #%d\n", utils.Green("✓"), id)
	return nil
}

func (cli *CLI) handlePriority(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: priority <id> <low|medium|high|urgent>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	priorityStr := strings.ToLower(args[1])
	var priority models.Priority
	
	switch priorityStr {
	case "low":
		priority = models.Low
	case "medium":
		priority = models.Medium
	case "high":
		priority = models.High
	case "urgent":
		priority = models.Urgent
	default:
		return fmt.Errorf("invalid priority: %s. Use: low, medium, high, urgent", priorityStr)
	}
	
	if err := cli.service.SetPriority(id, priority); err != nil {
		return err
	}
	
	fmt.Printf("%s Set priority of todo #%d to %s\n", 
		utils.Green("✓"), id, utils.Bold(priority.String()))
	return nil
}

func (cli *CLI) handleDue(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: due <id> <date> [time]")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	dateStr := args[1]
	timeStr := "12:00"
	if len(args) > 2 {
		timeStr = args[2]
	}
	
	// Parse date and time
	datetimeStr := dateStr + " " + timeStr
	dueDate, err := time.Parse("2006-01-02 15:04", datetimeStr)
	if err != nil {
		// Try alternative formats
		formats := []string{
			"2006-01-02",
			"2006/01/02",
			"01/02/2006",
			"Jan 2, 2006",
			"January 2, 2006",
		}
		
		for _, format := range formats {
			if dueDate, err = time.Parse(format, dateStr); err == nil {
				break
			}
		}
		
		if err != nil {
			return fmt.Errorf("invalid date format: %s", dateStr)
		}
	}
	
	if err := cli.service.SetDueDate(id, dueDate); err != nil {
		return err
	}
	
	fmt.Printf("%s Set due date of todo #%d to %s\n", 
		utils.Green("✓"), id, dueDate.Format("2006-01-02 15:04"))
	return nil
}

func (cli *CLI) handleTag(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: tag <id> <add|remove> <tag>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	action := strings.ToLower(args[1])
	if action != "add" && action != "remove" {
		return fmt.Errorf("invalid action: %s. Use: add, remove", action)
	}
	
	if len(args) < 3 {
		return fmt.Errorf("usage: tag <id> <add|remove> <tag>")
	}
	
	tag := args[2]
	
	if action == "add" {
		if err := cli.service.AddTag(id, tag); err != nil {
			return err
		}
		fmt.Printf("%s Added tag '%s' to todo #%d\n", utils.Green("✓"), tag, id)
	} else {
		if err := cli.service.RemoveTag(id, tag); err != nil {
			return err
		}
		fmt.Printf("%s Removed tag '%s' from todo #%d\n", utils.Green("✓"), tag, id)
	}
	
	return nil
}

func (cli *CLI) handleCategory(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: category <list|add|set|delete> [args...]")
	}
	
	action := strings.ToLower(args[0])
	
	switch action {
	case "list":
		cli.handleCategoryList()
	case "add":
		return cli.handleCategoryAdd(args[1:])
	case "set":
		return cli.handleCategorySet(args[1:])
	case "delete":
		return cli.handleCategoryDelete(args[1:])
	default:
		return fmt.Errorf("unknown category action: %s", action)
	}
	
	return nil
}

func (cli *CLI) handleCategoryList() {
	categories := cli.service.GetAllCategories()
	
	if len(categories) == 0 {
		fmt.Println(utils.Yellow("No categories found."))
		return
	}
	
	fmt.Println(utils.Bold("Categories:"))
	for _, category := range categories {
		fmt.Printf("  %d. %s %s\n", 
			category.ID, 
			utils.Colorize(category.Name, category.Color),
			utils.Yellow(fmt.Sprintf("(%s)", category.Description)))
	}
}

func (cli *CLI) handleCategoryAdd(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: category add <name> [description] [color]")
	}
	
	name := args[0]
	description := ""
	color := utils.ColorWhite
	
	if len(args) > 1 {
		description = args[1]
	}
	if len(args) > 2 {
		color = args[2]
	}
	
	category, err := cli.service.AddCategory(name, description, color)
	if err != nil {
		return err
	}
	
	fmt.Printf("%s Created category #%d: %s\n", 
		utils.Green("✓"), category.ID, utils.Bold(category.Name))
	return nil
}

func (cli *CLI) handleCategorySet(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: category set <todo_id> <category_id>")
	}
	
	todoID, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid todo ID: %s", args[0])
	}
	
	categoryID, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid category ID: %s", args[1])
	}
	
	if err := cli.service.SetCategory(todoID, categoryID); err != nil {
		return err
	}
	
	fmt.Printf("%s Set category of todo #%d to category #%d\n", 
		utils.Green("✓"), todoID, categoryID)
	return nil
}

func (cli *CLI) handleCategoryDelete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: category delete <id>")
	}
	
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %s", args[0])
	}
	
	if err := cli.service.DeleteCategory(id); err != nil {
		return err
	}
	
	fmt.Printf("%s Deleted category #%d\n", utils.Green("✓"), id)
	return nil
}

func (cli *CLI) handleSearch(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: search <query>")
	}
	
	query := strings.Join(args, " ")
	todos := cli.service.SearchTodos(query)
	
	if len(todos) == 0 {
		fmt.Printf("%s No todos found matching '%s'\n", utils.Yellow("No results:"), query)
		return nil
	}
	
	fmt.Printf("%s Found %d todos matching '%s':\n", 
		utils.Green("Search results:"), len(todos), query)
	cli.printTodoList(todos)
	return nil
}

func (cli *CLI) handleFilter(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: filter <field> <value>")
	}
	
	field := strings.ToLower(args[0])
	value := args[1]
	
	var todos []*models.Todo
	
	switch field {
	case "priority":
		var priority models.Priority
		switch strings.ToLower(value) {
		case "low":
			priority = models.Low
		case "medium":
			priority = models.Medium
		case "high":
			priority = models.High
		case "urgent":
			priority = models.Urgent
		default:
			return fmt.Errorf("invalid priority: %s", value)
		}
		todos = cli.service.GetTodosByPriority(priority)
	case "status":
		var status models.Status
		switch strings.ToLower(value) {
		case "pending":
			status = models.Pending
		case "in-progress", "progress":
			status = models.InProgress
		case "completed":
			status = models.Completed
		case "cancelled":
			status = models.Cancelled
		default:
			return fmt.Errorf("invalid status: %s", value)
		}
		todos = cli.service.GetTodosByStatus(status)
	case "category":
		categoryID, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid category ID: %s", value)
		}
		todos = cli.service.GetTodosByCategory(categoryID)
	case "tag":
		todos = cli.service.SearchTodosByTag(value)
	default:
		return fmt.Errorf("unknown filter field: %s. Use: priority, status, category, tag", field)
	}
	
	if len(todos) == 0 {
		fmt.Printf("%s No todos found with %s=%s\n", utils.Yellow("No results:"), field, value)
		return nil
	}
	
	fmt.Printf("%s Found %d todos with %s=%s:\n", 
		utils.Green("Filter results:"), len(todos), field, value)
	cli.printTodoList(todos)
	return nil
}

func (cli *CLI) handleSort(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: sort <field> [asc|desc]")
	}
	
	field := strings.ToLower(args[0])
	ascending := true
	
	if len(args) > 1 {
		switch strings.ToLower(args[1]) {
		case "asc", "ascending":
			ascending = true
		case "desc", "descending":
			ascending = false
		default:
			return fmt.Errorf("invalid sort order: %s. Use: asc, desc", args[1])
		}
	}
	
	todos := cli.service.GetAllTodos()
	todos = cli.service.SortTodos(todos, field, ascending)
	
	order := "ascending"
	if !ascending {
		order = "descending"
	}
	
	fmt.Printf("%s Sorted todos by %s (%s):\n", 
		utils.Green("Sort results:"), field, order)
	cli.printTodoList(todos)
	return nil
}

func (cli *CLI) handleStats() {
	stats := cli.service.GetStats()
	
	fmt.Println(utils.Bold("Todo Statistics:"))
	fmt.Printf("  Total todos: %d\n", stats["total"])
	fmt.Printf("  Pending: %d\n", stats["pending"])
	fmt.Printf("  In Progress: %d\n", stats["in_progress"])
	fmt.Printf("  Completed: %d\n", stats["completed"])
	fmt.Printf("  Overdue: %d\n", stats["overdue"])
	fmt.Printf("  Due Today: %d\n", stats["due_today"])
	
	// Additional statistics
	statusCounts := cli.service.GetTodoCountByStatus()
	priorityCounts := cli.service.GetTodoCountByPriority()
	
	fmt.Println("\n" + utils.Bold("By Status:"))
	for status, count := range statusCounts {
		fmt.Printf("  %s: %d\n", status.String(), count)
	}
	
	fmt.Println("\n" + utils.Bold("By Priority:"))
	for priority, count := range priorityCounts {
		fmt.Printf("  %s: %d\n", priority.String(), count)
	}
	
	// Most used tags
	mostUsedTags := cli.service.GetMostUsedTags(5)
	if len(mostUsedTags) > 0 {
		fmt.Println("\n" + utils.Bold("Most Used Tags:"))
		for _, tagCount := range mostUsedTags {
			fmt.Printf("  %s\n", tagCount.String())
		}
	}
}

func (cli *CLI) handleExport(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: export <format> <filepath>")
	}
	
	format := strings.ToLower(args[0])
	filepath := args[1]
	
	switch format {
	case "json":
		if err := cli.service.ExportToJSON(filepath); err != nil {
			return err
		}
		fmt.Printf("%s Exported todos to JSON: %s\n", utils.Green("✓"), filepath)
	case "csv":
		if err := cli.service.ExportToCSV(filepath); err != nil {
			return err
		}
		fmt.Printf("%s Exported todos to CSV: %s\n", utils.Green("✓"), filepath)
	default:
		return fmt.Errorf("unsupported format: %s. Use: json, csv", format)
	}
	
	return nil
}

func (cli *CLI) handleImport(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: import <filepath>")
	}
	
	filepath := args[0]
	
	if err := cli.service.ImportFromJSON(filepath); err != nil {
		return err
	}
	
	fmt.Printf("%s Imported todos from: %s\n", utils.Green("✓"), filepath)
	return nil
}

func (cli *CLI) handleClear(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: clear <completed|overdue>")
	}
	
	target := strings.ToLower(args[0])
	
	switch target {
	case "completed":
		if err := cli.service.ClearCompleted(); err != nil {
			return err
		}
		fmt.Printf("%s Cleared all completed todos\n", utils.Green("✓"))
	case "overdue":
		if err := cli.service.ClearOverdue(); err != nil {
			return err
		}
		fmt.Printf("%s Cleared all overdue todos\n", utils.Green("✓"))
	default:
		return fmt.Errorf("invalid target: %s. Use: completed, overdue", target)
	}
	
	return nil
}

func (cli *CLI) handleArchive() error {
	if err := cli.service.ArchiveCompleted(); err != nil {
		return err
	}
	
	fmt.Printf("%s Archived all completed todos\n", utils.Green("✓"))
	return nil
}

// Display functions

func (cli *CLI) printWelcome() {
	fmt.Println(utils.Bold(utils.Cyan("📝 CLI Todo App")))
	fmt.Println(utils.Yellow("A comprehensive todo management system"))
	fmt.Println(utils.Yellow("Type 'help' for available commands"))
	fmt.Println()
}

func (cli *CLI) printHelp() {
	fmt.Println(utils.Bold("Available Commands:"))
	fmt.Println()
	
	commands := []struct {
		command string
		usage   string
		desc    string
	}{
		{"add, a", "add <title> [description]", "Add a new todo"},
		{"list, l", "list [filter]", "List todos (pending, in-progress, completed, overdue, today, all)"},
		{"show, s", "show <id>", "Show detailed information about a todo"},
		{"edit, e", "edit <id> <title> [description]", "Edit a todo"},
		{"delete, d", "delete <id>", "Delete a todo"},
		{"complete, c", "complete <id>", "Mark a todo as completed"},
		{"start", "start <id>", "Mark a todo as in progress"},
		{"priority, p", "priority <id> <level>", "Set priority (low, medium, high, urgent)"},
		{"due", "due <id> <date> [time]", "Set due date"},
		{"tag", "tag <id> <add|remove> <tag>", "Add or remove a tag"},
		{"category, cat", "category <action> [args...]", "Manage categories (list, add, set, delete)"},
		{"search", "search <query>", "Search todos by title or description"},
		{"filter", "filter <field> <value>", "Filter todos by field"},
		{"sort", "sort <field> [asc|desc]", "Sort todos by field"},
		{"stats", "stats", "Show statistics"},
		{"export", "export <format> <filepath>", "Export todos (json, csv)"},
		{"import", "import <filepath>", "Import todos from JSON"},
		{"clear", "clear <target>", "Clear todos (completed, overdue)"},
		{"archive", "archive", "Archive completed todos"},
		{"help, h", "help", "Show this help message"},
		{"quit, q, exit", "quit", "Exit the application"},
	}
	
	for _, cmd := range commands {
		fmt.Printf("  %-20s %-30s %s\n", 
			utils.Bold(cmd.command), 
			utils.Yellow(cmd.usage), 
			cmd.desc)
	}
	fmt.Println()
}

func (cli *CLI) printTodoList(todos []*models.Todo) {
	if len(todos) == 0 {
		fmt.Println(utils.Yellow("No todos found."))
		return
	}
	
	fmt.Printf("%-4s %-8s %-8s %-30s %-15s %-10s\n",
		utils.Bold("ID"),
		utils.Bold("Status"),
		utils.Bold("Priority"),
		utils.Bold("Title"),
		utils.Bold("Due Date"),
		utils.Bold("Category"))
	
	fmt.Println(strings.Repeat("-", 80))
	
	for _, todo := range todos {
		statusIcon := todo.Status.GetStatusIcon()
		priorityColor := todo.Priority.GetPriorityColor()
		
		category := "None"
		if todo.Category != nil {
			category = todo.Category.Name
		}
		
		dueDate := "No due date"
		if todo.DueDate != nil {
			dueDate = todo.DueDate.Format("2006-01-02 15:04")
		}
		
		title := utils.TruncateString(todo.Title, 30)
		
		fmt.Printf("%-4d %-8s %-8s %-30s %-15s %-10s\n",
			todo.ID,
			statusIcon+" "+todo.Status.String(),
			priorityColor+" "+todo.Priority.String(),
			title,
			dueDate,
			category)
	}
	fmt.Println()
}

func (cli *CLI) printTodoDetail(todo *models.Todo) {
	fmt.Println(utils.Bold("Todo Details:"))
	fmt.Printf("  ID: %d\n", todo.ID)
	fmt.Printf("  Title: %s\n", utils.Bold(todo.Title))
	fmt.Printf("  Description: %s\n", todo.Description)
	fmt.Printf("  Status: %s %s\n", todo.Status.GetStatusIcon(), todo.Status.String())
	fmt.Printf("  Priority: %s %s\n", todo.Priority.GetPriorityColor(), todo.Priority.String())
	
	if todo.Category != nil {
		fmt.Printf("  Category: %s (%s)\n", todo.Category.Name, todo.Category.Description)
	} else {
		fmt.Printf("  Category: None\n")
	}
	
	if len(todo.Tags) > 0 {
		fmt.Printf("  Tags: %s\n", strings.Join(todo.Tags, ", "))
	} else {
		fmt.Printf("  Tags: None\n")
	}
	
	fmt.Printf("  Due Date: %s\n", todo.GetFormattedDueDate())
	fmt.Printf("  Created: %s\n", todo.GetFormattedCreatedAt())
	fmt.Printf("  Updated: %s\n", todo.GetFormattedUpdatedAt())
	
	if todo.CompletedAt != nil {
		fmt.Printf("  Completed: %s\n", todo.GetFormattedCompletedAt())
	}
	
	if todo.IsOverdue() {
		fmt.Printf("  %s\n", utils.Red("⚠️  OVERDUE"))
	}
	
	daysUntilDue := todo.DaysUntilDue()
	if daysUntilDue >= 0 {
		if daysUntilDue == 0 {
			fmt.Printf("  %s\n", utils.Yellow("⚠️  Due today"))
		} else if daysUntilDue == 1 {
			fmt.Printf("  %s\n", utils.Yellow("⚠️  Due tomorrow"))
		} else if daysUntilDue <= 3 {
			fmt.Printf("  %s\n", utils.Yellow(fmt.Sprintf("⚠️  Due in %d days", daysUntilDue)))
		}
	}
}

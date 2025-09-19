package services

import (
	"fmt"
	"sort"
	"sync"

	"task_management/models"
	"task_management/utils"
)

// Project Service
type ProjectService struct {
	projects map[string]*models.Project
	teams    map[string]*models.Team
	users    map[string]*models.User
	mu       sync.RWMutex
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		projects: make(map[string]*models.Project),
		teams:    make(map[string]*models.Team),
		users:    make(map[string]*models.User),
	}
}

// User Management
func (ps *ProjectService) RegisterUser(username, email, passwordHash, firstName, lastName string) (*models.User, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	// Check if user already exists
	for _, user := range ps.users {
		if user.Username == username || user.Email == email {
			return nil, fmt.Errorf("user already exists")
		}
	}
	
	user := models.NewUser(username, email, passwordHash, firstName, lastName)
	ps.users[user.ID] = user
	return user, nil
}

func (ps *ProjectService) GetUser(userID string) *models.User {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	return ps.users[userID]
}

func (ps *ProjectService) GetUserByUsername(username string) *models.User {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	for _, user := range ps.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}

func (ps *ProjectService) GetAllUsers() []*models.User {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var users []*models.User
	for _, user := range ps.users {
		users = append(users, user)
	}
	
	// Sort by username
	sort.Slice(users, func(i, j int) bool {
		return users[i].Username < users[j].Username
	})
	
	return users
}

// Team Management
func (ps *ProjectService) CreateTeam(name, description string, createdBy *models.User) (*models.Team, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	team := models.NewTeam(name, description, createdBy)
	ps.teams[team.ID] = team
	
	// Add creator as team member
	team.AddMember(createdBy)
	
	return team, nil
}

func (ps *ProjectService) GetTeam(teamID string) *models.Team {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	return ps.teams[teamID]
}

func (ps *ProjectService) AddTeamMember(teamID, userID string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	team := ps.teams[teamID]
	user := ps.users[userID]
	
	if team == nil {
		return fmt.Errorf("team not found")
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}
	
	return team.AddMember(user)
}

func (ps *ProjectService) RemoveTeamMember(teamID, userID string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	team := ps.teams[teamID]
	if team == nil {
		return fmt.Errorf("team not found")
	}
	
	return team.RemoveMember(userID)
}

func (ps *ProjectService) GetAllTeams() []*models.Team {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var teams []*models.Team
	for _, team := range ps.teams {
		teams = append(teams, team)
	}
	
	// Sort by name
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Name < teams[j].Name
	})
	
	return teams
}

// Project Management
func (ps *ProjectService) CreateProject(name, description string, owner *models.User) (*models.Project, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	// Check if user has permission to create projects
	if !owner.HasPermission("create_project") {
		return nil, fmt.Errorf("user does not have permission to create projects")
	}
	
	project := models.NewProject(name, description, owner)
	ps.projects[project.ID] = project
	
	// Add owner as project member
	project.AddMember(owner)
	
	return project, nil
}

func (ps *ProjectService) GetProject(projectID string) *models.Project {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	return ps.projects[projectID]
}

func (ps *ProjectService) AddProjectMember(projectID, userID string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	project := ps.projects[projectID]
	user := ps.users[userID]
	
	if project == nil {
		return fmt.Errorf("project not found")
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}
	
	return project.AddMember(user)
}

func (ps *ProjectService) RemoveProjectMember(projectID, userID string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	project := ps.projects[projectID]
	if project == nil {
		return fmt.Errorf("project not found")
	}
	
	return project.RemoveMember(userID)
}

func (ps *ProjectService) UpdateProjectStatus(projectID string, status models.ProjectStatus, user *models.User) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	project := ps.projects[projectID]
	if project == nil {
		return fmt.Errorf("project not found")
	}
	
	// Check if user has permission to update project
	if !ps.canUserModifyProject(project, user) {
		return fmt.Errorf("user does not have permission to update this project")
	}
	
	project.UpdateStatus(status)
	return nil
}

func (ps *ProjectService) UpdateProjectSettings(projectID string, settings models.ProjectSettings, user *models.User) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	project := ps.projects[projectID]
	if project == nil {
		return fmt.Errorf("project not found")
	}
	
	// Check if user has permission to update project
	if !ps.canUserModifyProject(project, user) {
		return fmt.Errorf("user does not have permission to update this project")
	}
	
	project.UpdateSettings(settings)
	return nil
}

func (ps *ProjectService) UpdateProjectDetails(projectID, name, description string, user *models.User) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	
	project := ps.projects[projectID]
	if project == nil {
		return fmt.Errorf("project not found")
	}
	
	// Check if user has permission to update project
	if !ps.canUserModifyProject(project, user) {
		return fmt.Errorf("user does not have permission to update this project")
	}
	
	project.UpdateDetails(name, description)
	return nil
}

func (ps *ProjectService) GetAllProjects() []*models.Project {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var projects []*models.Project
	for _, project := range ps.projects {
		projects = append(projects, project)
	}
	
	// Sort by name
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})
	
	return projects
}

func (ps *ProjectService) GetUserProjects(userID string) []*models.Project {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var userProjects []*models.Project
	for _, project := range ps.projects {
		if project.CanUserAccess(userID) {
			userProjects = append(userProjects, project)
		}
	}
	
	// Sort by name
	sort.Slice(userProjects, func(i, j int) bool {
		return userProjects[i].Name < userProjects[j].Name
	})
	
	return userProjects
}

// Search and Filtering
func (ps *ProjectService) SearchProjects(query string) []*models.Project {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var results []*models.Project
	for _, project := range ps.projects {
		if utils.ContainsIgnoreCase(project.Name, query) || utils.ContainsIgnoreCase(project.Description, query) {
			results = append(results, project)
		}
	}
	
	// Sort by name
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})
	
	return results
}

func (ps *ProjectService) GetProjectsByStatus(status models.ProjectStatus) []*models.Project {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var projects []*models.Project
	for _, project := range ps.projects {
		if project.Status == status {
			projects = append(projects, project)
		}
	}
	
	// Sort by name
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})
	
	return projects
}

func (ps *ProjectService) GetProjectsByOwner(ownerID string) []*models.Project {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	
	var projects []*models.Project
	for _, project := range ps.projects {
		if project.Owner.ID == ownerID {
			projects = append(projects, project)
		}
	}
	
	// Sort by name
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})
	
	return projects
}

// Helper Methods
func (ps *ProjectService) canUserModifyProject(project *models.Project, user *models.User) bool {
	// Admin can modify any project
	if user.Role == models.Admin {
		return true
	}
	
	// Project owner can modify their own projects
	if project.Owner.ID == user.ID {
		return true
	}
	
	// Project managers can modify projects they own
	if user.Role == models.ProjectManager && project.Owner.ID == user.ID {
		return true
	}
	
	return false
}


package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// =============================================================================
// CORE ENTITIES
// =============================================================================

// User Roles
type UserRole int

const (
	RegularUser UserRole = iota
	Moderator
	Admin
)

func (ur UserRole) String() string {
	switch ur {
	case RegularUser:
		return "Regular User"
	case Moderator:
		return "Moderator"
	case Admin:
		return "Admin"
	default:
		return "Unknown"
	}
}

// Vote Types
type VoteType int

const (
	Upvote VoteType = iota
	Downvote
)

func (vt VoteType) String() string {
	switch vt {
	case Upvote:
		return "Upvote"
	case Downvote:
		return "Downvote"
	default:
		return "Unknown"
	}
}

// Content Types
type ContentType int

const (
	QuestionContent ContentType = iota
	AnswerContent
)

func (ct ContentType) String() string {
	switch ct {
	case QuestionContent:
		return "Question"
	case AnswerContent:
		return "Answer"
	default:
		return "Unknown"
	}
}

// Badge Types
type BadgeType int

const (
	QuestionBadge BadgeType = iota
	AnswerBadge
	VotingBadge
	ReputationBadge
)

func (bt BadgeType) String() string {
	switch bt {
	case QuestionBadge:
		return "Question Badge"
	case AnswerBadge:
		return "Answer Badge"
	case VotingBadge:
		return "Voting Badge"
	case ReputationBadge:
		return "Reputation Badge"
	default:
		return "Unknown"
	}
}

// =============================================================================
// USER SYSTEM
// =============================================================================

// User
type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	Reputation   int
	Role         UserRole
	Badges       []*Badge
	JoinDate     time.Time
	IsActive     bool
	mu           sync.RWMutex
}

func NewUser(username, email, passwordHash string) *User {
	return &User{
		ID:           fmt.Sprintf("U%d", time.Now().UnixNano()),
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Reputation:   1, // Starting reputation
		Role:         RegularUser,
		Badges:       make([]*Badge, 0),
		JoinDate:     time.Now(),
		IsActive:     true,
	}
}

func (u *User) AddReputation(points int) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Reputation += points
	if u.Reputation < 1 {
		u.Reputation = 1 // Minimum reputation
	}
}

func (u *User) GetReputation() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Reputation
}

func (u *User) AddBadge(badge *Badge) {
	u.mu.Lock()
	defer u.mu.Unlock()
	
	// Check if user already has this badge
	for _, existingBadge := range u.Badges {
		if existingBadge.ID == badge.ID {
			return
		}
	}
	
	u.Badges = append(u.Badges, badge)
}

func (u *User) HasPrivilege(privilege string) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	
	switch privilege {
	case "moderate":
		return u.Role == Moderator || u.Role == Admin
	case "admin":
		return u.Role == Admin
	case "vote":
		return u.Reputation >= 15
	case "comment":
		return u.Reputation >= 50
	case "edit_others":
		return u.Reputation >= 2000
	default:
		return false
	}
}

func (u *User) GetBadges() []*Badge {
	u.mu.RLock()
	defer u.mu.RUnlock()
	
	badges := make([]*Badge, len(u.Badges))
	copy(badges, u.Badges)
	return badges
}

// =============================================================================
// CONTENT SYSTEM
// =============================================================================

// Content Interface
type Content interface {
	GetID() string
	GetAuthor() *User
	GetContentType() ContentType
	GetVotes() []*Vote
	GetScore() int
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	IsDeleted() bool
	SetDeleted(deleted bool)
}

// Question
type Question struct {
	ID          string
	Title       string
	Body        string
	Tags        []*Tag
	Author      *User
	Answers     []*Answer
	Votes       []*Vote
	Views       int
	IsClosed    bool
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	mu          sync.RWMutex
}

func NewQuestion(title, body string, author *User, tags []*Tag) *Question {
	return &Question{
		ID:        fmt.Sprintf("Q%d", time.Now().UnixNano()),
		Title:     title,
		Body:      body,
		Tags:      tags,
		Author:    author,
		Answers:   make([]*Answer, 0),
		Votes:     make([]*Vote, 0),
		Views:     0,
		IsClosed:  false,
		Deleted:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (q *Question) GetID() string {
	return q.ID
}

func (q *Question) GetAuthor() *User {
	return q.Author
}

func (q *Question) GetContentType() ContentType {
	return QuestionContent
}

func (q *Question) GetVotes() []*Vote {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.Votes
}

func (q *Question) GetScore() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	
	score := 0
	for _, vote := range q.Votes {
		if vote.Type == Upvote {
			score++
		} else {
			score--
		}
	}
	return score
}

func (q *Question) GetCreatedAt() time.Time {
	return q.CreatedAt
}

func (q *Question) GetUpdatedAt() time.Time {
	return q.UpdatedAt
}

func (q *Question) IsDeleted() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.Deleted
}

func (q *Question) SetDeleted(deleted bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Deleted = deleted
}

func (q *Question) AddAnswer(answer *Answer) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Answers = append(q.Answers, answer)
}

func (q *Question) GetAnswers() []*Answer {
	q.mu.RLock()
	defer q.mu.RUnlock()
	
	answers := make([]*Answer, len(q.Answers))
	copy(answers, q.Answers)
	return answers
}

func (q *Question) IncrementViews() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Views++
}

func (q *Question) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.IsClosed = true
}

func (q *Question) UpdateTitle(title string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Title = title
	q.UpdatedAt = time.Now()
}

func (q *Question) UpdateBody(body string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Body = body
	q.UpdatedAt = time.Now()
}

// Answer
type Answer struct {
	ID          string
	Body        string
	Author      *User
	Question    *Question
	Votes       []*Vote
	IsAccepted  bool
	Deleted     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	mu          sync.RWMutex
}

func NewAnswer(body string, author *User, question *Question) *Answer {
	return &Answer{
		ID:         fmt.Sprintf("A%d", time.Now().UnixNano()),
		Body:       body,
		Author:     author,
		Question:   question,
		Votes:      make([]*Vote, 0),
		IsAccepted: false,
		Deleted:    false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (a *Answer) GetID() string {
	return a.ID
}

func (a *Answer) GetAuthor() *User {
	return a.Author
}

func (a *Answer) GetContentType() ContentType {
	return AnswerContent
}

func (a *Answer) GetVotes() []*Vote {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Votes
}

func (a *Answer) GetScore() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	
	score := 0
	for _, vote := range a.Votes {
		if vote.Type == Upvote {
			score++
		} else {
			score--
		}
	}
	return score
}

func (a *Answer) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *Answer) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a *Answer) IsDeleted() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Deleted
}

func (a *Answer) SetDeleted(deleted bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Deleted = deleted
}

func (a *Answer) Accept() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.IsAccepted = true
}

func (a *Answer) UpdateBody(body string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Body = body
	a.UpdatedAt = time.Now()
}

// =============================================================================
// VOTING SYSTEM
// =============================================================================

// Vote
type Vote struct {
	ID        string
	User      *User
	Content   Content
	Type      VoteType
	Timestamp time.Time
}

func NewVote(user *User, content Content, voteType VoteType) *Vote {
	return &Vote{
		ID:        fmt.Sprintf("V%d", time.Now().UnixNano()),
		User:      user,
		Content:   content,
		Type:      voteType,
		Timestamp: time.Now(),
	}
}

// =============================================================================
// TAG SYSTEM
// =============================================================================

// Tag
type Tag struct {
	ID          string
	Name        string
	Description string
	UsageCount  int
	IsActive    bool
	mu          sync.RWMutex
}

func NewTag(name, description string) *Tag {
	return &Tag{
		ID:          fmt.Sprintf("T%d", time.Now().UnixNano()),
		Name:        name,
		Description: description,
		UsageCount:  0,
		IsActive:    true,
	}
}

func (t *Tag) IncrementUsage() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.UsageCount++
}

func (t *Tag) DecrementUsage() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.UsageCount > 0 {
		t.UsageCount--
	}
}

func (t *Tag) GetUsageCount() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.UsageCount
}

// =============================================================================
// BADGE SYSTEM
// =============================================================================

// Badge
type Badge struct {
	ID          string
	Name        string
	Description string
	Type        BadgeType
	Criteria    string
	Icon        string
}

func NewBadge(name, description string, badgeType BadgeType, criteria, icon string) *Badge {
	return &Badge{
		ID:          fmt.Sprintf("B%d", time.Now().UnixNano()),
		Name:        name,
		Description: description,
		Type:        badgeType,
		Criteria:    criteria,
		Icon:        icon,
	}
}

// Badge Manager
type BadgeManager struct {
	badges map[string]*Badge
	mu     sync.RWMutex
}

func NewBadgeManager() *BadgeManager {
	bm := &BadgeManager{
		badges: make(map[string]*Badge),
	}
	bm.initializeDefaultBadges()
	return bm
}

func (bm *BadgeManager) initializeDefaultBadges() {
	// Question badges
	bm.addBadge(NewBadge("First Question", "Asked your first question", QuestionBadge, "questions >= 1", "🎯"))
	bm.addBadge(NewBadge("Questioner", "Asked 10 questions", QuestionBadge, "questions >= 10", "❓"))
	bm.addBadge(NewBadge("Inquisitive", "Asked 100 questions", QuestionBadge, "questions >= 100", "🔍"))
	
	// Answer badges
	bm.addBadge(NewBadge("First Answer", "Answered your first question", AnswerBadge, "answers >= 1", "💡"))
	bm.addBadge(NewBadge("Answerer", "Answered 10 questions", AnswerBadge, "answers >= 10", "📝"))
	bm.addBadge(NewBadge("Expert", "Answered 100 questions", AnswerBadge, "answers >= 100", "🎓"))
	
	// Voting badges
	bm.addBadge(NewBadge("Voter", "Voted 10 times", VotingBadge, "votes >= 10", "🗳️"))
	bm.addBadge(NewBadge("Active Voter", "Voted 100 times", VotingBadge, "votes >= 100", "📊"))
	
	// Reputation badges
	bm.addBadge(NewBadge("Rising Star", "Reached 100 reputation", ReputationBadge, "reputation >= 100", "⭐"))
	bm.addBadge(NewBadge("Expert", "Reached 1000 reputation", ReputationBadge, "reputation >= 1000", "🏆"))
	bm.addBadge(NewBadge("Legend", "Reached 10000 reputation", ReputationBadge, "reputation >= 10000", "👑"))
}

func (bm *BadgeManager) addBadge(badge *Badge) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	bm.badges[badge.ID] = badge
}

func (bm *BadgeManager) GetBadge(id string) *Badge {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	return bm.badges[id]
}

func (bm *BadgeManager) GetAllBadges() []*Badge {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	badges := make([]*Badge, 0, len(bm.badges))
	for _, badge := range bm.badges {
		badges = append(badges, badge)
	}
	return badges
}

// =============================================================================
// STACK OVERFLOW SERVICE
// =============================================================================

// Stack Overflow Service
type StackOverflowService struct {
	Users        map[string]*User
	Questions    map[string]*Question
	Answers      map[string]*Answer
	Tags         map[string]*Tag
	Votes        map[string]*Vote
	BadgeManager *BadgeManager
	mu           sync.RWMutex
}

func NewStackOverflowService() *StackOverflowService {
	return &StackOverflowService{
		Users:        make(map[string]*User),
		Questions:    make(map[string]*Question),
		Answers:      make(map[string]*Answer),
		Tags:         make(map[string]*Tag),
		Votes:        make(map[string]*Vote),
		BadgeManager: NewBadgeManager(),
	}
}

// User Management
func (sos *StackOverflowService) RegisterUser(username, email, passwordHash string) (*User, error) {
	sos.mu.Lock()
	defer sos.mu.Unlock()
	
	// Check if user already exists
	for _, user := range sos.Users {
		if user.Username == username || user.Email == email {
			return nil, fmt.Errorf("user already exists")
		}
	}
	
	user := NewUser(username, email, passwordHash)
	sos.Users[user.ID] = user
	return user, nil
}

func (sos *StackOverflowService) GetUser(userID string) *User {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	return sos.Users[userID]
}

func (sos *StackOverflowService) GetUserByUsername(username string) *User {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	
	for _, user := range sos.Users {
		if user.Username == username {
			return user
		}
	}
	return nil
}

// Question Management
func (sos *StackOverflowService) PostQuestion(title, body string, author *User, tagNames []string) (*Question, error) {
	sos.mu.Lock()
	defer sos.mu.Unlock()
	
	// Get or create tags
	tags := make([]*Tag, 0)
	for _, tagName := range tagNames {
		tag := sos.getOrCreateTag(tagName)
		tags = append(tags, tag)
		tag.IncrementUsage()
	}
	
	question := NewQuestion(title, body, author, tags)
	sos.Questions[question.ID] = question
	
	// Award badge for first question
	if author.GetReputation() < 100 {
		badge := sos.BadgeManager.GetBadge("B1") // First Question badge
		if badge != nil {
			author.AddBadge(badge)
		}
	}
	
	return question, nil
}

func (sos *StackOverflowService) GetQuestion(questionID string) *Question {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	return sos.Questions[questionID]
}

func (sos *StackOverflowService) SearchQuestions(query string) []*Question {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	
	var results []*Question
	queryLower := strings.ToLower(query)
	
	for _, question := range sos.Questions {
		if question.IsDeleted() {
			continue
		}
		
		titleMatch := strings.Contains(strings.ToLower(question.Title), queryLower)
		bodyMatch := strings.Contains(strings.ToLower(question.Body), queryLower)
		
		if titleMatch || bodyMatch {
			results = append(results, question)
		}
	}
	
	// Sort by score (votes) and recency
	sort.Slice(results, func(i, j int) bool {
		scoreI := results[i].GetScore()
		scoreJ := results[j].GetScore()
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		return results[i].CreatedAt.After(results[j].CreatedAt)
	})
	
	return results
}

// Answer Management
func (sos *StackOverflowService) PostAnswer(body string, author *User, question *Question) (*Answer, error) {
	sos.mu.Lock()
	defer sos.mu.Unlock()
	
	answer := NewAnswer(body, author, question)
	sos.Answers[answer.ID] = answer
	question.AddAnswer(answer)
	
	// Award badge for first answer
	if author.GetReputation() < 100 {
		badge := sos.BadgeManager.GetBadge("B4") // First Answer badge
		if badge != nil {
			author.AddBadge(badge)
		}
	}
	
	return answer, nil
}

func (sos *StackOverflowService) GetAnswers(questionID string) []*Answer {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	
	question := sos.Questions[questionID]
	if question == nil {
		return nil
	}
	
	answers := question.GetAnswers()
	
	// Sort by acceptance status, then by score
	sort.Slice(answers, func(i, j int) bool {
		if answers[i].IsAccepted != answers[j].IsAccepted {
			return answers[i].IsAccepted
		}
		return answers[i].GetScore() > answers[j].GetScore()
	})
	
	return answers
}

func (sos *StackOverflowService) AcceptAnswer(answerID string) error {
	sos.mu.Lock()
	defer sos.mu.Unlock()
	
	answer := sos.Answers[answerID]
	if answer == nil {
		return fmt.Errorf("answer not found")
	}
	
	// Unaccept all other answers for this question
	for _, existingAnswer := range answer.Question.GetAnswers() {
		if existingAnswer.ID != answerID {
			existingAnswer.mu.Lock()
			existingAnswer.IsAccepted = false
			existingAnswer.mu.Unlock()
		}
	}
	
	answer.Accept()
	return nil
}

// Voting System
func (sos *StackOverflowService) VoteContent(userID, contentID string, voteType VoteType) error {
	sos.mu.Lock()
	defer sos.mu.Unlock()
	
	user := sos.Users[userID]
	if user == nil {
		return fmt.Errorf("user not found")
	}
	
	if !user.HasPrivilege("vote") {
		return fmt.Errorf("insufficient reputation to vote")
	}
	
	var content Content
	if question := sos.Questions[contentID]; question != nil {
		content = question
	} else if answer := sos.Answers[contentID]; answer != nil {
		content = answer
	} else {
		return fmt.Errorf("content not found")
	}
	
	// Check if user already voted
	for _, existingVote := range content.GetVotes() {
		if existingVote.User.ID == userID {
			// User already voted, update the vote
			existingVote.Type = voteType
			return nil
		}
	}
	
	// Create new vote
	vote := NewVote(user, content, voteType)
	sos.Votes[vote.ID] = vote
	
	// Add vote to content
	switch c := content.(type) {
	case *Question:
		c.mu.Lock()
		c.Votes = append(c.Votes, vote)
		c.mu.Unlock()
	case *Answer:
		c.mu.Lock()
		c.Votes = append(c.Votes, vote)
		c.mu.Unlock()
	}
	
	// Update reputation
	reputationChange := sos.calculateReputationChange(content, voteType)
	content.GetAuthor().AddReputation(reputationChange)
	
	return nil
}

func (sos *StackOverflowService) calculateReputationChange(content Content, voteType VoteType) int {
	switch content.GetContentType() {
	case QuestionContent:
		if voteType == Upvote {
			return 5
		} else {
			return -1
		}
	case AnswerContent:
		if voteType == Upvote {
			return 10
		} else {
			return -2
		}
	default:
		return 0
	}
}

// Tag Management
func (sos *StackOverflowService) getOrCreateTag(name string) *Tag {
	tag := sos.Tags[name]
	if tag == nil {
		tag = NewTag(name, fmt.Sprintf("Tag for %s", name))
		sos.Tags[name] = tag
	}
	return tag
}

func (sos *StackOverflowService) GetTag(name string) *Tag {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	return sos.Tags[name]
}

func (sos *StackOverflowService) GetPopularTags() []*Tag {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	
	var tags []*Tag
	for _, tag := range sos.Tags {
		tags = append(tags, tag)
	}
	
	// Sort by usage count
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].GetUsageCount() > tags[j].GetUsageCount()
	})
	
	return tags
}

// Statistics
func (sos *StackOverflowService) GetUserStats(userID string) map[string]interface{} {
	sos.mu.RLock()
	defer sos.mu.RUnlock()
	
	user := sos.Users[userID]
	if user == nil {
		return nil
	}
	
	questionCount := 0
	answerCount := 0
	voteCount := 0
	
	for _, question := range sos.Questions {
		if question.Author.ID == userID && !question.IsDeleted() {
			questionCount++
		}
	}
	
	for _, answer := range sos.Answers {
		if answer.Author.ID == userID && !answer.IsDeleted() {
			answerCount++
		}
	}
	
	for _, vote := range sos.Votes {
		if vote.User.ID == userID {
			voteCount++
		}
	}
	
	return map[string]interface{}{
		"reputation":    user.GetReputation(),
		"questions":     questionCount,
		"answers":       answerCount,
		"votes":         voteCount,
		"badges":        len(user.GetBadges()),
		"join_date":     user.JoinDate,
	}
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== STACK OVERFLOW SYSTEM DEMONSTRATION ===\n")

	// Create Stack Overflow service
	so := NewStackOverflowService()
	
	// Register users
	user1, _ := so.RegisterUser("alice", "alice@example.com", "hash1")
	user2, _ := so.RegisterUser("bob", "bob@example.com", "hash2")
	user3, _ := so.RegisterUser("charlie", "charlie@example.com", "hash3")
	
	fmt.Println("1. USER REGISTRATION:")
	fmt.Printf("User 1: %s (ID: %s)\n", user1.Username, user1.ID)
	fmt.Printf("User 2: %s (ID: %s)\n", user2.Username, user2.ID)
	fmt.Printf("User 3: %s (ID: %s)\n", user3.Username, user3.ID)
	
	// Give users some reputation to enable voting
	user1.AddReputation(50)
	user2.AddReputation(100)
	user3.AddReputation(25)
	
	fmt.Println()
	
	// Post questions
	fmt.Println("2. POSTING QUESTIONS:")
	question1, _ := so.PostQuestion(
		"How to implement a binary tree in Go?",
		"I'm learning Go and want to implement a binary tree. Can someone help me with the basic structure?",
		user1,
		[]string{"go", "data-structures", "binary-tree"},
	)
	fmt.Printf("Question 1 posted: %s\n", question1.Title)
	
	question2, _ := so.PostQuestion(
		"What are design patterns?",
		"I keep hearing about design patterns but don't understand what they are. Can someone explain?",
		user2,
		[]string{"design-patterns", "programming", "architecture"},
	)
	fmt.Printf("Question 2 posted: %s\n", question2.Title)
	
	question3, _ := so.PostQuestion(
		"Best practices for database design",
		"What are some best practices I should follow when designing a database schema?",
		user3,
		[]string{"database", "sql", "design", "best-practices"},
	)
	fmt.Printf("Question 3 posted: %s\n", question3.Title)
	
	fmt.Println()
	
	// Post answers
	fmt.Println("3. POSTING ANSWERS:")
	answer1, _ := so.PostAnswer(
		"Here's a basic binary tree implementation in Go:\n\n```go\ntype TreeNode struct {\n    Value int\n    Left  *TreeNode\n    Right *TreeNode\n}\n```",
		user2,
		question1,
	)
	fmt.Printf("Answer 1 posted for question 1\n")
	
	answer2, _ := so.PostAnswer(
		"Design patterns are reusable solutions to common problems in software design. They provide templates for solving specific design issues.",
		user3,
		question2,
	)
	fmt.Printf("Answer 2 posted for question 2\n")
	
	_, _ = so.PostAnswer(
		"Some key best practices:\n1. Normalize your data\n2. Use appropriate indexes\n3. Design for scalability\n4. Consider data integrity",
		user1,
		question3,
	)
	fmt.Printf("Answer 3 posted for question 3\n")
	
	fmt.Println()
	
	// Voting
	fmt.Println("4. VOTING SYSTEM:")
	
	// Users vote on questions
	so.VoteContent(user2.ID, question1.ID, Upvote)
	so.VoteContent(user3.ID, question1.ID, Upvote)
	fmt.Printf("Question 1 received 2 upvotes (Score: %d)\n", question1.GetScore())
	
	so.VoteContent(user1.ID, question2.ID, Upvote)
	so.VoteContent(user3.ID, question2.ID, Upvote)
	fmt.Printf("Question 2 received 2 upvotes (Score: %d)\n", question2.GetScore())
	
	// Users vote on answers
	so.VoteContent(user1.ID, answer1.ID, Upvote)
	so.VoteContent(user3.ID, answer1.ID, Upvote)
	fmt.Printf("Answer 1 received 2 upvotes (Score: %d)\n", answer1.GetScore())
	
	so.VoteContent(user1.ID, answer2.ID, Upvote)
	so.VoteContent(user2.ID, answer2.ID, Upvote)
	fmt.Printf("Answer 2 received 2 upvotes (Score: %d)\n", answer2.GetScore())
	
	// Accept an answer
	so.AcceptAnswer(answer1.ID)
	fmt.Printf("Answer 1 accepted for question 1\n")
	
	fmt.Println()
	
	// Display user reputations
	fmt.Println("5. USER REPUTATIONS:")
	fmt.Printf("Alice: %d reputation\n", user1.GetReputation())
	fmt.Printf("Bob: %d reputation\n", user2.GetReputation())
	fmt.Printf("Charlie: %d reputation\n", user3.GetReputation())
	
	fmt.Println()
	
	// Search functionality
	fmt.Println("6. SEARCH FUNCTIONALITY:")
	searchResults := so.SearchQuestions("design")
	fmt.Printf("Search results for 'design': %d questions found\n", len(searchResults))
	for i, question := range searchResults {
		if i < 3 { // Show first 3 results
			fmt.Printf("  %d. %s (Score: %d)\n", i+1, question.Title, question.GetScore())
		}
	}
	
	fmt.Println()
	
	// Popular tags
	fmt.Println("7. POPULAR TAGS:")
	popularTags := so.GetPopularTags()
	for i, tag := range popularTags {
		if i < 5 { // Show top 5 tags
			fmt.Printf("  %s: %d uses\n", tag.Name, tag.GetUsageCount())
		}
	}
	
	fmt.Println()
	
	// User statistics
	fmt.Println("8. USER STATISTICS:")
	stats1 := so.GetUserStats(user1.ID)
	fmt.Printf("Alice's stats: %+v\n", stats1)
	
	stats2 := so.GetUserStats(user2.ID)
	fmt.Printf("Bob's stats: %+v\n", stats2)
	
	fmt.Println()
	
	// Badge system
	fmt.Println("9. BADGE SYSTEM:")
	fmt.Printf("Alice's badges: %d\n", len(user1.GetBadges()))
	for _, badge := range user1.GetBadges() {
		fmt.Printf("  %s: %s\n", badge.Name, badge.Description)
	}
	
	fmt.Printf("Bob's badges: %d\n", len(user2.GetBadges()))
	for _, badge := range user2.GetBadges() {
		fmt.Printf("  %s: %s\n", badge.Name, badge.Description)
	}
	
	fmt.Println()
	
	// Question details
	fmt.Println("10. QUESTION DETAILS:")
	fmt.Printf("Question 1: %s\n", question1.Title)
	fmt.Printf("  Score: %d\n", question1.GetScore())
	fmt.Printf("  Answers: %d\n", len(question1.GetAnswers()))
	fmt.Printf("  Tags: ")
	for _, tag := range question1.Tags {
		fmt.Printf("%s ", tag.Name)
	}
	fmt.Println()
	
	// Show answers for question 1
	answers := so.GetAnswers(question1.ID)
	fmt.Printf("  Answers for question 1:\n")
	for i, answer := range answers {
		fmt.Printf("    %d. Score: %d, Accepted: %t\n", i+1, answer.GetScore(), answer.IsAccepted)
	}
	
	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}

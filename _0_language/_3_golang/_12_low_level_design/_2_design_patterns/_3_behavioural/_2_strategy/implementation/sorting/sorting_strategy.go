package sorting


type SortingStrategy interface {
	Sort(data []int) []int
	GetName() string
	GetTimeComplexity() string
}
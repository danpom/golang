package school

// Define the Grade and School types here.

type Grade struct {
	student []string
	level   int
}

type School map[int]Grade

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	(*s)[grade] = append((*s)[grade], student)
}

func (s *School) Grade(level int) []string {
	return (*s)[level]
}

func (s *School) Enrollment() []Grade {
	return
}

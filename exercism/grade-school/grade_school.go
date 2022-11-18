package school

import "sort"

// Define the Grade and School types here.

type Grade struct {
	level    int
	students []string
}

//type School map[int]*Grade
type School map[int]*Grade

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	if _, ok := (*s)[grade]; !ok {
		(*s)[grade] = &Grade{
			level:    grade,
			students: []string{student},
		}
	} else {
		(*s)[grade].students = append((*s)[grade].students, student)
	}
}

func (s *School) Grade(level int) []string {
	for k := range *s {
		if k == level {
			return (*s)[level].students
		}
	}
	return []string{}

}

func (s *School) Enrollment() []Grade {
	var res []Grade
	for _, v := range *s {
		//sort students alphabetically
		sortStudents := v.students
		sort.Slice(sortStudents, func(i, j int) bool {
			return sortStudents[i] < sortStudents[j]
		})
		g := Grade{
			students: sortStudents,
			level:    v.level,
		}
		res = append(res, g)
	}
	//sort grades ascending
	sort.Slice(res, func(i, j int) bool {
		return res[i].level < res[j].level
	})

	return res

}

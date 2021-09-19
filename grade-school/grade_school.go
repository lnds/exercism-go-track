package school

type Grade struct {
	grade    int
	students []string
}

//
type School struct {
	grades map[int]Grade
}

func New() *School {
	return &School{
		grades: map[int]Grade{},
	}
}

func (s *School) Add(student string, i int) {
	if grade, ok := s.grades[i]; ok {
		grade.students = append(grade.students, student)
		for i := len(grade.students) - 1; i > 0; i-- {
			if grade.students[i] < grade.students[i-1] {
				grade.students[i], grade.students[i-1] = grade.students[i-1], grade.students[i]
			}
		}
		s.grades[i] = grade
	} else {
		s.grades[i] = Grade{grade: i, students: []string{student}}
	}
}

func (s *School) Grade(grade int) []string {
	return s.grades[grade].students
}

func (s *School) Enrollment() []Grade {
	grades := []Grade{}
	for _, g := range s.grades {
		grades = append(grades, g)
		for i := len(grades) - 1; i > 0; i-- {
			if grades[i].grade < grades[i-1].grade {
				grades[i], grades[i-1] = grades[i-1], grades[i]
			}
		}
	}
	return grades
}

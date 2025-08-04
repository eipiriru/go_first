package webapps

var students []StudentMaster

type StudentMaster struct {
	Id   int64
	Name string
	Sex  string
}

func getStudents(id int64) []StudentMaster {
	var result StudentMaster
	if id != 0 {
		for _, s := range students {
			if s.Id == id {
				result = s
			}
		}
		return []StudentMaster{result}
	}
	return students
}

func students_init() {
	s1 := StudentMaster{Name: "Alice", Id: 1, Sex: "Perempuan"}
	s2 := StudentMaster{Name: "Bob", Id: 2, Sex: "Laki-laki"}

	// Correct usage: Assign the result of append back to 'students'
	students = append(students, StudentMaster{Id: 3, Name: "Charlie", Sex: "Laki-laki"})
	students = append(students, s1)
	students = append(students, s2)
}

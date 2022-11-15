package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "ning",
			LastName:  "huang",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
			},
		},
		{
			ID:        2,
			FirstName: "xiao",
			LastName:  "wang",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
			},
		}
	}
}

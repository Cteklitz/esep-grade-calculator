package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}


func (gc *GradeCalculator) GetPassing() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 70 {
		return "Pass"
	} else {
		return "Fail"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	weighted_grade := computeAverage(gc.grades)
	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	assignment_weight := 0.5
	exam_weight := 0.35
	essay_weight := 0.15

	assignment_sum := 0
	exam_sum := 0
	essay_sum := 0

	assignment_count := 0
	exam_count := 0
	essay_count := 0

	for _, grade := range grades {
		switch grade.Type.String() {
		case "assignment":
			assignment_sum += grade.Grade
			assignment_count++
		case "exam":
			exam_sum += grade.Grade
			exam_count++
		case "essay":
			essay_sum += grade.Grade
			essay_count++
		}
	}

	out := (float64(assignment_sum / assignment_count) * assignment_weight) + (float64(exam_sum / exam_count) * exam_weight) + (float64(essay_sum / essay_count) * essay_weight)

	return int(out)
}

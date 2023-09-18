package resolver

type Student struct {
	Index           int    `json:"index"`
	StudentId       int    `json:"studentId"`
	Name            string `json:"name"`
	EnglishName     string `json:"englishName"`
	Gender          string `json:"gender"`
	Grade           int    `json:"grade"`
	School          string `json:"school"`
	Major           string `json:"major"`
	IsInternational bool   `json:"isInternational"`
}

type Course struct {
	Name       string    `json:"name"`
	Id         string    `json:"id"`
	Code       string    `json:"code"`
	StudentNum int       `json:"studentNum"`
	Class      string    `json:"class"`
	School     string    `json:"school"`
	Teacher    string    `json:"teacher"`
	Time       string    `json:"time"`
	Students   []Student `json:"students"`
}

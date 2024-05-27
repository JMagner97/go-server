package response

type EnrollmentResponse struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	CourseName string `json:"coursename"`
	StartYear  int    `json:"startyear"`
	EndYear    int    `json:"endyear"`
}

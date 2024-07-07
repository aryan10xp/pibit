package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Defining the structure for the resume output
type Resume struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Experience []string `json:"experience"`
	Education  []string `json:"education"`
	Skills     []string `json:"skills"`
}

func main() {
	// Sample resume text
	resumeText := `
	John Doe
	Email: aryan.mantri@gmail.com
	Phone: 7349039203

	Experience:
	- Software Engineer at OLA 
	- Machine Learning Intern at Harrier Information Systems

	Education:
	- B.Tech in computer science engineering

	Skills:
	- Go, Python, JavaScript
	- SQL, NoSQL
	- Golang, HTML
	`

	// Parse the resume text
	resume := parseResume(resumeText)

	// Convert the resume to JSON
	jsonData, err := json.Marshal(resume)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}

// Function to parse the resume text
func parseResume(text string) Resume {
	lines := strings.Split(text, "\n")
	var resume Resume

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		// Parse name
		if i == 0 {
			resume.Name = line
		}

		// Parse email
		if strings.HasPrefix(line, "Email:") {
			resume.Email = strings.TrimSpace(strings.TrimPrefix(line, "Email:"))
		}

		// Parse phone
		if strings.HasPrefix(line, "Phone:") {
			resume.Phone = strings.TrimSpace(strings.TrimPrefix(line, "Phone:"))
		}

		// Parse experience
		if strings.HasPrefix(line, "Experience:") {
			for j := i + 1; j < len(lines); j++ {
				expLine := strings.TrimSpace(lines[j])
				if expLine == "" {
					break
				}
				resume.Experience = append(resume.Experience, expLine)
			}
		}

		// Parse education
		if strings.HasPrefix(line, "Education:") {
			for j := i + 1; j < len(lines); j++ {
				eduLine := strings.TrimSpace(lines[j])
				if eduLine == "" {
					break
				}
				resume.Education = append(resume.Education, eduLine)
			}
		}

		// Parse skills
		if strings.HasPrefix(line, "Skills:") {
			for j := i + 1; j < len(lines); j++ {
				skillLine := strings.TrimSpace(lines[j])
				if skillLine == "" {
					break
				}
				skills := strings.Split(skillLine, ",")
				for _, skill := range skills {
					resume.Skills = append(resume.Skills, strings.TrimSpace(skill))
				}
			}
		}
	}

	return resume
}

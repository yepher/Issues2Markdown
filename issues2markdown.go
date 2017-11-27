package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

/**
BitbucketIssues the structure of Bitbucket issues json export
**/
type BitbucketIssues struct {
	Milestones  []interface{} `json:"milestones"`
	Attachments []struct {
		Path     string `json:"path"`
		Issue    int    `json:"issue"`
		User     string `json:"user"`
		Filename string `json:"filename"`
	} `json:"attachments"`

	Versions []interface{} `json:"versions"`
	Comments []struct {
		Content   interface{} `json:"content"`
		CreatedOn time.Time   `json:"created_on"`
		User      string      `json:"user"`
		UpdatedOn interface{} `json:"updated_on"`
		Issue     int         `json:"issue"`
		ID        int         `json:"id"`
	} `json:"comments"`
	Meta struct {
		DefaultMilestone interface{} `json:"default_milestone"`
		DefaultAssignee  interface{} `json:"default_assignee"`
		DefaultKind      string      `json:"default_kind"`
		DefaultComponent interface{} `json:"default_component"`
		DefaultVersion   interface{} `json:"default_version"`
	} `json:"meta"`
	Components []struct {
		Name string `json:"name"`
	} `json:"components"`
	Issues []struct {
		Status           string        `json:"status"`
		Priority         string        `json:"priority"`
		Kind             string        `json:"kind"`
		ContentUpdatedOn interface{}   `json:"content_updated_on"`
		Voters           []interface{} `json:"voters"`
		Title            string        `json:"title"`
		Reporter         string        `json:"reporter"`
		Component        string        `json:"component"`
		Watchers         []string      `json:"watchers"`
		Content          string        `json:"content"`
		Assignee         interface{}   `json:"assignee"`
		CreatedOn        time.Time     `json:"created_on"`
		Version          interface{}   `json:"version"`
		EditedOn         interface{}   `json:"edited_on"`
		Milestone        interface{}   `json:"milestone"`
		UpdatedOn        time.Time     `json:"updated_on"`
		ID               int           `json:"id"`
	} `json:"issues"`
	Logs []struct {
		Comment     int       `json:"comment"`
		ChangedTo   string    `json:"changed_to"`
		Field       string    `json:"field"`
		CreatedOn   time.Time `json:"created_on"`
		User        string    `json:"user"`
		Issue       int       `json:"issue"`
		ChangedFrom string    `json:"changed_from"`
	} `json:"logs"`
}

func main() {
	fileArg := flag.String("file", "./db-1.0.json", "The exported issues list to be rendered")
	flag.Parse()

	if *fileArg == "" {
		fmt.Println("Error: -file is a required field\n\n ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, e := ioutil.ReadFile(*fileArg)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))

	var issuesType BitbucketIssues
	json.Unmarshal(file, &issuesType)
	//fmt.Printf("Results: %v\n", issuesType)

	output := ""
	toc := "| Title |Component |Priority| Id |\n|---|---|---|---|\n"

	sort.Slice(issuesType.Issues, func(i, j int) bool {
		return statusToInt(issuesType.Issues[i].Priority) < statusToInt(issuesType.Issues[j].Priority)
	})

	sort.Slice(issuesType.Comments, func(i, j int) bool {
		return issuesType.Comments[i].CreatedOn.Before(issuesType.Comments[j].CreatedOn)
	})

	for _, element := range issuesType.Issues {
		if element.Status == "resolved" {
			continue
		}

		//fmt.Printf("TODO: Render: [%d]: %s\n", index, element.Title)
		toc = fmt.Sprintf("%s|%s | %s | %s |<a href='#anchor_%d'>%d</a>|\n", toc, element.Title, element.Component, element.Priority, element.ID, element.ID)

		output = fmt.Sprintf("%s<a name=\"anchor_%d\"></a>\n", output, element.ID)
		output = fmt.Sprintf("%s## Issue %d - %s\n\n", output, element.ID, element.Title)

		output = fmt.Sprintf("%s%s\n\n", output, element.Title)

		output = fmt.Sprintf("%s| Field | Value |\n", output)
		output = fmt.Sprintf("%s|---|---|\n", output)
		output = fmt.Sprintf("%s|%s|%s|\n", output, "Component", element.Component)
		output = fmt.Sprintf("%s|%s|%s|\n", output, "Kind", element.Kind)
		output = fmt.Sprintf("%s|%s|%s|\n", output, "Status", element.Status)
		output = fmt.Sprintf("%s|%s|%s|\n", output, "Priority", element.Priority)
		output = fmt.Sprintf("%s|%s|%s|\n", output, "Created", element.CreatedOn)
		output = fmt.Sprintf("%s|%s|%s|\n", output, "Updated", element.UpdatedOn)
		output = fmt.Sprintf("%s\n\n", output)

		output = fmt.Sprintf("%s### Details\n\n", output)
		output = fmt.Sprintf("%s%s\n\n", output, element.Content)

		comments := loadComment(element.ID, issuesType)
		output = fmt.Sprintf("%s\n%s\n\n", output, comments)

		output = fmt.Sprintf("%s------------------\n<a href='#top'>top<a>\n\n", output)

	}

	fmt.Printf("<a name=\"top\"></a>\n\n")
	fmt.Println(toc)
	fmt.Printf("\n\n")
	fmt.Println(output)
}

func loadComment(id int, issues BitbucketIssues) string {
	comments := ""

	for _, element := range issues.Comments {
		//fmt.Printf("Enter Comment(%d, %d)\n", element.Issue, id)
		if element.Issue == id {
			if element.Content != nil && element.Content != "" {
				comments = fmt.Sprintf("%s\n**Comment**\n\n", comments)

				comments = fmt.Sprintf("%s| Field | Value |\n|---|---|\n", comments)
				comments = fmt.Sprintf("%s|**By:**| %s |\n", comments, element.User)
				comments = fmt.Sprintf("%s|**Created:** | %s|\n", comments, element.CreatedOn)

				comments = fmt.Sprintf("%s\n%s\n", comments, element.Content)
				comments = fmt.Sprintf("%s\n\n", comments)

			}
		}
	}
	return comments
}

func statusToInt(priority string) int {
	switch priority {
	case "blocker":
		return 1
	case "critical":
		return 2
	case "major":
		return 3
	case "minor":
		return 4
	case "trivial":
		return 5
	default:
		fmt.Printf("ERROR: Unknown Priority: %s\n", priority)
		return 0
	}
}

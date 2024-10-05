package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
)

var dictionary = []string{"literally", "all", "the", "words"}

type ShamelessSelfPromoter interface {
	promote(application *JobApplication) error
}

type Job struct {
	role           string
	companyName    string
	applicationUrl string
}

type JobApplication struct {
	Name     string   `json:"name"`
	Skills   []string `json:"skills"`
	Photo    os.File  `json:"photo,omitempty"`
	URL      string   `json:"url"`
	JobTitle string   `json:"jobTitle"`
}

type DencoSeca struct {
	firstName string
	lastName  string
	scaredOf  string
	headshots map[string]os.File
	Portfolio string
	Home      string
	Skills    []string
}

func New(headshotsDir string) *DencoSeca {
	headshots := make(map[string]os.File)

	files, err := os.ReadDir(headshotsDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			f, err := os.Open(file.Name())
			if err != nil {
				log.Fatalf("Failed to open file: %v", err)
			}
			headshots[file.Name()] = *f
		}
	}

	return &DencoSeca{
		headshots: headshots,
		firstName: "Leon",
		lastName:  "Brown",
		Home:      "Edinburgh",
		scaredOf:  "Spiders that disappear when you look away for like, *one* second.",
	}
}

func (d *DencoSeca) AddSkill(skill string) {
	if !slices.Contains(d.Skills, strings.ToLower(skill)) {
		d.Skills = append(d.Skills, strings.ToLower(skill))
	}
}

func (d *DencoSeca) IsCompatibleWithJob(requiredSkills []string, description, companyAddress string) bool {
	if strings.Contains(description, d.scaredOf) {
		log.Fatalf("%s is now a NO-GO zone", companyAddress)
	}

	var matchedSkills []string
	for _, skill := range requiredSkills {
		if slices.Contains(d.Skills, skill) {
			matchedSkills = append(matchedSkills, skill)
		}
	}

	if len(matchedSkills) >= 4 {
		log.Println("Bills paid 💷")
	} else {
		log.Println("Apply anyway and learn fast! 👍")
	}

	return true

}

func (d *DencoSeca) ApplyForJob(job Job) error {
	useTooMuchHairGel := slices.Contains(dictionary, strings.Replace(job.companyName, "s", "z", -1))

	var photo os.File
	if useTooMuchHairGel {
		photo = d.headshots["naughtiesBoyband"]
	} else {
		photo = d.headshots["tastefulCableknit"]
	}

	application := &JobApplication{
		Name:     d.firstName + " " + d.lastName,
		Skills:   d.Skills,
		Photo:    photo,
		URL:      job.applicationUrl,
		JobTitle: job.role,
	}

	err := d.promote(application)
	if err != nil {
		return err
	}

	return nil

}

func (d *DencoSeca) promote(application *JobApplication) error {
	jsonData, err := json.Marshal(application)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", application.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New(response.Status)
	}

	return nil
}

func main() {
	headshotsDir := "./stock-photos-of-random-people-smiling"

	var yourNewBestFriend = New(headshotsDir)
	yourNewBestFriend.AddSkill("Kansas City Shuffle")

	job := Job{
		role:           "Transpondster",
		companyName:    "InstantSolutionz",
		applicationUrl: "https://www.instant-solutionz.com/apply",
	}

	if err := yourNewBestFriend.ApplyForJob(job); err != nil {
		log.Fatal(err)
	}
}

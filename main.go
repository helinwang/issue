package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
)

func main() {
	url := flag.String("repo", "https://github.com/PaddlePaddle/Paddle", "repo URL")
	issues := flag.String("issues", "", "comma separated list of issues")
	flag.Parse()

	if *issues == "" {
		flag.Usage()
		return
	}

	ss := strings.Split(*url, "github.com/")
	sp := strings.Split(ss[1], "/")
	owner := sp[0]
	repo := sp[1]

	client := github.NewClient(nil)
	is := strings.Split(*issues, ",")

	var log string
	waitFeedback := 0
	moreDetail := 0
	closed := 0
	for _, issue := range is {
		number, err := strconv.Atoi(issue)
		if err != nil {
			panic(err)
		}

		i, _, err := client.Issues.Get(context.TODO(), owner, repo, number)
		if err != nil {
			panic(err)
		}

		comments, _, err := client.Issues.ListComments(context.TODO(), owner, repo, number, nil)
		if err != nil {
			panic(err)
		}

		var replyUsers string
		idx := 0
		dup := make(map[string]bool)
		for _, c := range comments {
			name := *c.User.Login
			if name == *i.User.Login {
				continue
			}

			if dup[name] {
				continue
			}

			if idx == 0 {
				replyUsers = name
			} else {
				replyUsers += ", " + name
			}

			idx++
			dup[name] = true
		}

		var labels string
		checked := false

		state := *i.State
		if state == "closed" {
			closed++
			state += " ✅"
			checked = true
		}

		for i, l := range i.Labels {
			name := *l.Name
			switch name {
			case "Waiting for User Feedback":
				if !checked {
					waitFeedback++
					checked = true
				}
				name += " 🔵"
			case "Need More Details":
				if !checked {
					moreDetail++
					checked = true
				}
				name += " 🔵"
			}

			if i == 0 {
				labels = name
			} else {
				labels += ", " + name
			}
		}

		link := fmt.Sprintf(`https://github.com/%s/%s/issues/%s`, owner, repo, issue)
		log += fmt.Sprintf(`Title: %s
Link: %s
Labels: %s
Reply: %s
State: %s

`, *i.Title, link, labels, replyUsers, state)
	}

	stat := fmt.Sprintf("Total: %d, Closed: %d", len(is), closed)
	if waitFeedback > 0 {
		stat += fmt.Sprintf(", Waiting user feedback: %d", waitFeedback)
	}

	if moreDetail > 0 {
		stat += fmt.Sprintf(", Need more detail: %d", moreDetail)
	}

	fmt.Printf("Github Issues (%s)\n\n", stat)
	fmt.Println(log)
}

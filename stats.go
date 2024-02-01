package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const sixMonthsInDays int = 182

var now = time.Now()

func stats(email string, repos []string) {
	commits := processRepos(repos, email)
	fmt.Println()
	printTable(commits)
}

func fillCommits(path, email string, commits map[int]int) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	commitIterator, err := repo.Log(&git.LogOptions{Since: &since, Until: &until})
	if err != nil {
		return err
	}

	err = commitIterator.ForEach(func(c *object.Commit) error {
		if c.Author.Email != email {
			return nil
		}

		days := daysAgo(c.Author.When)
		if days < 0 {
			return nil
		}
		commits[days]++
		return nil
	})
	return err
}

func processRepos(repos []string, email string) map[int]int {
	m := map[int]int{}
	var err error
	for _, repo := range repos {
		err = fillCommits(repo, email, m)
		if err != nil {
			fmt.Printf("failed to fill commits in %q: %v", repo, err)
		}
	}
	return m
}

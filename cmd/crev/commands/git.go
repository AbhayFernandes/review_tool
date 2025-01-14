package commands

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func getCurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		// Let's think about a better way to do this
		return "", err
	}
	return dir, nil
}

func getRepository(dir string) *git.Repository {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		fmt.Println("Something went wrong opening your git directory. Are you in a git repo?")
	}
	return repo
}

func getRemoteCommit(repo *git.Repository) (*object.Commit, error) {
	// make this dynamic in future and default to either main/master
	remoteRefName := plumbing.ReferenceName("refs/remotes/origin/main")
	remoteRef, err := repo.Reference(remoteRefName, true)

	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(remoteRef.Hash())

	return commit, nil
}

func getPatchDiffs(repository *git.Repository) (string, error) {
	head, err := repository.Head()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	local_commit, _ := repository.CommitObject(head.Hash())
	remote_commit, err := getRemoteCommit(repository)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	local_tree, err := local_commit.Tree()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	remote_tree, err := remote_commit.Tree()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	changes, _ := remote_tree.Diff(local_tree)
	patch, _ := changes.Patch()

	return patch.String(), nil
}

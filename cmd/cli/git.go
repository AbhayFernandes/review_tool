package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func getCurrentDir() string {
    dir, err := os.Getwd()
    if err != nil {
        // Let's think about a better way to do this
        panic(err)
    }
    return dir 
}

func getWorkTree(dir string) *git.Worktree {
    r, err := git.PlainOpen(dir)
    if err != nil {
        fmt.Println("Something went wrong opening your git directory. Are you in a git repo?")
    }

    worktree, err := r.Worktree()
    if err != nil {
        fmt.Println("Something went wrong opening your git directory. Are you in a git repo?")
    }

    return worktree
}

func getRepository(dir string) *git.Repository {
    repo, err := git.PlainOpen(dir)
    if err != nil {
        fmt.Println("Something went wrong opening your git directory. Are you in a git repo?")
    }
  return repo
}

func getRemoteCommit(repo *git.Repository) (*object.Commit, error) {
    // Clone a bare repository into memory
    storer := memory.NewStorage()

    remotes, _ := repo.Remotes()
    remote := remotes[0]

    mem_repo, err := git.Clone(storer, nil, &git.CloneOptions{
        URL: remote.Config().URLs[0],
    })
    if (err != nil) {
        fmt.Println("There was an error obtaining information from the remote repository")
        return nil, err
    }

    head, err := mem_repo.Head()
    if (err != nil) {
        fmt.Println("There was an error getting the remote repo HEAD.")
        return nil, err
    }

    commit, err := mem_repo.CommitObject(head.Hash())
    if (err != nil) {
        fmt.Println("There was an error converting the remote repo's HEAD reference.")
        return nil, err
    }

    return commit, nil
}

func getPatchDiffs(repository *git.Repository) (string, error) {
    head, _ := repository.Head()
    local_commit,_ := repository.CommitObject(head.Hash())
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

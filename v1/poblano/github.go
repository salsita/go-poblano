// Copyright (c) 2014 The go-poblano AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package poblano

import (
	"errors"
	"fmt"
	"net/http"
)

type GitHubService struct {
	client *Client
}

func newGitHubService(client *Client) *GitHubService {
	return &GitHubService{client}
}

func (srv *GitHubService) GetPoblanoProject(repoOwner, repoName string) (*Project, *http.Response, error) {
	u := fmt.Sprintf("/api/projects?where[services.github.fullName]=%v/%v", repoOwner, repoName)
	req, err := srv.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var projects []*Project
	resp, err := srv.client.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}

	switch len(projects) {
	case 0:
		return nil, resp, errors.New("Poblano project record not found")
	case 1:
		return projects[0], resp, nil
	default:
		return nil, resp, errors.New("Poblano returned multiple project records")
	}
}

func (srv *GitHubService) GetPoblanoUser(login string) (*User, *http.Response, error) {
	u := fmt.Sprintf("/api/users?where[services.github.username]=%v", login)
	req, err := srv.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := srv.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	switch len(users) {
	case 0:
		return nil, resp, errors.New("Poblano user record not found")
	case 1:
		return users[0], resp, nil
	default:
		return nil, resp, errors.New("Poblano returned multiple user records")
	}
}

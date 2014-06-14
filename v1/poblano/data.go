// Copyright (c) 2014 The go-poblano AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package poblano

type Project struct {
	Name        string
	Slug        string
	Description string
	Services    struct {
		GitHub *struct {
			Id        int
			Name      string
			URL       string
			Connected bool
		} `json:"github"`
		PivotalTracker *struct {
			Id        int
			URL       string
			Connected bool
		} `json:"pivotalTracker"`
	}
}

type User struct {
	Name     string
	Email    string
	Services struct {
		GitHub *struct {
			Username    string
			AccessToken string
			Connected   bool
		} `json:"github"`
		PivotalTracker *struct {
			Id          int
			Username    string
			AccessToken string
			Connected   bool
		} `json:"pivotalTracker"`
	}
}

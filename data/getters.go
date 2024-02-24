package data

import (
	"errors"
	"slices"
)

func GetBlogs() ([]Blog, error) {
	return blogs.Get(), nil
}

func GetBlogByPublicId(id string) (Blog, error) {
	index := slices.IndexFunc(blogs.Get(), func(b Blog) bool {
		return b.PublicId == id
	})
	if index < 0 {
		return Blog{}, errors.New("blog was not found")
	}
	return blogs.Get()[index], nil
}

func GetProjectGroups() ([]ProjectGroup, error) {
	return projects.Get(), nil
}

func GetInfo() (Info, error) {
	return info.Get(), nil
}

func GetContactLinks() ([]ContactLink, error) {
	return nil, nil
}

func GetWorkXP() ([]WorkExperience, error) {
	return work.Get(), nil
}

func GetVolunteeringXP() ([]VolunteeringExperience, error) {
	return volunteering.Get(), nil
}

package data

import (
	"errors"
	"slices"
)

func IncrementBlogReads(publicId string) error {
	index := slices.IndexFunc(blogs.Get(), func(b Blog) bool {
		return b.PublicId == publicId
	})
	if index < 0 {
		return errors.New("blog was not found")
	}
	blogs.Get()[index].VisitTimes++
	return nil
}

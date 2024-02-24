package data

import (
	"encoding/json"
	"errors"
	"os"
	"slices"

	"github.com/mbaraa/mbaraa.com/config"
	"github.com/mbaraa/mbaraa.com/log"
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

func UpdateBlogsMeta() error {
	log.Infoln("Updating blogs files...")
	fileNames, err := getBlogsFileNames()
	if err != nil {
		return err
	}

	publicIdMemoryBlogIndexMap := map[string]Blog{}
	for _, fileName := range fileNames {
		index := slices.IndexFunc(blogs.Get(), func(b Blog) bool {
			return b.PublicId == fileName
		})
		if index < 0 {
			continue
		}
		publicIdMemoryBlogIndexMap[config.Config().FilesDir+"/blogs/"+fileName+".json"] = blogs.Get()[index]
	}

	for fileName, blog := range publicIdMemoryBlogIndexMap {
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			continue
		}
		blog.Content = ""
		formattedJson, err := json.MarshalIndent(blog, "", "  ")
		if err != nil {
			continue
		}
		_, err = f.Write(formattedJson)
		if err != nil {
			continue
		}
	}

	log.Infoln("Finished updating blogs files  ✓")
	return nil
}

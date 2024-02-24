package data

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/mbaraa/mbaraa.com/config"
	"github.com/mbaraa/mbaraa.com/log"
)

type MutexWrapper[T any] struct {
	data T
	mu   sync.RWMutex
}

func (w *MutexWrapper[T]) Get() T {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.data
}

func (w *MutexWrapper[T]) Set(data T) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.data = data
}

var (
	blogs        = &MutexWrapper[[]Blog]{}
	info         = &MutexWrapper[Info]{}
	projects     = &MutexWrapper[[]ProjectGroup]{}
	work         = &MutexWrapper[[]WorkExperience]{}
	volunteering = &MutexWrapper[[]VolunteeringExperience]{}
)

func init() {
	err := loadBlogs()
	if err != nil {
		log.Fatalln(string(log.ErrorLevel), err)
	}
	err = loadVolunteeringXP()
	if err != nil {
		log.Fatalln(string(log.ErrorLevel), err)
	}
	err = loadWorkXP()
	if err != nil {
		log.Fatalln(string(log.ErrorLevel), err)
	}
	err = loadProjects()
	if err != nil {
		log.Fatalln(string(log.ErrorLevel), err)
	}
	err = loadInfo()
	if err != nil {
		log.Fatalln(string(log.ErrorLevel), err)
	}
}

func loadVolunteeringXP() error {
	vols, err := loadJsonData[[]VolunteeringExperience](config.Config().FilesDir + "/data/volunteering.json")
	if err != nil {
		return err
	}
	volunteering.Set(vols)
	return nil
}

func loadWorkXP() error {
	works, err := loadJsonData[[]WorkExperience](config.Config().FilesDir + "/data/work.json")
	if err != nil {
		return err
	}
	work.Set(works)
	return nil
}

func loadProjects() error {
	prjcts, err := loadJsonData[[]ProjectGroup](config.Config().FilesDir + "/data/projects.json")
	if err != nil {
		return err
	}
	slices.SortFunc(prjcts, func(pgI, pgJ ProjectGroup) int {
		return int(pgI.Order) - int(pgJ.Order)
	})
	projects.Set(prjcts)
	return nil
}

func loadInfo() error {
	inf, err := loadJsonData[Info](config.Config().FilesDir + "/data/about.json")
	if err != nil {
		return err
	}
	info.Set(inf)
	return nil
}

func loadJsonData[T any](filePath string) (out T, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	err = json.NewDecoder(file).Decode(&out)
	if err != nil {
		return
	}
	return
}

func loadBlogs() error {
	fileNames, err := getBlogsFileNames()
	if err != nil {
		return err
	}

	var blgs []Blog
	for _, fileName := range fileNames {
		blog, err := readBlogMeta(fileName)
		if err != nil {
			log.Errorf("error reading blog's metadata for %s: %s", fileName, err)
			continue
		}
		blog.Content, err = readBlogContent(fileName)
		if err != nil {
			log.Errorf("error reading blog's content for %s: %s", fileName, err)
			continue
		}
		blgs = append(blgs, blog)
	}
	slices.SortFunc(blgs, func(a, b Blog) int {
		return int(b.WrittenAt) - int(a.WrittenAt)
	})
	blogs.Set(blgs)

	return nil
}

func readBlogContent(fileName string) (string, error) {
	blogContentFile, err := os.ReadFile(fmt.Sprintf("%s/blogs/%s.md", config.Config().FilesDir, fileName))
	if err != nil {
		return "", err
	}
	return string(blogContentFile), nil
}

func readBlogMeta(fileName string) (blog Blog, err error) {
	blogMetaFile, err := os.Open(fmt.Sprintf("%s/blogs/%s.json", config.Config().FilesDir, fileName))
	if err != nil {
		return
	}
	err = json.NewDecoder(blogMetaFile).Decode(&blog)
	if err != nil {
		return
	}
	return
}

func getBlogsFileNames() (fileNames []string, err error) {
	fileNamesMap := map[string]bool{}
	err = filepath.WalkDir(config.Config().FilesDir+"/blogs", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileNamesMap[filepath.Base(path[:strings.LastIndex(path, ".")])] = true
		}
		return nil
	})
	if err != nil {
		return
	}
	for k := range fileNamesMap {
		fileNames = append(fileNames, k)
	}

	return
}

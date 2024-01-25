package db

import (
	"regexp"
	"slices"
	"strings"
	"time"
)

func InsertBlog(blog Blog) error {
	blog.PublicId = toKebab(blog.Title)
	blog.WrittenAt = time.Now().Unix()

	_, err := blogColl.InsertOne(ctx, blog)
	if err != nil {
		return err
	}
	return nil
}

func InsertProjectGroup(pg ProjectGroup) error {
	_, err := projectGroupColl.InsertOne(ctx, pg)
	if err != nil {
		return err
	}
	return nil
}

func InsertVolunteeringXP(xp VolunteeringExperience) error {
	_, err := volunteeringExperienceColl.InsertOne(ctx, xp)
	if err != nil {
		return err
	}
	return nil
}

func InsertWorkXP(xp WorkExperience) error {
	_, err := workExperienceColl.InsertOne(ctx, xp)
	if err != nil {
		return err
	}
	return nil
}

func InsertInfo(info Info) error {
	_, err := infoColl.InsertOne(ctx, info)
	if err != nil {
		return err
	}
	return nil
}

func toKebab(s string) string {
	patt := regexp.MustCompile(`^[a-z0-9-]+$`)
	return strings.Join(slices.DeleteFunc(strings.Split(strings.ReplaceAll(strings.ToLower(s), " ", "-"), ""), func(s string) bool {
		return !patt.MatchString(s)
	}), "")
}

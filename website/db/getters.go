package db

import (
	"slices"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlogs() ([]Blog, error) {
	blogs, err := getDocuments[Blog](blogColl)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(blogs, func(a, b Blog) int {
		return int(a.WrittenAt) - int(b.WrittenAt)
	})

	return blogs, nil
}

func GetBlog(id string) (Blog, error) {
	return getDocumentById[Blog](blogColl, id)
}

func GetProjectGroups() ([]ProjectGroup, error) {
	return getDocuments[ProjectGroup](projectGroupColl)
}

func GetProjectGroup(id string) (ProjectGroup, error) {
	return getDocumentById[ProjectGroup](projectGroupColl, id)
}

func GetInfo() (Info, error) {
	row := infoColl.FindOne(ctx, bson.M{})
	if row.Err() != nil {
		return Info{}, row.Err()
	}
	var info Info
	err := row.Decode(&info)
	if err != nil {
		return Info{}, err
	}

	return info, nil
}

func GetContactLinks() ([]ContactLink, error) {
	return getDocuments[ContactLink](contactLinkColl)
}

func GetWorkXP() ([]WorkExperience, error) {
	return getDocuments[WorkExperience](workExperienceColl)
}

func GetVolunteeringXP() ([]VolunteeringExperience, error) {
	return getDocuments[VolunteeringExperience](volunteeringExperienceColl)
}

func getDocuments[T any](coll *mongo.Collection) ([]T, error) {
	rows, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var docs []T
	for rows.Next(ctx) {
		var doc T
		err = rows.Decode(&doc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func getDocumentById[T any](coll *mongo.Collection, id string) (T, error) {
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return *new(T), err
	}
	row := coll.FindOne(ctx, bson.M{"_id": bson.M{"$eq": primitiveId}})
	if row.Err() != nil {
		return *new(T), row.Err()
	}
	var doc T
	err = row.Decode(&doc)
	if err != nil {
		return *new(T), err
	}
	return doc, nil
}

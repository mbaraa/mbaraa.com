package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteBlog(id string) error {
	return deleteDocumentById(blogColl, id)
}

func DeleteProjectGroup(id string) error {
	return deleteDocumentById(projectGroupColl, id)
}

func DeleteWorkXP(id string) error {
	return deleteDocumentById(workExperienceColl, id)
}

func DeleteVolunteeringXP(id string) error {
	return deleteDocumentById(volunteeringExperienceColl, id)
}

func deleteDocumentById(coll *mongo.Collection, id string) error {
	primitveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = coll.DeleteOne(ctx, bson.M{"_id": bson.M{"$eq": primitveId}})
	if err != nil {
		return err
	}
	return nil
}

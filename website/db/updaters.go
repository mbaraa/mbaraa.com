package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateBlog(id string, blog Blog) error {
	return updateDocumentById(blogColl, id, blog)
}

func UpdateProjectGroup(id string, pg ProjectGroup) error {
	return updateDocumentById(projectGroupColl, id, pg)
}

func UpdateInfo(info Info) error {
	row := infoColl.FindOne(ctx, bson.M{})
	if row.Err() != nil {
		_, err := infoColl.InsertOne(ctx, Info{})
		if err != nil {
			return err
		}
	}
	var _info Info
	err := row.Decode(&_info)
	if err != nil {
		return err
	}
	return updateDocumentById(infoColl, _info.Id, info)
}

func UpdateWorkXP(id string, xp WorkExperience) error {
	return updateDocumentById(workExperienceColl, id, xp)
}

func UpdateVolunteeringXP(id string, xp VolunteeringExperience) error {
	return updateDocumentById(volunteeringExperienceColl, id, xp)
}

func updateDocumentById(coll *mongo.Collection, id string, newThing any) error {
	primitveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = coll.UpdateByID(ctx, primitveId, newThing)
	if err != nil {
		return err
	}
	return nil
}

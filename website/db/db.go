package db

import (
	"context"
	"mbaraacom/config"
	"mbaraacom/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx                        context.Context
	client                     *mongo.Client
	blogColl                   *mongo.Collection
	contactLinkColl            *mongo.Collection
	projectGroupColl           *mongo.Collection
	workExperienceColl         *mongo.Collection
	volunteeringExperienceColl *mongo.Collection
	infoColl                   *mongo.Collection
)

func init() {
	ctx = context.Background()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Config().DbUri))
	if err != nil {
		log.Errorln(err.Error())
	}

	blogColl = client.Database("mbaraacom").Collection("blogs")
	contactLinkColl = client.Database("mbaraacom").Collection("contact-links")
	projectGroupColl = client.Database("mbaraacom").Collection("projects")
	workExperienceColl = client.Database("mbaraacom").Collection("work-experience")
	volunteeringExperienceColl = client.Database("mbaraacom").Collection("volunteering-experience")
	infoColl = client.Database("mbaraacom").Collection("info")
}

type Blog struct {
	Id          string `bson:"_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Content     string `bson:"content,omitempty"`
	PublicId    string `bson:"public_id,omitempty"`
	VisitTimes  uint   `bson:"visited_times,omitempty"`
	WrittenAt   int64  `bson:"written_at,omitempty"`
}

type ContactLink struct {
	Id       string `bson:"_id,omitempty"`
	Title    string `bson:"title,omitempty"`
	Link     string `bson:"link,omitempty"`
	Target   string `bson:"target,omitempty"`
	IconPath string `bson:"icon_path,omitempty"`
}

type Project struct {
	Id          string `bson:"_id,omitempty"`
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	LogoUrl     string `bson:"logo_url"`
	SourceCode  string `bson:"source_code"`
	Website     string `bson:"website"`
	StartYear   string `bson:"start_year"`
	EndYear     string `bson:"end_year"`
}

type ProjectGroup struct {
	Id          string `bson:"_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Order       uint   `bson:"order,omitempty"`
	Projects    []Project
}

type WorkExperience experience

type VolunteeringExperience experience

type experience struct {
	Id          string   `bson:"_id,omitempty"`
	Title       string   `bson:"title,omitempty"`
	Description string   `bson:"description,omitempty"`
	Location    string   `bson:"location"`
	StartDate   int64    `bson:"start_date,omitempty"`
	EndDate     int64    `bson:"end_date,omitempty"`
	Roles       []string `bson:"roles,omitempty"`
}

type Info struct {
	Id           string   `bson:"_id,omitempty"`
	FullAbout    string   `bson:"full_about,omitempty"`
	BriefAbout   string   `bson:"brief_about,omitempty"`
	BlogIntro    string   `bson:"blog_intro,omitempty"`
	Technologies []string `bson:"technologies,omitempty"`
}

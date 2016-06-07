package main

import (
	"encoding/json"
	"github.com/zebresel-com/mongodm"
	tuxmongo "github.com/zebresel-com/mongodm"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Innerstruct struct {
	Sliced []string `json:"sliced" bson:"sliced"`
}

type TestSliceModel struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	Id                   string      `json:"id" bson:"id"`
	InnerStruct          Innerstruct `json:"innerStruct" bson:"innerStruct"`
}

type TestSliceModelTux struct {
	tuxmongo.DocumentBase `json:",inline" bson:",inline"`
	Id                    string      `json:"id" bson:"id"`
	InnerStruct           Innerstruct `json:"innerStruct" bson:"innerStruct"`
}

var (
	localeEnJson = `{
    "en-US": {
        "validation.field_required": "Field '%s' is required.",
        "validation.field_invalid": "Field '%s' has an invalid value.",
        "validation.field_invalid_id": "Field '%s' contains an invalid object id value.",
        "validation.field_minlen": "Field '%s' must be at least %v characters long.",
        "validation.field_maxlen": "Field '%s' can be maximum %v characters long.",
        "validation.entry_exists": "%s already exists for value '%v'.",
        "validation.field_not_exclusive": "Only one of both fields can be set: '%s'' or '%s'.",
        "validation.field_required_exclusive": "Field '%s' or '%s' required.",
        "validation.field_invalid_relation11": "Field '%s' has wrong relation. Expected an array.",
        "validation.field_invalid_relation1n": "Field '%s' has wrong relation. No array expected."
    }
}
`
	// LocaleMap is the locale map for validating struct in mongodm.
	LocaleMap map[string]map[string]string
)

func main() {
	TestSliceTuxago()
	TestSliceZebresel()
}

func TestSliceZebresel() {
	json.Unmarshal([]byte(localeEnJson), &LocaleMap)
	dbConfig := &mongodm.Config{
		DatabaseHost: "mongo",
		DatabaseName: "test",
		Locals:       LocaleMap["en-US"],
	}

	dbConnection, err := mongodm.Connect(dbConfig)

	if err != nil {
		return
	}
	defer dbConnection.Close()
	log.Println("DB: Connected to database")
	dbConnection.Register(&TestSliceModel{}, "testslices")

	Test := dbConnection.Model("testslicemodel")

	testSlice := &TestSliceModel{}

	Test.New(testSlice)

	err = Test.FindOne(bson.M{"deleted": false}).Exec(testSlice)

	if _, ok := err.(*mongodm.NotFoundError); ok {
		log.Println("DB: FindOne failed, minimum one result was expected", err)
	} else if err != nil {
		log.Println("DB: FindOne failed", err)
	}

	testSlice.InnerStruct.Sliced = append(testSlice.InnerStruct.Sliced, []string{"Some other test", "And another one."}...)

	err = testSlice.Save()

	if err != nil {
		log.Println("DB: update on model failed, save error", err)
	}

	//rollback

	testSliceUpdate := &TestSliceModel{}

	Test.New(testSliceUpdate)
	err = Test.FindOne(bson.M{"id": "testSlice"}).Exec(testSliceUpdate)

	if err != nil {
		log.Println("DB: find with updated data failed", err)
	}

	testSliceUpdate.InnerStruct.Sliced = append(testSliceUpdate.InnerStruct.Sliced, "Some test")

	err = testSliceUpdate.Save()

	if err != nil {
		log.Println("DB: rollback on model failed, save error", err)
	}
}

func TestSliceTuxago() {
	json.Unmarshal([]byte(localeEnJson), &LocaleMap)
	dbConfig := &tuxmongo.Config{
		DatabaseHost: "mongo",
		DatabaseName: "test",
		Locals:       LocaleMap["en-US"],
	}

	dbConnection, err := tuxmongo.Connect(dbConfig)

	if err != nil {
		return
	}
	defer dbConnection.Close()
	log.Println("DB: Connected to database")
	dbConnection.Register(&TestSliceModel{}, "testslices")

	Test := dbConnection.Model("testslicemodel")

	testSlice := &TestSliceModel{}

	Test.New(testSlice)

	err = Test.FindOne(bson.M{"deleted": false}).Exec(testSlice)

	if _, ok := err.(*tuxmongo.NotFoundError); ok {
		log.Println("DB: FindOne failed, minimum one result was expected", err)
	} else if err != nil {
		log.Println("DB: FindOne failed", err)
	}

	testSlice.InnerStruct.Sliced = append(testSlice.InnerStruct.Sliced, []string{"Some other test", "And another one."}...)

	err = testSlice.Save()

	if err != nil {
		log.Println("DB: update on model failed, save error", err)
	}

	//rollback

	testSliceUpdate := &TestSliceModel{}

	Test.New(testSliceUpdate)
	err = Test.FindOne(bson.M{"id": "testSlice"}).Exec(testSliceUpdate)

	if err != nil {
		log.Println("DB: find with updated data failed", err)
	}

	testSliceUpdate.InnerStruct.Sliced = append(testSliceUpdate.InnerStruct.Sliced, "Some test")

	err = testSliceUpdate.Save()

	if err != nil {
		log.Println("DB: rollback on model failed, save error", err)
	}
}

package app

import (
	//	"fmt"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"log"
	"net/http"
	"reflect"
	"strings"
	"table8/models"
	//	"time"
	"appengine"
	"appengine/datastore"
)

const DEBUG_MODE = true
const RESTAURANT_UID = "@test-mcdonalds"

func Register(rpc_v2 *rpc.Server) {
	log.Printf("... REGISTERING...")

	rpc_v2.RegisterService(new(Region), "")
	rpc_v2.RegisterService(new(RestaurantTag), "")
	rpc_v2.RegisterService(new(RestaurantDetail), "")
	rpc_v2.RegisterService(new(System), "")

	// prints a concise summary of the exported API calls
	if DEBUG_MODE {
		list_methods := func(m interface{}) {
			typ := reflect.TypeOf(m)
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}
			fooType := reflect.TypeOf(m)
			for i := 0; i < fooType.NumMethod(); i++ {
				method := fooType.Method(i)
				args := reflect.New(method.Type.In(2).Elem()).Elem().Interface()
				resp := reflect.New(method.Type.In(3).Elem()).Elem().Interface()
				log.Printf("response = api.call('%s.%s', %+v) # response: %+v", typ.Name(), method.Name, args, resp)
			}

		}

		log.Printf("\n")
		log.Printf("EXPORTED METHOD NAMES:\n")
		list_methods(new(Region))
		list_methods(new(RestaurantDetail))
		list_methods(new(RestaurantTag))
		list_methods(new(System))
		log.Printf("\n")
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Printf("ERROR: %s %s", msg, err)
		panic(err)
	}
}

func uidlist(args ...string) string {
	return strings.Join(args, ";")
}

type GeneralArgs struct {
}

type GeneralReply struct {
	Warnings []string `json:",omitempty"`
}

/*
-------------------------------------------------
INTERNAL CALL: Some general system tests, checking/showing how errors and other conditions are handled
-------------------------------------------------
*/

type System struct{}

func (h *System) Test(r *http.Request, args *struct{ Test string }, reply *struct{}) error {
	if args.Test == "" {
		return &json2.Error{Code: json2.E_INVALID_REQ, Message: "Required parameter: Test"}
	}
	if args.Test == "fatal" {
		// fatal, programming error
		x := 0
		y := 0
		x = x / y
	}
	return nil
}

type SystemArgs struct {
	GeneralArgs
	CleanupOnly bool
}

/*
This call deletes and creates unit-testing data.
*/
func (h *System) CreateTestData(r *http.Request, args *SystemArgs, reply *GeneralReply) error {
	c := appengine.NewContext(r)
	var err error

	// delete any previous testing data
	{
		var keys_to_delete []*datastore.Key

		REGION_UIDS_TO_DELETE := []string{"@test-san-francisco", "@test-los-angeles"}
		RESTAURANT_UIDS_TO_DELETE := []string{"@test-mcdonalds", "@test-In-N-Out", "@test-Wendys"}
		TAGS_UIDS_TO_DELETE := []string{"@test-american", "@test-french"}

		for _, uid := range REGION_UIDS_TO_DELETE {
			q := datastore.NewQuery("Tags").Filter("Uid =", uid).KeysOnly()
			keys_to_delete, err = q.GetAll(c, nil)
			checkErr(err, "DB error1")
			err = datastore.DeleteMulti(c, keys_to_delete)
			checkErr(err, "DB error2")
		}
		for _, uid := range RESTAURANT_UIDS_TO_DELETE {
			q := datastore.NewQuery("Region").Filter("Uid =", uid).KeysOnly()
			keys_to_delete, err = q.GetAll(c, nil)
			checkErr(err, "DB error3")
			err = datastore.DeleteMulti(c, keys_to_delete)
			checkErr(err, "DB error4")
		}

		for _, uid := range TAGS_UIDS_TO_DELETE {
			q := datastore.NewQuery("Restaurant").Filter("Uid =", uid).KeysOnly()
			_, err = q.GetAll(c, &keys_to_delete)
			checkErr(err, "DB error5")
			err = datastore.DeleteMulti(c, keys_to_delete)
			checkErr(err, "DB error6")
		}
		log.Printf("... cleanup done")
	}

	// re-create the data (unless this was a cleanup operation only)
	if !args.CleanupOnly {
		region1 := models.Region{Uid: "@test-san-francisco", Title: "TEST San Francisco"}
		_, err = datastore.Put(c, datastore.NewKey(c, "Region", region1.Uid, 0, nil), &region1)
		checkErr(err, "fail trying to insert")
		region2 := models.Region{Uid: "@test-los-angeles", Title: "TEST Los Angeles"}
		_, err = datastore.Put(c, datastore.NewKey(c, "Region", region2.Uid, 0, nil), &region2)
		checkErr(err, "fail trying to insert")

		restaurant1 := models.Restaurant{Uid: "@test-mcdonalds", Title: "TEST McDonalds", Tags: []string{"French Cuisine", "American"}, Regions: []string{region1.Uid}}
		_, err = datastore.Put(c, datastore.NewKey(c, "Restaurant", restaurant1.Uid, 0, nil), &restaurant1)
		checkErr(err, "fail trying to insert")
		restaurant2 := models.Restaurant{Uid: "@test-In-N-Out", Tags: []string{"American"}, Regions: []string{region1.Uid, region2.Uid}}
		_, err = datastore.Put(c, datastore.NewKey(c, "Restaurant", restaurant2.Uid, 0, nil), &restaurant2)
		checkErr(err, "fail trying to insert")
		restaurant3 := models.Restaurant{Uid: "@test-Wendys", Tags: []string{"American"}, Regions: []string{region2.Uid}}
		_, err = datastore.Put(c, datastore.NewKey(c, "Restaurant", restaurant3.Uid, 0, nil), &restaurant3)
		checkErr(err, "fail trying to insert")

		log.Printf("... creation done")
	}
	return nil
}

/*
-------------------------------------------------
RestaurantTag handlers
-------------------------------------------------
*/
type RestaurantDetail struct{}

type RestaurantDetailUpdateArgs struct {
	models.Restaurant
}

type RestaurantDetailReply struct {
	Count int
	Items []models.Restaurant
}

func (h *RestaurantDetail) Retrieve(r *http.Request, args *GeneralArgs, reply *RestaurantDetailReply) error {
	c := appengine.NewContext(r)
	var err error

	restaurants := []models.Restaurant{}
	q := datastore.NewQuery("Restaurant").Filter("Uid =", RESTAURANT_UID)
	key, err := q.GetAll(c, &restaurants)
	checkErr(err, "fail trying to select")
	if key == nil {
		checkErr(nil, "fail trying to select")
	}
	reply.Count = 1
	reply.Items = restaurants
	return nil
}

func (h *RestaurantDetail) Update(r *http.Request, args *RestaurantDetailUpdateArgs, reply *RestaurantDetailReply) error {
	c := appengine.NewContext(r)
	var err error
	restaurants := []models.Restaurant{}
	q := datastore.NewQuery("Restaurant").Filter("Uid =", RESTAURANT_UID)
	restaurant_keys, err := q.GetAll(c, &restaurants)
	checkErr(err, "fail trying to select")
	if restaurant_keys == nil || restaurant_keys[0] == nil {
		checkErr(nil, "fail trying to select")
	}
	restaurant := &restaurants[0]
	if args.Title != "" {
		restaurant.Title = args.Title
	}
	if args.Details != "" {
		restaurant.Details = args.Details
	}
	if args.Tags != nil {
		restaurant.Tags = args.Tags
	}
	if args.Regions != nil {
		restaurant.Regions = args.Regions
	}
	_, err = datastore.Put(c, restaurant_keys[0], restaurant)
	checkErr(err, "fail trying to update")
	reply.Count = 1
	reply.Items = restaurants
	return nil
}

/*
-------------------------------------------------
RestaurantTag handlers
-------------------------------------------------
*/
type RestaurantTag struct {
	Tag string
}

type RestaurantTagListReply struct {
	Count int
	Items []models.RestaurantTag
}

func (h *RestaurantTag) Retrieve(r *http.Request, args *GeneralArgs, reply *RestaurantTagListReply) error {
	c := appengine.NewContext(r)

	var items []models.RestaurantTag
	q := datastore.NewQuery("Tag")
	_, err := q.GetAll(c, &items)
	if err != nil {
		return err
	}
	reply.Items = items
	return nil
}

/*
-------------------------------------------------
Region handlers
-------------------------------------------------
*/
type Region struct {
	Tag string
}

type RegionListReply struct {
	Count int
	Items []models.Region
}

func (h *Region) Retrieve(r *http.Request, args *GeneralArgs, reply *RegionListReply) error {
	c := appengine.NewContext(r)

	var items []models.Region
	q := datastore.NewQuery("Region")
	_, err := q.GetAll(c, &items)
	if err != nil {
		return err
	}
	reply.Items = items
	return nil
}

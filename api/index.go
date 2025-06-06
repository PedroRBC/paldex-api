package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	. "github.com/tbxark/g4vercel"
)

type Pal struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Wiki        string `json:"wiki"`
	Types       []struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"types"`
	ImageWiki   string `json:"imageWiki"`
	Suitability []struct {
		Type  string `json:"type"`
		Image string `json:"image"`
		Level int    `json:"level"`
	} `json:"suitability"`
	Drops       []string `json:"drops"`
	Aura        struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Tech        interface{} `json:"tech"`
	} `json:"aura"`
	Description string `json:"description"`
	Skills      []struct {
		Level       int    `json:"level"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Cooldown    int    `json:"cooldown"`
		Power       int    `json:"power"`
		Description string `json:"description"`
	} `json:"skills"`
	Stats struct {
		HP      int `json:"hp"`
		Attack  struct {
			Melee  int `json:"melee"`
			Ranged int `json:"ranged"`
		} `json:"attack"`
		Defense int `json:"defense"`
		Speed   struct {
			Ride int `json:"ride"`
			Run  int `json:"run"`
			Walk int `json:"walk"`
		} `json:"speed"`
		Stamina int `json:"stamina"`
		Support int `json:"support"`
		Food    int `json:"food"`
	} `json:"stats"`
	Asset    string `json:"asset"`
	Genus    string `json:"genus"`
	Rarity   int    `json:"rarity"`
	Price    int    `json:"price"`
	Size     string `json:"size"`
	Maps     struct {
		Day   string `json:"day"`
		Night string `json:"night"`
	} `json:"maps"`
	Breeding struct {
		Rank            int     `json:"rank"`
		Order           int     `json:"order"`
		ChildEligble    bool    `json:"child_eligble"`
		MaleProbability float64 `json:"male_probability"`
	} `json:"breeding"`
}

var pals []Pal

func init() {
	// Read and parse the JSON file
	data, err := os.ReadFile("api/pals.json")
	if err != nil {
		fmt.Printf("Error reading pals.json: %v\n", err)
		return
	}

	err = json.Unmarshal(data, &pals)
	if err != nil {
		fmt.Printf("Error parsing pals.json: %v\n", err)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	// Get all pals
	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"pals": pals,
		})
	})

	// Get pal by ID
	server.GET("/:id", func(context *Context) {
		idStr := context.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.JSON(400, H{
				"error": "Invalid ID format",
			})
			return
		}

		for _, pal := range pals {
			if pal.ID == id {
				context.JSON(200, H{
					"pal": pal,
				})
				return
			}
		}

		context.JSON(404, H{
			"error": "Pal not found",
		})
	})

	// Search pals by name
	server.GET("/search/:name", func(context *Context) {
		name := context.Param("name")
		var results []Pal

		for _, pal := range pals {
			if pal.Name == name {
				results = append(results, pal)
			}
		}

		if len(results) > 0 {
			context.JSON(200, H{
				"pals": results,
			})
		} else {
			context.JSON(404, H{
				"error": "No pals found with that name",
			})
		}
	})

	server.Handle(w, r)
}
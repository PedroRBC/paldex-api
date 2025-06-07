package handler

import (
	"net/http"
	"strconv"

	pkg "github.com/PedroRBC/paldex-api/api/_pkg"
	vercel "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := vercel.New()

	// Get all pals
	server.GET("/", func(context *vercel.Context) {
		context.JSON(200, vercel.H{
			"pals": pkg.Pals,
		})
	})

	// Get pal by ID
	server.GET("/:id", func(context *vercel.Context) {
		idStr := context.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.JSON(400, vercel.H{
				"error": "Invalid ID format",
			})
			return
		}

		for _, pal := range pkg.Pals {
			if pal.ID == id {
				context.JSON(200, vercel.H{
					"pal": pal,
				})
				return
			}
		}

		context.JSON(404, vercel.H{
			"error": "Pal not found",
		})
	})

	// Search pals by name
	server.GET("/search/:name", func(context *vercel.Context) {
		name := context.Param("name")
		var results []pkg.Pal

		for _, pal := range pkg.Pals {
			if pal.Name == name {
				results = append(results, pal)
			}
		}

		if len(results) > 0 {
			context.JSON(200, vercel.H{
				"pals": results,
			})
		} else {
			context.JSON(404, vercel.H{
				"error": "No pals found with that name",
			})
		}
	})

	server.Handle(w, r)
}
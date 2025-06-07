package handler

import (
	"net/http"
	"strconv"

	pkg "github.com/PedroRBC/paldex-api/api/_pkg"
	vercel "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := vercel.New()

	server.GET("/", func(context *vercel.Context) {
		context.JSON(200, vercel.H{
			"message": "Paldex API",
			"version": "1.0.0",
		})
	})

	// Get all pals
	server.GET("/api", func(context *vercel.Context) {
		context.JSON(200, vercel.H{
			"content": pkg.Pals,
		})
	})

	// Get pal by ID
	server.GET("/api/:id", func(context *vercel.Context) {
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
					"content": pal,
				})
				return
			}
		}

		context.JSON(404, vercel.H{
			"error": "Pal not found",
		})
	})

	// Search pals by name
	server.GET("/api/search/:name", func(context *vercel.Context) {
		name := context.Param("name")
		var results []pkg.Pal

		for _, pal := range pkg.Pals {
			if pal.Name == name {
				results = append(results, pal)
			}
		}

		if len(results) > 0 {
			context.JSON(200, vercel.H{
				"content": results,
			})
		} else {
			context.JSON(404, vercel.H{
				"error": "No pals found with that name",
			})
		}
	})

	server.Handle(w, r)
}
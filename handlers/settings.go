package handlers

import (
	"net/http"
	"strconv"

	"github.com/angelofallars/htmx-go"
	"github.com/bayadev/bayahtmx/db"
	"github.com/bayadev/bayahtmx/models"
	"github.com/bayadev/bayahtmx/templates"
	"github.com/bayadev/bayahtmx/templates/components"
	"github.com/bayadev/bayahtmx/templates/pages"
	"github.com/gin-gonic/gin"
)

func GetSettings(c *gin.Context) {

	settings := models.NewSettingRepository(db.DB)

	config, store, social, err := settings.GetAllForSettingsPage()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}


	// Define template meta tags.
	metaTags := pages.MetaTags(
		"gowebly, htmx example page, go with htmx",               // define meta keywords
		"Welcome to example! You're here because it worked out.", // define meta description
	)

	// Define template body content.
	bodyContent := pages.Settings(config, store, social)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Welcome to example!", // define title text
		metaTags,              // define meta tags
		bodyContent,           // define body content
	)

	// Render index page template.
	if err := htmx.NewResponse().RenderTempl(c.Request.Context(), c.Writer, indexTemplate); err != nil {
		// If not, return HTTP 500 error.
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}


func GetQuotations(c *gin.Context) {



	page := c.Query("page")
	var pageInt = 1
	if page != "" {
	pageInt,_ = strconv.Atoi(page)
	}



	// Define template body content.
	bodyContent := components.TableRows(pageInt)


	// Render index page template.
	if err := htmx.NewResponse().RenderTempl(c.Request.Context(), c.Writer, bodyContent); err != nil {
		// If not, return HTTP 500 error.
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}


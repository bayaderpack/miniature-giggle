package main

import (
	// "context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	// "github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
	// "github.com/gin-gonic/gin/render"

	"github.com/gin-gonic/gin"

	gowebly "github.com/gowebly/helpers"
	"github.com/bayadev/bayahtmx/handlers"
	"github.com/bayadev/bayahtmx/utils"
)



// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "8000"))
	if err != nil {
		return err
	}

	// Create a new Fiber server.
	router := gin.Default()

	// Define HTML renderer for template engine.
	router.HTMLRender = &utils.TemplRender{}

	// Handle static files.
	router.Static("/static", "./static")

	// Handle index page view.
	router.GET("/", handlers.IndexViewHandler)

	router.GET("/settings", handlers.GetSettings)

	router.GET("/table", handlers.TableViewHandler)
	router.GET("/quotation", handlers.GetQuotations)

	router.GET("/tab1", tab1)
	router.GET("/tab2", tab2)
	router.GET("/tab3", tab3)

	// Handle API endpoints.
	router.GET("/api/hello-world", handlers.ShowContentAPIHandler)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}

func tab1(c *gin.Context) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request) {
		// If not, return HTTP 400 error.
		c.AbortWithError(http.StatusBadRequest, errors.New("non-htmx request"))
		return
	}

	// Write HTML content.
	c.Writer.Write([]byte("<p>🎉 Yes</p>"))

	// Send htmx response.
	htmx.NewResponse().Write(c.Writer)
}


func tab2(c *gin.Context) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request) {
		// If not, return HTTP 400 error.
		c.AbortWithError(http.StatusBadRequest, errors.New("non-htmx request"))
		return
	}

	// Write HTML content.
	c.Writer.Write([]byte("<p>🎉 test 23</p>"))

	// Send htmx response.
	htmx.NewResponse().Write(c.Writer)
}

func tab3(c *gin.Context) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request) {
		// If not, return HTTP 400 error.
		c.AbortWithError(http.StatusBadRequest, errors.New("non-htmx request"))
		return
	}

	// Write HTML content.
	c.Writer.Write([]byte("<p>🎉 test 23532</p>"))

	// Send htmx response.
	htmx.NewResponse().Write(c.Writer)
}

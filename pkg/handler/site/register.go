package site

import (
	"github.com/aandrku/portfolio-v2/pkg/handler/blog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(e *echo.Echo) {
	// index page
	e.GET("/", getIndex)

	e.GET("/login", getLogin)
	e.POST("/login", postLogin)

	// home window
	e.GET("/home/window", getHomeWindow)

	// about window
	e.GET("/about/window", getAboutWindow)

	// projects window
	e.GET("/projects/window", getProjectsWindow)
	e.GET("/projects/:id", getProject)

	// contact window
	e.GET("/contact/window", getContactWindow)

	e.POST("/contact", postContact, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1)))

	// delete endpoint, used to delete elements from a webpage using HTMX
	e.GET("/delete", getDelete)

	// blog group
	blog.Register(e.Group("/blog"))

}

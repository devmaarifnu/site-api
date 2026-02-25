package routes

import (
	"lpmaarifnu-site-api/internal/config"
	"lpmaarifnu-site-api/internal/database"
	"lpmaarifnu-site-api/internal/handlers"
	"lpmaarifnu-site-api/internal/middleware"
	"lpmaarifnu-site-api/internal/repositories"
	"lpmaarifnu-site-api/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Initialize repositories
	newsRepo := repositories.NewNewsRepository(database.GetDB())
	opinionRepo := repositories.NewOpinionRepository(database.GetDB())
	documentRepo := repositories.NewDocumentRepository(database.GetDB())
	heroRepo := repositories.NewHeroRepository(database.GetDB())
	orgRepo := repositories.NewOrganizationRepository(database.GetDB())
	pageRepo := repositories.NewPageRepository(database.GetDB())
	categoryRepo := repositories.NewCategoryRepository(database.GetDB())
	// tagRepo := repositories.NewTagRepository(database.GetDB()) // Not used yet
	settingRepo := repositories.NewSettingRepository(database.GetDB())
	pengurusRepo := repositories.NewPengurusRepository(database.GetDB())
	editorialRepo := repositories.NewEditorialRepository(database.GetDB())
	contactRepo := repositories.NewContactRepository(database.GetDB())
	eventRepo := repositories.NewEventRepository(database.GetDB())

	// Initialize services
	newsService := services.NewNewsService(newsRepo)
	opinionService := services.NewOpinionService(opinionRepo)
	documentService := services.NewDocumentService(documentRepo)
	heroService := services.NewHeroService(heroRepo)
	orgService := services.NewOrganizationService(orgRepo)
	pageService := services.NewPageService(pageRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	settingService := services.NewSettingService(settingRepo)
	pengurusService := services.NewPengurusService(pengurusRepo)
	editorialService := services.NewEditorialService(editorialRepo)
	contactService := services.NewContactService(contactRepo)
	eventService := services.NewEventService(eventRepo)

	// Initialize handlers
	newsHandler := handlers.NewNewsHandler(newsService)
	opinionHandler := handlers.NewOpinionHandler(opinionService)
	documentHandler := handlers.NewDocumentHandler(documentService)
	heroHandler := handlers.NewHeroHandler(heroService)
	orgHandler := handlers.NewOrganizationHandler(orgService)
	pageHandler := handlers.NewPageHandler(pageService)
	categoryHandler := handlers.NewCategoryHandler(categoryService, newsService)
	settingHandler := handlers.NewSettingHandler(settingService)
	pengurusHandler := handlers.NewPengurusHandler(pengurusService)
	editorialHandler := handlers.NewEditorialHandler(editorialService)
	contactHandler := handlers.NewContactHandler(contactService)
	eventHandler := handlers.NewEventHandler(eventService)

	// Apply global middleware
	router.Use(middleware.Logger())
	// router.Use(middleware.CORS(&cfg.CORS)) // Disabled: CORS handled by Nginx
	router.Use(middleware.RateLimit(&cfg.RateLimit))

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// News routes
		news := v1.Group("/news")
		{
			news.GET("", newsHandler.GetAllNews)
			news.GET("/featured", newsHandler.GetFeaturedNews)
			news.GET("/:slug", newsHandler.GetNewsBySlug)
		}

		// Opinion routes
		opinions := v1.Group("/opinions")
		{
			opinions.GET("", opinionHandler.GetAllOpinions)
			opinions.GET("/:slug", opinionHandler.GetOpinionBySlug)
		}

		// Document routes
		documents := v1.Group("/documents")
		{
			documents.GET("", documentHandler.GetAllDocuments)
			documents.GET("/:id", documentHandler.GetDocumentByID)
		}

		// Hero slides routes
		v1.GET("/hero-slides", heroHandler.GetActiveSlides)

		// Organization routes
		organization := v1.Group("/organization")
		{
			organization.GET("/structure", orgHandler.GetOrganizationStructure)
			organization.GET("/board-members", orgHandler.GetBoardMembers)
			organization.GET("/pengurus", pengurusHandler.GetAllPengurus)
		}

		// Page routes
		v1.GET("/pages/:slug", pageHandler.GetPageBySlug)

		// Category routes
		categories := v1.Group("/categories")
		{
			categories.GET("", categoryHandler.GetAllCategories)
			categories.GET("/:slug", categoryHandler.GetCategoryBySlug)
		}

		// Settings route
		v1.GET("/settings", settingHandler.GetPublicSettings)

		// Editorial routes
		editorial := v1.Group("/editorial")
		{
			editorial.GET("/team", editorialHandler.GetEditorialTeam)
		}

		// Contact routes
		contact := v1.Group("/contact")
		{
			contact.POST("/submit", contactHandler.SubmitContact)
		}

		// Events routes
		events := v1.Group("/events")
		{
			events.GET("/flayers", eventHandler.GetAllEvents)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "LP Maarif API is running",
		})
	})
}

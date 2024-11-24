package handler

import (
	"github.com/DavidG9999/DMS/DMS_api_gateway/internal/service"

	_ "github.com/DavidG9999/DMS/DMS_api_gateway/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", h.getUser)
			user.PUT("/name", h.updateName)
			user.PUT("/password", h.updatePassword)
			user.DELETE("/", h.deleteUser)
		}
		autos := api.Group("/autos")
		{
			autos.POST("/", h.createAuto)
			autos.GET("/", h.getAutos)
			autos.PUT("/:id", h.updateAuto)
			autos.DELETE("/:id", h.deleteAuto)
		}
		contragents := api.Group("/contragents")
		{
			contragents.POST("/", h.createContragent)
			contragents.GET("/", h.getContragents)
			contragents.PUT("/:id", h.updateContragent)
			contragents.DELETE("/:id", h.deleteContragent)

		}
		dispetchers := api.Group("/dispetchers")
		{
			dispetchers.POST("/", h.createDispetcher)
			dispetchers.GET("/", h.getDispetchers)
			dispetchers.PUT("/:id", h.updateDispetcher)
			dispetchers.DELETE("/:id", h.deleteDispetcher)
		}
		drivers := api.Group("/drivers")
		{
			drivers.POST("/", h.createDriver)
			drivers.GET("/", h.getDrivers)
			drivers.PUT("/:id", h.updateDriver)
			drivers.DELETE("/:id", h.deleteDriver)
		}
		mehanics := api.Group("/mehanics")
		{
			mehanics.POST("/", h.createMehanic)
			mehanics.GET("/", h.getMehanics)
			mehanics.PUT("/:id", h.updateMehanic)
			mehanics.DELETE("/:id", h.deleteMehanic)
		}
		organiations := api.Group("/organizations")
		{
			organiations.POST("/", h.createOrganization)
			organiations.GET("/", h.getOrganizations)
			organiations.PUT("/:id", h.updateOrganization)
			organiations.DELETE("/:id", h.deleteOrganization)

			accounts := organiations.Group(":id/bank_accounts")
			{
				accounts.POST("/", h.createBankAccount)
				accounts.GET("/", h.getBankAccounts)
				accounts.PUT("/:bank_account_id", h.updateBankAccount)
				accounts.DELETE("/:bank_account_id", h.deleteBankAccount)
			}
		}
		putlists := api.Group("/putlists")
		{
			putlists.POST("/", h.createPutlist)
			putlists.GET("/", h.getPutlists)
			putlists.GET("/:number", h.getPutlistByNumber)
			putlists.PUT("/:number", h.updatePutlist)
			putlists.DELETE("/:number", h.deletePutlist)

			putlist_bodies := putlists.Group(":number/putlist_bodies")
			{
				putlist_bodies.POST("/", h.createPutlistBody)
				putlist_bodies.GET("/", h.getPutlistBodies)
				putlist_bodies.PUT("/:putlist_body_id", h.updatePutlistBody)
				putlist_bodies.DELETE("/:putlist_body_id", h.deletePutlistBody)

			}
		}
	}
	return router
}

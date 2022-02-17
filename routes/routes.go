package routes

import (
	"github.com/Rafipratama22/mnc_test.git/controller"
	"github.com/Rafipratama22/mnc_test.git/middleware"
	"github.com/Rafipratama22/mnc_test.git/repository"

	"github.com/Rafipratama22/mnc_test.git/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	_ "github.com/Rafipratama22/mnc_test.git/docs"
)

type Server struct {
	router *gin.Engine
}

var (
	db *gorm.DB = config.SetUpDatabase()
	userRepository repository.UserRepository = repository.NewUserRepo(db)
	authMiddleware middleware.AuthMiddleware = middleware.NewAuthMiddleware(db)
	userController controller.UserController = controller.NewUserController(userRepository)
	postRepository repository.PostRepository = repository.NewPostRepo(db)
	postController controller.PostController = controller.NewPostController(postRepository)
	companyRepository repository.CompanyRepository = repository.NewCompanyRepo(db)
	companyController controller.CompanyController = controller.NewCompanyController(companyRepository)
	loginRepository repository.LoginRepository = repository.NewLoginRepo(db)
	loginController controller.LoginController = controller.NewLoginController(loginRepository)
)

func MainSever() *Server {
	return &Server{
		router: gin.New(),
	}
}

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authurization
func (server *Server) Start() *gin.Engine {
	// Gin instance
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	apiName := "/api/v1/"
	route.GET(apiName + "ping", authMiddleware.ValidateTokenUser, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})	
	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	// Login Routes
	loginRoute := route.Group("form")
	{
		loginRoute.POST("/register/user", func(ctx *gin.Context) {
			loginController.RegisterUser(ctx)
		})
		loginRoute.POST("/register/company", func(ctx *gin.Context) {
			loginController.RegisterCompany(ctx)
		})
		loginRoute.POST("/login", func(ctx *gin.Context) {
			loginController.Login(ctx)
		})
	}

	// User Routes
	userRoute := route.Group("user")
	userRoute.Use(authMiddleware.ValidateTokenUser)
	{
		userRoute.GET("/point", func(ctx *gin.Context) {
			userController.GetAllPoint(ctx)
		})		
	}

	// Post Routes
	postRoute := route.Group("post")
	postRoute.Use(authMiddleware.ValidateTokenUser)
	{
		postRoute.POST("/", func(ctx *gin.Context) {
			postController.CreatePost(ctx)
		})
		postRoute.GET("/:id", func(ctx *gin.Context) {
			postController.DetailPost(ctx)
		})
		postRoute.PUT("/:id", func(ctx *gin.Context) {
			postController.UpdatePost(ctx)
		})
	}

	// Company Routes
	companyRoute := route.Group("company")
	companyRoute.Use(authMiddleware.ValidateTokenCompany)
	{
		companyRoute.GET("/user/register", func(ctx *gin.Context) {
			companyController.AllUserRegister(ctx)
		})
		companyRoute.GET("/user/login", func(ctx *gin.Context) {
			companyController.AllUserLogin(ctx)
		})
		companyRoute.GET("/user/:id", func(ctx *gin.Context) {
			companyController.DetailUser(ctx)
		})
		companyRoute.GET("/user/point", func(c *gin.Context) {
			companyController.AllUserPoint(c)
		})
		companyRoute.GET("/user/point/:id", func(ctx *gin.Context) {
			companyController.DetailUserPoint(ctx)
		})
		companyRoute.PATCH("/user/:id", func(ctx *gin.Context) {
			companyController.PostPoint(ctx)
		})
		companyRoute.GET("/post", func(ctx *gin.Context) {
			companyController.AllPost(ctx)
		})
		companyRoute.GET("/post/:id", func(ctx *gin.Context) {
			companyController.DetailPost(ctx)
		})
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return route
}

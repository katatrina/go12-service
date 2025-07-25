package sharedinfras

import (
	"context"
	"io"
	
	"github.com/katatrina/go12-service/middleware"
	sharecomponent "github.com/katatrina/go12-service/shared/component"
	"github.com/katatrina/go12-service/shared/datatype"
	sharedrpc "github.com/katatrina/go12-service/shared/infras/rpc"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IMiddlewareProvider interface {
	Auth() gin.HandlerFunc
	CheckRoles(roles ...datatype.UserRole) gin.HandlerFunc
}

type IDbContext interface {
	GetMainConnection() *gorm.DB
}

type IUploader interface {
	SaveFileUpload(ctx context.Context, ioReader io.Reader, dst, contentType string, length int64) error
	GetDomain() string
}

type IAppContext interface {
	MiddlewareProvider() IMiddlewareProvider
	DbContext() IDbContext
	Uploader() IUploader
	GetConfig() *datatype.Config
}

type appContext struct {
	mldProvider IMiddlewareProvider
	dbContext   IDbContext
	uploader    IUploader
	config      *datatype.Config
}

func NewAppContext(db *gorm.DB) IAppContext {
	dbCtx := NewDbContext(db)
	
	config := datatype.NewConfig()
	introspectRpcClient := sharedrpc.NewIntrospectRPCClient(config.UserServiceURL)
	
	provider := middleware.NewMiddlewareProvider(introspectRpcClient)
	
	uploader, err := sharecomponent.NewS3Uploader(config.AWS.AccessKey, config.AWS.BucketName, config.AWS.Domain, config.AWS.Region, config.AWS.SecretKey)
	
	if err != nil {
		panic(err)
	}
	
	return &appContext{
		mldProvider: provider,
		dbContext:   dbCtx,
		uploader:    uploader,
		config:      config,
	}
}

func (c *appContext) MiddlewareProvider() IMiddlewareProvider {
	return c.mldProvider
}

func (c *appContext) DbContext() IDbContext {
	return c.dbContext
}

func (c *appContext) GetConfig() *datatype.Config {
	return c.config
}

func (c *appContext) Uploader() IUploader {
	return c.uploader
}

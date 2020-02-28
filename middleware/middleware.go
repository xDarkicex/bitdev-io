package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {

}

type Config struct {
	ContentSecurityPolicy string `yaml:"content_security_policy"`
	StrictTransfer        string
	// default middleware.DefaultSkipper

	Skipper            middleware.Skipper
	ContentTypeNosniff string `yaml:"content_type_nosniff"`
	XSSProtection      string `yaml:"xss_protection"`
}

// SecureDefaultConfig is the default SecureHeader middleware config
var SecureDefaultConfig = Config{
	Skipper:               middleware.DefaultSkipper,
	ContentTypeNosniff:    "nosniff",
	XSSProtection:         "1; mode=block",
	ContentSecurityPolicy: "block-all-mixed-content",
	StrictTransfer:        "max-age=31536000; includeSubDomains",
}

// Secure header middleware did I?
func SecureHeader() echo.MiddlewareFunc {
	return secureHeaderWithConfig(SecureDefaultConfig)
}

func secureHeaderWithConfig(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = SecureDefaultConfig.Skipper
	}
	if config.ContentTypeNosniff == "" {
		config.ContentTypeNosniff = SecureDefaultConfig.ContentTypeNosniff
	}
	if config.XSSProtection == "" {
		config.XSSProtection = SecureDefaultConfig.XSSProtection
	}
	if config.ContentSecurityPolicy == "" {
		config.ContentSecurityPolicy = SecureDefaultConfig.ContentSecurityPolicy
	}
	if config.StrictTransfer == "" {
		config.StrictTransfer = SecureDefaultConfig.StrictTransfer
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			c.Response().Header().Set(echo.HeaderStrictTransportSecurity, config.StrictTransfer)
			c.Response().Header().Set(echo.HeaderContentSecurityPolicy, config.ContentSecurityPolicy)
			c.Response().Header().Set(echo.HeaderXXSSProtection, config.XSSProtection)
			c.Response().Header().Set(echo.HeaderXContentTypeOptions, config.ContentTypeNosniff)
			return next(c)
		}
	}
}

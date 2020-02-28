package middleware

func init() {


}

type Config struct {
		// Skipper defines a function to skip middleware.
		// default middleware.DefaultSkipper
		Skipper middleware.Skipper

		// WhiteList is an allowed ip list.
		WhiteList []string

		// BlackList is a disallowed ip list.
		BlackList []string

		// Block by default.
		BlockByDefault bool
}

// SecureDefaultConfig is the default SecureHeader middleware config
var SecureDefaultConfig = Config{
		Skipper:        middleware.DefaultSkipper,
		BlockByDefault: false,
}

func SecureHeader() echo.MiddlewareFunc {
		return secureHeaderWithConfig(SecureDefaultConfig)
}

func secureHeaderWithConfig(config Config) echo.MiddlewareFunc {
		var err error
		if config.Skipper == nil {
				config.Skipper = DefaultConfig.Skipper
		}
		return func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
						if config.Skipper(c) {
								return next(c)
						}

						return next(c)
				}
		}
}

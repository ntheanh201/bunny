package middleware

import (
	"github.com/labstack/echo/v4"
	"org.idev.bunny/backend/api/enum"
	"org.idev.bunny/backend/app"
	"org.idev.bunny/backend/common/logger"
	tokenutil "org.idev.bunny/backend/common/util/token"
	userrepository "org.idev.bunny/backend/repository/user"
)

func AuthGuard(appCtx *app.AppContext, userRepo *userrepository.UserRepository) echo.MiddlewareFunc {

	log := logger.New("Middleware", "AuthGuard")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if ctx.Request().Header.Get("Authorization") == "" && appCtx.Config.Env == enum.Dev {
				return next(ctx)
			}
			tokenRaw := tokenutil.GetTokenFromHeader(*ctx.Request())

			isValid, err := tokenutil.Verify(ctx.Request().Context(), tokenRaw, appCtx.Config.JWKsUrl)
			if err != nil || !isValid {
				log.Errorf("Error verify token: %v", err)
				return echo.NewHTTPError(401, err.Error())
			}
			token, err := tokenutil.Parse(ctx.Request().Context(), tokenRaw, appCtx.Config.JWKsUrl)
			if err != nil {
				log.Errorf("Error verify token: %v", err)
				return echo.NewHTTPError(401, err.Error())
			}
			userId := token.Subject()
			user, err := userRepo.FindByUserId(ctx.Request().Context(), userId)
			if err != nil {
				log.Errorf("Error verify token: %v", err)
				return echo.NewHTTPError(401, err.Error())
			}
			ctx.Set("user", user)
			return next(ctx)
		}
	}
}
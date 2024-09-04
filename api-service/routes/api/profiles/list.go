package profiles

import (
	"net/http"

	"yabro.io/social-api/app"
	"yabro.io/social-api/apperror"

	"github.com/gin-gonic/gin"
)

type ListProfilesRequest struct {
	Cursor string `form:"cursor"`
	Limit  int    `form:"limit,default=10"`
}

type ListProfilesResponse struct {
	Profiles   []PublicProfile `json:"profiles"`
	NextCursor string          `json:"next_cursor,omitempty"`
}

func ListProfiles(appState *app.AppState) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ListProfilesRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		if req.Limit <= 0 || req.Limit > 100 {
			req.Limit = 10 // Default to 10 if invalid or not provided
		}

		dbProfiles, nextCursor, err := appState.Stores.Profile.List(req.Cursor, req.Limit)
		if err != nil {
			c.Error(apperror.ToAppError(err))
			return
		}

		publicProfiles := make([]PublicProfile, len(dbProfiles))
		for i, profile := range dbProfiles {
			publicProfiles[i] = FromEntity(&profile)
		}

		response := ListProfilesResponse{
			Profiles:   publicProfiles,
			NextCursor: nextCursor,
		}

		c.JSON(http.StatusOK, response)
	}
}

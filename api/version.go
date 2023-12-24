// This file is part of the happyDomain (R) project.
// Copyright (c) 2020-2024 happyDomain
// Authors: Pierre-Olivier Mercier, et al.
//
// This program is offered under a commercial and under the AGPL license.
// For commercial licensing, contact us at <contact@happydomain.org>.
//
// For AGPL licensing:
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	HDVersion Version
)

func DeclareVersionRoutes(router *gin.RouterGroup) {
	router.GET("/version", showVersion)
}

type Version struct {
	Version string `json:"version"`
}

// showVersion returns the current happyDomain version.
//
//	@Summary	Get happyDomain version
//	@Schemes
//	@Description	Retrieve the current happyDomain version.
//	@Tags			version
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Version
//	@Router			/version [get]
func showVersion(c *gin.Context) {
	c.JSON(http.StatusOK, HDVersion)
}

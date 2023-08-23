// Copyright or © or Copr. happyDNS (2021)
//
// contact@happydomain.org
//
// This software is a computer program whose purpose is to provide a modern
// interface to interact with DNS systems.
//
// This software is governed by the CeCILL license under French law and abiding
// by the rules of distribution of free software.  You can use, modify and/or
// redistribute the software under the terms of the CeCILL license as
// circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and rights to copy, modify
// and redistribute granted by the license, users are provided only with a
// limited warranty and the software's author, the holder of the economic
// rights, and the successive licensors have only limited liability.
//
// In this respect, the user's attention is drawn to the risks associated with
// loading, using, modifying and/or developing or reproducing the software by
// the user in light of its specific status of free software, that may mean
// that it is complicated to manipulate, and that also therefore means that it
// is reserved for developers and experienced professionals having in-depth
// computer knowledge. Users are therefore encouraged to load and test the
// software's suitability as regards their requirements in conditions enabling
// the security of their systems and/or data to be ensured and, more generally,
// to use and operate it in the same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"git.happydns.org/happydomain/admin"
	"git.happydns.org/happydomain/config"
)

type Admin struct {
	router *gin.Engine
	cfg    *config.Options
	srv    *http.Server
}

func NewAdmin(cfg *config.Options) Admin {
	if cfg.DevProxy == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	admin.DeclareRoutes(cfg, router)

	app := Admin{
		router: router,
		cfg:    cfg,
	}

	return app
}

func (app *Admin) Start() {
	app.srv = &http.Server{
		Addr:    app.cfg.AdminBind,
		Handler: app.router,
	}

	log.Printf("Admin interface listening on %s\n", app.cfg.AdminBind)
	if !strings.Contains(app.cfg.AdminBind, ":") {
		if _, err := os.Stat(app.cfg.AdminBind); !os.IsNotExist(err) {
			if err := os.Remove(app.cfg.AdminBind); err != nil {
				log.Fatal(err)
			}
		}

		unixListener, err := net.Listen("unix", app.cfg.AdminBind)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(app.srv.Serve(unixListener))
	} else if err := app.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("admin listen: %s\n", err)
	}
}
func (app *Admin) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.srv.Shutdown(ctx); err != nil {
		log.Fatal("Admin Server Shutdown:", err)
	}
}

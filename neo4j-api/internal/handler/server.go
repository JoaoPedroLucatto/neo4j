package handler

import (
	"context"
	"io"
	"neo4j-api/internal/db"
	"neo4j-api/internal/usecase"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type API struct {
	Router *gin.Engine
}

type Server struct {
	Usecase *usecase.Usecase
	Log     *zerolog.Logger
}

func NewServer(ctx context.Context, log *zerolog.Logger, repository db.Repository) *Server {
	return &Server{
		Usecase: usecase.NewUsecaseService(context.Background(), log, repository),
		Log:     log,
	}
}

func (s *Server) Server() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	r.GET("/", Home)

	r.GET("/vaccinations", GetCovidSummary(s))
	r.GET("/vaccinations/first_doses", GetFirstDoseCount(s))

	r.GET("/vaccines", GetVaccinesUsedByCountry(s))
	r.GET("/vaccines/authorization", GetVaccineApprovalDates(s))
	r.GET("/vaccines/:vaccineName/countries", GetCountriesByVaccine(s))

	return r
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "neo4j"})
}

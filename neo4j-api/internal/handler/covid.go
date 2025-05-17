package handler

import (
	"neo4j-api/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCovidSummary(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		s.Log.Info().Msg("handler get covid summary")
		country := ctx.Request.URL.Query().Get("country")
		dateStr := ctx.Request.URL.Query().Get("date")

		date, err := util.ParseDate(dateStr)
		if err != nil {
			s.Log.Err(err).Msg("erro convert string to date")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		stats, err := s.Usecase.GetCovidSummary(country, date)
		if err != nil {
			s.Log.Err(err).Msg("erro handler get covid summary")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, stats)
	})
}

func GetFirstDoseCount(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		s.Log.Info().Msg("handler get first dosage count")
		country := ctx.Request.URL.Query().Get("country")
		dateStr := ctx.Request.URL.Query().Get("date")

		date, err := util.ParseDate(dateStr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		count, err := s.Usecase.GetFirstDoseCount(country, date)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, count)
	})
}

func GetVaccinesUsedByCountry(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		s.Log.Info().Msg("handler get vaccines used by country")
		country := ctx.Request.URL.Query().Get("country")

		vaccines, err := s.Usecase.GetVaccinesUsedByCountry(country)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{"country": country, "vaccines": vaccines})
	})
}

func GetVaccineApprovalDates(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		s.Log.Info().Msg("handler get vaccines approval dates")
		country := ctx.Request.URL.Query().Get("country")

		approvals, err := s.Usecase.GetVaccineApprovalDates(country)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, approvals)
	})
}

func GetCountriesByVaccine(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		s.Log.Info().Msg("handler get contries by vaccine")
		vaccineName := ctx.Param("vaccineName")

		s.Log.Info().Msg(vaccineName)

		vaccineNameByCountry, err := s.Usecase.GetCountriesByVaccine(vaccineName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, vaccineNameByCountry)
	})
}

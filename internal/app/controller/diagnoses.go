package controller

import (
	"TopDoctorsBackendChallenge/internal/models"
	"TopDoctorsBackendChallenge/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *handler) searchDiagnoses(ctx *gin.Context) {
	patientName := ctx.Query("patientName")
	dateParam := ctx.Query("date")
	var date time.Time
	if len(dateParam) != 0 {
		var err error
		date, err = time.ParseInLocation(time.DateOnly, dateParam, time.Local)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "could not parse date",
			})
			return
		}
	}

	filters := models.SearchDiagnosesFilters{
		PatientName: patientName,
		Date:        date,
	}

	diagnoses, err := h.interactor.SearchDiagnoses(ctx, filters)
	if err != nil {
		logger.CtxErrorf(ctx, "error searching for diagnoses: %v", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	diagnosesDTO := mapToDiagnosesDTO(diagnoses)

	ctx.JSON(http.StatusOK, diagnosesDTO)
}

type CreateDiagnosisRequest struct {
	PatientID    string  `json:"patient_id"`
	Description  string  `json:"description"`
	Prescription *string `json:"prescription"`
}

func (h *handler) createDiagnosis(ctx *gin.Context) {
	var body CreateDiagnosisRequest

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "could not bind body",
		})
		return
	}

	diagnosis, err := h.interactor.CreateDiagnosis(ctx, body.PatientID, body.Description, body.Prescription)
	if err != nil {
		logger.CtxErrorf(ctx, "could not create diagnosis for patient %s: %v", body.PatientID, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	diagnosisDTO := mapToDiagnosisDTO(diagnosis)

	ctx.JSON(http.StatusOK, diagnosisDTO)
}

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Store *DataStore
}

func (h *Handlers) getID(c *gin.Context) string {
	id := c.Param("id")
	return strings.TrimSuffix(id, ".json")
}

// GetProvinces godoc
// @Summary Get all provinces
// @Description Get list of all provinces in Indonesia
// @Tags provinces
// @Accept  json
// @Produce  json
// @Success 200 {array} Province
// @Router /provinces [get]
// @Router /provinces.json [get]
func (h *Handlers) GetProvinces(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.Provinces)
}

// GetRegencies godoc
// @Summary Get regencies by province ID
// @Description Get list of regencies in a specific province
// @Tags regencies
// @Accept  json
// @Produce  json
// @Param id path string true "Province ID"
// @Success 200 {array} Regency
// @Router /regencies/{id} [get]
func (h *Handlers) GetRegencies(c *gin.Context) {
	id := h.getID(c)
	regencies, ok := h.Store.Regencies[id]
	if !ok {
		c.JSON(http.StatusOK, []Regency{})
		return
	}
	c.JSON(http.StatusOK, regencies)
}

// GetDistricts godoc
// @Summary Get districts by regency ID
// @Description Get list of districts in a specific regency
// @Tags districts
// @Accept  json
// @Produce  json
// @Param id path string true "Regency ID"
// @Success 200 {array} District
// @Router /districts/{id} [get]
func (h *Handlers) GetDistricts(c *gin.Context) {
	id := h.getID(c)
	districts, ok := h.Store.Districts[id]
	if !ok {
		c.JSON(http.StatusOK, []District{})
		return
	}
	c.JSON(http.StatusOK, districts)
}

// GetVillages godoc
// @Summary Get villages by district ID
// @Description Get list of villages in a specific district
// @Tags villages
// @Accept  json
// @Produce  json
// @Param id path string true "District ID"
// @Success 200 {array} Village
// @Router /villages/{id} [get]
func (h *Handlers) GetVillages(c *gin.Context) {
	id := h.getID(c)
	villages, ok := h.Store.Villages[id]
	if !ok {
		c.JSON(http.StatusOK, []Village{})
		return
	}
	c.JSON(http.StatusOK, villages)
}

// GetProvince godoc
// @Summary Get province details by ID
// @Description Get specific province information
// @Tags provinces
// @Accept  json
// @Produce  json
// @Param id path string true "Province ID"
// @Success 200 {object} Province
// @Router /province/{id} [get]
func (h *Handlers) GetProvince(c *gin.Context) {
	id := h.getID(c)
	province, ok := h.Store.ProvincesMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, province)
}

// GetRegency godoc
// @Summary Get regency details by ID
// @Description Get specific regency information
// @Tags regencies
// @Accept  json
// @Produce  json
// @Param id path string true "Regency ID"
// @Success 200 {object} Regency
// @Router /regency/{id} [get]
func (h *Handlers) GetRegency(c *gin.Context) {
	id := h.getID(c)
	regency, ok := h.Store.RegenciesMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, regency)
}

// GetDistrict godoc
// @Summary Get district details by ID
// @Description Get specific district information
// @Tags districts
// @Accept  json
// @Produce  json
// @Param id path string true "District ID"
// @Success 200 {object} District
// @Router /district/{id} [get]
func (h *Handlers) GetDistrict(c *gin.Context) {
	id := h.getID(c)
	district, ok := h.Store.DistrictsMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, district)
}

// GetVillage godoc
// @Summary Get village details by ID
// @Description Get specific village information
// @Tags villages
// @Accept  json
// @Produce  json
// @Param id path string true "Village ID"
// @Success 200 {object} Village
// @Router /village/{id} [get]
func (h *Handlers) GetVillage(c *gin.Context) {
	id := h.getID(c)
	village, ok := h.Store.VillagesMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, village)
}

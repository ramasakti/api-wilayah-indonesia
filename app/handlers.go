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

func (h *Handlers) GetProvinces(c *gin.Context) {
	c.JSON(http.StatusOK, h.Store.Provinces)
}

func (h *Handlers) GetRegencies(c *gin.Context) {
	id := h.getID(c)
	regencies, ok := h.Store.Regencies[id]
	if !ok {
		c.JSON(http.StatusOK, []Regency{})
		return
	}
	c.JSON(http.StatusOK, regencies)
}

func (h *Handlers) GetDistricts(c *gin.Context) {
	id := h.getID(c)
	districts, ok := h.Store.Districts[id]
	if !ok {
		c.JSON(http.StatusOK, []District{})
		return
	}
	c.JSON(http.StatusOK, districts)
}

func (h *Handlers) GetVillages(c *gin.Context) {
	id := h.getID(c)
	villages, ok := h.Store.Villages[id]
	if !ok {
		c.JSON(http.StatusOK, []Village{})
		return
	}
	c.JSON(http.StatusOK, villages)
}

func (h *Handlers) GetProvince(c *gin.Context) {
	id := h.getID(c)
	province, ok := h.Store.ProvincesMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, province)
}

func (h *Handlers) GetRegency(c *gin.Context) {
	id := h.getID(c)
	regency, ok := h.Store.RegenciesMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, regency)
}

func (h *Handlers) GetDistrict(c *gin.Context) {
	id := h.getID(c)
	district, ok := h.Store.DistrictsMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, district)
}

func (h *Handlers) GetVillage(c *gin.Context) {
	id := h.getID(c)
	village, ok := h.Store.VillagesMap[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, village)
}

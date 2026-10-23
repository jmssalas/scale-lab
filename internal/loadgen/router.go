package loadgen

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Initialize() (*gin.Engine, *Generator) {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	g := CreateGenerator(CreateGeneratorMetrics(reg))
	r := gin.Default()

	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))
	r.PUT("/rate", g.handleRate)

	return r, g
}

func (g *Generator) handleRate(c *gin.Context) {
	var request struct {
		Rate uint64 `json:"rate" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.SetRate(request.Rate)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

package v1

import (
	"github.com/gin-gonic/gin"

)

// @Summary Get line information
// @Tags Line
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/line [get]
func Line(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	legendData := []string{
		"A列車：開始吧 觀光開發計畫",
		"集合啦！動物森友會",
		"Fitness Boxing 2",
		"紙片瑪利歐：摺紙國王",
		"皮克敏™３ 豪華版",
		"ZELDA無雙 災厄啟示錄"}
	xAxisData := []string{
		"http://localhost:8081/70010000030778.jpg",
		"http://localhost:8081/animal_crossing.jpg",
		"http://localhost:8081/fitness_boxing_2.jpg",
		"http://localhost:8081/paper_mario.jpg",
		"http://localhost:8081/pikumi3.jpg",
		"http://localhost:8081/zelda_musou.jpg"}
	c.JSON(200, gin.H{
		"legend_data": legendData,
		"xAxis_data":  xAxisData,
	})
}
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuzjing/assets-management/database"
	"github.com/yuzjing/assets-management/models"
)

// 支持模糊搜索 和 范围查询 (>, >=, <, <=)
func GetAllAssets(c *gin.Context) {
	// --- 白名单: 模糊搜索字段 ---
	likeFilters := map[string]bool{
		"gdzc_bh": true, "gdzc_lb": true, "gdzc_ggxh": true, "gdzc_sccj": true,
		"gdzc_zjfs": true, "gdzc_userbm": true, "gdzc_zt": true, "gdzc_gjxx": true,
		"gdzc_user": true, "gdzc_gsname": true, "gdzc_dw": true, "gdzc_djuser": true,
	}

	// --- 白名单: 范围查询字段 ---
	rangeFilters := map[string]bool{
		"gdzc_je": true, "gdzc_nx": true, "gdzc_rzdate": true,
		"gdzc_lydate": true, "gdzc_djdate": true,
	}

	// --- 白名单 3: 允许排序的字段 ---
	allowedSorts := map[string]bool{
		"gdzc_id": true, "gdzc_je": true, "gdzc_nx": true, "gdzc_rzdate": true,
		"gdzc_lydate": true, "gdzc_djdate": true, // 通常只对数字和日期/时间字段进行排序
	}

	query := database.DB.Model(&models.Asset{})

	// --- 1. 处理筛选 ---
	for key, values := range c.Request.URL.Query() {
		if len(values) == 0 || values[0] == "" {
			continue
		}
		value := values[0]

		// 检查范围查询
		isRangeQuery := false
		operators := map[string]string{"_gte": ">=", "_lte": "<=", "_gt": ">", "_lt": "<"}
		for suffix, operator := range operators {
			if strings.HasSuffix(key, suffix) {
				isRangeQuery = true
				fieldName := strings.TrimSuffix(key, suffix)
				if rangeFilters[fieldName] {
					query = query.Where(fmt.Sprintf("%s %s ?", fieldName, operator), value)
				}
				break
			}
		}

		// 检查模糊搜索
		if !isRangeQuery {
			if likeFilters[key] {
				query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+value+"%")
			}
		}
	}

	// --- 2. 核心新逻辑: 处理排序 ---
	sortBy := c.DefaultQuery("sortBy", "gdzc_id") // 默认按 ID 排序
	order := c.DefaultQuery("order", "asc")       // 默认升序

	// 安全检查：确保 sortBy 字段在我们的白名单内
	if allowedSorts[sortBy] {
		// 安全检查：确保 order 方向是 asc 或 desc
		if strings.ToLower(order) == "desc" {
			order = "DESC"
		} else {
			order = "ASC" // 默认为升序
		}

		// 动态构建 Order By 子句
		// 例如: "gdzc_je DESC"
		orderClause := fmt.Sprintf("%s %s", sortBy, order)
		query = query.Order(orderClause)
	}

	// --- 3. 执行查询 ---
	var assets []models.Asset
	result := query.Find(&assets)

	// --- 4. 返回结果 ---
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, assets)
}

func GetAllLogs(c *gin.Context) {
	var logs []models.AssetChangeLog // 使用 AssetChangeLog 模型

	// 目前只做最简单的查询，未来也可以像 GetAllAssets 一样增加筛选和排序
	result := database.DB.Find(&logs)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func main() {
	router := gin.Default()
	database.ConnectDB()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	router.GET("/assets", GetAllAssets)
	// router.POST("/assets", CreateAsset)
	router.GET("/logs", GetAllLogs)

	// router.PUT("/assets/:id", UpdateAsset)
	router.Run(":8123")
}

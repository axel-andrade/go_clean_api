package utils

import (
	"go_clean_api/api/entities"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GenerateUUIDV4() string {
	return uuid.NewV4().String()
}

func GetPaginationOptionsFromURL(c *gin.Context) entities.PaginationOptions {
	var options entities.PaginationOptions

	options.Limit, _ = strconv.Atoi(c.Query("limit"))
	options.Page, _ = strconv.Atoi(c.Query("page"))
	options.Sort = c.Query("order")

	return options
}

func Paginate(db *gorm.DB, model interface{}, pagination *entities.PaginationOptions, output *entities.PaginateResult) func(db *gorm.DB) *gorm.DB {
	var totalDocs int64
	db.Model(model).Count(&totalDocs)
	output.TotalDocs = totalDocs

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func FormatPaginateOutput(pagination *entities.PaginationOptions, output *entities.PaginateResult) *entities.PaginateResult {

	totalPages := int(math.Ceil(float64(output.TotalDocs) / float64(pagination.Limit)))

	output.TotalPages = totalPages
	output.Limit = int64(pagination.Limit)
	output.Page = int64(pagination.Page)
	output.HasPrevPage = int64(pagination.Page) > 1
	output.HasNextPage = int64(pagination.Page) < int64(totalPages)

	if output.HasPrevPage {
		output.PrevPage = int64(pagination.Page) - 1
	}

	if output.HasNextPage {
		output.NextPage = int64(pagination.Page) + 1
	}

	return output
}

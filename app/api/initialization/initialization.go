package initialization

import (
	"imp-goswagger/app/api/transport"
	"imp-goswagger/app/model"
	"imp-goswagger/app/registry"
	"imp-goswagger/helper/constant"
	"imp-goswagger/helper/database"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetConfigString(conf string) string {
	if strings.HasPrefix(conf, "${") && strings.HasSuffix(conf, "}") {
		return os.Getenv(strings.TrimSuffix(strings.TrimPrefix(conf, "${"), "}"))
	}

	return conf
}

func GetConfigInt(conf string) int {
	if strings.HasPrefix(conf, "${") && strings.HasSuffix(conf, "}") {
		result, _ := strconv.Atoi(os.Getenv(strings.TrimSuffix(strings.TrimPrefix(conf, "${"), "}")))
		return result
	}

	result, _ := strconv.Atoi(conf)
	return result
}

func DbInit() (*gorm.DB, error) {
	db, err := database.NewConnectionDB(constant.DBNAME, constant.HOST, constant.DBUSERNAME, constant.DBPASSWORD, constant.PORT)
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(&model.Product{})
	return db, nil
}

func InitRouting(db *gorm.DB, logger log.Logger) *http.ServeMux {
	swagHttp := transport.SwaggerHttpHandler(log.With(logger, "SwaggerTransportLayer", "HTTP"))
	globalHttp := GlobalHttpHandler(log.With(logger, "GlobalTransportLayer", "HTTP"), db)

	// Routing path
	mux := http.NewServeMux()
	mux.Handle("/", swagHttp)
	mux.Handle("/api/v1/", globalHttp)

	return mux
}

func GlobalHttpHandler(logger log.Logger, db *gorm.DB) http.Handler {
	productService := registry.RegisterProductService(db)
	productHttp := transport.ProductHttpHandler(productService, log.With(logger, "ProductTransportLayer", "HTTP"))

	pr := mux.NewRouter()
	pr.PathPrefix("/api/v1/product").Handler(productHttp)

	return pr
}

package routes

import (
	"database/sql"
	"net/http"

	"money-management-be/controllers"
	"money-management-be/middleware"
	"money-management-be/repositories"
	"money-management-be/services"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	incomeRepo := repositories.NewIncomeRepositories(db)
	incomeService := services.NewIncomeServices(incomeRepo)
	incomeController := controllers.NewIncomeController(incomeService)
	mux.Handle("/pemasukan/", http.StripPrefix("/pemasukan", middleware.CORSMiddleware(InitIncomesRoutes(incomeController))))

	outcomeRepo := repositories.NewOutcomeRepositories(db)
	outcomeService := services.NewOutcomeServices(outcomeRepo)
	outcomeController := controllers.NewOutcomeController(outcomeService)
	mux.Handle("/pengeluaran/", http.StripPrefix("/pengeluaran", middleware.CORSMiddleware(InitOutcomesRoutes(outcomeController))))

	assetRepo := repositories.NewAssetRepositories(db)
	assetService := services.NewAssetService(assetRepo)
	assetController := controllers.NewAssetController(assetService)
	mux.Handle("/aset/", http.StripPrefix("/aset", middleware.CORSMiddleware(InitAssetsRoutes(assetController))))

	debtRepo := repositories.NewDebtRepositories(db)
	debtService := services.NewDebtServices(debtRepo)
	debtController := controllers.NewDebtController(debtService)
	mux.Handle("/tagihan/", http.StripPrefix("/tagihan", middleware.CORSMiddleware(InitDebtsRoutes(debtController))))

	return mux
}

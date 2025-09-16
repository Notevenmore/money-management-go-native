package routes

import (
	"net/http"

	"money-management-be/middleware"
	"money-management-be/routes/assets"
	"money-management-be/routes/debts"
	"money-management-be/routes/incomes"
	"money-management-be/routes/outcomes"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/pemasukan/", http.StripPrefix("/pemasukan", middleware.CORSMiddleware(incomes.InitIncomesRoutes())))
	mux.Handle("/pengeluaran/", http.StripPrefix("/pengeluaran", middleware.CORSMiddleware(outcomes.InitOutcomesRoutes())))
	mux.Handle("/aset/", http.StripPrefix("/aset", middleware.CORSMiddleware(assets.InitAssetsRoutes())))
	mux.Handle("/tagihan/", http.StripPrefix("/tagihan", middleware.CORSMiddleware(debts.InitDebtsRoutes())))

	return mux
}

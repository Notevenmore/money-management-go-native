package migrations

func Migrate() {
	CreateIncomesTable()
	CreateDebtTable()
	CreateOutcomesTable()
	CreateAssetTable()
}

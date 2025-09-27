package migrations

func Migrate() {
	CreateIncomesTable()
	CreateDebtTable()
	CreateOutcomesTable()
	CreateDebtOutcomeTable()
	CreateAssetTable()
}

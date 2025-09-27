package migrations

import (
	"fmt"
	"money-management-be/database"
)

func CreateDebtOutcomeTable() {
	query := `
	CREATE TABLE IF NOT EXISTS debt_outcome (
		id SERIAL PRIMARY KEY,
		id_debt INTEGER NOT NULL UNIQUE,
		id_outcome INTEGER NOT NULL UNIQUE,
		CONSTRAINT fk_debt FOREIGN KEY (id_debt) REFERENCES debts(id),
		CONSTRAINT fk_outcome FOREIGN KEY (id_outcome) REFERENCES outcomes(id)
	);
	
	DROP TRIGGER IF EXISTS trg_delete_outcome_after_debt_outcome ON debt_outcome;

	CREATE OR REPLACE FUNCTION delete_outcome_if_link_deleted()
	RETURNS TRIGGER AS $$
	BEGIN
		DELETE FROM outcomes WHERE id = OLD.id_outcome;
		RETURN OLD;
	END;
	$$ LANGUAGE plpgsql;

	CREATE TRIGGER trg_delete_outcome_after_debt_outcome
	AFTER DELETE ON debt_outcome
	FOR EACH ROW
	EXECUTE FUNCTION delete_outcome_if_link_deleted();
	`

	_, err := database.DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration: debts outcome table created successfully")
}

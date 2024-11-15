package supplier

import "time"

type Supplier struct {
	SupplierID  int64  `db:"supplier_id" json:"supplier_id"`
	CompanyName string `db:"company_name" json:"company_name"`
	ChiefName   string `db:"chief_name" json:"chief_name"`
	Address     string `db:"address" json:"address"`
	Email       string `db:"email" json:"email"`
}

type Supply struct {
	SupplyID      int64     `db:"supply_id"`
	SupplierID    int64     `db:"supplier_id"`
	TransactionID int64     `db:"transaction_id"`
	DateFrom      time.Time `db:"date_from"`
	Price         float64   `db:"price"`
}

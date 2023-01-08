package sqldb

import "database/sql"

// ConnectDB opens a connection to the database
func ConnectDB(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

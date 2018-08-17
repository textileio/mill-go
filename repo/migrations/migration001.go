package migrations

import (
	"database/sql"
	_ "github.com/mutecomm/go-sqlcipher"
	"os"
	"path"
)

type Migration001 struct{}

func (Migration001) Up(repoPath string, dbPassword string, testnet bool) error {
	var dbPath string
	if testnet {
		dbPath = path.Join(repoPath, "datastore", "testnet.db")
	} else {
		dbPath = path.Join(repoPath, "datastore", "mainnet.db")
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	if dbPassword != "" {
		p := "pragma key='" + dbPassword + "';"
		if _, err := db.Exec(p); err != nil {
			return err
		}
	}

	// add column for encrypted metadata to blocks
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("alter table blocks add column dataMetadataCipher blob;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	// update version
	f2, err := os.Create(path.Join(repoPath, "repover"))
	if err != nil {
		return err
	}
	defer f2.Close()
	if _, err = f2.Write([]byte("2")); err != nil {
		return err
	}
	return nil
}

func (Migration001) Down(repoPath string, dbPassword string, testnet bool) error {
	return nil
}

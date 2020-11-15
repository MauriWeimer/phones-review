package gateway

import (
	"phones-review/gadgets/smartphones/models"
	"phones-review/internal/database"
	"phones-review/internal/logs"
)

type SmartphoneStorageGateway interface {
	Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)

	// Delete, Find, Update ...
}

type SmartphoneStorage struct {
	*database.MySQLClient
}

func (this *SmartphoneStorage) Add(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	tx, err := this.MySQLClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec("insert into smartphone (name, price, country_origin, os) values (?, ?, ?, ?)",
		cmd.Name, cmd.Price, cmd.CountryOrigin, cmd.Os,
	)

	// Las transacciones se deben cerrar con Rollback o Commit

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fech last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &models.Smartphone{
		Id:            id,
		Name:          cmd.Name,
		Price:         cmd.Price,
		CountryOrigin: cmd.CountryOrigin,
		Os:            cmd.Os,
	}, nil
}

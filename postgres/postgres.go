package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"wb-l0/domain"
	logs "wb-l0/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetParamsForDB() string {
	host := "postgres"
	port := "5432"
	user := "postgres"
	pass := "123"
	dbname := "wb-task"
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	return params
}

func Connect(ctx context.Context, params string) *pgxpool.Pool {

	dbpool, err := pgxpool.New(ctx, params)
	if err != nil {
		logs.LogFatal(logs.Logger, "postgres_connector", "Connect", err, err.Error())
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		logs.LogFatal(logs.Logger, "postgres_connector", "Connect", err, err.Error())
	}
	logs.Logger.Info("Connected to postgres")
	logs.Logger.Debug("postgres client :", dbpool)

	return dbpool
}

func GetOrdersToCache(db domain.PgxPoolIface, ctx context.Context) (domain.Order, error) {
	row := db.QueryRow(ctx, "SELECT * FROM delivery")
	var delivery domain.Delivery
	err := row.Scan(nil, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City,
		&delivery.Address, &delivery.Region, &delivery.Email)

	if err == pgx.ErrNoRows {
		return domain.Order{}, err
	}

	if err != nil {
		return domain.Order{}, err
	}

	row = db.QueryRow(ctx, "SELECT * FROM payment")

	var payment domain.Payment
	err = row.Scan(nil, &payment.Transaction, &payment.RequestID, &payment.Currency,
		&payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank,
		&payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)

	if err == pgx.ErrNoRows {
		return domain.Order{}, err
	}
	if err != nil {
		return domain.Order{}, err
	}

	rows, err := db.Query(ctx, "SELECT * FROM item")
	if err != nil {
		return domain.Order{}, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		var item domain.Item
		err = rows.Scan(nil, &item.ChrtID, &item.TrackNumber, &item.Price, &item.RID,
			&item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			return domain.Order{}, err
		}
		items = append(items, item)
	}

	row = db.QueryRow(ctx, "SELECT * FROM orders")

	var order domain.Order
	err = row.Scan(nil, &order.OrderUID, &order.TrackNumber,
		&order.Entry, &order.Locale, &order.InternalSign,
		&order.CustomerID, &order.ShardKey, &order.SmID,
		&order.DateCreated, &order.OofShard, nil, nil, nil, nil)

	order.Delivery = delivery
	order.Items = items
	order.Payment = payment

	return order, nil
}

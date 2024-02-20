package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type Order struct {
	OrderUID        string    `json:"order_uid"`
	TrackNumber     string    `json:"track_number"`
	Entry           string    `json:"entry"`
	Delivery        Delivery  `json:"delivery"`
	Payment         Payment   `json:"payment"`
	Items           []Item    `json:"items"`
	Locale          string    `json:"locale"`
	InternalSign    string    `json:"internal_signature"`
	CustomerID      string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	ShardKey        string    `json:"shardkey"`
	SmID            int       `json:"sm_id"`
	DateCreated     time.Time `json:"date_created"`
	OofShard        string    `json:"oof_shard"`
}

func GetParamsForDB() string {
	host := "postgres"
	port := "5432"
	user := "postgres"
	pass := "123"
	dbname := "wb-task"
	params := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	return params
}

func main() {

	db, err := sql.Open("postgres", GetParamsForDB())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	delivery := Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
	payment := Payment{
		Transaction:  "b563feb7b2b84b6test",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDT:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}
	items := []Item{
		{ChrtID: 9934930,
			TrackNumber: "WBILMTESTTRACK",
			Price:       453,
			RID:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NmID:        2389212,
			Brand:       "Vivienne Sabo",
			Status:      202},
	}
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal(err)
	}

	order := Order{
		OrderUID:        "b563feb7b2b84b6test",
		TrackNumber:     "WBILMTESTTRACK",
		Entry:           "WBIL",
		Delivery:        delivery,
		Payment:         payment,
		Items:           items,
		Locale:          "en",
		InternalSign:    "test",
		CustomerID:      "test",
		DeliveryService: "meest",
		ShardKey:        "9",
		SmID:            99,
		DateCreated:     time.Date(2021, 11, 26, 6, 22, 19, 23, loc),
		OofShard:        "1",
	}

	_, err = db.Exec(`INSERT INTO delivery (name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt,bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		items[0].ChrtID, items[0].TrackNumber, items[0].Price, items[0].RID, items[0].Name, items[0].Sale, items[0].Size, items[0].TotalPrice, items[0].NmID, items[0].Brand, items[0].Status)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSign, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Данные успешно добавлены в базу данных.")
}

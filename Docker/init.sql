CREATE TABLE delivery (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255),
                          phone VARCHAR(20),
                          zip VARCHAR(10),
                          city VARCHAR(255),
                          address VARCHAR(255),
                          region VARCHAR(255),
                          email VARCHAR(255)
);

CREATE TABLE payment (
                         id SERIAL PRIMARY KEY,
                         transaction VARCHAR(255),
                         request_id VARCHAR(255),
                         currency VARCHAR(10),
                         provider VARCHAR(255),
                         amount INTEGER,
                         payment_dt BIGINT,
                         bank VARCHAR(255),
                         delivery_cost INTEGER,
                         goods_total INTEGER,
                         custom_fee INTEGER
);

CREATE TABLE item (
                      id SERIAL PRIMARY KEY,
                      chrt_id INTEGER,
                      track_number VARCHAR(255),
                      price INTEGER,
                      rid VARCHAR(255),
                      name VARCHAR(255),
                      sale INTEGER,
                      size VARCHAR(10),
                      total_price INTEGER,
                      nm_id INTEGER,
                      brand VARCHAR(255),
                      status INTEGER
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        order_uid VARCHAR(255),
                        track_number VARCHAR(255),
                        entry VARCHAR(10),
                        locale VARCHAR(10),
                        internal_signature VARCHAR(255),
                        customer_id VARCHAR(255),
                        delivery_service VARCHAR(255),
                        shardkey VARCHAR(10),
                        sm_id INT,
                        date_created TIMESTAMP,
                        oof_shard VARCHAR(10),
                        delivery_id INT,
                        payment_id INT,
                        item_id INT,
                        FOREIGN KEY (delivery_id) REFERENCES delivery(id),
                        FOREIGN KEY (payment_id) REFERENCES payment(id),
                        FOREIGN KEY (item_id) REFERENCES item(id)
);

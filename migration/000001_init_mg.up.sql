CREATE TABLE orders
(
    id                 serial primary key,
    order_uid          varchar(255),
    track_number       varchar(255) not null,
    entry              varchar(255),
    locale             varchar(255),
    internal_signature varchar(255),
    customer_id        varchar(255),
    delivery_service   varchar(255),
    shardkey           varchar(255),
    sm_id              int,
    date_created       timestamp default current_timestamp,
    oof_shard          varchar(255)
);

CREATE TABLE deliveries
(
    id       serial primary key,
    order_id int references orders (id),
    name     varchar(255),
    phone    varchar(255),
    zip      varchar(255),
    city     varchar(255),
    address  varchar(255),
    region   varchar(255),
    email    varchar(255)
);

CREATE TABLE payments
(
    id            serial primary key,
    order_id      int references orders (id),
    transaction   varchar(255),
    request_id    varchar(255),
    currency      varchar(255),
    provider      varchar(255),
    amount        int,
    payment_dt    int,
    bank          varchar(255),
    delivery_cost int,
    goods_total   int,
    custom_fee    int
);

CREATE TABLE items
(
    id           serial primary key,
    order_id     int references orders (id),
    chrt_id      int,
    track_number varchar(255),
    price        int,
    rid          varchar(255),
    name         varchar(255),
    sale         int,
    size         varchar(255),
    total_price  int,
    nm_id        int,
    brand        varchar(255),
    status       int
);
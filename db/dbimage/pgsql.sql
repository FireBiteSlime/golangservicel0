DROP TABLE payments, orders, items, deliveries;

CREATE TABLE IF NOT EXISTS public.orders
(
    "Order_uid" text COLLATE pg_catalog."default" NOT NULL,
    "Track_number" text COLLATE pg_catalog."default",
    "Entry" text COLLATE pg_catalog."default",
    "Locale" text COLLATE pg_catalog."default",
    "Internal_signature" text COLLATE pg_catalog."default",
    "Customer_id" text COLLATE pg_catalog."default",
    "Delivery_service" text COLLATE pg_catalog."default",
    "Shardkey" text COLLATE pg_catalog."default",
    "Sm_id" smallint,
    "Date_created" text COLLATE pg_catalog."default",
    "Oof_shard" text COLLATE pg_catalog."default",
    CONSTRAINT orders_pkey PRIMARY KEY ("Order_uid"),
    CONSTRAINT "Track_number" UNIQUE ("Track_number")
        INCLUDE("Track_number")
);

CREATE TABLE IF NOT EXISTS public.payments
(
    "Transaction" text COLLATE pg_catalog."default" NOT NULL,
    "Request_id" text COLLATE pg_catalog."default",
    "Currency" text COLLATE pg_catalog."default",
    "Provider" text COLLATE pg_catalog."default",
    "Amount" integer,
    "Payment_dt" bigint,
    "Bank" text COLLATE pg_catalog."default",
    "Delivery_cost" integer,
    "Goods_total" smallint,
    "Custom_fee" smallint,
    CONSTRAINT payments_pkey PRIMARY KEY ("Transaction"),
    CONSTRAINT "Order_uid" FOREIGN KEY ("Transaction")
        REFERENCES public.orders ("Order_uid") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE IF NOT EXISTS public.items
(
    "Chrt_id" integer NOT NULL,
    "Track_number" text COLLATE pg_catalog."default",
    "Price" integer,
    "Rid" text COLLATE pg_catalog."default",
    "Name" text COLLATE pg_catalog."default",
    "Sale" smallint,
    "Size" text COLLATE pg_catalog."default",
    "Total_price" integer,
    "Nm_id" integer,
    "Brand" text COLLATE pg_catalog."default",
    "Status" smallint,
    CONSTRAINT items_pkey PRIMARY KEY ("Chrt_id"),
    CONSTRAINT "Track_number" FOREIGN KEY ("Track_number")
        REFERENCES public.orders ("Track_number") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE IF NOT EXISTS public.deliveries
(
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "Phone" text COLLATE pg_catalog."default",
    "Zip" text COLLATE pg_catalog."default",
    "City" text COLLATE pg_catalog."default",
    "Address" text COLLATE pg_catalog."default",
    "Region" text COLLATE pg_catalog."default",
    "Email" text COLLATE pg_catalog."default",
    "Order_id" text COLLATE pg_catalog."default",
    CONSTRAINT deliveries_pkey PRIMARY KEY ("Name"),
    CONSTRAINT "Order_id" FOREIGN KEY ("Order_id")
        REFERENCES public.orders ("Order_uid") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);


ALTER TABLE IF EXISTS public.deliveries
    OWNER to "default";

ALTER TABLE IF EXISTS public.items
    OWNER to "default";

ALTER TABLE IF EXISTS public.payments
    OWNER to "default";

ALTER TABLE IF EXISTS public.orders
    OWNER to "default";
# order-management

Run
`docker-compose build`
`docker-compose up`

The containers are up you are ready to go.

Now you have to insert an user in `user` table then you can login

cURL For Login API

```
curl --location 'http://localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "01901901901@mailinator.com",
    "password": "321dsaf"
}'
```

cURL For Order Creation API

```
curl --location 'http://localhost:8080/api/v1/orders' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{TOKEN}}' \
--data '{
    "store_id": 131172,
    "merchant_order_id": "gg",
    "recipient_name": "ironman",
    "recipient_phone": "0171700074",
    "recipient_address": "banani, gulshan 2, dhaka, bangladesh",
    "recipient_city": 2,
    "recipient_zone": 1,
    "recipient_area": 1,
    "delivery_type": 48,
    "item_type": 2,
    "special_instruction": "nai",
    "item_quantity": 1,
    "item_weight": 2,
    "amount_to_collect": 2000,
    "item_description": "nai"
}'
```

cURL For Fetching Orders API

```
curl --location 'http://localhost:8080/api/v1/orders/all?limit=10&page=1' \
--header 'Authorization: Bearer {{TOKEN}}' \
```

cURL For Order Cancelling API

```
curl --location --request PUT 'http://localhost:8080/api/v1/orders/17329744464O8HIH/cancel' \
--header 'Authorization: Bearer {{TOKEN}}' \
```

# Golang Backend Test

### Requirement:

-   Buat sebuah Rest API product dengan bahasa pemrograman go dengan RDMS mysql atau postgre
-   Buat 2 endpoint.
    -   API Add product
    -   API List product dengan sorting

### Detail Requirement

-   Dengan spesifikasi product:
    -   product id
    -   product name
    -   product price
    -   product description
    -   product quantity
-   Rest API bisa melakukan tambah product
-   Rest API bisa mengurutkan berdasarkan berikut:
    -   product terbaru
    -   product harga termurah
    -   product harga termahal
    -   sort by product name (A-Z) dan (Z-A)

### Architecture

-   Beritahu kita desain arsitekture seperti apa yang Anda buat dan disertai alasan-nya di dalam Readme.md

### Bonus points

-   Implement redis cache
-   Implement docker

Code di publish di github/bitbucket, sertakan readme dan dokumentasi API didalam readme, sehingga mudah digunakan.

Good Luck!

## Endpoints

Get all products

```bash
[GET] /product

Response:
[
  {
    "name": "Whimzees XXS 113pcs",
    "price": 250000,
    "description": "Dog treats with natural ingredients",
    "quantity": 5,
    "created_at": "2011-07-06T10:35:24Z"
  },
  {
    "name": "Skintific sunscreen",
    "price": 120000,
    "description": "Sunscreen with spf 50 and PA ++++",
    "quantity": 7,
    "created_at": "2011-02-28T13:34:09Z"
  },
  ...
]

```

Get All Product sort by field

```bash
[GET] /product?sort_by=name

Response:
[
  {
    "name": "3D Embellishment Art Lamp",
    "price": 200000,
    "description": "3D led lamp sticker Wall sticker 3d wall art light on/off button  cell operated (included)",
    "quantity": 54,
    "created_at": "2011-10-05T05:32:55Z"
  },
  {
    "name": "ikan bilis",
    "price": 50000,
    "description": "a great food",
    "quantity": 3,
    "created_at": "2023-08-16T20:10:56Z"
  },
  ...
]
```

Get All Products sort by field and sort type

```bash
[GET] /product?sort_by=price&sort_type=desc

Response:
[
  {
    "name": "MacBook Pro",
    "price": 34500000,
    "description": "MacBook Pro 2021 with mini-LED display may launch between September, November",
    "quantity": 83,
    "created_at": "2011-12-05T19:23:01Z"
  },
  {
    "name": "Iphone 14",
    "price": 25000000,
    "description": "High class phone type from apple",
    "quantity": 15,
    "created_at": "2012-03-21T03:58:07Z"
  },
  ...
]
```

Insert product

```bash
[POST] /product

Body:
{
    "name": "Bike",
    "price": 7500000,
    "description": "a famous bike",
    "quantity": 17
}

Response:
{
    "message": "product inserted"
}
```

## Architecture Design

This REST API using N-Tier Architecture where we divide into several layers such as layer handler, service, and repository. this is because the scope of the API is still very small with the functionality of displaying products and creating products only. using this architecture has several advantages such as:

-   Easy to manage. Because it is divided into several layers, it does not affect other layers if there is a change in one of the layers.
-   High scalability. If we want to add something, it can be added to the intended layer without affecting other layers.
-   Very efficient in the development process. because the architecture is easy to understand.
-   Reusable. because the application is divided into several independent layers, we can use it in other places. for example, we can replicate it for other needs by only changing the desired output data or data layer without changing the logic.
-   Easy to test. Since each layer is independent, we can perform isolated tests without involving other layers.

## Tech Stacks

| Tools | Version |
| ----- | ------- |
| Go    | 1.20    |
| MySQL | 8.1     |

## How to use

-   run database seed MySQL to initialize Database
-   change config.example.yaml to config.yaml
-   change configuration in config.yaml like database config or app config
-   Run "go get"
-   Run local development

```bash
  go run main.go
```

-   to run from distribution file, copy config.yaml to bin folder and run as usually

# Product-api

# Cara Menjalankan Program
- Pastikan bahasa pemrograman golang sudah terinstall di komputer kamu, jika belum untuk lebih lengkapnya maka bisa mengunjungi https://go.dev/doc/
- Tambahkan file .env dengan cara menduplikasi .env-example dan merubah namanya ke '.env'
- Sesuaikan nilai yang ada di .env dengan environment local
- Ketikan pada terminal `go get`
- Ketikan pada terminal `make build_and_run`

# Deskripsi

Di dalam project ini saya menggunakan gin sebagai framework.

Dalam project ini terdapat 2 endpoint :

  1. [GET] `http://localhost:8080/index?limit=10&sort=id&order=desc`

     - Endpoint ini berfungsi untuk mengambil data product secara keseluruhan dengan pagination
     - Untuk mengurutkan berdasarkan product terbaru maka tambahkan query param seperti ini `http://localhost:8080/index?limit=10&sort=id&order=desc`
     - Untuk mengurutkan berdasarkan product dengan harga termurah maka tambahkan query param seperti ini `http://localhost:8080/index?limit=10&sort=price&order=asc`
     - Untuk mengurutkan berdasarkan product dengan harga terendah maka tambahkan query param seperti ini `http://localhost:8080/index?limit=10&sort=price&order=asc`
     - Untuk mengurutkan berdasarkan product dengan quantity maka tambahkan query param seperti ini `http://localhost:8080/index?limit=10&sort=price&order=asc`
     
     - response example
     ```
         {
            "data": {
                "TotalData": 10,
                "FilteredData": 2,
                "Data": [
                    {
                        "id": 3,
                        "name": "Magic Mouse",
                        "price": 2000000,
                        "description": "Series 2",
                        "quantity": 4,
                        "created_at": "2022-01-12T22:11:37+07:00",
                        "updated_at": "2022-01-12T22:11:37+07:00",
                        "deleted_at": null
                    },
                    {
                        "id": 1,
                        "name": "iPhone 13 Pro",
                        "price": 14000000,
                        "description": "iPhone 13 Blue Ocean",
                        "quantity": 12,
                        "created_at": "2022-01-12T11:59:26+07:00",
                        "updated_at": "2022-01-12T11:59:26+07:00",
                        "deleted_at": null
                    }
                ]
            },
            "meta": {
                "message": "Get Product List Success.",
                "code": 200,
                "status": "success"
            }
         }
         


  2. [POST] `http://localhost:8080/products`

     Endpoint ini berfungsi untuk menambahkan data product

    request example
    {
      "name": "Macbook Air 2022",
      "price": 15000000,
      "description": "Ram 512",
      "quantity": 10
     }
     
     response example
     {
        "data": {
            "id": 10,
            "name": "Macbook Air 2022",
            "price": 15000000,
            "description": "Ram 512",
            "quantity": 10,
            "created_at": "2022-01-13T20:49:58.954674+07:00",
            "updated_at": "2022-01-13T20:49:58.954674+07:00",
            "deleted_at": null
        },
        "meta": {
            "message": "Get Product List Success.",
            "code": 200,
            "status": "success"
        }
     }
     
Struktur Projek
```
  common
  ├── dto
      ├── base_response_dto.go
      ├── pagination_dto.go
  ├── util
      ├── utils.go
  database
  ├── database.go
  product
  ├── product.go
  ├── product_api.go
  ├── product_dto.go
  ├── product_mapper.go
  ├── product_repository.go
  └── product_service.go
  go.mod
  go.sum
  main.go
  wire.go
  wire_gen.go

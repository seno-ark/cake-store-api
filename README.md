# cake-store-api

cake-store-api adalah Rest API dengan fungsi CRUD untuk menyimpan, mengubah, menghapus, dan menampilkan data cakes.

## Development
Menjalankan db server di container:
- make docker-db

Menjalankan API server di local:
- make go-dev

Menjalankan test:
- make go-test


## Run
Untuk menjalankan API server & db server di container
- make docker-up


## .Env

Sebelum menjalankan server API, buat file .env dengan cara meng-copy .env-example.
Kemudian sesuaikan isi variabel-variabel yang tersedia.

```
Jika server API dijalankan di dalam container, maka `MYSQL_HOST = db` (sesuai dengan nama service container mysql)
```


## Migration

Pada project ini, migration dilakukan menggunakan https://github.com/golang-migrate/migrate

__Create migration file__

Untuk membuat file migration baru, perintahnya adalah:

```migrate create -ext sql -dir database/migration -seq <file_name>```

Kemudian sesuaikan <file_name>.up.sql & <file_name>.down.sql di dalam direktori `database/migration`

__Run Migration__

- make migrate-up
- make migrate-down


## Makefile

Berikut adalah command yang dapat dijalankan dengan perintah `make` <command>:

- __go-dev__: untuk menjalankan server API di host
- __go-test__: untuk menjalankan fil test
- __go-build__: untuk build api server menjadi binary file

- __docker-up__: untuk menjalankan API server & Mysql Server di dalam container
- __docker-db__: untuk menjalankan Mysql Server di dalam container
- __docker-down__: untuk menghentikan container
- __migrate-up__: menjalankan db migration
- __migrate-down__: membatalkan db migration


## Api Documentation
Dokumentasi API endpoints dapat dilihat di file: `api-doc.md`

# MyTask.App


Readme in english [click here]().

MyTask app adalah sebuah aplikasi untuk manajemen pengelolaan task atau sebuah project untuk user. Dengan aplikasi ini, pengguna / user akan lebih mudah dan felxible dalam manajemen task dan mengatur projectnya. 


## Fitur Users 

- Registers
- Login Users
- CRUD Users

## Fitur Task & Project

- CRUD Task
- CRUD Project

## Teknologi
- Database (PostgreSQL, DBeaver)
- Golang 
- Framework (GIN)
- ORM (GORM)

## Menjalankan Lokal

Cloning project

```bash
  $ 
```

Masuk ke direktori project

```bash
  $ cd ~/nama project kamu
```
Buat `database` baru

Buat sebuah file dengan nama di dalam folder root project `.env` dengan format dibawah ini. Sesuaikan configurasi di komputer lokal

```bash
export DBUSER='postgres'
export DBPASS='masukkan password kamu'
export DBHOST='localhost'
export DBPORT='5432'
export DBNAME='my_task_app'
export JWTSECRET='......'

```

Jalankan aplikasi 

```bash
  $ go run main.go
```


## Authors

- [@royanqodri](https://github.com/royanqodri)


 
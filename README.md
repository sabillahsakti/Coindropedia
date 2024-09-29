# CoinDropedia Backend

CoinDropedia Backend adalah API yang menyediakan informasi tentang airdrop cryptocurrency dan fitur favorit pengguna. Proyek ini dibangun menggunakan Go dan PostgreSQL, dan dilengkapi dengan autentikasi pengguna menggunakan JWT.

## Daftar Isi

1. [Fitur](#fitur)
2. [Teknologi yang Digunakan](#teknologi-yang-digunakan)
3. [Prasyarat](#prasyarat)
4. [Instalasi](#instalasi)
5. [Konfigurasi](#konfigurasi)
6. [Rute API](#rute-api)
7. [Penggunaan](#penggunaan)
8. [Lisensi](#lisensi)

## Fitur

- Autentikasi pengguna menggunakan JWT.
- Mendapatkan daftar airdrop yang tersedia.
- Menandai airdrop sebagai favorit.

## Teknologi yang Digunakan

- **Golang**: Bahasa pemrograman untuk membangun API.
- **PostgreSQL**: Basis data yang digunakan untuk menyimpan data pengguna dan airdrop.
- **GORM**: ORM (Object-Relational Mapping) untuk interaksi dengan basis data.
- **JWT**: JSON Web Tokens untuk autentikasi.

## Prasyarat

Sebelum menjalankan proyek ini, pastikan Anda memiliki:

- **Go** (minimal v1.16)
- **PostgreSQL**
- **Golang modules** diaktifkan (`GO111MODULE=on`)

## Instalasi

1. **Clone repositori ini**:
   ```bash
   git clone https://github.com/username/coindropedia-backend.git
   cd coindropedia-backend

2. **Instal dependensi**:
    ```bash
    go mod tidy
3. **Buat basis data di PostgreSQL dan sesuaikan konfigurasi di file .env.**

4. **Jalankan aplikasi**:
    ```bash
    go run main.go
    
## Konfigurasi
Buat file .env di root proyek Anda dan tambahkan variabel berikut:
  ```
  DATABASE_URL=postgres://username:password@localhost:5432/database_name
  JWT_KEY=your_jwt_secret_key
  ```
Gantilah username, password, dan database_name sesuai dengan pengaturan PostgreSQL Anda.

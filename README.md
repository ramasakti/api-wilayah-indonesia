API Data Wilayah Indonesia (Golang Rewrite)
==========================================

> [!IMPORTANT]
> Proyek ini adalah hasil **rewrite ke Go (Golang)** dari repositori asli [emsifa/api-wilayah-indonesia](https://github.com/emsifa/api-wilayah-indonesia). 
> Berbeda dengan versi aslinya yang menggunakan API statis, versi ini menggunakan server Go dinamis yang jauh lebih efisien untuk dijalankan di lingkungan Docker.

## Dokumentasi API (Swagger)

Aplikasi ini dilengkapi dengan dokumentasi interaktif menggunakan Swagger UI. 
Setelah menjalankan server, Anda dapat mengakses dokumentasi di:
`http://localhost:8080/swagger/index.html`

## Pengembangan

Aplikasi dikembangkan menggunakan:
* **Go (Golang)**: Bahasa pemrograman utama.
* **Gin**: Framework web untuk performa tinggi.
* **Swagger**: Untuk dokumentasi API otomatis.
* **Docker**: Untuk isolasi dan kemudahan deployment.

## Cara Menjalankan

### Menggunakan Docker Compose (Direkomendasikan)
Cukup jalankan perintah berikut di direktori utama:
```bash
docker compose up -d
```
API akan tersedia di port `8080`.

### Menjalankan Lokal
1. Pastikan Anda telah memasang Go (versi 1.25 atau lebih baru).
2. Masuk ke direktori aplikasi: `cd app`
3. Jalankan aplikasi: `go run .`

## Versi Asli (Static API)
Informasi di bawah ini adalah keterangan cara kerja dari versi asli (static API). Pada versi Go ini, CSV dibaca langsung ke memori saat startup.

#### Apa yang dimaksud API statis? 
...

API statis adalah API yang _endpoint_-nya terdiri dari file statis.

#### Keuntungan API statis?

* Dapat dihosting pada _static file hosting_ seperti Github Page, Netlify, dsb.
* Proses lebih cepat karena tidak membutuhkan server-side scripting.

#### Bagaimana cara kerjanya?

* Daftar provinsi, kab/kota, kecamatan, kelurahan/desa disimpan pada folder `data` berupa file `csv` (agar mudah diedit).
* Kemudian script `generate.php` dijalankan. Script ini akan membaca file `csv` didalam folder `data`, kemudian men-generate ribuan endpoint (file) kedalam folder `static/api`.
* API siap 'dihidangkan'.

#### Saya mau hosting di Github saya sendiri, bagaimana caranya?

* Klik fork di pojok kanan atas.
* Pada halaman forking, **HAPUS CENTANG** "Copy the master branch only".
* Klik "Create Fork".
* Setelah selesai di Fork, klik Settings (bukan setting account, tapi setting repository).
* Klik menu "Pages" untuk masuk ke menu pengaturan GitHub Pages.
* Pada menu pengaturan GitHub Pages:
  * Pilih Source: Deploy from a Branch
  * Branch: `gh-pages`
  * Direktori: `/root`
  * Klik Save
* Tunggu beberapa menit (5-10 menitan), kembali ke halaman home repository (https://github.com/usernamekamu/api-wilayah-indonesia).
* Kalau halaman sudah terdeploy, di bagian kanan halaman, akan muncul informasi "Environments". Kalau belum tunggu lagi beberapa menit, lalu refresh.
* Kalau sudah muncul informasi Environmentsnya, klik bagian "🚀 github-pages".
* Di halaman Deployments, klik "View Deployment" untuk melihat halaman yang berhasil terdeploy.

## ENDPOINTS

#### 1. Mengambil Daftar Provinsi

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json
```

Contoh Response:

```
[
  {
    "id": "11",
    "name": "ACEH"
  },
  {
    "id": "12",
    "name": "SUMATERA UTARA"
  },
  ...
]
```

#### 2. Mengambil Daftar Kab/Kota pada Provinsi Tertentu

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/regencies/{provinceId}.json
```

Contoh untuk mengambil daftar kab/kota di provinsi Aceh (ID = 11):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/regencies/11.json
```

Contoh Response:

```
[
  {
    "id": "1101",
    "province_id": "11",
    "name": "KABUPATEN SIMEULUE"
  },
  {
    "id": "1102",
    "province_id": "11",
    "name": "KABUPATEN ACEH SINGKIL"
  },
  ...
]
```

#### 3. Mengambil Daftar Kecamatan pada Kab/Kota Tertentu

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/districts/{regencyId}.json
```

Contoh untuk mengambil daftar kecamatan di Aceh Selatan (ID = 1103):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/districts/1103.json
```

Contoh Response:

```
[
  {
    "id": "1103010",
    "regency_id": "1103",
    "name": "TRUMON"
  },
  {
    "id": "1103011",
    "regency_id": "1103",
    "name": "TRUMON TIMUR"
  },
  ...
]
```

#### 4. Mengambil Daftar Kelurahan pada Kecamatan Tertentu

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/villages/{districtId}.json
```

Contoh untuk mengambil daftar kelurahan di Trumon (ID = 1103010):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/villages/1103010.json
```

Contoh Response:

```
[
  {
    "id": "1103010001",
    "district_id": "1103010",
    "name": "KUTA PADANG"
  },
  {
    "id": "1103010002",
    "district_id": "1103010",
    "name": "RAKET"
  },
  ...
]
```

#### 5. Mengambil Data Provinsi berdasarkan ID Provinsi

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/province/{provinceId}.json
```

Contoh untuk mengambil data provinsi Aceh (ID = 11):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/province/11.json
```

Contoh Response:

```
{
  "id": "11",
  "name": "ACEH"
}
```

#### 6. Mengambil Data Kab/Kota berdasarkan ID Kab/Kota

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/regency/{regencyId}.json
```

Contoh untuk mengambil data kabupaten Aceh Selatan (ID = 1103):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/regency/1103.json
```

Contoh Response:

```
{
  "id": "1103",
  "province_id": "11",
  "name": "KABUPATEN ACEH SELATAN"
}
```

#### 7. Mengambil Data Kecamatan berdasarkan ID Kecamatan

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/district/{districtId}.json
```

Contoh untuk mengambil data kecamatan Trumon Timur (ID = 1103011):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/district/1103011.json
```

Contoh Response:

```
{
  "id": "1103011",
  "regency_id": "1103",
  "name": "TRUMON TIMUR"
}
```

#### 8. Mengambil Data Kelurahan berdasarkan ID Kelurahan

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/village/{villageId}.json
```

Contoh untuk mengambil data kelurahan Jambo Dalem (ID = 1103011010):

```
GET https://emsifa.github.io/api-wilayah-indonesia/api/village/1103011010.json
```

Contoh Response:

```
{
  "id": "1103011010",
  "district_id": "1103011",
  "name": "JAMBO DALEM"
}
```

## LIMITASI

Karena API ini dihosting di Github Page, Github Page sendiri memberikan batasan bandwith 100GB/bulan. Rata-rata endpoint disini memiliki ukuran 1KB/endpoint, jadi kurang lebih request yang dapat digunakan adalah 100.000.000 request per bulan, atau sekitar 3.000.000 request/hari.

Karena limitasi ini, disarankan untuk hosting API ini di github kamu sendiri.

Untuk lebih detail tentang limitasi Github Page, bisa dilihat [disini](https://help.github.com/en/articles/about-github-pages#usage-limits).

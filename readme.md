# Money Management Backend

Backend service untuk aplikasi **Money Management**, dibangun menggunakan **Golang** dengan arsitektur berbasis `controllers`, `services`, dan `repositories`.  
Database yang digunakan adalah **PostgreSQL**.

---

## 🚀 Tech Stack
- [Go](https://golang.org/) v1.25+
- [PostgreSQL](https://www.postgresql.org/) v15+
- [net/http](https://pkg.go.dev/net/http) (routing)
- `database/sql` + [lib/pq](https://pkg.go.dev/github.com/lib/pq) (Postgres driver)

---

## 📂 Project Structure
```bash
money-management-be/
├── controllers/      # Handler untuk request (HTTP layer)
├── database/         # Config koneksi DB & migrasi
│   ├── migrations/   # File migrasi database
├── middleware/       # Middleware (contoh: CORS)
├── models/           # Struct data (DTO / entity)
├── repositories/     # Query database (CRUD)
├── responses         # Template Response REST API
├── routes/           # Routing per modul (assets, incomes, outcomes, debts)
├── services/         # Business logic
├── .env.example      # Konfigurasi environment
├── go.mod
├── go.sum
└── main.go           # Entry point aplikasi
```

## ⚙️ Setup & Installation
```bash
# 1. Clone repository
git clone https://github.com/Notevenmore/money-management-go-native.git
cd money-management-go-native

# 2. Install dependencies
go mod tidy

# 3. Copy File .env.example
APP_PORT= #Port REST API
APP_SERVER= #Server Domain REST API (Default localhost)
APP_ORIGIN= #Allowed Request Domain port
APP_TYPE= # "http:// or https://"

DB_CONNECTION=pgsql
DB_HOST= # Database Domain
DB_PORT= # Database Port
DB_DATABASE=money_management_db # Database Name
DB_USERNAME= # Database Username
DB_PASSWORD= # Database Password

# 4. Database Migration + Run Server
go run main.go
```

## 📡 API Endpoints
### Assets
```bash
GET  /aset/     # Get All Assets
POST /aset/     # Create New Assets
```
#### 🧪 Example Request (cURL Get Assets)
```bash
curl 
  -X GET 
  -H "Content-Type: application/json" 
  http://10.4.169.118:8080/aset/ 
```
#### 🧪 Example Response (cURL Get Assets)
```bash
{
  "status":"Success",
  "message":"Daftar Aset berhasil didapatkan",
  "data":[
    {
      "id":1,
      "name":"Emas 1 Gram",
      "is_reusable":true
    }
  ],
  "status_code":202
}
```

#### 🧪 Example Request (cURL Create Assets)
```bash
curl 
  -X POST 
  -H "Content-Type: application/json"  
  -d '{
    "name":"Emas 1 Gram", 
    "is_reusable": true
  }' 
  http://10.4.169.118:8080/aset/
```
#### 🧪 Example Response (cURL Create Assets)
```bash
{
  "status":"Success",
  "message":"Data Aset berhasil ditambahkan",
  "status_code":202
}
```

### Incomes
```bash
GET  /pemasukan/     # Get All Incomes
POST /pemasukan/     # Create New Incomes
```
#### 🧪 Example Request (cURL Get Incomes)
```bash
curl 
  -X GET 
  -H "Content-Type: application/json"
/pemasukan/
```
#### 🧪 Example Response (cURL Get Incomes)
```bash
{
  "status":"Success",
  "message":"Daftar Pemasukan berhasil didapatkan",
  "data":[
    {
      "id":1,
      "name":"Tabungan Mingguan",
      "nominal":1000000,
      "date":"2025-06-23T14:18:08Z",
      "category":"Insentif"
    }
  ],
  "status_code":202
}
```

#### 🧪 Example Request (cURL Create Incomes)
```bash
curl 
  -X POST 
  -H "Content-Type: application/json"  
  -d '{
      "name":"Tabungan Bulanan", 
      "nominal": 3200000, 
      "date":"2025-06-23T14:18:08Z", 
      "category":"Insentif"
    }' 
/pemasukan/
```
#### 🧪 Example Response (cURL Create Incomes)
```bash
{
  "status":"Success",
  "message":"Data Pemasukan berhasil ditambahkan",
  "status_code":202
}
```

### Outcomes
```bash
GET  /pengeluaran/     # Get All Outcomes
POST /pengaluaran/     # Create New Outcomes
```
#### 🧪 Example Request (cURL Get Outcomes)
```bash
curl 
  -X GET 
  -H "Content-Type: application/json"
/pengeluaran
```
#### 🧪 Example Response (cURL Get Outcomes)
```bash
{
  "status":"Success",
  "message":"Daftar Pengeluaran berhasil didapatkan",
  "data":[
    {
      "id":1,
      "name":"Pembayaran Tagihan",
      "nominal":200000,
      "date":"2025-06-23T14:18:08Z",
      "category":"Tagihan"
    }
  ],
  "status_code":202
}
```

#### 🧪 Example Request (cURL Create Outcomes)
```bash
curl 
  -X POST 
  -H "Content-Type: application/json"  
  -d '{
      "name":"Pembayaran Tagihan", 
      "nominal": 200000, 
      "date":"2025-06-23T14:18:08Z", 
      "category":"Tagihan"
    }' 
/pengeluaran/
```
#### 🧪 Example Response (cURL Create Outcomes)
```bash
{
  "status":"Success",
  "message":"Data Pengeluaran berhasil ditambahkan","status_code":202
}
```

### Debts
```bash
GET  /tagihan/     # Get All Debts
POST /tagihan/     # Create New Debts
```
#### 🧪 Example Request (cURL Get Debts)
```bash
curl 
  -X GET 
  -H "Content-Type: application/json" 
  /tagihan
```
#### 🧪 Example Response (cURL Get Debts)
```bash
{
  "status":"Success",
  "message":"Daftar Tagihan berhasil didapatkan",
  "data":[
    {
      "id":1,"name":"Shopee Pay Later",
      "nominal":200000,
      "deadline":"2025-06-23T14:18:08Z",
      "is_finish":false
    }
  ],
  "status_code":202
}
```

#### 🧪 Example Request (cURL Create Debts)
```bash
curl 
  -X POST 
  -H "Content-Type: application/json"  
  -d '{
    "name":"Shopee Pay Later", 
    "nominal":200000,
    "deadline":"2025-06-23T14:18:08Z", 
    "is_finish": false
  }' 
  /tagihan
```
#### 🧪 Example Response (cURL Create Debts)
```bash
{
  "status":"Success",
  "message":"Data Tagihan berhasil ditambahkan",
  "status_code":202
}
```

#### 🧪 Example Request (cURL Update Debts)
```bash
curl 
  -X PUT 
  -H "Content-Type: application/json" 
  -d '{ 
    "name": "Shopee Pay Later", 
    "nominal": 200000, 
    "is_finish": true 
  }' /tagihan/1
```
#### 🧪 Example Response (cURL Update Debts)
```bash
{
  "status":"Success",
  "message":"Data Tagihan berhasil diupdate",
  "status_code":202
}
```

## 🛠️ Architecture Flow
```bash
flowchart TD
    Client -->|HTTP Request| Controller
    Controller --> Service
    Service --> Repository
    Repository -->|SQL Query| Database[(PostgreSQL)]
    Database --> Repository
    Repository --> Service
    Service --> Controller
    Controller -->|HTTP Response (JSON)| Client
```

## 📜 License

Proyek ini dilisensikan di bawah **MIT License** – lihat file [LICENSE](./LICENSE) untuk detail lebih lanjut.

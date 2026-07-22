# go-service-client

**Concurrent REST API Client** — Sebuah CLI berbasis Go yang mengonsumsi layanan aritmetika HTTP secara konkuren.

Project ini mendemonstrasikan **pola concurrent fan-out** di Go: mengirim 3 request HTTP secara paralel ke endpoint aritmetika, mengumpulkan hasilnya melalui channel, dan menampilkan respons agregat.

## Arsitektur

```
cmd/main.go                          # Entry point — orchestrator goroutine + channel
  │
  ├── internal/repositories/math.go  # Business logic — panggil API dengan timeout
  │       │
  │       └── internal/commons/client.go  # HTTP transport — RestClient wrapper
  │               │
  │               └── internal/entities/math.go  # DTO (NumbersReq, NumbersResp)
  │
  └── internal/entities/http_method.go  # Interface generik HttpMethod[T]
```

Lapisan arsitektur:

| Lapisan | Paket | Tanggung Jawab |
|---|---|---|
| **Entry Point** | `cmd/main` | Orchestrasi goroutine, aggregasi hasil via channel |
| **Repository** | `internal/repositories` | Logika bisnis, timeout per-request, logging timing |
| **Client** | `internal/commons` | HTTP client, serialisasi/deserialisasi JSON |
| **Entities** | `internal/entities` | Data transfer objects (DTO) + interface generik |
| **Interface** | `internal/entities` | Abstraksi `HttpMethod[T any]` untuk HTTP POST |

## Prerequisites

- **Go 1.26.5** atau lebih baru
- **Matematika API Server** — project ini membutuhkan server yang menyediakan endpoint:
  - `POST /math/addition`
  - `POST /math/subtraction`
  - `POST /math/multiply`

  Server diharapkan berjalan di `http://localhost:1323` (dapat diubah di `cmd/main.go`).

  Contoh server yang kompatibel: [go-service-diini](https://github.com/diiniS/go-service-diini)

## Instalasi

```bash
# Clone repositori
git clone https://github.com/diiniS/go-service-client.git
cd go-service-client

# Pastikan dependencies tersedia (tidak ada dependency eksternal)
go mod tidy
```

## Menjalankan

Pastikan server matematika sudah berjalan, lalu:

```bash
go run cmd/main.go
```

Output yang dihasilkan (contoh):

```
{3 2 6}
```

Format: `{Addition Subtraction Multiply}` — hasil dari operasi `3 + 2`, `3 - 2`, dan `3 * 2`.

## Struktur Project

```
go-service-client/
├── README.md
├── go.mod                    # Module: go-service-client (Go 1.26.5, zero external deps)
├── go.sum                    # Lockfile (kosong — tanpa dependency eksternal)
├── cmd/
│   └── main.go               # Entry point + goroutine orchestrator
└── internal/
    ├── commons/
    │   └── client.go         # RestClient — HTTP POST dengan context dan JSON
    ├── entities/
    │   ├── http_method.go    # Interface generik HttpMethod[T any]
    │   └── math.go           # NumbersReq / NumbersResp DTO
    └── repositories/
        └── math.go           # Fungsi Addition, Subtraction, Multiply + callTimeout
```

## Detail Teknis

### ⚡ Concurrency Pattern — Fan-Out

`cmd/main.go` menggunakan **fan-out pattern**:

1. Buat 1 channel bertipe `allRespChannel`
2. Launch 3 goroutine — masing-masing memanggil endpoint berbeda
3. Induk goroutine membaca 3 kali dari channel untuk mengumpulkan semua hasil
4. Hasil digabungkan ke struct `allResponse` dan dicetak

### ⏱ Timeout Management

Setiap panggilan API memiliki timeout independen:

| Endpoint | Timeout |
|---|---|
| `POST /math/addition` | 30 detik |
| `POST /math/subtraction` | 30 detik |
| `POST /math/multiply` | 10 detik |

Timeout diimplementasikan dengan `context.WithTimeout`, bukan sekadar HTTP client timeout — memastikan pembatalan yang tepat di seluruh lapisan.

### 🔧 Komponen Utama

#### RestClient (`internal/commons/client.go`)
- Membungkus `*http.Client` standar Go dengan timeout 30 detik
- Method `PostJSON` menangani: marshaling request → HTTP POST → membaca body → unmarshaling response
- Menggunakan `context.Context` untuk mendukung pembatalan

#### callTimeout (`internal/repositories/math.go`)
- Private function yang membungkus `RestClient.PostJSON` dengan timeout per-panggilan
- Mencetak log timing ke stdout (elapsed time, timeout, path, hasil/error)

#### HttpMethod Interface (`internal/entities/http_method.go`)
- Interface generik `HttpMethod[T any]` — mendefinisikan kontrak `PostJSON`
- Tersedia untuk injeksi ketergantungan dan pengujian di masa depan

## Development

### Menambahkan Endpoint Baru

1. Tambahkan function baru di `internal/repositories/math.go`:
   ```go
   func Divide(ctx context.Context, client *commons.RestClient, payload entities.NumbersReq) (int32, error) {
       return callTimeout(ctx, client, "math/divide", payload, 30*time.Second)
   }
   ```
2. Panggil dari goroutine baru di `cmd/main.go`

### Menjalankan dengan Data Berbeda

Ubah nilai `A` dan `B` di `cmd/main.go`:

```go
payload := entities.NumbersReq{A: 10, B: 5}
```

### Testing

```bash
go test ./...
```

## Dependencies

**Zero external dependencies.** Project ini hanya menggunakan Go standard library:

| Package | Penggunaan |
|---|---|
| `net/http` | HTTP client & request |
| `encoding/json` | Serialisasi/deserialisasi JSON |
| `context` | Timeout & cancellation per-request |
| `io` / `bytes` | Baca response body |
| `time` | Konfigurasi timeout |
| `fmt` / `errors` | Output dan error handling |

## Lisensi

MIT

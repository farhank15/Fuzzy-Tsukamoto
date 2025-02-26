# Project fuzzy

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
###
# Sistem Penentuan Predikat Kelulusan Mahasiswa 🎓
> Menggunakan Metode Fuzzy Tsukamoto

## 📝 Deskripsi
Sistem ini dikembangkan untuk menentukan predikat kelulusan mahasiswa berdasarkan berbagai parameter akademik dan non-akademik menggunakan metode Fuzzy Tsukamoto.

## 🎯 Fitur Utama
- Perhitungan otomatis predikat kelulusan
- Mempertimbangkan 7 variabel penilaian
- Implementasi logika fuzzy untuk hasil yang lebih akurat
- Bobot yang dapat disesuaikan untuk setiap variabel

## 📊 Variabel dan Bobot Penilaian

### 1. IPK (30%)
| Kategori | Rentang | Nilai |
|----------|---------|-------|
| Sangat Rendah | [0 - 2.00) | 1 |
| Rendah | [2.00 - 2.75) | 2 |
| Cukup | [2.75 - 3.50) | 3 |
| Sedang | [3.50 - 3.79) | 4 |
| Tinggi | [3.79 - 4.00) | 5 |
| Sangat Tinggi | [4.00] | 6 |

### 2. Lama Studi (20%)
| Kategori | Rentang | Nilai |
|----------|---------|-------|
| Sangat Cepat | [7 semester) | 5 |
| Cepat | [7 - 8 semester) | 4 |
| Sedang | [9 - 10 semester) | 3 |
| Lama | [11 - 12 semester) | 2 |
| Sangat Lama | [12+ semester] | 1 |

### 3. SKS per Semester (10%)
| Kategori | Rentang | Nilai |
|----------|---------|-------|
| Sangat Rendah | [0 - 12 SKS) | 1 |
| Rendah | [12 - 15 SKS) | 2 |
| Sedang | [16 - 18 SKS) | 3 |
| Tinggi | [19 - 21 SKS) | 4 |
| Sangat Tinggi | [21+ SKS] | 5 |

### 4. Skripsi (10%)
| Kategori | Nilai Huruf | Bobot | Nilai |
|----------|-------------|-------|-------|
| Rendah | C | 2.00 - 2.99 | 2.00 |
| Sedang | B | 3.00 - 3.49 | 3.00 |
| Tinggi | A | 3.50 - 4.00 | 4.00 |

### 5. Publikasi Ilmiah (10%)
| Kategori | Deskripsi | Nilai |
|----------|-----------|-------|
| Tidak Ada | 0 publikasi | 0 |
| Rendah | Publikasi internal | 1 |
| Sedang | Seminar nasional | 2 |
| Tinggi | Jurnal nasional | 3 |
| Sangat Tinggi | Jurnal internasional | 4 |

### 6. Aktivitas Organisasi (5%)
| Kategori | Deskripsi | Nilai |
|----------|-----------|-------|
| Sangat Rendah | Tidak aktif | 0 |
| Rendah | Anggota pasif | 1 |
| Sedang | Anggota aktif | 2 |
| Tinggi | Pengurus | 3 |
| Sangat Tinggi | Ketua/Wakil | 4 |

### 7. Mata Kuliah Mengulang (15%)
| Kategori | Rentang | Nilai |
|----------|---------|-------|
| Sangat Rendah | [0 - 1] | 5 |
| Rendah | [2 - 3] | 4 |
| Sedang | [4 - 5] | 3 |
| Banyak | [6 - 7] | 2 |
| Sangat Banyak | [8+] | 1 |

## 🎓 Predikat Kelulusan
| Predikat | Nilai | Kriteria Minimum |
|----------|--------|------------------|
| Summa Cum Laude | 6 | IPK = 4.00, Lama Studi ≤ 8 semester, Nilai Skripsi A, Tidak ada nilai < B |
| Magna Cum Laude | 5 | IPK 3.79 - 3.99, Lama Studi ≤ 8 semester |
| Cum Laude | 4 | IPK 3.50 - 3.79, Lama Studi ≤ 9 semester |
| Sangat Memuaskan | 3 | IPK 2.75 - 3.50 |
| Memuaskan | 2 | IPK 2.00 - 2.75 |
| Cukup | 1 | IPK ≤ 2.00 |
## 🔄 Proses Perhitungan

2. Proses fuzzifikasi untuk setiap variabel
3. Terapkan aturan fuzzy
4. Lakukan defuzzifikasi
5. Tentukan predikat akhir

## 📝 Catatan Penting
- Semua nilai input harus dalam rentang yang ditentukan
- Bobot variabel dapat disesuaikan sesuai kebijakan institusi
- Predikat akhir ditentukan berdasarkan hasil defuzzifikasi

## 📄 Lisensi
MIT License - lihat file [LICENSE.md](LICENSE.md) untuk detail lengkap.
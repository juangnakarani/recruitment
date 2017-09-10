#### Problem 4
##### Concurrency Task Worker

Implementasi sebuah program untuk merangkum data museum di indonesia berdasarkan lokasi kabupaten/kota, simpan dalam file csv nama yang sesuai.

```
Kota Jakarta Pusat.csv
Kota Malang.csv
``` 

* Untuk sumber informasi gunakan API yang disediakan oleh **Open Data Indonesia** khusus data [museum](http://data.go.id/dataset/museum-indonesia)
* Gunakan *net/http* package untuk mengambil data dari API yang disediakan.
* Olah data mentah dari API jika diperlukan.
* Implementasi *concurrent* process untuk process mengambilan data dari API menggunakan *goroutine* . untuk ilustrasi [ilustrasi](https://talks.golang.org/2012/concurrency.slide)
* Batasi jumlah *concurrent* process, dengan mengaplikasikan *Queue* dan *Worker*.

ilustrasi:
```
grab -concurrent_limit=2 -output=/home/bayu/museum 
```
aplikasi ini terdiri dari dua macam struct:
1. Data -> terdiri dari array of Museum
2. Museum -> detail field dari API

data di ambil menggunakan http client request, hasil dari JSON di decode menjadi string.
hal spesial yang saya temui disini adalah tidak dapat langsung unmarshal json dengan method ```json.Unmarshal``` .
Awalnya saya sudah menyerah :) tapi saya coba untuk membuat server sendiri dengan content-type json yang persis dengan API di museum (file: concurrenserver.go)
Alhasil content json yang saya serve sendiri bisa di decode oleh standart paket golang namun akhirnya ketemu dengan Byte Order Mark untuk decode API museum yg mungkin saja menggunakan BOM unicode.
flow process dari function ```getDataMuseum(kodeKabKota string)``` :
1. get json dari API
2. olahdata untuk bisa di write di paket ```"encoding/csv"```
3. write csv to file sesuai kota masing-masing.

Saya belum bisa menggunakan os argument untuk penerapan concurrency ini






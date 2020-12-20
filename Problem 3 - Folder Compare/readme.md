## **Problem 3 - Folder Compare**

---

Folder `no 1` berisi jawaban dari soal no 1 (membandingkan folder).

Sedangkan folder `no 2` berisi jawaban dari soal no 2 (membandingkan file).

Pada masing-masing folder, telah dibuat 2 folder (sorce dan target) dengan susunan seperti yang ada pada soal.

---

Untuk menjalankan program, dilakukan dengan:

`go run compareFolder.go source/ target/`

- setiap argument harus diakhiri dengan slash

---

**Compare folder** memiliki 2 fungsi utama, yaitu fungsi rekursif untuk mendapatkan nama file yang ada di dalam folder (`recursiveDir`), dan fungsi untuk membandingkan kedua folder tersebut (`searchTheDiff`).

untuk `recursiveDir` digunakanlah `ioutil.ReadDir()` untuk membaca isi sebuah direktori dan menelusuri setiap hasil yang didapatkan menggunakan `for`. Di dalam for tersebut akan dicek apakah item tersebut merupakan direktori (`isDir()`) jika benar maka fungsi `recursiveDir` akan dipanggil lagi dengan path-nya adalah `path sekarang + nama direktori yang ditemukan`. Dan apabila bukan direktori, data tersebut akan disimpan di array string yang pada akhirnya akan berisi data-data file yang ada di folder tsb.

untuk `searchTheDiff` hanya menggunakan 2 perulangan. Perulangan pertama akan mengambil setiap data pada `array pertama`, lalu data tersebut akan dibandingkan dengan data yang ada di `array kedua` dengan perulangan kedua. Apabila data pada `array pertama` tidak ditemukan di `array kedua`, maka program akan mengeluarkan `nama data tersebut + labelnya`.

---

**Compare file** memiliki syntax yang sama dengan `compare folder`, hanya ada penambahan di fungsi `searchTheDiff`,

yaitu apabila data pada `array pertama` ditemukan juga pada `array kedua`, maka akan dilakukan pembacaan pada kedua data tersebut menggunakan `ioutil.ReadFile()` dan hasilnya akan dibandingkan menggunakan `bytes.Equal()`. Apabila hasilnya tidak sama, maka program akan menampilkan `nama data tersebut + "MODIFIED"`.

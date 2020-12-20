## **Problem 2 - Queue**

---

Solusi yang diminta soal adalah solusi yang ada di folder `using interface`,

tetapi saya gagal melewati test case yang ukuran queue:

> `if i < 5 && q.Len() != (i+1) { ... } else if i >= 5 && q.Len() != 5 { ... }`

karena ketika saya membatasi array yang dimasukkan hanya sejumlah `size`, maka akan timbul error di bagian `Contains`

> `if !q.Contains(v) { ... }`

---

**using interface**

Pada `using interface` ini terdapat sebuah struct dan beberapa function yang digunakan untuk queue, yaitu:

- QQ struct -> struct yang terdiri dari 2 data, yiatu:
  - data -> array data yang berupa interface untuk menampung data queue
  - size -> nilai untuk mencatat batas maksimal queue yang sebenarnya akan digunakan di fungsi Push
- New -> membuat queue baru dengan tipe data QQ
- Push -> insert data ke queue
- Keys -> print semua data di queue
- Contains -> mengecek apakah data tertentu ada di dalam queue
- Pop -> mengeluarkan nilai pertama dari queue

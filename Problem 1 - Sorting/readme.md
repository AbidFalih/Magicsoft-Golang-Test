## **Problem 1 - Sorting**

---

masing-masing folder berisi jawaban dari problem 1.

Inputan / masukan dari user berbentuk array dengan nilai integer di dalamnya, seperti:
`[4, 3, 1, 6, 7]`

- setiap angka dipisahkan dengan koma lalu spasi
- diawali dan diakhiri dengan kurung siku

---

Folder no 1 berisi code untuk menvisualisasikan array masukan.

Folder no 2 berisi code untuk menvisualisasikan step-by-step insertion sort.

Folder no 3 berisi code untuk menvisualisasikan step-by-step reverse insertion sort

---

**Visualisasi array** dilakukan dengan menggunakan perulangan di dalam perulangan.

Perulangan pertama dilakukan sebanyak nilai maksimal dari data array + 1 _(digunakan untuk membentuk tinggi visualisasi dan nilai arraynya)_, lalu perulangan kedua _(berada di dalam perulangan pertama)_ dilakukan sejumlah banyaknya data yang ada di array _(mewakili setiap nilai array)_

di setiap perulangan kedua dilakukan pengecekan apakah data tersebut >= nilai max - i. Apabila benar, maka akan ditampilkan `|` untuk membentuk bar, dan apabila salah, program hanya akan menampilkan spasi ` `.

---

**Insertion sort** dilakukan dengan menggunakan perulangan di dalam perulangan.

Perulangan pertama dilakukan sebanyak data yang ada di dalam array - 1 _(karena data pertama tidak perlu dicek)_. Sedangkan perulangan kedua dilakukan sebanyak `i` _(`i` merupakan iterator untuk perulangan pertama)_.

Perulangan kedua ini berfungsi untuk membandingkan nilai pada index tersebut dengan nilai di index sebelumnya. Ketika diurutkan secara asc, maka apabila nilai di index-1 lebih besar dari nilai di index, nilai tersebut akan ditukar.

Untuk pengurutan secara desc, hanya perlu melakukan hal yang sebaliknya _(membalik operator menjadi kurang dari)_

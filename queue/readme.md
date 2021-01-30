#### Solution 2
##### Unique Queue

Terdapat struct bernama List yang digunakan untuk menampung data array, dan jumlah maksimum data yang dapat ditampung

1. Fungsi **New(size int)**

    Fungsi ini digunakan untuk membuat 
    
2. Fungsi **Push(key interface{})**

    Fungsi ini digunakan untuk menambah data kedalam array. Didalam fungsi ini terdapat kondisi dimana jika jumlah data dalam array sekarang sama dengan jumlah maksimum data yang dapat ditampung, maka kita akan memanggil fungsi **pop()** untuk mendelete satu elemen dari array sehingga data baru dapat ditambahkan ke array yang sudah ada
    
3. Fungsi **Pop()**

    Fungsi ini digunakan untuk mendelete suatu data dari array yang sudah ada. Fungsi ini akan mengembalikan data value data yang dihapus. Cara kerja fungsi ini pertama, fungsi ini akan menginisialisasikan array baru yang digunakan untuk memasukkan data array yang baru dan mereplace data array yang lama. Lalu kita akan mengecek jumlah data didalam array, jika jumlah data > 0, maka value yang dikembalikan nantinya adalah data array index ke-0. Proses selanjutnya adalah mengecek jika jumlah data didalam array > 1, maka akan ada perulangan yang berguna untuk memasukkan data ke array yang baru. perulangan ini dimulai dari data array ke-2 (index ke-1)
    
4. Fungsi **Contains(key interface{})**

    Fungsi ini digunakan untuk mengecek apakah suatu value terdapat didalam array yang sudah ada. Didalam fungsi ini terdapat perulangan yang digunakan untuk membandingkan satu demi satu elemen didalam array dengan value yang didapat dari parameter. Jika ada elemen yang valuenya sama, sama fungsi ini akan mengembalikan kondisi true, jika tidak akan mengembalikan kondisi false
    
5. Fungsi **Len()**

    Fungsi ini digunakan untuk menghitung jumlah data didalam array.
    
6. Fungsi **Keys()**

    Fungsi ini mengambil semua data yang ada didalam array
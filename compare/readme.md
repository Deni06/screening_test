#### Solution 3
##### Compare Folder

Contoh compile :
```
Windows:
compare.exe source/ target/

Linux:
compare source/ target/
```

##### Penjelasan Program

Difile compare.go terdapat fungsi **compare(source, target string)**. Fungsi tersebut digunakan untuk mengcompare dua directory dan file dari parameter yang dikirim. Cara kerja fungsi ini adalah :
- Pertama, buat array untuk masing-masing directory source dan directory target. array ini nanti digunakan untuk menampung list file dari masing-masing directory
- Lalu kita menyiapkan satu variabel yang bertipe **map[string]int** dimana variabel tersebut nantinya berguna untuk mengecek data-data yang didelete didirectory source 
- Lalu program akan mengambil list file yang berada dimasing-masing directory source dan target untuk dicompare.
- Ketika program melakukan perulangan untuk mengambil list file, ada kondisi untuk mengecek apakah data yang kita dapatkan itu sebuah file atau directory. Jika data yang kita dapat adalah file dan nama file yang kita dapat != nama masing-masing directory source atau target, proses akan dilanjutkan, jika tidak perulangan untuk data file tersebut akan diskip 
- Jika data yang kita dapat adalah file dan nama file yang kita dapat != nama masing-masing directory source atau target, maka proses akan dilanjutkan ke kondisi selanjutnya dimana jika directory parameternya == "." || directory parameternya == "" maka nama file yang didapatkan akan langsung dimasukkan ke masing-masing array. Jika tidak, maka kita harus melakukan split terhadap nama file yang kita dapat, dengan menghilangkan prefix directory target kita sebelum dimasukkan ke masing-masing array dan jika kita melakukan perulangan untuk mengambil list file directory target, maka kita harus menambahkan index ke variabel bertipe map tadi dengan key nama filenya dan value = 0
- Setelah mendapat semua list file, kita melakukan perulangan untuk directory source terlebih dahulu. didalam perulangan tersebut, kita juga menyiapkan suatu variabel bertipe **boolean** yang berfungsi untuk mengecek apakah ada file yang sama didirectory target nanti
- Didalam perulangan directory source, kita juga melakukan perulangan untuk directory target. Jika value dari perulangan directory target sama dengan value dari perulangan directory source, maka variabel bertipe **boolean** yang kita buat tadi akan menjadi true dan value dari variabel map dengan index yang sama seperti value dari perulangan directory target akan menjadi 1.
- Lalu jika ada file yang sama, kita akan mengcompare isi dari masing masing file dengan cara menyiapkan variabel bertipe array string guna menampung data dari salah satu file yang akan dicompare. Kita juga menyiapkan variabel int yang berfungsi sebagai counter ketika dua data sedang dicompare
- Lalu kita harus mengambil seluruh isi dari 2 file yang kita akan compare
- Untuk file pertama, ketika kita melakukan perulangan untuk mengambil isi dari file pertama, kita juga akan mengpush isi file tersebut ke variabel bertipe array string tadi
- Ketika kita melakukan perulangan untuk mengambil isi dari file kedua, kita juga sekaligus mengcompare dengan file pertama yang sudah kita ambil isinya. Ketika kita melakukan perulangan file kedua, variabel counter bertipe int tadi akan mengcounter jumlah perulangan yang dilakukan. Jika isi kedua file berbeda atau nilai counter == jumlah total isi dari file pertama, program akan melakukan print "(nama file) MODIFIED"
- Lalu perulangan untuk directory target akan dihentikan (selesai)
- Setelah perulangan directory target selesai, kita mengecek value dari variabel bertipe **boolean** tadi yang digunakan untuk mengecek apakah ada file yang sama atau tidak. Jika value variabel tersebut bernilai false, maka program akan melakukan print "(nama file) NEW"
- Setelah semua perulangan diatas selesai, kita melakukan perulangan lagi untuk variabel map yang tadi dibuat. Didalam perulangan tersebut terdapat kondisi dimana jika valuenya == 0 maka akan program akan melakukan print "(nama file) DELETED"
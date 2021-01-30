#### Solution 1
##### Sorting and visualization

Untuk code yang digunakan untuk pembuatan logic program ada difolder util didalam file util.go. Sementara file **createChart.go**, **InsertionSort.go**, dan **ReverseSort.go** digunakan untuk mengcompile program 
 
1. Visualisasi *data array* sederhana dalam bentuk *vertical barcharts*.
    
    Disini saya membuat fungsi **createChart(charts []int)** yang berada difile util.go didalam folder util yang nantinya dipanggil oleh file **createChart.go** untuk bisa dicompile. Difungsi tersebut terdapat 3 perulangan antara lain :
    - Perulangan pertama digunakan untuk mencari nilai terbesar dari array yang didapat dari parameter, dimana nilai terbesar ini digunakan untuk limit perulangan diperulangan selanjutnya
    - Perulangan kedua digunakan untuk membuat row dengan jumlah maksimal row yang dibuat sesuai dengan nilai terbesar yang didapat dari perulangan sebelumnya  
    - Perulangan ketiga berada didalam perulangan kedua, dimana perulangan ini digunakan untuk membuat kolom yang jumlahnya sesuai dengan jumlah data yang ada diarray. Didalam perulangan ini juga terdapat kondisi dimana jika index dari perulangan kedua == 0, maka program akan melakukan print value dari perulangan ketiga. Jika index perulangan kedua > 0 dan index dari perulangan kedua lebih kecil atau sama dengan value dari perulangan ketiga, maka program akan melakukan print "| ". Jika tidak maka program akan melakukan print "  "
        
2. Implementasikan algoritma *insertion sort*, dan memvisualisasikannya 

   Disini saya membuat fungsi **InsertionSort(charts []int)** yang berada difile util.go didalam folder util yang nantinya dipanggil oleh file **InsertionSort.go** untuk bisa dicompile. Difungsi tersebut terdapat dua perulangan antara lain :
   - Perulangan pertama digunakan untuk meloop array yang didapat dari parameter, dimana perulangan dimulai dari data kedua array (index 1)
   - Perulangan kedua berada didalam perulangan pertama, nantinya perulangan ini digunakan untuk membandingkan data satu dengan data lainnya. Didalam perulangan ini terdapat kondisi dimana jika elemen data sekarang lebih kecil dari elemen data sebelumnya, maka akan terjadi pertukaran nilai antar dua elemen ini dan akan memanggil fungsi **createChart(charts []int)** untuk memvisualisasikan setiap langkah/*steps* *sorting*

3. Implementasikan algoritma *reverse sort*, dan memvisualisasikannya 

   Disini saya membuat fungsi **ReverseSort(charts []int)** yang berada difile util.go didalam folder util yang nantinya dipanggil oleh file **ReverseSort.go** untuk bisa dicompile. Difungsi tersebut terdapat dua perulangan antara lain :
   - Perulangan pertama digunakan untuk meloop array yang didapat dari parameter, dimana perulangan dimulai dari data kedua array (index 1)
   - Perulangan kedua berada didalam perulangan pertama, nantinya perulangan ini digunakan untuk membandingkan data satu dengan data lainnya. Didalam perulangan ini terdapat kondisi dimana jika elemen data sekarang lebih besar dari elemen data sebelumnya, maka akan terjadi pertukaran nilai antar dua elemen ini dan akan memanggil fungsi **createChart(charts []int)** untuk memvisualisasikan setiap langkah/*steps* *sorting*
cara membuat unit tests di module golang.

1. beri nama dengan akhiran _test, usahakan nama depannya sama dengan file yang ingin anda test, agar konsisten.

2. beri nama function test dengan diawali Test. untuk menandakan itu adalah fungsi unit test.

3. menjalankan unit test : 
    - go test : untuk menjalankan function test pada folder/package yang ada file _test.

    - go test -v : untuk melihat detail verbose dari function test yang dijalankan.

    - go test -v -run=<nama_function_test> : untuk menjalanakan function test  secara sepesifik.

4. jika ingin menjalankan unit test pada root module gunakan perintah : 
    
    - go test ./...

5. penggagalan/error pada unit test (jangan gunakan panic) :

    - gunakan t.Fail() : untuk menggagalkan unit test apabila terjadi kegagalan tetapi kan melanjutkan kode berikutnya, hingga function dirinya sendiri selesai.

    - gunakan t.FailNow() : untuk menggagalkan function test pada saat itu juga pada function test yang sedang berjalan, tetapi akan menjalankan function test yang lainnya/ function berikutnya.

    - gunakan t.Error() : sama seperti function t.Fail(), hanya saja kita dapat menambahkan pesan kesalahan.

    - gunakan t.Fatal() : sama seperti t.FailNow(), hanya saja kita dapat menambahkan pesan kesalahan.

6. jangan gunakan if else untuk pengecekan, gunakan assert/require dari library testify
    - assert : jika error mengembalikan Fail()
    - require : jika error mengembalikan FailNow()

7. Skip test : 
    - gunakan t.Skip() untuk melewatkan test pada kondisi tertentu.

8. Membuat TestMain : 
    - adalah fungsi yang akan  dijalankan pertama kali sebelum semua test dijalankan.
    - hanya dapat disetting untuk 1 package, jadi disetiap package harus buat TestMain lagi.
    - nama function dari TestMain harus "TestMain" tidak boleh yang lain.

9. membuat Sub Test

10. table test :
    - adalah membuat test dari data yang ada (struct/array)

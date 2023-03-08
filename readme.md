# Self payroll System

Fitur utama dari web service ini adalah employee dapat melakukan penarikan gaji secara mandiri setiap bulannya.

Lebih detail nya, fitur yang harus dikerjakan adalah:

1. Melakukan manage data jabatan berupa operasi CRUD (create, read, update, delete)
2. Melakukan manage data employee berupa operasi CRUD (Create, Read, Update, Delete)
3. Admin dapat melakukan topup balance perusahaan
4. Melakukan penarikan sallary dengan menyertakan employee ID dan secret ID, besaran salary berdasarkan jabatan yang dimiliki oleh tiap employee
5. Terdapat riwayat topup dan pengurangan balance perusahaan

Untuk menjalankan aplikasi lakukan perintah berikut

1. `cp .env.example .env` dan isi sesuai dengan environtment yang ada di PC kalian
2. `go mod tudy && go mod vendor`
3. `go run *.go`

Daftar endpoint ada di postman documenter yang [disertakan](https://documenter.getpostman.com/view/4080490/2s83Ychhk4)

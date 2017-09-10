#### Problem 3
##### Compare Folder

TASK #1. Implementasikan sebuah *program* yang membandingkan isi dari dua direktori melalui *parameter*. 
1. Jika file ada di source dan target, abaikan
2. Jika file ada di source tapi tidak ada di target berikan keterangan **NEW**
3. Jika file tidak ada di source tapi ada di target berikan keterangan **DELETED**

ILUSTRASI: membandingkan *source* direktori terhadap *target* direktori
> $ compare source/ target/

```
source
│   README.md               
│   file001.txt    
│
└───folder1
│   │   file011.txt
│   │   file012.txt
│   │
│   └───subfolder1
│       │   file111.txt
│       │   file112.txt       
│   
└───folder2
    │   file021.txt
    │   file022.txt
```

```
target
│   README.md
│   file001.txt    
│
└───folder1
│   │   file011.txt
│   │   file012.txt
│   │   file013.txt
│   │
│   └───subfolder1
│       │   file111.txt
│       │   file113.txt       
│   
└───folder2
    │   file021.txt
```

OUPUT:
```
folder1/file013.txt DELETED
folder1/subfolder1/file112.txt NEW
folder1/subfolder1/file113.txt DELETED
folder2/file022.txt NEW

```

TASK #2. Modifikasi program #1 untuk *compare* file content untuk rule (1), jika ada perbedaan beri keterangan **MODIFIED** 


Dokumentasi:
aplikasi mengcompare 2 folder yang sudah dibuat sebelumnya yaitu source dan target.
flow process aplikasi ini adalah sebagai berikut:
1. mebandingkan isi file yang berbeda
2, jika hanya terdapat pada folder source saja maka ditandai sebagai NEW dan jika hanya terdapat di folder target ditandai DELETED
3. method yang bekerja pada diff diatas adalah ```getDiffFiles``` dengan parameter berupa slice.
4. untuk MODIFIED menggunakan method yang berbeda yaitu ```getMatchingFiles``` lalu saya bandingkan secara checksum dengan sh1sum.
5. jika hasil checksum berneda maka file tersebut ditandai MODIFIED

berikut gambaran folder tree di local saya:
juang@devbook:~/GOPATH/src/recruitment/compare$ tree
.
├── compare.go
├── readme.md
├── source
│   └── folder1
│       ├── file1.txt
│       ├── file3.txt
│       └── file4.txt
└── target
    └── folder1
        ├── file1.txt
        ├── file2.txt
        ├── file3.txt
        └── file5.txt

4 directories, 9 files

dan berikut hasilnya:
juang@devbook:~/GOPATH/src/recruitment/compare$ go run compare.go 
/folder1/file4.txt NEW 
/folder1/file2.txt DELETED 
/folder1/file5.txt DELETED 
/folder1/file1.txt MODIFIED 


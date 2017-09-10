#### Problem 1
##### Sorting and visualization

1. *Design* dan implementasikan sebuah *program* atau *subprogram* yang akan menampilan visualisasi *data array* sederhana dalam bentuk *vertical barcharts*, dan sebagai tambahan tampilkan setiap nilai data di sumbu *horizontal*.
    
    ```
    INPUT: Numerical array
    [1, 4, 5, 6, 8, 2]

    OUTPUT: Vertical Barcharts

            |   
            |   
          | |  
        | | |   
      | | | |  
      | | | |  
      | | | | |
    | | | | | | 
    1 4 5 6 8 2 

    ```
2. Implementasikan algoritma *insertion sort*, dan gunakan *subprogram* (1) untuk memvisualisasikan setiap langkah/*steps* *sorting* 

    ```
    INPUT: Numerical array

    [1, 4, 5, 6, 8, 2]

    OUTPUT:

    - Sorted array (ascending)
    - Steps visualization

            |   
            |   
          | |  
        | | |   
      | | | |   
      | | | |   
      | | | | | 
    | | | | | | 
    1 4 5 6 8 2 

              | 
              | 
          |   | 
        | |   | 
      | | |   | 
      | | |   | 
      | | | | | 
    | | | | | | 
    1 4 5 6 2 8 

    ... dan seterusnya ...

    ```

3. Modifikasi *subprogram* (2) untuk *reverse sorting* dan lakukan juga visualisasi dengan *subprogram* (1)

Dokumentasi:
1. menggunakan dua algoritma sort insertion yaitu ascending dan descending
2. chart di gambar pada method ```drawBarChart(param []int)``` 

berikut capture run program:
juang@devbook:~/GOPATH/src/recruitment/sorting$ go run sorting.go 
+-draw vertical barchart--+
    |  
    |  
   ||  
  |||  
 ||||  
 ||||  
 ||||| 
|||||| 
145682
+-start sort asc--+
     | 
     | 
   | | 
  || | 
 ||| | 
 ||| | 
 ||||| 
|||||| 
145628
[1 4 5 6 2 8]
     | 
     | 
    || 
  | || 
 || || 
 || || 
 ||||| 
|||||| 
145268
[1 4 5 2 6 8]
     | 
     | 
    || 
   ||| 
 | ||| 
 | ||| 
 ||||| 
|||||| 
142568
[1 4 2 5 6 8]
     | 
     | 
    || 
   ||| 
  |||| 
  |||| 
 ||||| 
|||||| 
124568
[1 2 4 5 6 8]
+--result sort asc--+
[1 2 4 5 6 8]
+--start sort desc-+
     | 
     | 
    || 
   ||| 
  |||| 
  |||| 
| |||| 
|||||| 
214568
[2 1 4 5 6 8]
     | 
     | 
    || 
   ||| 
 | ||| 
 | ||| 
|| ||| 
|||||| 
241568
[2 4 1 5 6 8]
     | 
     | 
    || 
   ||| 
|  ||| 
|  ||| 
|| ||| 
|||||| 
421568
[4 2 1 5 6 8]
     | 
     | 
    || 
  | || 
| | || 
| | || 
||| || 
|||||| 
425168
[4 2 5 1 6 8]
     | 
     | 
    || 
 |  || 
||  || 
||  || 
||| || 
|||||| 
452168
[4 5 2 1 6 8]
     | 
     | 
    || 
|   || 
||  || 
||  || 
||| || 
|||||| 
542168
[5 4 2 1 6 8]
     | 
     | 
   | | 
|  | | 
|| | | 
|| | | 
|||| | 
|||||| 
542618
[5 4 2 6 1 8]
     | 
     | 
  |  | 
| |  | 
|||  | 
|||  | 
|||| | 
|||||| 
546218
[5 4 6 2 1 8]
     | 
     | 
 |   | 
||   | 
|||  | 
|||  | 
|||| | 
|||||| 
564218
[5 6 4 2 1 8]
     | 
     | 
|    | 
||   | 
|||  | 
|||  | 
|||| | 
|||||| 
654218
[6 5 4 2 1 8]
    |  
    |  
|   |  
||  |  
||| |  
||| |  
|||||  
|||||| 
654281
[6 5 4 2 8 1]
   |   
   |   
|  |   
|| |   
||||   
||||   
|||||  
|||||| 
654821
[6 5 4 8 2 1]
  |    
  |    
| |    
|||    
||||   
||||   
|||||  
|||||| 
658421
[6 5 8 4 2 1]
 |     
 |     
||     
|||    
||||   
||||   
|||||  
|||||| 
685421
[6 8 5 4 2 1]
|      
|      
||     
|||    
||||   
||||   
|||||  
|||||| 
865421
[8 6 5 4 2 1]
+--result sort asc--+
[8 6 5 4 2 1]


package main

import (
    "net/http"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "bytes"
    "log"
    "encoding/csv"
    "os"
    "strings"
)
/*
malang: 051800
jakbar: 016200
jaktim: 016400
jaksel: 016300
kseribu: 010100
jakpus: 016000
jakut: 016100
  */
type Data struct{
    Data []Museum `json:"data"`
}

type Museum struct{
    MuseumId string `json:"museum_id"`
    Kode_pengelolaan string `json:"kode_pengelolaan"`
    Nama string `json:"nama"`
    Sdm string `json:"sdm"`
    Alamat_jalan string `json:"alamat_jalan"`
    Desa_kelurahan string `json:"desa_kelurahan"`
    Kecamatan string  `json:"kecamatan"`
    Kabupaten_kota string `json:"kabupaten_kota"`
    Propinsi string `json:"propinsi"`
    Lintang string `json:"lintang"`
    Bujur string `json:"bujur"`
    Koleksi string `json:"koleksi"`
    Sumber_dana string `json:"sumber_dana"`
    Pengelola string `json:"pengelola"`
    Tipe string `json:"tipe"`
    Standar string `json:"standar"`
    Tahun_berdiri string `json:"tahun_berdiri"`
    Bangunan string `json:"bangunan"`
    Luas_tanah string `json:"luas_tanah"`
    Status_kepemilikan string `json:"status_kepemilikan"`
}

func main() {
    c := make(chan int,3)
    go getDataMuseum("051800", c) //malang
    go getDataMuseum("016000", c) //jakpus
    go getDataMuseum("016400", c) //jaktim
    <-c
    os.Exit(0)
    /*
    getDataMuseum("051800")
    getDataMuseum("016000")
    getDataMuseum("016400")
    */
}

func getDataMuseum(kodeKabKota string, c chan int){
    
    url := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET?kode_kab_kota="+kodeKabKota
    //test on local json is run well but somehow get error on museum API
    //url := "http://localhost:8080/museum"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	
    jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	var jsonData Data
    //err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData) // not working!
    err = json.Unmarshal(bytes.TrimPrefix(jsonDataFromHttp, []byte("\xef\xbb\xbf")), &jsonData) // BOM sound is good :)
      if err != nil {
        panic(err)
    }
    data := [][]string{}
    var kotaAsfileName string
    for _, m := range jsonData.Data{
        isi := []string{"museum_id: " + m.MuseumId + ", kode_pengelolaan: " + m.Kode_pengelolaan + ", nama: " + m.Nama + ", sdm: " + m.Sdm + ", alamat_jalan: " + m.Alamat_jalan + ", desa_kelurahan: " + m.Desa_kelurahan + ", kecamatan: " + m.Kecamatan + ", kabupaten_kota: " + m.Kabupaten_kota + ", propinsi: " + m.Propinsi +", lintang: " + m.Lintang + ", bujur: " + m.Bujur + ", koleksi: " + m.Koleksi + ", sumber_dana: " + m.Sumber_dana + ", pengelola: " + m.Pengelola + ", tipe: " + m.Pengelola + ", standar: " + m.Standar + ", tahun_berdiri: " + m.Tahun_berdiri + ", bangunan: " + m.Bangunan + ", luas_tanah: " + m.Luas_tanah + ", status_kepemilikan: " + m.Status_kepemilikan}
        data = append(data, isi)
        kotaAsfileName = m.Kabupaten_kota
    }

    file, err := os.Create(strings.Replace(kotaAsfileName, " ","", -1)+".csv")
    checkError("Cannot create file", err)
    defer file.Close()
    
    w := csv.NewWriter(file)
    defer w.Flush()
    
    fmt.Println("writing to file" + kotaAsfileName)
    for _, value := range data {
        err := w.Write(value)
        checkError("Cannot write to file", err)
    }
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

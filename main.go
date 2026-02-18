package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


type tugas struct{
	Judul string `json:"judul"`
	Body string `json:"body"`
}

func ambilInput(prompt string, scanner *bufio.Scanner) string{
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func TugasDone(namaFile ... string) error{
	os.MkdirAll("Done", 0755)

	for _, nama := range namaFile{
		oldLocation := nama
		newLocation := "Done/" + nama
		err := os.Rename(oldLocation, newLocation)
		
		if err != nil{
			fmt.Println("gagal memindahkan tugas yang sudah dikerjakan", err)
		}
	}
	return nil
}

func ListDone(){
	files, err := os.ReadDir("Done/")
	if err != nil{
		fmt.Println("Gagal membaca tugas yang sudah selesai", err)
		return
	}
	fmt.Println("### Tugas Yang Sudah Selesai")
	found := false

	for _, file:= range files{

		if !file.IsDir() && filepath.Ext(file.Name()) == ".json"{
			fullpath := filepath.Join("Done", file.Name())
			content, err := os.ReadFile(fullpath)
			if err != nil{
				fmt.Println("gagal membaca tugas yang sudah selesai", fullpath, err)
				continue
			}
			var t tugas
			if err := json.Unmarshal(content, &t); err == nil{
				fmt.Printf("- [%s]: %s\n", t.Judul, t.Body)
				found = true
			}
		}
	}
	if !found{
		fmt.Println("Tidak ada tugas yang di temukan")
	}

}


func (t *tugas) save() error{
	todo := t.Judul + ".json"
	jsonData, err := json.MarshalIndent(t, "", "	")
	if err != nil{
		return err
	}
	return os.WriteFile(todo, jsonData, 0600)

}

func load(judul string) (*tugas, error){
	data, err := os.ReadFile(judul + ".json")
	if err != nil{
		return nil, err
	}

	var t tugas
	err = json.Unmarshal(data, &t)
	return &t, err
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Println("#### Menu Tugas ####")
		fmt.Println("1. Tambah/edit Tugas")
		fmt.Println("2. Lihat Tugas Yang Belum")
		fmt.Println("3. Lihat Tugas Yang Sudah")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Tugas Selesai")
		fmt.Println("6. keluar")
		pilihan := ambilInput("pilih menu (1-4): ", scanner)
		switch pilihan{
		case "1":
				judul := ambilInput("Masukan Nama Mata Pelajaran : ", scanner)
				isi := ambilInput("Masukkan deskripsi Tugas: ", scanner)
				t := &tugas{Judul: judul, Body: isi}
				if err := t.save(); err ==nil{
					fmt.Println("Tugas Berhasil Disimpan")
				}
		case "2":
				judul := ambilInput("Masukkan Nama Mata Pelajaran : ", scanner)
				t, err := load(judul)

				if err != nil{
					fmt.Println("Mata Pelajaran Tidak Di Temukan")
				}else{
					
					fmt.Printf("\n[Mata Pelajaran]: %s\n[Deskripsi Tugas]: %s\n", t.Judul, t.Body)
				}
		case "3":
				ListDone()
		case "4":
				judul := ambilInput("Masukan Mata Pelajaran : ", scanner)
				err := os.Remove(judul + ".json")
				if err != nil{
					fmt.Println("Gagal Menghapus Tugas", err)
				}else{
					fmt.Println("Tugas Berhasil Di Hapus")
				}
		case "5":
				judul := ambilInput("Masukkan Mata Pelajaran yang sudah selesai: ", scanner)
				TugasDone(judul + ".json")
				
		case "6":
				fmt.Println("Sayonara!")
				return
		default:
				fmt.Println("Inputan Tidak Valid")
		}
	}
}



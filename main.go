package main

import "fmt"

const Nmax = 1000

type film struct {
	judul     string
	rating    float64
	jmlrating int
}
type datafilm [Nmax]film

func main() {
	fmt.Println("                                                                              ")
	fmt.Println("  ██╗███████╗██████╗ ███████╗                                                 ")
	fmt.Println("  ██║██╔════╝██╔══██╗██╔════╝                                                 ")
	fmt.Println("  ██║█████╗  ██████╔╝███████╗                                                 ")
	fmt.Println("  ██║██╔══╝  ██╔══██╗╚════██║                                                 ")
	fmt.Println("  ██║██║     ██║  ██║███████║                                                 ")
	fmt.Println("  ╚═╝╚═╝     ╚═╝  ╚═╝╚══════╝                                                 ")
	fmt.Println("                                                                              ")
	fmt.Println("--------------------------------------")
	fmt.Println("INDONESIAN FILM RATING SYSTEM                          ")
	fmt.Println("")

	var opsi int
	for opsi != 4 {
		fmt.Printf("1. Rating Film\n2. Cari Film\n3. Top Film\n4. Keluar\n")
		fmt.Println("Gunakan Underscore(_) Untuk Menggantikan Spasi")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&opsi)
		fmt.Println("")

		switch opsi {
		case 1:
			ratingfilm(&daftarfilm, n)
		case 2:
			carifilm(&daftarfilm, &n)
		case 3:
			sortfilm(daftarfilm, n)
		case 4:
			fmt.Println("Terimakasih")
		default:
			fmt.Println("Opsi Tidak Valid")
		}
	}
	fmt.Println("Terimakasih")
}

func ratingfilm(a *datafilm, n int) {
	var targetfilm string
	var ratingtambahan float64
	var ketemu bool

	fmt.Print("Cari judul film: ")
	fmt.Scan(&targetfilm)

	for i := 0; i < n; i++ {
		if a[i].judul == targetfilm {
			fmt.Print("Berikan Rating: ")
			fmt.Scan(&ratingtambahan)
			a[i].jmlrating += 1
			a[i].rating = (a[i].rating*float64(a[i].jmlrating-1) + ratingtambahan) / float64(a[i].jmlrating)
			fmt.Printf("%s %.1f(%d)\n\n", a[i].judul, a[i].rating, a[i].jmlrating)
			ketemu = true
			i = n
		}
	}

	if !ketemu {
		fmt.Println("Film Tidak Ditemukan")
	}
}

func carifilm(a *datafilm, n *int) {
	var targetfilm, filmtambahan string
	var ketemu bool
	var opsi int
	var ratingtambahan float64

	fmt.Print("Cari Film: ")
	fmt.Scan(&targetfilm)

	for i := 0; i < *n; i++ {
		if a[i].judul == targetfilm {
			fmt.Printf("%s %.1f(%d)\n\n", a[i].judul, a[i].rating, a[i].jmlrating)
			ketemu = true
			i = *n
		}
	}

	if ketemu {
		fmt.Println("Film Ditemukan")
		fmt.Printf("1. Cari Film Lain\n2. Kembali Ke Menu\n")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&opsi)
		fmt.Println("")
		if opsi == 1 {
			carifilm(a, n)
		}
	} else {
		fmt.Println("Film Tidak Ditemukan")
		fmt.Printf("1. Tambahkan Film\n2. Kembali Ke Menu\n")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&opsi)
		fmt.Println("")
		if opsi == 1 {
			fmt.Print("Masukan Judul: ")
			fmt.Scan(&filmtambahan)
			a[*n].judul = filmtambahan

			fmt.Print("Masukan Rating: ")
			fmt.Scan(&ratingtambahan)
			a[*n].rating = ratingtambahan

			a[*n].jmlrating = 1
			*n++
			fmt.Println("Film berhasil ditambahkan")
		}
	}
}

func sortfilm(a datafilm, n int) {
	var datasort datafilm = a
	var opsi, max, min int
	var temp film

	fmt.Printf("1. Rating Tertinggi\n2. Rating Terendah\n3. Menu\n")
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&opsi)
	fmt.Println("")

	switch opsi {
	case 1:
		for i := 0; i < n-1; i++ {
			max = i
			for j := i + 1; j < n; j++ {
				if datasort[max].rating < datasort[j].rating || (datasort[max].rating == datasort[j].rating && datasort[max].jmlrating < datasort[j].jmlrating) {
					max = j
				}
			}
			temp = datasort[i]
			datasort[i] = datasort[max]
			datasort[max] = temp
		}
		fmt.Println("--- TOP FILM (RATING TERTINGGI) ---")
		for i := 0; i < n; i++ {
			fmt.Printf("%-4d %-50s %.1f(%d rating)\n", i+1, datasort[i].judul, datasort[i].rating, datasort[i].jmlrating)
		}
		fmt.Println("")
	case 2:
		for i := 0; i < n-1; i++ {
			min = i
			for j := i + 1; j < n; j++ {
				if datasort[min].rating > datasort[j].rating || (datasort[min].rating == datasort[j].rating && datasort[min].jmlrating < datasort[j].jmlrating) {
					min = j
				}
			}
			temp = datasort[i]
			datasort[i] = datasort[min]
			datasort[min] = temp
		}
		fmt.Println("--- TOP FILM (RATING TERENDAH) ---")
		for i := 0; i < n; i++ {
			fmt.Printf("%-4d %-50s %.1f(%d rating)\n", i+1, datasort[i].judul, datasort[i].rating, datasort[i].jmlrating)
		}
		fmt.Println("")
	case 3:
		return
	default:
		fmt.Println("Opsi tidak Valid")
	}
}

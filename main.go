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
}

func ratingfilm(a *datafilm, n int) {
	var targetfilm string
	var ratingtambahan float64
	var ketemu bool
	fmt.Println("Cari judul film (Ketik 0 untuk kembali): ")
	fmt.Scan(&targetfilm)

	if targetfilm == "0" {
		fmt.Println("Kembali ke menu utama")
		return
	}

	for i := 0; i < n; i++ {
		//sequential search//
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

func sortJudulAscending(a *datafilm, n int) {
	for i := 1; i < n; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j].judul > key.judul {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = key
	}
}

func carifilm(a *datafilm, n *int) {
	var targetfilm, filmtambahan string
	var ketemu bool
	var opsi, indexKetemu int
	var ratingtambahan float64
	fmt.Print("Cari Film (Ketik 0 untuk kembali): ")
	fmt.Scan(&targetfilm)

	if targetfilm == "0" {
		fmt.Println("Kembali ke menu utama")
		return
	}

	sortJudulAscending(a, *n)
	low := 0
	high := *n - 1
	//binary search//
	for low <= high {
		mid := (low + high) / 2
		if a[mid].judul == targetfilm {
			ketemu = true
			indexKetemu = mid
			low = high + 1
		} else if a[mid].judul < targetfilm {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if ketemu {
		fmt.Printf("%s %.1f(%d)\n\n", a[indexKetemu].judul, a[indexKetemu].rating, a[indexKetemu].jmlrating)
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
			fmt.Print("Masukan Judul (Ketik 0 untuk batal): ")
			fmt.Scan(&filmtambahan)

			if filmtambahan == "0" {
				fmt.Println("Batal menambahkan film.")
				return
			}

			fmt.Print("Masukan Rating: ")
			fmt.Scan(&ratingtambahan)
			a[*n].judul = filmtambahan
			a[*n].rating = ratingtambahan
			a[*n].jmlrating = 1
			*n++
			fmt.Println("Film berhasil ditambahkan")
		}
	}
}

func sortfilm(a datafilm, n int) {
	var datasort datafilm = a
	var opsi, min int
	var temp film
	fmt.Printf("1. Rating Tertinggi\n2. Rating Terendah\n0. Kembali Ke Menu\n")
	fmt.Print("Pilih Opsi: ")
	fmt.Scan(&opsi)
	fmt.Println("")
	if opsi == 0 {
		return
	}

	switch opsi {
	case 1:
		//insertion sort//
		for i := 1; i < n; i++ {
			key := datasort[i]
			j := i - 1
			for j >= 0 && (datasort[j].rating < key.rating || (datasort[j].rating == key.rating && datasort[j].jmlrating < key.jmlrating)) {
				datasort[j+1] = datasort[j]
				j = j - 1
			}
			datasort[j+1] = key
		}
		fmt.Println("--- TOP FILM (RATING TERTINGGI) ---")
		for i := 0; i < n; i++ {
			fmt.Printf("%-4d %-50s %.1f(%d rating)\n", i+1, datasort[i].judul, datasort[i].rating, datasort[i].jmlrating)
		}
		fmt.Println("")
	case 2:
		//selection sort//
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
	default:
		fmt.Println("Opsi tidak Valid")
	}
}

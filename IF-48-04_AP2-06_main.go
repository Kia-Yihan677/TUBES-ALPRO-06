package main

import "fmt"

const NMAX int = 99

type tubes struct {
	keterampilan string
	minat        string
}

type karir struct {
	nama         string
	industri     string
	keterampilan string
	minat        string
	gaji         int
	kecocokan    float64
}

type tabInt [NMAX]tubes

type betInt [NMAX]karir

func main() {
	var data tabInt
	var karir betInt
	var nData, nIsi, pilih int
	nData = 0

	daftarKarir(&karir)

	for pilih != 8 {
		fmt.Println("--- Aplikasi Rekomendasi Karier ---")
		fmt.Println("1. Tambah Keterampilan")
		fmt.Println("2. Hapus Keterampilan")
		fmt.Println("3. Tambah Minat")
		fmt.Println("4. Hapus Minat")
		fmt.Println("5. Lihat Keahlian dan Minat")
		fmt.Println("6. Lihat Rekomendasi Karier")
		fmt.Println("7. Cari Karir")
		fmt.Println("8. Keluar")
		fmt.Println("-----------------------------------")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&pilih)
		fmt.Println(" ")
		switch pilih {
		case 1:
			tambahK(&data, &nData)
		case 2:
			hapusK(&data, &nData)
		case 3:
			tambahM(&data, &nIsi)
		case 4:
			hapusM(&data, &nIsi)
		case 5:
			lihatMK(data, nData, nIsi)
		case 6:
			lihatR(data, nData, nIsi, &karir)
			if nData > 0 && nIsi > 0 {
				var pilih2 int
				fmt.Println("Urutkan berdasarkan:")
				fmt.Println("1. Gaji")
				fmt.Println("2. Kecocokan")
				fmt.Print("Pilih: ")
				fmt.Scan(&pilih2)

				if pilih2 == 1 {
					sortGaji(&karir)
				} else if pilih2 == 2 {
					sortKecocokan(&karir)
				}
				fmt.Println("\nHasil setelah diurutkan:\n")
				lihatR(data, nData, nIsi, &karir)
			}
		case 7:
			if nData > 0 && nIsi > 0 {
				cariKategori(karir, nData)
			} else {
				fmt.Println("Tambahkan keterampilan dan minat terlebih dahulu")
				fmt.Println()
			}
		}
	}
}

func tambahK(A *tabInt, n *int) {
	var i int
	fmt.Println("--- Pilih Keterampilan ---")
	fmt.Println("1. Programming")
	fmt.Println("2. Creativity")
	fmt.Println("3. Leadership")
	fmt.Println("--------------------------")
	if *n < 1 {
		fmt.Print("Pilih: ")
		fmt.Scan(&i)

		switch i {
		case 1:
			(*A)[*n].keterampilan = "Programming"
		case 2:
			(*A)[*n].keterampilan = "Creativity"
		case 3:
			(*A)[*n].keterampilan = "Leadership"
		}
		fmt.Println("Keterampilan berhasil ditambahkan")
		*n++
	} else {
		fmt.Println("Keterampilan Penuh")
	}
	fmt.Println(" ")
}

func hapusK(A *tabInt, n *int) {
	if *n > 0 {
		*n--
		fmt.Println("Keterampilan berhasil dihapus")
		fmt.Println()
	} else {
		fmt.Println("Keterampilan sudah kosong")
	}
}

func tambahM(A *tabInt, m *int) {
	var i int
	fmt.Println("--- Pilih Minat ---")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Seni & Desain")
	fmt.Println("3. Komunikasi")
	fmt.Println("-------------------")
	if *m < 1 {
		fmt.Print("Pilih: ")
		fmt.Scan(&i)

		switch i {
		case 1:
			(*A)[*m].minat = "Teknologi"
		case 2:
			(*A)[*m].minat = "Seni & Desain"
		case 3:
			(*A)[*m].minat = "Komunikasi"
		}
		fmt.Println("Minat berhasil ditambahkan")
		*m++
	} else {
		fmt.Println("Minat Penuh")
	}
	fmt.Println(" ")
}

func hapusM(A *tabInt, m *int) {
	if *m > 0 {
		*m--
		fmt.Println("Minat berhasil dihapus")
		fmt.Println()
	} else {
		fmt.Println("Minat sudah kosong")
	}
}

func lihatMK(A tabInt, n, m int) {
	var i, j int
	fmt.Println("--- Keterampilan ---")
	for i = 0; i < n; i++ {
		fmt.Println(A[i].keterampilan)
		fmt.Println()
	}
	fmt.Println("--- Minat ---")
	for j = 0; j < m; j++ {
		fmt.Println(A[j].minat)
		fmt.Println()
	}
}

func lihatR(A tabInt, n, m int, B *betInt) {
	var cocok, total float64
	var i int
	fmt.Println("------- Rekomendasi Karier -------")
	if n > 0 && m > 0 {
		for i = 0; i < 5; i++ {
			cocok = 0
			total = 0
			if B[i].keterampilan == A[0].keterampilan {
				cocok++
				total = total + 2
				if B[i].minat == A[0].minat {
					cocok++
				}
				B[i].kecocokan = (cocok / total) * 100
				fmt.Println("Nama Karier   :", B[i].nama)
				fmt.Println("Industri      :", B[i].industri)
				fmt.Println("Gaji Rata-rata:", B[i].gaji, "juta")
				fmt.Printf("Kecocokan     : %.0f%%\n", B[i].kecocokan)
				fmt.Println()
			}
		}
	}
	if n == 0 {
		fmt.Println("Data keterampilan masih kosong")
	}
	if m == 0 {
		fmt.Println("Data minat masih kosong")
	}
	fmt.Println("----------------------------------")
}

func daftarKarir(B *betInt) {
	B[0].nama = "Software Engineer"
	B[0].industri = "Teknologi"
	B[0].keterampilan = "Programming"
	B[0].minat = "Teknologi"
	B[0].gaji = 12

	B[1].nama = "UI/UX Designer"
	B[1].industri = "Teknologi"
	B[1].keterampilan = "Creativity"
	B[1].minat = "Seni & Desain"
	B[1].gaji = 10

	B[2].nama = "Project Manager"
	B[2].industri = "Manajemen"
	B[2].keterampilan = "Leadership"
	B[2].minat = "Komunikasi"
	B[2].gaji = 15

	B[3].nama = "Graphic Designer"
	B[3].industri = "Kreatif"
	B[3].keterampilan = "Creativity"
	B[3].minat = "Seni & Desain"
	B[3].gaji = 8

	B[4].nama = "Front-END Developer"
	B[4].industri = "Teknologi"
	B[4].keterampilan = "Programming"
	B[4].minat = "Seni & Desain"
	B[4].gaji = 14
}

func sortGaji(B *betInt) {
	var i, idx, pass int
	var temp karir

	for pass = 0; pass < 4; pass++ {
		idx = pass
		for i = pass + 1; i < 5; i++ {
			if (*B)[i].gaji > (*B)[idx].gaji {
				idx = i
			}
		}
		temp = (*B)[pass]
		(*B)[pass] = (*B)[idx]
		(*B)[idx] = temp
	}
}

func sortKecocokan(B *betInt) {
	var i, idx, pass int
	var temp karir

	for pass = 0; pass < 4; pass++ {
		idx = pass
		for i = pass + 1; i < 5; i++ {
			if (*B)[i].kecocokan > (*B)[idx].kecocokan {
				idx = i
			}
		}
		temp = (*B)[pass]
		(*B)[pass] = (*B)[idx]
		(*B)[idx] = temp
	}
}

func cariKategori(B betInt, n int) {
	var i, cocok int
	var x string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scan(&x)
	fmt.Println()
	for i = 0; i < 5; i++ {
		if B[i].industri == x {
			fmt.Println("Nama Karier   :", B[i].nama)
			fmt.Println("Industri      :", B[i].industri)
			fmt.Println("Gaji Rata-rata:", B[i].gaji, "juta")
			fmt.Printf("Kecocokan     : %.0f%%\n", B[i].kecocokan)
			fmt.Println()
			cocok++
		}
		if cocok == 0 {
			fmt.Println("Data tidak ada")
		}
	}
}

package main

import "fmt"

type detailStartup struct {
	nama         string
	bidangUsaha  string
	produk       string
	pendanaan    int
	tahunBerdiri int
	jumlahTim    int
	anggotaTim   [10]tim
}
type tim struct {
	nama string
	role string
}

type tabStartup [20]detailStartup

func main() {
	var startup tabStartup
	var jumlahStartup int
	var pilihan int

	startup[0] = detailStartup{"Tokopedia", "Commerce", "E-commerce", 850000000, 2014, 2, [10]tim{{"Meyneza", "CEO"}, {"Kiya", "CMO"}}}
	startup[1] = detailStartup{"Gojek", "Service", "App", 430000000, 2012, 3, [10]tim{{"Nisa", "CEO"}, {"Nadia", "CTO"}, {"Yuri", "CMO"}}}
	startup[2] = detailStartup{"Traveloka", "Commerce", "E-commerce", 120000000, 2018, 4, [10]tim{{"Ara", "CEO"}, {"Naufal", "CTO"}, {"Daffandri", "CMO"}, {"Adam", "CFO"}}}
	startup[3] = detailStartup{"Midtrans", "Fintech", "Payment Gateway", 940000000, 2021, 2, [10]tim{{"Rama", "CEO"}, {"Alvon", "CTO"}}}
	startup[4] = detailStartup{"Ruangguru", "Edu", "Video", 240000000, 2020, 3, [10]tim{{"Sendy", "CEO"}, {"Adira", "CMO"}, {"Akbar", "CTO"}}}
	jumlahStartup = 5

	for pilihan != 8 {
		menu(&pilihan)

		if pilihan == 1 {
			tampilDaftarStartup(startup, jumlahStartup)
		} else if pilihan == 2 {
			cariStartup(startup, jumlahStartup)
		} else if pilihan == 3 {
			urutkanStartup(&startup, &jumlahStartup)
		} else if pilihan == 4 {
			tampilJumlahPerKategori(startup, jumlahStartup)
		} else if pilihan == 5 {
			tambahStartup(&startup, &jumlahStartup)
		} else if pilihan == 6 {
			ubahStartup(&startup, jumlahStartup)
		} else if pilihan == 7 {
			hapusStartup(&startup, &jumlahStartup)
		} else if pilihan == 8 {
			fmt.Println("Terima kasih telah menggunakan program ini")
		}
	}
}

func menu(pilihan *int) {
	fmt.Print("\n\n=====SISTEM MANAJEMEN STARTUP=====",
		" \n1. Tampilkan semua daftar startup",
		" \n2. Cari startup (nama/bidang usaha)",
		" \n3. Urutkan startup (total pendanaan/tahun berdiri)",
		" \n4. Tampilkan jumlah startup per kategori bidang usaha",
		" \n5. Tambahkan data startup",
		" \n6. Ubah data startup",
		" \n7. Hapus data startup",
		" \n8. Exit",
		" \nPILIHAN: ",
	)
	fmt.Scan(pilihan)
}

func tampilDaftarStartup(startup tabStartup, jumlahStartup int) {
	var i int

	if jumlahStartup == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}

	fmt.Println("\n===== DAFTAR STARTUP =====")
	for i = 0; i < jumlahStartup; i++ {
		fmt.Printf("\n%d. Nama: %s\n", i+1, startup[i].nama)
		fmt.Printf("   Bidang Usaha: %s\n", startup[i].bidangUsaha)
		fmt.Printf("   Produk: %s\n", startup[i].produk)
		fmt.Printf("   Pendanaan: Rp %d\n", startup[i].pendanaan)
		fmt.Printf("   Tahun Berdiri: %d\n", startup[i].tahunBerdiri)
		fmt.Println("   Anggota Tim:")

		var j int = 0
		for j < startup[i].jumlahTim {
			fmt.Printf("     %d. %s (%s)\n", j+1, startup[i].anggotaTim[j].nama, startup[i].anggotaTim[j].role)
			j++
		}
		fmt.Println()
	}
}

// rujukan ke materi week 9-10
func cariStartup(startup tabStartup, jumlahStartup int) {
	var keyword string
	var ketemu bool
	var k int
	ketemu = false

	fmt.Print("\nMasukkan kata kunci (nama/bidang usaha): ")
	fmt.Scan(&keyword)

	for k = 0; k < jumlahStartup; k++ {
		if startup[k].nama == keyword || startup[k].bidangUsaha == keyword {
			fmt.Println("========================")
			fmt.Println("Data Startup ke-", k+1)
			fmt.Println("a. Nama Startup:", startup[k].nama)
			fmt.Println("b. Bidang Usaha:", startup[k].bidangUsaha)
			fmt.Println("c. Produk:", startup[k].produk)
			fmt.Println("d. Pendanaan:", startup[k].pendanaan)
			fmt.Println("e. Tahun Berdiri:", startup[k].tahunBerdiri)
			fmt.Println("f. Anggota Tim:")
			var j int
			for j = 0; j < startup[k].jumlahTim; j++ {
				fmt.Printf("     %d. %s (%s)\n", j+1, startup[k].anggotaTim[j].nama, startup[k].anggotaTim[j].role)
			}
			fmt.Println("========================")
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada startup yang cocok dengan kata kunci tersebut.")
	}
}

func urutkanStartup(startup *tabStartup, jumlahStartup *int) {
	if *jumlahStartup <= 1 {
		fmt.Print("Data startup kosong")
		return
	}

	fmt.Print("\nUrutkan berdasarkan:\n1. Pendanaan\n2. Tahun Berdiri\nPilihan: ")
	var pilihan int
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Print("\n\nIngin berdasarkan total pendanaan asecending atau descending?",
			" \n1. Asecending",
			" \n2. Descending",
			" \nPILIHAN: ",
		)
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tampilPendanaan(startup, jumlahStartup, 1)
		} else if pilihan == 2 {
			tampilPendanaan(startup, jumlahStartup, 2)
		}
	} else if pilihan == 2 {
		fmt.Print("\n\nIngin berdasarkan tahun pendirian asecending atau descending?",
			" \n1. Asecending",
			" \n2. Descending",
			" \nPILIHAN: ",
		)
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tampilTahunberdiri(startup, jumlahStartup, 1)
		} else if pilihan == 2 {
			tampilTahunberdiri(startup, jumlahStartup, 2)
		}
	} else {
		fmt.Println("Pilihan tidak tersedia")
	}
}

func tampilPendanaan(startup *tabStartup, jumlahStartup *int, a int) {
	var startups tabStartup
	startups = *startup
	var i, idx, pass int
	var temp detailStartup

	//selection sort rukujan tp modul 12
	if a == 1 {
		pass = 1
		for pass < *jumlahStartup {
			idx = pass - 1
			i = pass
			for i < *jumlahStartup {
				if startups[idx].pendanaan > startups[i].pendanaan {
					idx = i
				}
				i += 1
			}
			temp = startups[pass-1]
			startups[pass-1] = startups[idx]
			startups[idx] = temp
			pass += 1
		}
	} else {
		pass = 1
		for pass < *jumlahStartup {
			idx = pass - 1
			i = pass
			for i < *jumlahStartup {
				if startups[idx].pendanaan < startups[i].pendanaan {
					idx = i
				}
				i += 1
			}
			temp = startups[pass-1]
			startups[pass-1] = startups[idx]
			startups[idx] = temp
			pass += 1
		}
	}
	tampilDaftarStartup(startups, *jumlahStartup)
}

func tampilTahunberdiri(startup *tabStartup, jumlahStartup *int, a int) {
	var startups tabStartup
	startups = *startup
	var i, idx, pass int
	var temp detailStartup

	//selection sort rukujan tp modul 12
	if a == 1 {
		pass = 1
		for pass < *jumlahStartup {
			idx = pass - 1
			i = pass
			for i < *jumlahStartup {
				if startups[idx].tahunBerdiri > startups[i].tahunBerdiri {
					idx = i
				}
				i += 1
			}
			temp = startups[pass-1]
			startups[pass-1] = startups[idx]
			startups[idx] = temp
			pass += 1
		}
	} else {
		pass = 1
		for pass < *jumlahStartup {
			idx = pass - 1
			i = pass
			for i < *jumlahStartup {
				if startups[idx].tahunBerdiri < startups[i].tahunBerdiri {
					idx = i
				}
				i += 1
			}
			temp = startups[pass-1]
			startups[pass-1] = startups[idx]
			startups[idx] = temp
			pass += 1
		}
	}
	tampilDaftarStartup(startups, *jumlahStartup)
}

// Misal startupnya:

// Tokopedia -> Commerce
// Gojek -> Service
// Traveloka -> Commerce

// Maka:

// Pertama: kategori = [“Commerce”], jumlah = [1]
// Kedua: kategori = [“Commerce”, “Service”], jumlah = [1, 1]
// Ketiga: Commerce ditemukan lagi → jumlah[0]++

func tampilJumlahPerKategori(startup tabStartup, jumlahStartup int) {
	var kategori [20]string
	var jumlah [20]int
	var banyakKategori, i, j int
	var ketemu bool

	for i = 0; i < jumlahStartup; i++ {
		ketemu = false
		for j = 0; j < banyakKategori; j++ {
			if startup[i].bidangUsaha == kategori[j] {
				jumlah[j]++
				ketemu = true
			}
		}
		if !ketemu {
			kategori[banyakKategori] = startup[i].bidangUsaha
			jumlah[banyakKategori] = 1
			banyakKategori++
		}
	}

	fmt.Println("\n===== JUMLAH STARTUP PER KATEGORI =====")
	for i = 0; i < banyakKategori; i++ {
		fmt.Printf("%s: %d startup\n", kategori[i], jumlah[i])
	}
}

func tambahStartup(startup *tabStartup, jumlahStartup *int) {
	var jumlahTim, i int

	if *jumlahStartup >= 20 {
		fmt.Println("Kapasitas penyimpanan startup penuh!")
		return
	}

	fmt.Println("\n===== TAMBAH STARTUP BARU =====")

	fmt.Print("Nama Startup: ")
	fmt.Scan(&startup[*jumlahStartup].nama)

	fmt.Print("Bidang Usaha: ")
	fmt.Scan(&startup[*jumlahStartup].bidangUsaha)

	fmt.Print("Produk: ")
	fmt.Scan(&startup[*jumlahStartup].produk)

	fmt.Print("Pendanaan (Rp): ")
	fmt.Scan(&startup[*jumlahStartup].pendanaan)

	fmt.Print("Tahun Berdiri: ")
	fmt.Scan(&startup[*jumlahStartup].tahunBerdiri)

	fmt.Print("Jumlah Anggota Tim: ")
	fmt.Scan(&jumlahTim)

	startup[*jumlahStartup].jumlahTim = jumlahTim
	startup[*jumlahStartup].anggotaTim = [10]tim{}

	for i = 0; i < jumlahTim; i++ {
		fmt.Printf("Anggota Tim %d:\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&startup[*jumlahStartup].anggotaTim[i].nama)

		fmt.Print("Role: ")
		fmt.Scan(&startup[*jumlahStartup].anggotaTim[i].role)

	}

	*jumlahStartup += 1
	fmt.Println("Data startup berhasil ditambahkan!")
}

func ubahStartup(startup *tabStartup, jumlahStartup int) {
	var pilihData int
	var pilihan string
	var kiri, kanan, tengah, idx int
	idx = -1
	kiri = 0
	kanan = jumlahStartup - 1

	tampilDaftarStartup(*startup, jumlahStartup)
	fmt.Print("Data startup keberapa yang ingin diubah: ")
	fmt.Scan(&pilihData)
	pilihData--

	if pilihData > jumlahStartup || pilihData <= 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
// binary search
	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2
		if pilihData == tengah {
			idx = tengah
		} else if pilihData < tengah {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if idx == pilihData {
		fmt.Println("========================")
		fmt.Println("Data Startup ke-", idx+1)
		fmt.Println("a. Nama Startup:", startup[idx].nama)
		fmt.Println("b. Bidang Usaha:", startup[idx].bidangUsaha)
		fmt.Println("c. Produk:", startup[idx].produk)
		fmt.Println("d. Pendanaan:", startup[idx].pendanaan)
		fmt.Println("e. Tahun Berdiri:", startup[idx].tahunBerdiri)
		fmt.Println("f. Anggota Tim:")
		var j int = 0
		for j < startup[idx].jumlahTim {
			fmt.Printf("     %d. %s (%s)\n", j+1, startup[idx].anggotaTim[j].nama, startup[idx].anggotaTim[j].role)
			j++
		}
		fmt.Println("========================")
		fmt.Print("Pilih data yang ingin diubah (a/b/c/d/e/f): ")
		fmt.Scan(&pilihan)

		if pilihan == "a" {
			fmt.Print("Masukan nama startup baru: ")
			fmt.Scan(&startup[idx].nama)
		} else if pilihan == "b" {
			fmt.Print("Masukan bidang usaha baru: ")
			fmt.Scan(&startup[idx].bidangUsaha)
		} else if pilihan == "c" {
			fmt.Print("Masukan produk baru: ")
			fmt.Scan(&startup[idx].produk)
		} else if pilihan == "d" {
			fmt.Print("Masukan pendanaan baru: ")
			fmt.Scan(&startup[idx].pendanaan)
		} else if pilihan == "e" {
			fmt.Print("Masukan tahun berdiri baru: ")
			fmt.Scan(&startup[idx].tahunBerdiri)
		} else if pilihan == "f" {
			var jumlahTimBaru, j int
			fmt.Print("Masukkan jumlah anggota tim baru: ")
			fmt.Scan(&jumlahTimBaru)
			startup[idx].jumlahTim = jumlahTimBaru

			for j = 0; j < jumlahTimBaru; j++ {
				fmt.Printf("Anggota Tim %d:\n", j+1)
				fmt.Print("Nama: ")
				fmt.Scan(&startup[idx].anggotaTim[j].nama)
				fmt.Print("Role: ")
				fmt.Scan(&startup[idx].anggotaTim[j].role)
			}

		} else {
			fmt.Println("Pilihan tidak tersedia.")
			return
		}
		fmt.Println("Data startup berhasil diubah!")
	}
}

func hapusStartup(startup *tabStartup, jumlahStartup *int) {
	var idx, i int

	if *jumlahStartup == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}

	tampilDaftarStartup(*startup, *jumlahStartup)

	fmt.Print("\nPilih nomor startup yang akan dihapus: ")
	fmt.Scan(&idx)
	idx--

	if idx < 0 || idx >= *jumlahStartup {
		fmt.Println("Nomor startup tidak valid.")
		return
	}

	for i = idx; i < *jumlahStartup-1; i++ {
		startup[i] = startup[i+1]
	}
	*jumlahStartup--
	fmt.Println("Data startup berhasil dihapus!")
}

package main

import "fmt"

//tipe bentukan struct menu
type MenuCafe struct {
	idMenu   int
	namaMenu string
	harga    int
	stok     int
}

//tipe bentukan struct pesanan
type Transaksi struct {
	namaMenu   string
	jumlahBeli int
	subTotal   int
}

//array dan counter data
var daftarMenu [100]MenuCafe
var jumlahMenu int = 0

//array untuk menampung struk pesanan aktif
var keranjangBelanja [50]Transaksi
var jumlahItemPesanan int = 0

func inisialisasiData() {
	daftarMenu[0] = MenuCafe{idMenu: 103, namaMenu: "Espresso", harga: 15000, stok: 50}
	daftarMenu[1] = MenuCafe{idMenu: 101, namaMenu: "Croissant", harga: 25000, stok: 20}
	daftarMenu[2] = MenuCafe{idMenu: 105, namaMenu: "Matcha-Latte", harga: 22000, stok: 35}
	daftarMenu[3] = MenuCafe{idMenu: 102, namaMenu: "Red-Velvet", harga: 24000, stok: 15}
	daftarMenu[4] = MenuCafe{idMenu: 104, namaMenu: "Ice-Americano", harga: 18000, stok: 40}
	jumlahMenu = 5
}

func tambahMenu() {
	fmt.Println("\n--- TAMBAH MENU BARU ---")
	fmt.Print("Masukkan ID Menu (Angka): ")
	fmt.Scanln(&daftarMenu[jumlahMenu].idMenu)
	fmt.Print("Masukkan Nama Menu      : ")
	fmt.Scanln(&daftarMenu[jumlahMenu].namaMenu)
	fmt.Print("Masukkan Harga          : ")
	fmt.Scanln(&daftarMenu[jumlahMenu].harga)
	fmt.Print("Masukkan Stok           : ")
	fmt.Scanln(&daftarMenu[jumlahMenu].stok)

	jumlahMenu++
	fmt.Println("\nMenu berhasil ditambahkan!")
}

func tampilkanTabel() {
	fmt.Println("\n=======================================================")
	fmt.Printf("%-8s %-20s %-15s %-10s\n", "ID", "Nama Menu", "Harga", "Stok")
	fmt.Println("=======================================================")
	for i := 0; i < jumlahMenu; i++ {
		fmt.Printf("%-8d %-20s Rp %-12d %-10d\n", 
			daftarMenu[i].idMenu, daftarMenu[i].namaMenu, daftarMenu[i].harga, daftarMenu[i].stok)
	}
	fmt.Println("=======================================================")
}

func buatPesanan() {
	selectionSortID()
	tampilkanTabel()

	var cariID, qty int
	fmt.Println("\n--- INPUT PESANAN PELANGGAN ---")
	fmt.Print("Masukkan ID Menu yang dipesan: ")
	fmt.Scanln(&cariID)

	//cari id menu menggunakan binary search
	low := 0
	high := jumlahMenu - 1
	indexDitemukan := -1

	for low <= high {
		mid := low + (high-low)/2
		if daftarMenu[mid].idMenu == cariID {
			indexDitemukan = mid
			break
		}
		if daftarMenu[mid].idMenu < cariID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	//validasi jika ID menu tidak ditemukan
	if indexDitemukan == -1 {
		fmt.Println("[ERROR] ID Menu tidak ditemukan. Transaksi dibatalkan.")
		return
	}

	//validasi jumlah beli dan stok
	fmt.Printf("Membeli %s. Masukkan Jumlah Porsi: ", daftarMenu[indexDitemukan].namaMenu)
	fmt.Scanln(&qty)

	if qty <= 0 {
		fmt.Println("[ERROR] Jumlah porsi tidak valid.")
		return
	}

	if qty > daftarMenu[indexDitemukan].stok {
		fmt.Printf("[ERROR] Stok tidak cukup! Sisa stok %s saat ini: %d\n", 
			daftarMenu[indexDitemukan].namaMenu, daftarMenu[indexDitemukan].stok)
		return
	}

	//proses pengurangan stok & memasukkan ke keranjang belanja
	daftarMenu[indexDitemukan].stok -= qty
	subTotalItem := daftarMenu[indexDitemukan].harga * qty

	keranjangBelanja[jumlahItemPesanan] = Transaksi{
		namaMenu:   daftarMenu[indexDitemukan].namaMenu,
		jumlahBeli: qty,
		subTotal:   subTotalItem,
	}
	jumlahItemPesanan++

	//cetak struk dan diskon fungsi rekursif
	fmt.Println("\n=======================================================")
	fmt.Println("                  STRUK PEMBAYARAN CAFE                ")
	fmt.Println("=======================================================")
	totalKotor := 0
	for i := 0; i < jumlahItemPesanan; i++ {
		fmt.Printf("%-20s x%-4d  Rp %-10d\n", 
			keranjangBelanja[i].namaMenu, keranjangBelanja[i].jumlahBeli, keranjangBelanja[i].subTotal)
		totalKotor += keranjangBelanja[i].subTotal
	}
	fmt.Println("-------------------------------------------------------")
	
	//hitung diskon
	diskon := hitungDiskonRekursif(totalKotor)
	totalBersih := totalKotor - diskon

	fmt.Printf("Total Kotor           : Rp %d\n", totalKotor)
	fmt.Printf("Diskon 			      : Rp %d\n", diskon)
	fmt.Printf("Total Bayar           : Rp %d\n", totalBersih)
	fmt.Println("=======================================================")
	fmt.Println("        Pesanan Berhasil Disimpan & Stok Dikurangi     ")
	
	//reset keranjang setelah transaksi selesai dicetak
	jumlahItemPesanan = 0 
}

func hitungDiskonRekursif(totalBelanja int) int {
	if totalBelanja < 50000 {
		return 0
	} else {
		return 3000 + hitungDiskonRekursif(totalBelanja-50000)
	}
}

func cariNilaiEkstrim() {
	if jumlahMenu == 0 {
		fmt.Println("Data menu kosong!")
		return
	}
	indexMax := 0
	indexMin := 0
	for i := 1; i < jumlahMenu; i++ {
		if daftarMenu[i].harga > daftarMenu[indexMax].harga {
			indexMax = i
		}
		if daftarMenu[i].harga < daftarMenu[indexMin].harga {
			indexMin = i
		}
	}
	fmt.Println("\n--- ANALISIS NILAI EKSTRIM HARGA ---")
	fmt.Printf("Menu Termahal : %s (Rp %d)\n", daftarMenu[indexMax].namaMenu, daftarMenu[indexMax].harga)
	fmt.Printf("Menu Termurah : %s (Rp %d)\n", daftarMenu[indexMin].namaMenu, daftarMenu[indexMin].harga)
}

func selectionSortID() {
	for i := 0; i < jumlahMenu-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahMenu; j++ {
			if daftarMenu[j].idMenu < daftarMenu[minIdx].idMenu {
				minIdx = j
			}
		}
		daftarMenu[i], daftarMenu[minIdx] = daftarMenu[minIdx], daftarMenu[i]
	}
}

func insertionSortHarga() {
	for i := 1; i < jumlahMenu; i++ {
		key := daftarMenu[i]
		j := i - 1
		for j >= 0 && daftarMenu[j].harga < key.harga {
			daftarMenu[j+1] = daftarMenu[j]
			j = j - 1
		}
		daftarMenu[j+1] = key
	}
}

func sequentialSearchNama(cari string) {
	ditemukan := false
	fmt.Println("\n--- HASIL PENCARIAN (SEQUENTIAL SEARCH) ---")
	for i := 0; i < jumlahMenu; i++ {
		if daftarMenu[i].namaMenu == cari {
			fmt.Printf("Ditemukan! ID: %d | %s | Harga: Rp %d | Stok: %d\n", 
				daftarMenu[i].idMenu, daftarMenu[i].namaMenu, daftarMenu[i].harga, daftarMenu[i].stok)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Printf("Menu dengan nama '%s' tidak ditemukan (Exact match).\n", cari)
	}
}

func binarySearchID(cariID int) {
	selectionSortID()
	low := 0
	high := jumlahMenu - 1
	ditemukan := false

	for low <= high {
		mid := low + (high-low)/2
		if daftarMenu[mid].idMenu == cariID {
			fmt.Println("\n--- HASIL PENCARIAN (BINARY SEARCH) ---")
			fmt.Printf("Ditemukan! ID: %d | %s | Harga: Rp %d | Stok: %d\n", 
				daftarMenu[mid].idMenu, daftarMenu[mid].namaMenu, daftarMenu[mid].harga, daftarMenu[mid].stok)
			ditemukan = true
			break
		}
		if daftarMenu[mid].idMenu < cariID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !ditemukan {
		fmt.Printf("\nMenu dengan ID %d tidak ditemukan.\n", cariID)
	}
}

func main() {
	inisialisasiData()
	var pilihan int

	for {
		fmt.Println("\n==================================================")
		fmt.Println("         SISTEM MANAJEMEN CAFE COFFEE-IN          ")
		fmt.Println("==================================================")
		fmt.Println("1. Buat Pesanan Baru")
		fmt.Println("2. Tambah Menu Baru")
		fmt.Println("3. Daftar Menu")
		fmt.Println("4. Menu Termahal")
		fmt.Println("5. Menu berdasarkan Nama")
		fmt.Println("6. Menu berdasarkan ID")
		fmt.Println("7. Cek Menu Termahal & Termurah")
		fmt.Println("8. Keluar/Exit")
		fmt.Println("==================================================")
		fmt.Print("Pilih menu (1-8): ")
		fmt.Scanln(&pilihan)

		if pilihan == 8 {
			fmt.Println("\nTerima kasih! Program selesai.")
			break
		}

		switch pilihan {
		case 1:
			buatPesanan()
		case 2:
			tambahMenu()
		case 3:
			selectionSortID()
			fmt.Println("\n[INFO] Data diurutkan berdasarkan ID (Ascending).")
			tampilkanTabel()
		case 4:
			insertionSortHarga()
			fmt.Println("\n[INFO] Data diurutkan berdasarkan Harga (Descending).")
			tampilkanTabel()
		case 5:
			var cariNama string
			fmt.Print("\nMasukkan nama menu (Contoh: Espresso / Matcha-Latte): ")
			fmt.Scanln(&cariNama)
			sequentialSearchNama(cariNama)
		case 6:
			var cariID int
			fmt.Print("\nMasukkan ID Menu yang dicari: ")
			fmt.Scanln(&cariID)
			binarySearchID(cariID)
		case 7:
			cariNilaiEkstrim()
		default:
			fmt.Println("\nPilihan tidak valid! Silakan masukkan angka 1-8.")
		}
	}
}
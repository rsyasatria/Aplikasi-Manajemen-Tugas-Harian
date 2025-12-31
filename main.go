package main

import "fmt"

//KAMUS GLOBAL
const NMAX int = 1000

type dataTugas struct {
	idTugas  int
	judul    string
	kategori string
	deadline int
	status   string
}
type tabTugas [NMAX]dataTugas

var data tabTugas
var nData int

//PROGRAM UTAMA
func main() {
	var pilihan int

	for {
		clrScreen()
		menuUtama()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			addTugas(&data, &nData)
		case 2:
			changeTugas(&data, nData)
		case 3:
			delTugas(&data, &nData)
		case 4:
			cetakData(data, nData)
		case 5:
			cariKategori()
		case 6:
			cariDeadline()
		case 7:
			urutDeadline()
		case 8:
			urutStatus()
		case 9:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}

		fmt.Print("\nTekan Enter untuk kembali ke menu...")
		fmt.Scanln(&pilihan)
		fmt.Scanln(&pilihan)
	}
}

func menuUtama() {
	fmt.Println("-----------------------")
	fmt.Println("         MENU          ")
	fmt.Println("-----------------------")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Ubah Tugas")
	fmt.Println("3. Hapus Tugas")
	fmt.Println("4. Cetak Data Tugas")
	fmt.Println("5. Cari Tugas Berdasarkan Kategori")
	fmt.Println("6. Cari Tugas Berdasarkan Deadline")
	fmt.Println("7. Urut Tugas Berdasarkan Deadline")
	fmt.Println("8. Urut Tugas Berdasarkan Status")
	fmt.Println("9. Exit")
	fmt.Println("-----------------------")
}

func clrScreen() {
	fmt.Print("\033[H\033[2J")
}

//PENAMBAHAN, PERUBAHAN, DAN PENGHAPUSAN (CRUD)
func addTugas(A *tabTugas, n *int) {
	if *n < NMAX {
		fmt.Println("\n=== TAMBAH TUGAS ===")
		fmt.Print("Masukkan ID Tugas: ")
		fmt.Scan(&A[*n].idTugas)

		fmt.Print("Masukkan Judul Tugas: ")
		fmt.Scan(&A[*n].judul)

		fmt.Print("Masukkan Kategori: ")
		fmt.Scan(&A[*n].kategori)

		var tgl string
		fmt.Print("Masukkan Deadline (dd/mm/yyyy): ")
		fmt.Scan(&tgl)
		A[*n].deadline = konversiTanggalKeInt(tgl)

		fmt.Print("Masukkan Status (Selesai/Belum): ")
		fmt.Scan(&A[*n].status)

		*n = *n + 1
		fmt.Println("\nTugas berhasil ditambahkan!")
	} else {
		fmt.Println("Data Tugas sudah penuh")
	}
}

func changeTugas(A *tabTugas, n int) {
	var id, idx int
	var tgl string

	if n == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== UBAH TUGAS ===")
	fmt.Print("Masukkan ID Tugas yang akan diubah: ")
	fmt.Scan(&id)

	idx = cariByID(*A, n, id)

	if idx != -1 {
		fmt.Println("Data ditemukan!")
		fmt.Print("Judul baru: ")
		fmt.Scan(&A[idx].judul)
		fmt.Print("Kategori baru: ")
		fmt.Scan(&A[idx].kategori)
		fmt.Print("Deadline baru: ")
		fmt.Scan(&tgl)
		A[idx].deadline = konversiTanggalKeInt(tgl)
		fmt.Print("Status baru (Selesai/Belum): ")
		fmt.Scan(&A[idx].status)
		fmt.Println("Data berhasil diubah!")
	} else {
		fmt.Println("ID Tugas tidak ditemukan!")
	}
}

func delTugas(A *tabTugas, n *int) {
	var id, idx int
	var konfirmasi string

	if *n == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== HAPUS TUGAS ===")
	fmt.Print("Masukkan ID Tugas yang akan dihapus: ")
	fmt.Scan(&id)

	idx = cariByID(*A, *n, id)

	if idx != -1 {
		fmt.Printf("Tugas ditemukan: %s\n", A[idx].judul)
		fmt.Print("Yakin ingin menghapus? (ya/tidak): ")
		fmt.Scan(&konfirmasi)

		if konfirmasi == "ya" {
			hapusDataByIndex(A, n, idx)
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	} else {
		fmt.Println("ID Tugas tidak ditemukan!")
	}
}
func konversiTanggalKeInt(tgl string) int {
	var dd, mm, yy int
	fmt.Sscanf(tgl, "%d/%d/%d", &dd, &mm, &yy)
	return yy*10000 + mm*100 + dd
}

func hapusDataByIndex(A *tabTugas, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func cariByID(A tabTugas, n int, id int) int {
	for i := 0; i < n; i++ {
		if A[i].idTugas == id {
			return i
		}
	}
	return -1
}

func cariKategori() {
	var kategori string

	if nData == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== CARI BERDASARKAN KATEGORI ===")
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&kategori)

	fmt.Println("\nHasil pencarian:")
	found := false
	for i := 0; i < nData; i++ {
		if data[i].kategori == kategori {
			cetakSatuTugas(data[i])
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada tugas dengan kategori tersebut.")
	}
}

// PENCARIAN BINARY SEARCH BERDASARKAN DEADLINE
func cariDeadline() {
	var deadline int

	if nData == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== CARI BERDASARKAN DEADLINE ===")
	fmt.Println("Data akan diurutkan terlebih dahulu...")

	// Sort data dulu untuk binary search
	selectionSort_deadline(&data, nData)

	fmt.Print("Masukkan deadline: ")
	fmt.Scan(&deadline)

	idx := pencarianBerdasarkan_tglDeadline(data, nData, deadline)

	if idx != -1 {
		fmt.Println("\nTugas ditemukan:")
		cetakSatuTugas(data[idx])
	} else {
		fmt.Println("Tidak ada tugas dengan deadline tersebut.")
	}
}

func pencarianBerdasarkan_tglDeadline(A tabTugas, n int, x int) int {
	var left, right, mid int
	var idx int

	left = 0
	right = n - 1
	idx = -1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < A[mid].deadline {
			right = mid - 1
		} else if x > A[mid].deadline {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

// ALGORITMA SORTING
func urutDeadline() {
	if nData == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== URUT BERDASARKAN DEADLINE ===")
	selectionSort_deadline(&data, nData)
	fmt.Println("Data berhasil diurutkan!")
	cetakData(data, nData)
}

func urutStatus() {
	if nData == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== URUT BERDASARKAN STATUS ===")
	insertionSort_status(&data, nData)
	fmt.Println("Data berhasil diurutkan!")
	cetakData(data, nData)
}

func selectionSort_deadline(A *tabTugas, n int) {
	var pass, idx, i int
	var temp dataTugas

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].deadline > A[i].deadline {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func insertionSort_status(A *tabTugas, n int) {
	var pass, i int
	var temp dataTugas

	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]

		for i > 0 && temp.status < A[i-1].status {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

// PENCETAKAN DATA
func cetakData(A tabTugas, n int) {
	if n == 0 {
		fmt.Println("\nTidak ada data tugas!")
		return
	}

	fmt.Println("\n=== DATA TUGAS ===")
	for i := 0; i < n; i++ {
		cetakSatuTugas(A[i])
	}
}

func cetakSatuTugas(tugas dataTugas) {
	fmt.Printf("ID Tugas    : %d\n", tugas.idTugas)
	fmt.Printf("Judul Tugas : %s\n", tugas.judul)
	fmt.Printf("Kategori    : %s\n", tugas.kategori)
	fmt.Printf("Deadline    : %d\n", tugas.deadline)
	fmt.Printf("Status      : %s\n", tugas.status)
	fmt.Println("-----------------------")
}

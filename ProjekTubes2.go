package main

import "fmt"

const USERMAX_M int = 100
const USERMAX_D int = 4

type Data struct {
	Ketua      string
	Anggota    jumlah
	Prodi      string
	Judul      string
	SumberDana string
	Luaran     string
	Tahun      int
	Type       string
}

type arrData [999]Data
type jumlah struct {
	Mahasiswa [USERMAX_M]string
	Dosen     [USERMAX_D]string
}
type frekuensi struct {
	tahun, n int
}
type tabFrek struct {
	arrFrekuensi [1000]frekuensi
	nFrek        int
}

func main() {
	menuAwal()
}
func start() {
	fmt.Println("Selamat datang di aplikasi tri dharma")
	fmt.Println("1. Mulai")
	fmt.Println("2. Keluar")
}
func quit() {
	fmt.Println("Terima kasih sudah menggunakan aplikasi ini")
}
func dataDummy(k *arrData, i *int) {

	// Data 1
	k[*i].Ketua = "John Doe"
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Alice", "Bob"}
	k[*i].Prodi = "Teknik_Informatika"
	k[*i].Judul = "Penelitian A"
	k[*i].SumberDana = "Internal"
	k[*i].Luaran = "Publikasi"
	k[*i].Tahun = 2022
	k[*i].Type = "Data_penelitian"
	*i++

	// Data 2
	k[*i].Ketua = "Jane Smith"
	k[*i].Anggota.Dosen = [USERMAX_D]string{"Dr. David", "Prof. Emily"}
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Charlie"}
	k[*i].Prodi = "Teknik_Elektro"
	k[*i].Judul = "Abdimas B"
	k[*i].SumberDana = "Eksternal"
	k[*i].Luaran = "Seminar"
	k[*i].Tahun = 2023
	k[*i].Type = "Abdimas"
	*i++

	// Data 3
	k[*i].Ketua = "Michael Johnson"
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Eve", "Frank"}
	k[*i].Prodi = "Teknik_Kimia"
	k[*i].Judul = "Penelitian C"
	k[*i].SumberDana = "Internal"
	k[*i].Luaran = "Publikasi"
	k[*i].Tahun = 2022
	k[*i].Type = "Data_penelitian"
	*i++

	// Data 4
	k[*i].Ketua = "Jessica Brown"
	k[*i].Anggota.Dosen = [USERMAX_D]string{"Dr. Sarah", "Prof. Thomas"}
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Grace"}
	k[*i].Prodi = "Teknik_Sipil"
	k[*i].Judul = "Abdimas D"
	k[*i].SumberDana = "Eksternal"
	k[*i].Luaran = "Seminar"
	k[*i].Tahun = 2023
	k[*i].Type = "Abdimas"
	*i++

	// Data 5
	k[*i].Ketua = "Andrew Wilson"
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Hannah", "Isaac"}
	k[*i].Prodi = "Teknik_Mesin"
	k[*i].Judul = "Penelitian E"
	k[*i].SumberDana = "Internal"
	k[*i].Luaran = "Publikasi"
	k[*i].Tahun = 2022
	k[*i].Type = "Data_penelitian"
	*i++

	// Data 6
	k[*i].Ketua = "Olivia Anderson"
	k[*i].Anggota.Dosen = [USERMAX_D]string{"Dr. Robert", "Prof. Samantha"}
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Jack"}
	k[*i].Prodi = "Teknik_Industri"
	k[*i].Judul = "Abdimas F"
	k[*i].SumberDana = "Eksternal"
	k[*i].Luaran = "Seminar"
	k[*i].Tahun = 2023
	k[*i].Type = "Abdimas"
	*i++

	// Data 7
	k[*i].Ketua = "Daniel Thompson"
	k[*i].Anggota.Mahasiswa = [USERMAX_M]string{"Lily", "Matthew"}
	k[*i].Prodi = "Teknik_Perikanan"
	k[*i].Judul = "Penelitian G"
	k[*i].SumberDana = "Internal"
	k[*i].Luaran = "Publikasi"
	k[*i].Tahun = 2022
	k[*i].Type = "Data_penelitian"
	*i++
}
func menuAwal() {
	var k, ks arrData
	var x, y, i, pilih int
	start()
	dataDummy(&k, &i)
	fmt.Print("Silakan pilih : ")
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		for x != 7 {
			fmt.Println("pilih salah satu untuk memulai :")
			fmt.Println("1. Penambahan data")
			fmt.Println("2. Edit data")
			fmt.Println("3. Delete data")
			fmt.Println("4. Tampilkan data")
			fmt.Println("5. Cari data")
			fmt.Println("6. Sorting data")
			fmt.Println("7. Keluar")
			fmt.Print("Pilihan: ")
			fmt.Scan(&x)
			for x != 1 && x != 2 && x != 3 && x != 4 && x != 5 && x != 6 && x != 7 {
				fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
				fmt.Scan(&x)
			}
			if x == 1 {
				penambahan(&k, &i)
			} else if x == 2 {
				editData(&k, i)
			} else if x == 3 {
				deleteData(&k, &i)
			} else if x == 4 {
				printData(k, i)
			} else if x == 5 {
				fmt.Println("Pilih pencarian")
				fmt.Println("1.Berdasarkan tahun")
				fmt.Println("2.Berdasarkan prodi")
				fmt.Println("3.Berdasarkan keduanya")
				fmt.Scan(&y)
				for y != 1 && y != 2 && y != 3 {
					fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
					fmt.Scan(&y)
				}
				if y == 1 {
					cariTahun(k, i)
				} else if y == 2 {
					cariProdi(k, i)
				} else if y == 3 {
					cariKeduanya(k, i)
				}
			} else if x == 6 {
				ks = k
				sorting(&ks, i)
			} else if x == 7 {
				quit()
			}
		}

	} else if pilih == 2 {
		quit()
	}
}
func penambahan(k *arrData, i *int) {
	var dana, n, j, x, m, tipe, luaran int
	j = 0
	temp := *i
	fmt.Print("Masukan berapa data yang mau ditambahkan: ")
	fmt.Scan(&n)
	for *i < n+temp {
		fmt.Print("Silahkan masukan data ketua: ")
		fmt.Scan(&k[*i].Ketua)
		fmt.Println("Pilih tipe data")
		fmt.Println("1. Data_penelitian")
		fmt.Println("2. Abdimas")
		fmt.Print("Pilihan: ")
		fmt.Scan(&tipe)
		for tipe != 1 && tipe != 2 {
			fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
			fmt.Scan(&tipe)
		}
		if tipe == 1 {
			k[*i].Type = "Data_penelitian"
		} else if tipe == 2 {
			k[*i].Type = "Abdimas"
		}
		fmt.Println("Masukan tipe anggota")
		fmt.Println("1. Dosen")
		fmt.Println("2. Mahasiswa")
		fmt.Println("3. Keduanya")
		fmt.Print("Pilihan: ")
		fmt.Scan(&x)
		for x != 1 && x != 2 && x != 3 {
			fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
			fmt.Scan(&x)
		}
		if x == 1 {
			fmt.Print("Masukan berapa banyak anggota dosen yang mau di input (maksimal 4): ")
			fmt.Scan(&m)
			for m != 1 && m != 2 && m != 3 && m != 4 {
				fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
				fmt.Scan(&m)
			}
			fmt.Println("Silahkan masukan data anggota")
			for j < m {
				fmt.Scan(&k[*i].Anggota.Dosen[j])
				j++
			}
		} else if x == 2 {
			fmt.Print("Masukan berapa banyak anggota mahasiswa yang mau di input: ")
			fmt.Scan(&m)
			fmt.Println("Silahkan masukan data anggota")

			for j < m {
				fmt.Scan(&k[*i].Anggota.Mahasiswa[j])
				j++
			}
		} else if x == 3 {
			fmt.Print("Masukan berapa banyak anggota dosen yang mau di input (maksimal 4): ")
			fmt.Scan(&m)
			for m != 1 && m != 2 && m != 3 && m != 4 {
				fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
				fmt.Scan(&m)
			}
			fmt.Println("Silahkan masukan data anggota")
			for j < m {
				fmt.Scan(&k[*i].Anggota.Dosen[j])
				j++
			}
			j = 0
			fmt.Println("Masukan berapa banyak anggota mahasiswa yang mau di input")
			fmt.Scan(&m)
			fmt.Println("Silahkan masukan data anggota")
			for j < m {
				fmt.Scan(&k[*i].Anggota.Mahasiswa[j])
				j++
			}
		}
		j = 0
		fmt.Println("Masukan data Prodi dan Judul Penelitian")
		fmt.Scan(&k[*i].Prodi)
		fmt.Scan(&k[*i].Judul)
		fmt.Println("Pilih sumber dana :")
		fmt.Println("1. Internal")
		fmt.Println("2. Eksternal")
		fmt.Print("Pilihan: ")
		fmt.Scan(&dana)
		for dana != 1 && dana != 2 {
			fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
			fmt.Scan(&dana)
		}
		if dana == 1 {
			k[*i].SumberDana = "Internal"
		} else if dana == 2 {
			k[*i].SumberDana = "Eksternal"
		}
		if tipe == 1 {
			fmt.Println("Pilih data luaran")
			fmt.Println("1.Publikasi")
			fmt.Println("2.Produk")
			fmt.Print("Pilihan: ")
			fmt.Scan(&luaran)
			for luaran != 1 && luaran != 2 {
				fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
				fmt.Scan(&luaran)
			}
			if luaran == 1 {
				k[*i].Luaran = "Publikasi"
			} else if luaran == 2 {
				k[*i].Luaran = "Produk"
			}
		} else if tipe == 2 {
			fmt.Println("Pilih data luaran")
			fmt.Println("1.Publikasi")
			fmt.Println("2.Produk")
			fmt.Println("3.Seminar")
			fmt.Println("4.Pelatihan")
			fmt.Print("Pilihan: ")
			fmt.Scan(&luaran)
			for luaran != 1 && luaran != 2 && luaran != 3 && luaran != 4 {
				fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
				fmt.Scan(&luaran)
			}
			if luaran == 1 {
				k[*i].Luaran = "Publikasi"
			} else if luaran == 2 {
				k[*i].Luaran = "Produk"
			} else if luaran == 3 {
				k[*i].Luaran = "Seminar"
			} else if luaran == 4 {
				k[*i].Luaran = "Pelatihan"
			}
		}
		fmt.Print("Masukan data tahun :")
		fmt.Scan(&k[*i].Tahun)
		*i++
	}
}
func printData(k arrData, i int) {
	for j := 0; j < i; j++ {
		data := k[j]
		fmt.Printf("Data %v\n", j+1)
		fmt.Printf("Tipe Data: %s\n", data.Type)
		fmt.Printf("Ketua: %s\n", data.Ketua)

		fmt.Println("Anggota Dosen:")
		for k := 0; k < USERMAX_D; k++ {
			if data.Anggota.Dosen[k] != "" {
				fmt.Printf("%v. %s\n", k+1, data.Anggota.Dosen[k])
			}
		}

		fmt.Println("Anggota Mahasiswa:")
		for k := 0; k < USERMAX_M; k++ {
			if data.Anggota.Mahasiswa[k] != "" {
				fmt.Printf("%v. %s\n", k+1, data.Anggota.Mahasiswa[k])
			}
		}

		fmt.Printf("Prodi: %s\n", data.Prodi)
		fmt.Printf("Judul: %s\n", data.Judul)
		fmt.Printf("Sumber Dana: %s\n", data.SumberDana)
		fmt.Printf("Luaran: %s\n", data.Luaran)
		fmt.Printf("Tahun: %d\n", data.Tahun)
		fmt.Println()
	}
}
func editData(k *arrData, i int) {
	var idx, x int
	var pilihan int
	var pilih string
	fmt.Print("Masukkan nomor data yang ingin diubah: ")
	fmt.Scan(&idx)
	for idx < 1 || idx > i {
		fmt.Print("Nomor data tidak valid, silakan masukkan kembali: ")
		fmt.Scan(&idx)
	}
	idx--
	data := &k[idx]

	fmt.Println("Data yang ingin diubah:")
	fmt.Println("1. Ketua")
	fmt.Println("2. Anggota Dosen")
	fmt.Println("3. Anggota Mahasiswa")
	fmt.Println("4. Prodi")
	fmt.Println("5. Judul")
	fmt.Println("6. Sumber Dana")
	fmt.Println("7. Luaran")
	fmt.Println("8. Tahun")
	fmt.Println("9. Kembali")
	fmt.Print("Pilih data yang ingin diubah: ")
	fmt.Scan(&pilihan)
	for pilihan != 1 && pilihan != 2 && pilihan != 3 && pilihan != 4 && pilihan != 5 && pilihan != 6 && pilihan != 7 && pilihan != 8 && pilihan != 9 {
		fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
		fmt.Scan(&pilihan)
	}
	if pilihan != 9 {
		if pilihan == 1 {
			fmt.Print("Masukkan nama ketua baru: ")
			fmt.Scan(&data.Ketua)
		} else if pilihan == 2 {
			fmt.Println("Anggota Dosen:")
			for j := 0; j < USERMAX_D; j++ {
				fmt.Printf("Masukkan anggota dosen ke-%d (Ketikan kosong jika data tidak di rubah): ", j+1)
				fmt.Scan(&pilih)
				if pilih != "kosong" {
					data.Anggota.Dosen[j] = pilih
				}
			}
		} else if pilihan == 3 {
			fmt.Println("Anggota Mahasiswa:")
			fmt.Print("Masukan berapa data yang akan di ubah atau ditambah:")
			fmt.Scan(&x)
			for j := 0; j < x; j++ {
				fmt.Printf("Masukkan anggota mahasiswa ke-%d (Ketikan kosong jika data tidak di rubah): ", j+1)
				fmt.Scan(&pilih)
				if pilih != "kosong" {
					data.Anggota.Mahasiswa[j] = pilih
				}
			}
		} else if pilihan == 4 {
			fmt.Print("Masukkan prodi baru: ")
			fmt.Scan(&data.Prodi)
		} else if pilihan == 5 {
			fmt.Print("Masukkan judul baru: ")
			fmt.Scan(&data.Judul)
		} else if pilihan == 6 {
			fmt.Print("Masukkan sumber dana baru (Internal/Eksternal): ")
			fmt.Scan(&data.SumberDana)
		} else if pilihan == 7 {
			fmt.Print("Masukkan luaran baru: ")
			fmt.Scan(&data.Luaran)
		} else if pilihan == 8 {
			fmt.Print("Masukkan tahun baru: ")
			fmt.Scan(&data.Tahun)
		}
		fmt.Println("Data berhasil diubah.")
	} else if pilihan == 9 {
		fmt.Println("Data tidak ada yang diubah")
	}
}
func deleteData(k *arrData, i *int) {
	var idx int
	fmt.Print("Masukkan nomor data yang ingin dihapus: ")
	fmt.Scan(&idx)
	for idx < 1 || idx > *i {
		fmt.Print("Nomor data tidak valid, silakan masukkan kembali: ")
		fmt.Scan(&idx)
	}
	idx--
	for j := idx; j < *i-1; j++ {
		k[j] = k[j+1]
	}

	*i--
}
func cariTahun(k arrData, i int) {
	var tahun int
	var n int = 0
	fmt.Scan(&tahun)
	for j := 0; j <= i; j++ {
		if k[j].Tahun == tahun {
			data := k[j]
			fmt.Printf("Data %v\n", j+1)
			fmt.Printf("Tipe Data: %s\n", data.Type)
			fmt.Printf("Ketua: %s\n", data.Ketua)

			fmt.Println("Anggota Dosen:")
			for k := 0; k < USERMAX_D; k++ {
				if data.Anggota.Dosen[k] != "" {
					fmt.Printf("%v. %s\n", k+1, data.Anggota.Dosen[k])
				}
			}

			fmt.Println("Anggota Mahasiswa:")
			for k := 0; k < USERMAX_M; k++ {
				if data.Anggota.Mahasiswa[k] != "" {
					fmt.Printf("%v. %s\n", k+1, data.Anggota.Mahasiswa[k])
				}
			}

			fmt.Printf("Prodi: %s\n", data.Prodi)
			fmt.Printf("Judul: %s\n", data.Judul)
			fmt.Printf("Sumber Dana: %s\n", data.SumberDana)
			fmt.Printf("Luaran: %s\n", data.Luaran)
			fmt.Printf("Tahun: %d\n", data.Tahun)
			fmt.Println()
			n++
		}
	}
	if n == 0 {
		fmt.Println("Data tidak ditemukan")
	}
}
func cariProdi(k arrData, i int) {
	var prodi string
	var n int = 0
	fmt.Scan(&prodi)
	for j := 0; j <= i; j++ {
		if k[j].Prodi == prodi {
			data := k[j]
			fmt.Printf("Data %v\n", j+1)
			fmt.Printf("Tipe Data: %s\n", data.Type)
			fmt.Printf("Ketua: %s\n", data.Ketua)

			fmt.Println("Anggota Dosen:")
			for k := 0; k < USERMAX_D; k++ {
				if data.Anggota.Dosen[k] != "" {
					fmt.Printf("%v. %s\n", k+1, data.Anggota.Dosen[k])
				}
			}

			fmt.Println("Anggota Mahasiswa:")
			for k := 0; k < USERMAX_M; k++ {
				if data.Anggota.Mahasiswa[k] != "" {
					fmt.Printf("%v. %s\n", k+1, data.Anggota.Mahasiswa[k])
				}
			}

			fmt.Printf("Prodi: %s\n", data.Prodi)
			fmt.Printf("Judul: %s\n", data.Judul)
			fmt.Printf("Sumber Dana: %s\n", data.SumberDana)
			fmt.Printf("Luaran: %s\n", data.Luaran)
			fmt.Printf("Tahun: %d\n", data.Tahun)
			fmt.Println()
			n++
		}
	}
	if n == 0 {
		fmt.Println("Data tidak ditemukan")
	}
}
func cariKeduanya(k arrData, i int) {
	var tahun int
	var prodi string
	var n int = 0
	fmt.Println("Masukan tahun dan prodi")
	fmt.Scan(&tahun, &prodi)
	for j := 0; j <= i; j++ {
		if k[j].Prodi == prodi && k[j].Tahun == tahun {
			data := k[j]
			fmt.Printf("Data %v\n", j+1)
			fmt.Printf("Tipe Data: %s\n", data.Type)
			fmt.Printf("Ketua: %s\n", data.Ketua)

			fmt.Println("Anggota Dosen:")
			for k := 0; k < USERMAX_D; k++ {
				if data.Anggota.Dosen[k] != "" {
					fmt.Printf("%v. %s\n", k+1, data.Anggota.Dosen[k])
				}
			}

			fmt.Println("Anggota Mahasiswa:")
			for k := 0; k < USERMAX_M; k++ {
				if data.Anggota.Mahasiswa[k] != "" {
					fmt.Printf("%v. %s\n", k+1, data.Anggota.Mahasiswa[k])
				}
			}

			fmt.Printf("Prodi: %s\n", data.Prodi)
			fmt.Printf("Judul: %s\n", data.Judul)
			fmt.Printf("Sumber Dana: %s\n", data.SumberDana)
			fmt.Printf("Luaran: %s\n", data.Luaran)
			fmt.Printf("Tahun: %d\n", data.Tahun)
			fmt.Println()
			n++
		}
	}
	if n == 0 {
		fmt.Println("Data tidak ditemukan")
	}
}
func sorting(k *arrData, i int) {
	var pilihan int
	var kemunculan tabFrek
	fmt.Println("Pilih jenis sorting:")
	fmt.Println("1. Berdasarkan tahun")
	fmt.Println("2. Berdasarkan jumlah kegiatan pertahunnya")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	for pilihan != 1 && pilihan != 2 {
		fmt.Print("Pilihan tidak valid, silahkan pilih kembali :")
		fmt.Scan(&pilihan)
	}

	if pilihan == 1 {
		for pass := 0; pass < i-1; pass++ {
			minIdx := pass
			for j := pass + 1; j < i; j++ {
				if k[j].Tahun < k[minIdx].Tahun {
					minIdx = j
				}
			}
			k[pass], k[minIdx] = k[minIdx], k[pass]
		}
		fmt.Println("Data telah diurutkan.")
		printData(*k, i)
	} else if pilihan == 2 {
		for j := 0; j < i; j++ {
			idx := searchTahun(kemunculan, k[j].Tahun)
			if idx == -1 {
				kemunculan.arrFrekuensi[kemunculan.nFrek].tahun = k[j].Tahun
				kemunculan.arrFrekuensi[kemunculan.nFrek].n = 1
				kemunculan.nFrek++
			} else {
				kemunculan.arrFrekuensi[idx].n++
			}
		}
		for pass := 0; pass < kemunculan.nFrek-1; pass++ {
			minIdx := pass
			for j := pass + 1; j < kemunculan.nFrek; j++ {
				if kemunculan.arrFrekuensi[j].n < kemunculan.arrFrekuensi[minIdx].n {
					minIdx = j
				}
			}
			kemunculan.arrFrekuensi[pass], kemunculan.arrFrekuensi[minIdx] = kemunculan.arrFrekuensi[minIdx], kemunculan.arrFrekuensi[pass]
		}
		fmt.Println("Data telah diurutkan.")
		printDataTahun(*k, i, kemunculan)
	}
}

func printDataTahun(k arrData, i int, kemunculan tabFrek) {
	for j := 0; j < kemunculan.nFrek; j++ {
		fmt.Printf("Tahun: %d, Jumlah kegiatan: %d\n", kemunculan.arrFrekuensi[j].tahun, kemunculan.arrFrekuensi[j].n)
		for l := 0; l < i; l++ {
			if k[l].Tahun == kemunculan.arrFrekuensi[j].tahun {
				data := k[l]
				fmt.Printf("Tipe Data: %s\n", data.Type)
				fmt.Printf("Ketua: %s\n", data.Ketua)

				fmt.Println("Anggota Dosen:")
				for k := 0; k < USERMAX_D; k++ {
					if data.Anggota.Dosen[k] != "" {
						fmt.Printf("%v. %s\n", k+1, data.Anggota.Dosen[k])
					}
				}

				fmt.Println("Anggota Mahasiswa:")
				for k := 0; k < USERMAX_M; k++ {
					if data.Anggota.Mahasiswa[k] != "" {
						fmt.Printf("%v. %s\n", k+1, data.Anggota.Mahasiswa[k])
					}
				}

				fmt.Printf("Prodi: %s\n", data.Prodi)
				fmt.Printf("Judul: %s\n", data.Judul)
				fmt.Printf("Sumber Dana: %s\n", data.SumberDana)
				fmt.Printf("Luaran: %s\n", data.Luaran)
				fmt.Printf("Tahun: %d\n", data.Tahun)
				fmt.Println()
			}
		}
		fmt.Println("--------------------")
	}
}

func searchTahun(k tabFrek, x int) int {
	var i int = 0
	for i < k.nFrek && k.arrFrekuensi[i].tahun != x {
		i++
	}
	if i == k.nFrek {
		return -1
	} else {
		return i
	}
}

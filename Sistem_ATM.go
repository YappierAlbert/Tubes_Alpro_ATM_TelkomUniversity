package main

import (
	"fmt"
)

const MAX int = 1000

type nasabah struct {
	nama, riwayat, bank        string
	norek, saldo, pin, nokartu int
}

type bank struct {
	namabank, cabang, alamat string
}

func isiNasabah(n *[MAX]nasabah, x, bnas int) {
	var i int
	for i = bnas; i < bnas+x; i++ {
		fmt.Print("Masukkan Nama Nasabah : ")
		fmt.Scan(&n[i].nama)
		fmt.Print("Masukkan Nomor Rekening Nasabah : ")
		fmt.Scan(&n[i].norek)
		fmt.Print("Masukkan Nomor Kartu Nasabah : ")
		fmt.Scan(&n[i].nokartu)
		fmt.Print("Masukkan Pin Nasabah : ")
		fmt.Scan(&n[i].pin)
		fmt.Print("Masukkan Saldo Nasabah : ")
		fmt.Scan(&n[i].saldo)
		fmt.Print("Masukkan Nama Bank : ")
		fmt.Scan(&n[i].bank)
	}
}

func isiBank(m *[MAX]bank, y, bbank int) {
	var i int
	for i = bbank; i < bbank+y; i++ {
		fmt.Print("Masukkan Nama Bank : ")
		fmt.Scan(&m[i].namabank)
		fmt.Print("Masukkan Cabang Bank : ")
		fmt.Scan(&m[i].cabang)
		fmt.Print("Masukkan Alamat Bank : ")
		fmt.Scan(&m[i].alamat)
	}
}

func tampilBank(m [MAX]bank, y, bbank int) {
	var i int
	for i = 0; i < bbank+y; i++ {
		fmt.Println(m[i].namabank, m[i].cabang, m[i].alamat)
	}
}

func tampilNasabah(n [MAX]nasabah, x, bnas int) {
	var i, j, idx int
	for i = 0; i < bnas+x-1; i++ {
		idx = i
		for j = i + 1; j < bnas+x; j++ {
			if n[j].nama < n[idx].nama {
				idx = j
			}
		}
		if idx != i {
			n[i], n[idx] = n[idx], n[i]
		}
	}
	for i = 0; i < bnas+x; i++ {
		fmt.Println(n[i].nama, n[i].norek, n[i].nokartu, n[i].pin, n[i].saldo, n[i].bank)
	}
}

func editNasabah(n *[MAX]nasabah, x, bnas int) {
	var i, cari int
	var yn string
	fmt.Print("Masukkan No Kartu Nasabah Yang Ingin Diedit : ")
	fmt.Scan(&cari)
	for i = 0; i < bnas+x; i++ {
		if n[i].nokartu == cari {
			fmt.Println(n[i].nama, n[i].norek, n[i].nokartu, n[i].pin, n[i].saldo, n[i].bank)
			fmt.Println("Apakah Data Diatas Benar? : y/n")
			fmt.Scan(&yn)
			if yn == "y" {
				fmt.Print("Masukkan Nama Baru Nasabah : ")
				fmt.Scan(&n[i].nama)
				fmt.Print("Masukkan Nomor Rekening Baru Nasabah : ")
				fmt.Scan(&n[i].norek)
				fmt.Print("Masukkan Nomor Kartu Baru Nasabah : ")
				fmt.Scan(&n[i].nokartu)
				fmt.Print("Masukkan Pin Baru Nasabah : ")
				fmt.Scan(&n[i].pin)
			} else {
				fmt.Println("Edit Dibatalkan!")
			}
			return
		} else {
			fmt.Println("Data Tidak Ditemukan!")
		}
	}
}

func cariNasabahBank(n [MAX]nasabah, m [MAX]bank, x, y, bnas, bbank int) {
	var s string
	var i, j int
	fmt.Print("Masukkan Nama Bank Yang Akan Dicari Nasabahnya : ")
	fmt.Scan(&s)
	for i = 0; i < bnas+x; i++ {
		for j = 0; j < bbank+y; j++ {
			if n[i].bank == s && m[j].namabank == s {
				fmt.Println(n[i].nama, n[i].norek, n[i].nokartu, m[j].namabank, m[j].cabang, m[j].alamat)
			}
		}
	}
}

func cariNasabah(n [MAX]nasabah, x, bnas int) {
	var i, carirek, carikar int
	fmt.Print("Masukkan No Kartu Dan No Rekening Nasabah Yang Ingin Dicari : ")
	fmt.Scan(&carikar, &carirek)
	for i = 0; i < bnas+x; i++ {
		if n[i].nokartu == carikar || n[i].norek == carirek {
			fmt.Println(n[i].nama, n[i].norek, n[i].nokartu, n[i].pin, n[i].saldo, n[i].bank)
		}
	}
}

func hapusNasabah(n *[MAX]nasabah, x, bnas *int) {
	var i, j, cari int
	var yn string
	fmt.Print("Masukkan No Kartu Nasabah Yang Ingin Dihapus : ")
	fmt.Scan(&cari)
	for i = 0; i < *bnas+*x; i++ {
		if n[i].nokartu == cari {
			fmt.Println(n[i].nama, n[i].norek, n[i].nokartu, n[i].pin, n[i].saldo, n[i].bank)
			fmt.Println("Apakah Anda Ingin Menghapus Data Diatas? : y/n")
			fmt.Scan(&yn)
			if yn == "y" {
				for j = i; j < *bnas+*x-1; j++ {
					n[j] = n[j+1]
				}
				*bnas--
				fmt.Println("Data nasabah berhasil dihapus!")
				return
			} else {
				fmt.Println("Delete Dibatalkan!")
				return
			}
		}
	}
}

func cekSaldo(n *[MAX]nasabah, x, bnas, noNas int) {
	var i int
	for i = 0; i < bnas+x; i++ {
		if n[i].norek == noNas {
			n[i].riwayat = "CekSaldo"
			fmt.Println("Saldo Anda : ", n[i].saldo)
		}
	}
}

func tranferRek(n *[MAX]nasabah, x, bnas, noNas int, rekTujuan, jumlahTf *int) {
	var i, j int
	fmt.Print("Masukkan Nomor Rekening Tujuan Anda : ")
	fmt.Scan(rekTujuan)
	fmt.Print("Masukkan Nominal Transfer : ")
	fmt.Scan(jumlahTf)
	for i = 0; i < bnas+x; i++ {
		if n[i].norek == noNas {
			if n[i].saldo < *jumlahTf {
				fmt.Print("Saldo Tidak Cukup")

			} else {
				n[i].saldo -= *jumlahTf
				n[i].riwayat = "Transfer"
				for j = 0; j < bnas+x; j++ {
					if n[j].norek == *rekTujuan {
						n[j].saldo += *jumlahTf

					}
				}
			}
		}
	}
	fmt.Println("Transfer Berhasil")
}

func tarikTunai(n *[MAX]nasabah, x, bnas, noNas int, jumlahTarik *int) {
	var i int
	fmt.Print("Masukkan Nominal Yang Ingin Ditarik : ")
	fmt.Scan(jumlahTarik)
	for i = 0; i < bnas+x; i++ {
		if n[i].norek == noNas {
			if n[i].saldo > *jumlahTarik {
				n[i].saldo -= *jumlahTarik
				n[i].riwayat = "Tarik Tunai"
			} else {
				fmt.Print("Saldo Tidak Cukup")
			}
		}
	}
	fmt.Println("Tarik Tunai Berhasil")
}

func depositTunai(n *[MAX]nasabah, x, bnas, noNas int, jumlahDeposit *int) {
	var i int
	fmt.Print("Masukkan Nominal Yang Ingin Dimasukkan : ")
	fmt.Scan(jumlahDeposit)
	for i = 0; i < bnas+x; i++ {
		if n[i].norek == noNas {
			n[i].saldo += *jumlahDeposit
			n[i].riwayat = "Deposit"
		}
	}
	fmt.Println("Deposit Berhasil")
}

func riwayat(n [MAX]nasabah, x, bnas, noNas, rekTujuan, jumlahTf, jumlahDeposit, jumlahTarik int) {
	var i int
	for i = 0; i < bnas+x; i++ {
		if n[i].norek == noNas {
			if n[i].riwayat == "CekSaldo" {
				fmt.Println(n[i].nama, "Melakukan Cek Saldo")
			} else if n[i].riwayat == "Transfer" {
				fmt.Println(n[i].nama, "Melakukan Transfer Sebanyak", jumlahTf, "Ke Nomor", rekTujuan)
			} else if n[i].riwayat == "Deposit" {
				fmt.Println(n[i].nama, "Melakukan Deposit Sebanyak", jumlahDeposit)
			} else if n[i].riwayat == "Tarik Tunai" {
				fmt.Println(n[i].nama, "Melakukan Tarik Tunai Sebanyak", jumlahTarik)
			} else {
				fmt.Println("Tidak Ada Riwayat")
			}
		}
	}
}

func main() {
	var n [MAX]nasabah
	var m [MAX]bank
	var aksi, x, bnas, y, bbank, noNas, pinNas, salah, cekRek, cekPin int
	var i, j, rekTujuan, jumlahTf, jumlahTarik, jumlahDeposit int
	bnas = 0
	bbank = 0
	for {
		fmt.Println("*************************************")
		fmt.Println("*                                   *")
		fmt.Println("*       Selamat Datang Di ATM       *")
		fmt.Println("*       ======================      *")
		fmt.Println("*       1. Manajemen Nasabah        *")
		fmt.Println("*       2. Manajemen Bank           *")
		fmt.Println("*       3. Transaksi Bank           *")
		fmt.Println("*       4. Keluar                   *")
		fmt.Println("*       ======================      *")
		fmt.Println("*                                   *")
		fmt.Println("*************************************")
		fmt.Print("Silahkan Tekan Angka : ")
		fmt.Scan(&aksi)
		fmt.Println()
		if aksi == 1 {
			fmt.Println("*************************************")
			fmt.Println("*                                   *")
			fmt.Println("*       Manajemen Nasabah           *")
			fmt.Println("*       ======================      *")
			fmt.Println("*       Silahkan Tekan Angka :      *")
			fmt.Println("*       1. Tambah Nasabah           *")
			fmt.Println("*       2. Tampilkan Nasabah        *")
			fmt.Println("*       3. Edit Nasabah             *")
			fmt.Println("*       4. Cari Nasabah             *")
			fmt.Println("*       5. Hapus Nasabah            *")
			fmt.Println("*       6. Kembali                  *")
			fmt.Println("*       ======================      *")
			fmt.Println("*                                   *")
			fmt.Println("*************************************")
			fmt.Print("Silahkan Tekan Angka : ")
			fmt.Scan(&aksi)
			fmt.Println()
			if aksi == 1 {
				if x > 0 {
					bnas = x
				}
				fmt.Print("Berapa banyak nasabah yang ingin ditambahkan : ")
				fmt.Scan(&x)
				if x > 0 {
					isiNasabah(&n, x, bnas)
				}else {
					aksi = 1
				}
			} else if aksi == 2 {
				tampilNasabah(n, x, bnas)
			} else if aksi == 3 {
				editNasabah(&n, x, bnas)
			} else if aksi == 4 {
				cariNasabah(n, x, bnas)
			} else if aksi == 5 {
				hapusNasabah(&n, &x, &bnas)
			} else if aksi == 6 {
				fmt.Println("Terima Kasih Telah Menggunakan Manajemen Nasabah")
			}

		} else if aksi == 2 {
			fmt.Println("*************************************")
			fmt.Println("*                                   *")
			fmt.Println("*       Manajemen Bank              *")
			fmt.Println("*       ======================      *")
			fmt.Println("*       Silahkan Tekan Angka :      *")
			fmt.Println("*       1. Tambah Bank              *")
			fmt.Println("*       2. Tampilkan Bank           *")
			fmt.Println("*       3. Cari Nasabah Bank        *")
			fmt.Println("*       4. Kembali                  *")
			fmt.Println("*       ======================      *")
			fmt.Println("*                                   *")
			fmt.Println("*************************************")
			fmt.Print("Silahkan Tekan Angka : ")
			fmt.Scan(&aksi)
			if aksi == 1 {
				if y > 0 {
					bbank = y
				}
				fmt.Print("Berapa banyak bank yang ingin ditambahkan : ")
				fmt.Scan(&y)
				isiBank(&m, y, bbank)
			} else if aksi == 2 {
				tampilBank(m, y, bbank)
			} else if aksi == 3 {
				cariNasabahBank(n, m, x, y, bnas, bbank)
			} else if aksi == 4 {
				fmt.Println("Terima Kasih Telah Menggunakan Manajemen Bank")
			}
		} else if aksi == 3 {
			fmt.Print("Masukkan Nomor Rekening : ")
			fmt.Scan(&noNas)
			fmt.Print("Masukkan Pin : ")
			fmt.Scan(&pinNas)
			salah = 0
			for i = 0; i < bnas+x; i++ {
				if n[i].norek != noNas && n[i].pin != pinNas {
					for j = 0; j < 3; j++ {
						salah += 1
						if salah >= 3 {
							break
						}
						fmt.Println("Nomor atau PIN salah")
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
					}
				} else if n[i].norek == noNas && n[i].pin == pinNas {
					cekRek = n[i].norek
					cekPin = n[i].pin
					fmt.Println("*************************************")
					fmt.Println("*                                   *")
					fmt.Println("*       Transaksi Bank              *")
					fmt.Println("*       ======================      *")
					fmt.Println("*       Silahkan Tekan Angka :      *")
					fmt.Println("*       1. Cek Saldo                *")
					fmt.Println("*       2. Transfer                 *")
					fmt.Println("*       3. Tarik Tunai              *")
					fmt.Println("*       4. Deposit                  *")
					fmt.Println("*       5. Riwayat                  *")
					fmt.Println("*       6. Kembali                  *")
					fmt.Println("*       ======================      *")
					fmt.Println("*                                   *")
					fmt.Println("*************************************")
					fmt.Print("Silahkan Tekan Angka : ")
					fmt.Scan(&aksi)
					if aksi == 1 {
						salah = 0
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
						if noNas != cekRek {
							fmt.Print("Nomor Rekening Salah! ")
						}
						if pinNas != cekPin {
							fmt.Println("Pin Salah!")
							fmt.Print("Masukkan kembali PIN :")
							fmt.Scan(&pinNas)
							salah += 1
							if salah == 3 {
								fmt.Println("Anda sudah salah memasukkan pin sebanyak 3 kali")
								break
							}
						} else if noNas == cekRek && pinNas == cekPin {
							cekSaldo(&n, x, bnas, noNas)
						}
					} else if aksi == 2 {
						salah = 0
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
						if noNas != cekRek {
							fmt.Print("Nomor Rekening Salah!")
						}
						if pinNas != cekPin {
							fmt.Println("Pin Salah!")
							fmt.Print("Masukkan kembali PIN :")
							fmt.Scan(&pinNas)
							salah += 1
							if salah == 3 {
								break
							}

						} else if noNas == cekRek && pinNas == cekPin {
							tranferRek(&n, x, bnas, noNas, &rekTujuan, &jumlahTf)
						}
					} else if aksi == 3 {
						salah = 0
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
						if noNas != cekRek {
							fmt.Print("Nomor Rekening Salah!")
						}
						if pinNas != cekPin {
							fmt.Println("Pin Salah!")
							fmt.Print("Masukkan kembali PIN :")
							fmt.Scan(&pinNas)
							salah += 1
							if salah == 3 {
								break
							}

						} else if noNas == cekRek && pinNas == cekPin {
							tarikTunai(&n, x, bnas, noNas, &jumlahTarik)
						}
					} else if aksi == 4 {
						salah = 0
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
						if noNas != cekRek {
							fmt.Print("Nomor Rekening Salah!")
						}
						if pinNas != cekPin {
							fmt.Println("Pin Salah!")
							fmt.Print("Masukkan kembali PIN :")
							fmt.Scan(&pinNas)
							salah += 1
							if salah == 3 {
								break
							}
						} else if noNas == cekRek && pinNas == cekPin {
							depositTunai(&n, x, bnas, noNas, &jumlahDeposit)
						}
					} else if aksi == 5 {
						salah = 0
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
						if noNas != cekRek {
							fmt.Print("Nomor Rekening Salah!")
						}
						if pinNas != cekPin {
							fmt.Println("Pin Salah!")
							fmt.Print("Masukkan kembali PIN :")
							fmt.Scan(&pinNas)
							salah += 1
							if salah == 3 {
								break
							}
						} else if noNas == cekRek && pinNas == cekPin {
							riwayat(n, x, bnas, noNas, rekTujuan, jumlahTf, jumlahDeposit, jumlahTarik)
						}

					} else if aksi == 6 {
						fmt.Println("Terima Kasih Sudah Menggunakan ATM")
					}
				} else if n[i].norek == noNas && n[i].pin != pinNas {
					fmt.Println("Pin Salah!")
					for j = 0; j < 3; j++ {
						salah += 1
						if salah >= 3 {
							fmt.Println("Anda sudah salah memasukkan pin sebanyak 3 kali")
							break
						}
						fmt.Println("Nomor atau PIN salah")
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
					}
				} else if n[i].norek != noNas && n[i].pin == pinNas {
					fmt.Println("Nomor Rekening Salah!")
					for j = 0; j < 3; j++ {
						fmt.Print("Masukkan Nomor Rekening : ")
						fmt.Scan(&noNas)
						fmt.Print("Masukkan Pin : ")
						fmt.Scan(&pinNas)
					}
				}
			}

		} else if aksi == 4 {
			break
		} else {
			fmt.Println("Tidak Valid!")
		}
	}
}

package main

import "fmt"

const NMAX = 100 //0...99

type ideProyek [NMAX]string
type kategoriProyek [NMAX]string
type tanggalBuat [NMAX]string
type upvote [NMAX]int

var jumlahData int

func inputIde(id *ideProyek, k *kategoriProyek, t *tanggalBuat, u *upvote) {
  if jumlahData >= NMAX {
    fmt.Println("Kapasitas penyimpanan penuh.")
    return
  }
  fmt.Print("Masukkan Ide Proyek: ")
  fmt.Scan(&id[jumlahData])
  fmt.Print("Masukkan Kategori Proyek: ")
  fmt.Scan(&k[jumlahData])
  fmt.Print("Masukkan Tanggal Buat (YYYY-MM-DD): ")
  fmt.Scan(&t[jumlahData])
  u[jumlahData] = 0
  jumlahData = jumlahData + 1
  fmt.Println("Ide berhasil ditambahkan.")
}

func tampilkanSemuaIde(id ideProyek, k kategoriProyek, t tanggalBuat, u upvote) {
  var i int
  if jumlahData == 0 {
    fmt.Println("Belum ada ide yang disimpan.")
    return
  }
  for i = 0; i < jumlahData; i++ {
    fmt.Println("---------------")
    fmt.Println("Nomor    :", i)
    fmt.Println("Ide      :", id[i])
    fmt.Println("Kategori :", k[i])
    fmt.Println("Tanggal  :", t[i])
    fmt.Println("Upvote   :", u[i])
  }
  fmt.Println("---------------")
}

// idx = 1 nfgakses array, idx 1 = 0, 
func editIde(id *ideProyek, k *kategoriProyek, t *tanggalBuat) {
  var idx int 

  if jumlahData == 0 {
	fmt.Println("Belum ada ide yang disimpan.")
	return
  }

  fmt.Print("Masukkan Nomor ide yang ingin diedit: ")
  fmt.Scan(&idx)
  if idx >= 0 && idx < jumlahData{
    fmt.Print("Masukkan Ide Proyek Baru: ")
    fmt.Scan(&id[idx])
    fmt.Print("Masukkan Kategori Baru: ")
    fmt.Scan(&k[idx])
    fmt.Print("Masukkan Tanggal Baru (YYYY-MM-DD): ")
    fmt.Scan(&t[idx])
    fmt.Println("Data berhasil diubah.")
  } else {
    fmt.Println("Nomor tidak valid.")
  }
}

func hapusIde(id *ideProyek, k *kategoriProyek, t *tanggalBuat, u *upvote) {
  var idx, i int
  fmt.Print("Masukkan Nomor ide yang ingin dihapus: ")
  fmt.Scan(&idx)
  if idx >= 0 && idx < jumlahData {
    for i = idx; i < jumlahData-1; i++ {
      id[i] = id[i+1]
      k[i] = k[i+1]
      t[i] = t[i+1]
      u[i] = u[i+1]
    }
    jumlahData = jumlahData - 1
    fmt.Println("Data berhasil dihapus.")
  } else {
    fmt.Println("Nomor tidak valid.")
  }
}

func upvoteIde(u *upvote) {
  var idx int
  fmt.Print("Masukkan Nomor ide yang ingin di-upvote: ")
  fmt.Scan(&idx)
  if idx >= 0 && idx < jumlahData {
    u[idx] = u[idx] + 1
    fmt.Println("Upvote berhasil.")
  } else {
    fmt.Println("Nomor tidak valid.")
  }
}

func cariIdeSequential(id *ideProyek, k *kategoriProyek, t *tanggalBuat, u *upvote) {
  var hasil [NMAX]int
  var jumlahHasil int
  var i, idx int
  var kategori string
  if jumlahData == 0 {
    fmt.Println("Belum ada data ide.")
    return
  }
  fmt.Print("Masukkan kategori yang ingin dicari: ")
  fmt.Scan(&kategori)
  for i = 0; i < jumlahData; i++ {
    if k[i] == kategori {
      hasil[jumlahHasil] = i
      jumlahHasil++
    }
  }

  if jumlahHasil == 0 {
    fmt.Println("Tidak ada ide dengan kategori tersebut.")
    return
  }

  fmt.Println("Hasil pencarian berdasarkan kategori:")

  for i = 0; i < jumlahHasil; i++ {
    idx = hasil[i]
    fmt.Println("---------------")
    fmt.Println("Nomor    :", i)
    fmt.Println("Ide      :", id[idx])
    fmt.Println("Kategori :", k[idx])
    fmt.Println("Tanggal  :", t[idx])
    fmt.Println("Upvote   :", u[idx])
  }
  fmt.Println("---------------")

  var pilihan int
  fmt.Print("Pilih Nomor ide yang ingin anda pilih: ")
  fmt.Scan(&pilihan)

  if pilihan >= 0 && pilihan < jumlahHasil {
    idx = hasil[pilihan]
    fmt.Println("Ide yang Anda pilih:")
    fmt.Println("Ide      :", id[idx])
    fmt.Println("Kategori :", k[idx])
    fmt.Println("Tanggal  :", t[idx])
    fmt.Println("Upvote   :", u[idx])
  } else {
    fmt.Println("Nomor tidak valid.")
  }
}

func urutkanKategori(k *kategoriProyek) {
  var i, j int
  for i = 0; i < jumlahData-1; i++ {
    for j = 0; j < jumlahData-i-1; j++ {
      if k[j] > k[j+1] {
        k[j], k[j+1] = k[j+1], k[j]
      }
    }
  }
}

func cariIdeBinary(id *ideProyek, k *kategoriProyek, t *tanggalBuat, u *upvote) {
  var kategori string
  var low int
  var high int
  var ketemu int
  var mid int
  var hasil [NMAX]int
  var jumlahHasil int
  var i int
  var pilihan int
  var idx int
  if jumlahData == 0 {
    fmt.Println("Belum ada data ide.")
    return
  }

  fmt.Print("Masukkan kategori yang ingin dicari: ")
  fmt.Scan(&kategori)
  urutkanKategori(k)
  low = 0
  high = jumlahData - 1
  ketemu = -1

  for low <= high {
    mid = (low + high) / 2
    if k[mid] == kategori {
      ketemu = mid
      high = mid - 1
    } else if k[mid] < kategori {
      low = mid + 1
    } else {
      high = mid - 1
    }
  }

  if ketemu == -1 {
    fmt.Println("Tidak ada ide dengan kategori tersebut.")
    return
  }

  jumlahHasil = 0

  i = ketemu
  for i >= 0 && k[i] == kategori {
    i--
  }
  i++
  for i < jumlahData && k[i] == kategori {
    hasil[jumlahHasil] = i
    jumlahHasil++
    i++
  }

  fmt.Println("Hasil pencarian berdasarkan kategori (Binary Search):")

  for i = 0; i < jumlahHasil; i++ {
    idx = hasil[i]
    fmt.Println("---------------")
    fmt.Println("Nomor    :", i)
    fmt.Println("Ide      :", id[idx])
    fmt.Println("Kategori :", k[idx])
    fmt.Println("Tanggal  :", t[idx])
    fmt.Println("Upvote   :", u[idx])
  }
  fmt.Println("---------------")

  fmt.Print("Pilih nomor ide yang ingin dilihat lebih lanjut: ")
  fmt.Scan(&pilihan)
  if pilihan >= 0 && pilihan < jumlahHasil {
    idx = hasil[pilihan]
    fmt.Println("Ide yang Anda pilih:")
    fmt.Println("Ide      :", id[idx])
    fmt.Println("Kategori :", k[idx])
    fmt.Println("Tanggal  :", t[idx])
    fmt.Println("Upvote   :", u[idx])
  } else {
    fmt.Println(" Nomor tidak valid.")
  }
}


func tampilkanUrutanUpvoteTerbanyak(id *ideProyek, k *kategoriProyek, t *tanggalBuat, u *upvote) {
  var tempIde [NMAX]string
  var tempKat [NMAX]string
  var tempTgl [NMAX]string
  var tempUp [NMAX]int
  var i int
  var j, maxIdx int
  var tStr string
  var tInt int
  if jumlahData == 0 {
    fmt.Println("Belum ada data ide.")
    return
  }
  for i = 0; i < jumlahData; i++ {
    tempIde[i] = id[i]
    tempKat[i] = k[i]
    tempTgl[i] = t[i]
    tempUp[i] = u[i]
  }

  for i = 0; i < jumlahData-1; i++ {
    maxIdx = i
    for j = i + 1; j < jumlahData; j++ {
      if tempUp[j] > tempUp[maxIdx] {
        maxIdx = j
      }
    }
    if maxIdx != i {
      tInt = tempUp[i]
      tempUp[i] = tempUp[maxIdx]
      tempUp[maxIdx] = tInt

      tStr = tempIde[i]
      tempIde[i] = tempIde[maxIdx]
      tempIde[maxIdx] = tStr

      tStr = tempKat[i]
      tempKat[i] = tempKat[maxIdx]
      tempKat[maxIdx] = tStr

      tStr = tempTgl[i]
      tempTgl[i] = tempTgl[maxIdx]
      tempTgl[maxIdx] = tStr
    }
  }

  fmt.Println(" Urutan Ide Berdasarkan Upvote Terbanyak:")
  for i = 0; i < jumlahData; i++ {
    fmt.Println("---------------")
    fmt.Println("Ide      :", tempIde[i])
    fmt.Println("Kategori :", tempKat[i])
    fmt.Println("Tanggal  :", tempTgl[i])
    fmt.Println("Upvote   :", tempUp[i])
  }
  fmt.Println("---------------")
}

func idePopuler(id *ideProyek, k *kategoriProyek, t *tanggalBuat, u *upvote) {
  var max int
  var idx, i int
  if jumlahData == 0 {
    fmt.Println("Tidak ada ide.")
    return
  }
  max = u[0]
  idx = 0
  for i = 1; i < jumlahData; i++ {
    if u[i] > max {
      max = u[i]
      idx = i
    }
  }
  fmt.Println("Ide Terpopuler:")
  fmt.Println("Ide      :", id[idx])
  fmt.Println("Kategori :", k[idx])
  fmt.Println("Tanggal  :", t[idx])
  fmt.Println("Upvote   :", u[idx])
}




func menu() {
  var id ideProyek
  var k kategoriProyek
  var t tanggalBuat
  var u upvote

  var pilihan int = -1
  for pilihan != 0 {
    fmt.Println(`   _____         .__  .__ __                  .__  __________                             .__         .__           .__    .___       `)
    fmt.Println(`  /  _  \ ______ |  | |__|  | _______    _____|__| \______   \ ____   ____    ____   ____ |  |   ____ |  | _____    |__| __| _/____   `)
    fmt.Println(` /  /_\  \\____ \|  | |  |  |/ /\__  \  /  ___/  |  |     ___// __ \ /    \  / ___\_/ __ \|  |  /  _ \|  | \__  \   |  |/ __ |/ __ \  `)
    fmt.Println(`/    |    \  |_> >  |_|  |    <  / __ \_\___ \|  |  |    |   \  ___/|   |  \/ /_/  >  ___/|  |_(  <_> )  |__/ __ \_ |  / /_/ \  ___/  `)
    fmt.Println(`\____|__  /   __/|____/__|__|_ \(____  /____  >__|  |____|    \___  >___|  /\___  / \___  >____/\____/|____(____  / |__\____ |\___  > `)
    fmt.Println(`        \/|__|                \/     \/     \/                    \/     \//_____/      \/                      \/          \/    \/  `)
    fmt.Println(`             __________               .__                 __                        .__                                               `)
    fmt.Println(`             \______   \____________  |__| ____   _______/  |________  ____   _____ |__| ____    ____                                `)
    fmt.Println(`              |    |  _/\_  __ \__  \ |  |/    \ /  ___/\   __\_  __ \/  _ \ /     \|  |/    \  / ___\                                `)
    fmt.Println(`              |    |   \ |  | \// __ \|  |   |  \\___ \  |  |  |  | \(  <_> )  Y Y  \  |   |  \/ /_/  >                               `)
    fmt.Println(`              |______  / |__|  (____  /__|___|  /____  > |__|  |__|   \____/|__|_|  /__|___|  /\___  /                                `)
    fmt.Println(`                     \/             \/        \/     \/                           \/        \//_____/                                `)
    fmt.Println()
    fmt.Println("                              🚀 SELAMAT DATANG DI APLIKASI PENGELOLA IDE BRAINSTORMING 🚀                                            ")
    fmt.Println("=======================================================================================================================================")
    fmt.Println("1. Tambah Ide")
    fmt.Println("2. Tampilkan Semua Ide")
    fmt.Println("3. Edit Ide")
    fmt.Println("4. Hapus Ide")
    fmt.Println("5. Upvote Ide")
    fmt.Println("6. Cari Ide Sequential")
    fmt.Println("7. Cari Ide Binary")
    fmt.Println("8. Urutan Ide Berdasarkan Upvote Terbanyak")
    fmt.Println("9. Ide Populer")
    fmt.Println("0. Keluar")
    fmt.Print("Pilih menu: ")
    fmt.Scan(&pilihan)

    switch pilihan {
    case 1:
      inputIde(&id, &k, &t, &u)
    case 2:
      tampilkanSemuaIde(id, k, t, u)
    case 3:
	if jumlahData == 0 {
		fmt.Println("Belum ada ide yang disimpan.")
	} else {
		 editIde(&id, &k, &t)
	}
     
    case 4:
	if jumlahData == 0 {
		fmt.Println("Belum ada ide yang disimpan.")
	} else {
		hapusIde(&id, &k, &t, &u)
	}
      
    case 5:
	if jumlahData == 0 {
		fmt.Println("Belum ada ide yang disimpan.")
	} else {
		upvoteIde(&u)
	}
    case 6:
      cariIdeSequential(&id, &k, &t, &u)	
    case 7:
      cariIdeBinary(&id, &k, &t, &u)
    case 8:
      tampilkanUrutanUpvoteTerbanyak(&id, &k, &t, &u)
    case 9:
      idePopuler(&id, &k, &t, &u)
    case 0:
      fmt.Println("Terima kasih telah menggunakan aplikasi.")
    default:
      fmt.Println("Pilihan tidak valid.")
    }
  }
}

func main() {
  menu()
}

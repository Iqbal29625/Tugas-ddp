package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Input nama pengguna
	fmt.Print("Masukkan nama Anda: ")
	reader := bufio.NewReader(os.Stdin)
	nama, _ := reader.ReadString('\n')

	// Kumpulan soal, pilihan jawaban, dan jawaban yang benar
	soal := []string{
		"Meong...Meong.. merupakan suara dari hewan?",
		"Gitar dimainkan dengan cara di?",
		"Tugas mahasiswa adalah?",
	}
	pilihan := [][]string{
		{"1. Kucing", "2. Buaya", "3. Kerbau", "4. Angsa"},
		{"1. Petik", "2. Injak", "3. Banting", "4. Tiup"},
		{"1. Belajar dengan sungguh - sugguh", "2. Rokokan dikantin dan bolos kelas", "3. Ikut demo papua merdeka", "4. Mengedarkan narkoba di lingkungan kampus agar dapat uang jajan"},
	}
	jawabanBenar := []string{"1", "1", "3"}

	// Menghitung jumlah soal
	jumlahSoal := len(soal)

	// Inisialisasi hasil jawaban
	jawabanBenarCount := 0
	jawabanSalahCount := 0

	// Mengolah setiap soal
	for i := 0; i < jumlahSoal; i++ {
		fmt.Printf("\nSoal %d: %s\n", i+1, soal[i])
		for _, pilihanJawaban := range pilihan[i] {
			fmt.Println(pilihanJawaban)
		}
		fmt.Print("Jawaban Anda: ")
		jawabanPengguna, _ := reader.ReadString('\n')

		// Membandingkan jawaban pengguna dengan jawaban yang benar
		jawabanPengguna = strings.TrimSpace(jawabanPengguna)
		if strings.EqualFold(jawabanPengguna, jawabanBenar[i]) {
			jawabanBenarCount++
		} else {
			jawabanSalahCount++
		}
	}
	// Menampilkan keterangan hasil
	fmt.Printf("\n%s, berikut adalah hasilnya:\n", strings.TrimSpace(nama))
	fmt.Println("Jumlah Soal: ", jumlahSoal)
	fmt.Println("Jawaban Benar: ", jawabanBenarCount)
	fmt.Println("Jawaban Salah: ", jawabanSalahCount)
}

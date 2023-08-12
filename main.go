package autotestGo

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// решение задачи (пример)
func solve(inFile string, outFile string) {

	// Открытие файла на чтение
	row := NewReadFile(inFile)
	// Открытие файла на запись
	file := NewOpenFile(outFile)

	n := row.ReadStringAsNumber()

	for i := 0; i < n; i++ {
		m := make(map[int]int)
		sum := 0

		count := row.ReadStringAsNumber()

		input := row.ReadString()
		reader := NewStringReader(input)

		for j := 0; j < count; j++ {
			x := reader.ReadNumber()

			m[x]++
		}

		for key, value := range m {
			sum = sum + value/3*2*key + value%3*key
		}

		file.WritelnNumber(sum)
	}
}

// StringReader сканнер строки
type StringReader struct {
	str string
	i   int
	len int
}

// NewStringReader конструктор, возвращает сканнер по строке
func NewStringReader(str string) *StringReader {

	return &StringReader{
		str,
		0,
		len(str),
	}
}

// ReadWord читает слово из строки до первого пробела
func (s *StringReader) ReadWord() string {

	begin := s.i

	for i := s.i; i < s.len; i++ {
		if s.str[i] == ' ' {
			s.i = i + 1
			break
		} else if i == s.len-1 && s.str[i] != ' ' {
			return s.str[begin:s.len]
		}
	}

	return s.str[begin : s.i-1]
}

// ReadNumber читает число из строки до первого пробела
func (s *StringReader) ReadNumber() int {
	begin := s.i

	for i := s.i; i < s.len; i++ {
		if s.str[i] == ' ' {
			s.i = i + 1
			break
		} else if i == s.len-1 && s.str[i] != ' ' {
			number, err := strconv.Atoi(s.str[begin:s.len])
			if err != nil {
				log.Fatalln(err)
			}
			return number
		}
	}

	number, err := strconv.Atoi(s.str[begin : s.i-1])
	if err != nil {
		log.Fatalln(err)
	}
	return number
}

// FileR хранит сканнер по открытому файлу
type FileR struct {
	data *bufio.Scanner
}

// NewReadFile конструктор, возвращает открытый файл
func NewReadFile(str string) *FileR {
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var maxSize int
	scanner := bufio.NewScanner(file)
	maxSize = int(info.Size())
	buffer := make([]byte, 0, maxSize)
	scanner.Buffer(buffer, maxSize)

	/*
		sc := bufio.NewScanner(file)
		const maxCapacity = 1024 * 1024 * 8
		buf := make([]byte, maxCapacity)
		sc.Buffer(buf, maxCapacity)
	*/

	return &FileR{scanner}
}

// ReadString читает строку из файла
func (f *FileR) ReadString() string {
	f.data.Scan()
	return f.data.Text()
}

// ReadStringAsNumber читает строку из файла и конвертирует в число
// использовать, когда подразумевается, что в строке находится только одно число
func (f *FileR) ReadStringAsNumber() int {
	f.data.Scan()

	number, err := strconv.Atoi(f.data.Text())

	if err != nil {
		log.Fatal(err)
	}

	return number
}

// FileW хранит дискриптор на открытый файл, готовый к записи
type FileW struct {
	file *os.File
}

// NewOpenFile конструктор, возвращает открытый файл, готовый к записи
func NewOpenFile(name string) *FileW {
	// Создание файла для запись
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	file.Close()

	// Открытие файла для записи
	file, err = os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	return &FileW{file: file}
}

// WritelnString записывает строку в файл, с переходом на новую строку \r\n
func (f *FileW) WritelnString(str string) {
	f.file.WriteString(str + "\r\n")
}

// WritelnNumber записывает число в файл, с переходом на новую строку \r\n
func (f *FileW) WritelnNumber(num int) {
	str := strconv.Itoa(num)
	f.file.WriteString(str + "\r\n")
}

// WriteString записывает строку в файл
func (f *FileW) WriteString(str string) {
	f.file.WriteString(str)
}

// WriteNumber записывает число в файл
func (f *FileW) WriteNumber(num int) {
	str := strconv.Itoa(num)
	f.file.WriteString(str)
}

// WriteStringWithSpace записывает строку в файл и добавляет пробел в конец
func (f *FileW) WriteStringWithSpace(str string) {
	f.file.WriteString(str + " ")
}

// WriteNumberWithSpace записывает число в файл и добавляет пробел в конец
func (f *FileW) WriteNumberWithSpace(num int) {
	str := strconv.Itoa(num)
	f.file.WriteString(str + " ")
}

// WriteSpace записывает число в файл
func (f *FileW) WriteSpace() {
	f.file.WriteString(" ")
}

/* --- Модуль сравнения файлов --- */

// CompareFile сравнивает n файлов, если хотя бы один из них отличается, возвращает false
// на вход принимает названия файлов
func CompareFile(names ...string) bool {
	files := make([]*os.File, 0, len(names))

	for _, name := range names {

		file, err := os.Open(name)
		if err != nil {
			fmt.Println("open file error")
			log.Fatal(err)
		}

		files = append(files, file)
	}

	checksums := []string{}
	for _, f := range files {
		f.Seek(0, 0) // Сброс курсора к началу файла
		sum, err := GetMD5SumString(f)
		if err != nil {
			panic(err)
		}
		// Добавляем контрольную сумму в общий массив
		checksums = append(checksums, sum)
	}

	for i := 0; i < len(checksums)-1; i++ {
		if !(checksums[i] == checksums[i+1]) {
			return false
		}
	}

	return true
}

// GetMD5SumString считает контрольную сумму для файла
func GetMD5SumString(f *os.File) (string, error) {
	file1Sum := md5.New()
	_, err := io.Copy(file1Sum, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", file1Sum.Sum(nil)), nil
}

// PrintTest печает в консоль результат теста
func PrintTest(b bool, testName string) {
	if b {
		fmt.Println(testName, "test passed")
	} else {
		fmt.Println(testName, "test failed")
	}
}

func main() {
	// пример использования
	fileName := []string{
		"01",
		"02",
		"03",
		"04",
		"05",
		"06",
		"07",
		"08",
		"09",
		"10",
	}

	for _, value := range fileName {

		nameInFile := "test/" + value
		nameOutFile := "test/" + value + ".res"
		nameAnswerFile := "test/" + value + ".a"

		solve(nameInFile, nameOutFile)
		PrintTest(CompareFile(nameAnswerFile, nameOutFile), value)
	}
}

# Программа для автоматического тестирования задач с Codeforce на go

## Структура StringReader сканнер строки
Поля:
- str string - строка, которую читаем
- i   int - текущий индекс, на который указывает курсор
- len int - длина строки
### Функция NewStringReader - конструктор, возвращает сканнер по строке
func NewStringReader(str string) *StringReader

### Метод ReadWord - читает слово из строки до первого пробела
func (s *StringReader) ReadWord() string

### Метод ReadNumber - читает число из строки до первого пробела
func (s *StringReader) ReadNumber() int

## Структура FileR - хранит сканнер по открытому файлу
Поля:
- data *bufio.Scanner сканнер по файлу

### Функция NewReadFile - конструктор, возвращает открытый файл
func NewReadFile(str string) *FileR

### Метод eadString - читает строку из файла
func (f *FileR) ReadString() string

### Метод ReadStringAsNumber - читает строку из файла и конвертирует в число. Использовать, когда подразумевается, что в строке находится только одно число
func (f *FileR) ReadStringAsNumber() int

## Структура FileW - хранит дискриптор на открытый файл, готовый к записи

### Функция NewOpenFile - конструктор, возвращает открытый файл, готовый к записи
func NewOpenFile(name string) *FileW

### Метод WritelnString - записывает строку в файл, с переходом на новую строку \r\n
func (f *FileW) WritelnString(str string)

### Метод WritelnNumber - записывает число в файл, с переходом на новую строку \r\n
func (f *FileW) WritelnNumber(num int)

## Модуль сравнения файлов

### Функция CompareFile сравнивает n файлов, если хотя бы один из них отличается, возвращает false. На вход принимает названия файлов

### Функция GetMD5SumString считает контрольную сумму для файла
func GetMD5SumString(f *os.File) (string, error)

### Функция PrintTest печает в консоль результат теста
func PrintTest(b bool, testName string)**
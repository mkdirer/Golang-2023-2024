package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

// Struktura danych użytkowników
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []*User  `xml:"user"`
}

// Struktura pojedynczego użytkownika
type User struct {
	XMLName  xml.Name `xml:"user"`
	Login    string   `xml:"login"`
	Password string   `xml:"password"`
	Role     int      `xml:"role"`
}

// Stałe reprezentujące uprawnienia
const (
	Nothing = 1 << iota
	Read
	Add
	Edit
	All = Read | Add | Edit
)

// Funkcja sprawdzająca poprawność wartości uprawnienia
func IsValid(value int) bool {
	return value >= Nothing && value <= All
}

// Struktura danych osobowych
type Person struct {
	XMLName    xml.Name `xml:"person"`
	Id         int      `xml:"id"`
	FirstName  string   `xml:"firstName"`
	LastName   string   `xml:"lastName"`
	Age        int      `xml:"age"`
	Birth      Data     `xml:"birth"`
	Death      Data     `xml:"death"`
	Pesel      int      `xml:"pesel"`
	CreditCard int      `xml:"creditcard"`
	Gender     rune     `xml:"gender"`
}

// Struktura reprezentująca datę
type Data struct {
	D, M, Y int
}

// Struktura zawierająca dane osób
type People struct {
	XMLName xml.Name  `xml:"persons"`
	People  []*Person `xml:"person"`
}

// Funkcja szyfrująca plik algorytmem AES
func EncryptFile(plainText []byte, filename, key string) error {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(cipherText)
	if err != nil {
		return err
	}
	return nil
}

func DecryptFile(filename string, key string) ([]byte, error) {
	file, err := os.ReadFile("encrypted.xml")
	if err != nil {
		fmt.Println("Error while reading", filename, err)
	}

	block, err := aes.NewCipher([]byte("The giraffes enter the wardrobe."))
	if err != nil {
		fmt.Println("Cipher error", err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Cipher GCM error", err)
		return nil, err
	}
	nonce := file[:gcm.NonceSize()]
	file = file[gcm.NonceSize():]
	decryptedfile, err := gcm.Open(nil, nonce, file, nil)
	if err != nil {
		fmt.Println("Decrypt file error", err)
		return nil, err
	}

	return decryptedfile, nil
}

// Funkcja sprawdzająca poprawność numeru PESEL
func IsValidPesel(person Person) bool {
	pesel := strconv.Itoa(person.Pesel)
	if len(pesel) != 11 {
		return false
	}
	peselYear, err := strconv.Atoi(pesel[:2])
	fmt.Println(peselYear)
	if err != nil {
		fmt.Println("Error while checking year")
		return false
	}
	birthYearLastTwoDigits := person.Birth.Y % 100
	fmt.Println(birthYearLastTwoDigits)
	var monthOffset int
	switch {
	case person.Birth.Y >= 1800 && person.Birth.Y <= 1899:
		monthOffset = 80
	case person.Birth.Y >= 1900 && person.Birth.Y <= 1999:
		monthOffset = 0
	case person.Birth.Y >= 2000 && person.Birth.Y <= 2099:
		monthOffset = 20
	case person.Birth.Y >= 2100 && person.Birth.Y <= 2199:
		monthOffset = 40
	case person.Birth.Y >= 2200 && person.Birth.Y <= 2299:
		monthOffset = 60
	default:
		fmt.Println("Error while checking year")
		return false
	}
	peselMonth, err := strconv.Atoi(pesel[2:4])
	if err != nil {
		fmt.Println("Error while checking month")
		return false
	}
	if peselMonth != (person.Birth.M+monthOffset)%100 {
		fmt.Println("Error while checking month")
		return false
	}
	peselDay, err := strconv.Atoi(pesel[4:6])
	if err != nil {
		fmt.Println("Error while checking day")
		return false
	}
	if peselDay != person.Birth.D {
		fmt.Println("Error while checking day")
		return false
	}
	female := false
	odd := []string{"1", "3", "5", "7", "9"}
	if person.Gender == 95 {
		female = true
	}
	if female && contains(odd, string(pesel[9])) {
		fmt.Println("Error while checking gender")
		return false
	}
	w := [11]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
	s := 0
	for i, v := range w {
		t, err := strconv.Atoi(string(pesel[i]))
		if err != nil {
			fmt.Println("Error while checking")
			return false
		}
		s += v * t
	}
	s = s % 10
	s = 10 - s
	if strconv.Itoa(s) != string(pesel[10]) {
		return false
	}
	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Funkcja sprawdzająca poprawność numeru karty kredytowej
func IsValidCreditCard(person Person) bool {
	creditCard := strconv.Itoa(person.CreditCard)
	if len(creditCard) != 16 {
		return false
	}
	weights := map[int]int{0: 0, 1: 2, 2: 4, 3: 6, 4: 8, 5: 1, 6: 3, 7: 5, 8: 7, 9: 9}
	sum := 0
	for i, r := range creditCard {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Println("Error while converting credit card digit")
			return false
		}
		weight, ok := weights[i]
		if !ok {
			fmt.Println("Error while getting weight for credit card digit")
			return false
		}
		if i%2 == 0 {
			sum += digit * weight
		} else {
			if digit*2 >= 10 {
				sum += (digit*2)%10 + (digit*2)/10
			} else {
				sum += digit * 2
			}
		}
	}
	return sum%10 == 0
}

// Funkcja sprawdzająca poprawność wieku
func IsValidAge(person Person) bool {
	if person.Death.Y != 0 {
		diff := person.Death.Y - person.Birth.Y
		if diff != person.Age {
			return false
		}
	} else {
		today := time.Now()
		diff := today.Year() - person.Birth.Y
		if diff != person.Age {
			return false
		}
	}
	return true
}

func main() {
	// Wczytanie danych użytkowników lub utworzenie nowej bazy
	var users Users
	filename := "users.xml"
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Println("File", filename, "doesn't exist")
	} else {
		xmlFile, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error while opening", filename, err)
		}
		defer xmlFile.Close()
		byteValue, err := io.ReadAll(xmlFile)
		if err != nil {
			fmt.Println("Error while reading", filename, err)
		}
		err = xml.Unmarshal(byteValue, &users)
		if err != nil {
			fmt.Println("Error while unmarshaling users", err)
			return
		}
	}

	fmt.Println("Nowy")

	// Dodawanie nowego użytkownika
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter login:")
	login, _ := reader.ReadString('\n')
	login = strings.TrimSpace(login)

	fmt.Println("Enter password:")
	password, _ := term.ReadPassword(int(syscall.Stdin))
	sha256Hash := sha256.Sum256(password)
	passh := hex.EncodeToString(sha256Hash[:])

	fmt.Println("Enter role (Nothing = 1, Read = 2, Add = 4, Edit = 8, All = 15):")
	roleStr, _ := reader.ReadString('\n')
	roleStr = strings.TrimSpace(roleStr)
	roleInt, _ := strconv.Atoi(roleStr)
	if !IsValid(roleInt) {
		roleInt = Nothing
	}

	user := User{Login: login, Password: passh, Role: roleInt}
	users.Users = append(users.Users, &user)

	// Zapisanie danych użytkowników do pliku
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Error while opening", filename, err)
	}
	defer file.Close()
	fileContent, err := xml.Marshal(users)
	if err != nil {
		fmt.Println("Error while marshaling users", err)
		return
	}
	_, err = file.Write(fileContent)
	if err != nil {
		fmt.Println("Error while writing to file", filename, err)
	}

	// Wczytanie zaszyfrowanych danych osobowych
	encryptionKey := "The giraffes enter the wardrobe."
	encryptedFilename := "encrypted.xml"
	var people People
	plainText, err := DecryptFile(encryptedFilename, encryptionKey)
	if err != nil {
		fmt.Println("Error while decrypting file", encryptedFilename, err)
	}
	if err = xml.Unmarshal(plainText, &people); err != nil {
		fmt.Println("Error while unmarshaling people", err)
	}

	fmt.Println(people.XMLName)
	fmt.Println(people.People)

	// Dodawanie nowej osoby do bazy danych osobowych
	var person Person
	person.Id = 3
	person.FirstName = "Name"
	person.LastName = "Surname"
	person.Age = 24
	person.Birth = Data{D: 11, M: 1, Y: 1999}
	person.Death = Data{}
	person.Pesel = 99011111954
	person.CreditCard = 242134123441
	person.Gender = 'W'

	if IsValidPesel(person) && IsValidCreditCard(person) && IsValidAge(person) {
		people.People = append(people.People, &person)
	}
	// Zapisanie danych osobowych do zaszyfrowanego pliku
	fmt.Println("Writing to file")
	peopleFileContent, err := xml.Marshal(people)
	if err != nil {
		fmt.Println("Error while marshaling people", err)
		return
	}
	err = EncryptFile(peopleFileContent, encryptedFilename, encryptionKey)
	if err != nil {
		fmt.Println("Error while writing decrypted content to file", encryptedFilename, err)
	}
}

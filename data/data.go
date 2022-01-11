package data

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitChat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// create a random UUID with from RFX 4122
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannnot generate UUID", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

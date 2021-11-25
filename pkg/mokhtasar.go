package pkg

import (
	"fmt"

	"gorm.io/gorm"
)

type Store interface {
	GetOrginalurl(key string) (string, error)
	Shorten(url string) (string, error)
}

type urls struct {
	// gorm.Model
	ID  int
	Url string
	Key string
}

type PostgresStore struct {
	DB              *gorm.DB
	RandomGenerator func(len int) string
}

func (m *PostgresStore) GetOrginalurl(key string) (string, error) {
	// query := `SELECT url FROM urls WHERE key=$1`
	// rows, err := m.DB.Query(query, key)
	// fmt.Println(key)
	davood := urls{}
	rows := m.DB.Where("key = ?", key).Last(&davood)
	// rows := m.DB.First(&i, 1)
	// ro := m.DB.Scan(rows)
	// fmt.Println("hi")
	// fmt.Println(m.DB.Table("urls").First(&i))
	// fmt.Println(davood)
	fmt.Println(rows)
	// fmt.Println("bye")
	return davood.Url, nil
}

func (m *PostgresStore) Shorten(url string) (string, error) {
	key := m.RandomGenerator(5)
	m.DB.Create(&urls{
		// ID:  50,
		Url: url,
		Key: key,
	})
	// _, err := m.DB.Exec(`INSERT INTO urls (url,key) VALUES ($1,$2)`, url, key)
	// if err != nil {
	// return "", err
	// }

	return key, nil
}

package domain

import (
	"hash/fnv"
	"time"
)

// ADR holds the data of a single Architecture Decision Record.
type ADR struct {
	ID         string    `gorm:"primaryKey"`
	Title      string    `gorm:"index"`
	Status     string
	Content    string `gorm:"type:text"`
	URL        string `gorm:"uniqueIndex"`
	SourceType string
	SourceURL  string
	SourceName string `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// GenerateID generates a Base62 encoded FNV-1a hash of the given URL
func GenerateID(url string) string {
	h := fnv.New64a()
	h.Write([]byte(url))
	hash := h.Sum64()
	return encodeBase62(hash)
}

// encodeBase62 encodes a uint64 to Base62 string
func encodeBase62(num uint64) string {
	const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return "0"
	}

	var result []byte
	for num > 0 {
		result = append([]byte{base62[num%62]}, result...)
		num /= 62
	}
	return string(result)
}

// TableName specifies the table name for GORM
func (ADR) TableName() string {
	return "adrs"
}

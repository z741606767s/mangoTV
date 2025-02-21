package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Md5 MD5加密
func Md5(str string, salt ...interface{}) (CryptStr string) {
	//判断salt盐是否存在
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func JsonMarshal(value interface{}) []byte {
	bs, err := json.Marshal(value)
	if err != nil {
		return []byte{}
	}
	return bs
}

// GenerateNickname 生成随机昵称的方法
func GenerateNickname() string {
	// 形容词列表
	var adjectives = []string{
		"Brave", "Clever", "Happy", "Bold", "Calm", "Daring", "Gentle", "Jolly", "Kind", "Quick",
		"Fierce", "Silent", "Mighty", "Noble", "Shy", "Fearless", "Wise", "Strong", "Loyal", "Playful",
		"Lucky", "Swift", "Bright", "Epic", "Majestic", "Serene", "Sly", "Charming", "Eager", "Graceful",
		"Curious", "Patient", "Vivid", "Radiant", "Witty", "Humble", "Vibrant", "Mystic", "Bold", "Dashing",
		"Heroic", "Glorious", "Stalwart", "Faithful", "Frosty", "Blazing", "Golden", "Silver", "Sapphire",
		"Emerald", "Crimson", "Shadow", "Sunny", "Stormy", "Courageous", "Sturdy", "Enigmatic", "Frosty",
		"Electric", "Whimsical", "Joyful", "Funky", "Zealous", "Funky", "Cheerful", "Passionate", "Proud",
		"Steady", "Crisp", "Diligent", "Sharp", "Stellar", "Radiant", "Graceful", "Fearless", "Funky",
	}

	// 名词列表
	var nouns = []string{
		"Lion", "Tiger", "Panda", "Eagle", "Shark", "Wolf", "Bear", "Falcon", "Hawk", "Whale",
		"Phoenix", "Dragon", "Leopard", "Panther", "Fox", "Griffin", "Turtle", "Dolphin", "Jaguar", "Cheetah",
		"Falcon", "Viper", "Serpent", "Stallion", "Raven", "Cougar", "Lynx", "Owl", "Cobra", "Scorpion",
		"Rhino", "Hippo", "Elephant", "Gazelle", "Chameleon", "Crane", "Otter", "Beetle", "Mantis", "Wolverine",
		"Giraffe", "Buffalo", "Jackal", "Pelican", "Seal", "Hedgehog", "Tarantula", "Iguana", "Zebra", "Gorilla",
		"Alligator", "Lobster", "Crocodile", "Meerkat", "Wombat", "Koala", "Penguin", "Moose", "Sparrow",
		"Unicorn", "Raptor", "Vulture", "Beaver", "Salamander", "Gecko", "Ram", "Antelope", "Parrot", "Orca",
	}

	rand.Seed(time.Now().UnixNano())
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return adjective + noun
}

// RandomNickname 生成随机符号昵称
func RandomNickname(length int) string {
	var sb strings.Builder
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()|~"
	charsetLength := len(charset)
	for i := 0; i < length; i++ {
		randomChar := charset[rand.Intn(charsetLength)]
		sb.WriteByte(randomChar)
	}
	return sb.String()
}

// HashEmail 处理并返回哈希后的邮箱
func HashEmail(email string) string {
	// 将邮箱转换为小写并去除首尾空格
	normalizedEmail := strings.TrimSpace(strings.ToLower(email))

	// 进行SHA-256哈希
	hash := sha256.New()
	hash.Write([]byte(normalizedEmail))
	return hex.EncodeToString(hash.Sum(nil))
}

func HashWithSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}

func GenerateOrderNumber() string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 生成一个6位的随机数
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000000)

	// 组合时间戳和随机数
	orderNumber := fmt.Sprintf("%d%06d", timestamp, randomNumber)
	return orderNumber
}

func GenerateUUIDOrderNumber() string {
	return uuid.New().String()
}

func RoundToTwoDecimal(value float64) float64 {
	return math.Round(value*100) / 100
}

func TruncateToTwoDecimal(value float64) float64 {
	factor := math.Pow(10, 2)
	return math.Trunc(value*factor) / factor
}

func ContainsInt64(slice []int64, element int64) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func GetManualOpenId() string {
	return fmt.Sprintf("manual-%s", uuid.New().String()[:23])
}

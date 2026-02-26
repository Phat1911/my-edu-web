package handlers

import (
	"edu-web-backend/internal/repository"
	"encoding/base64"
	"strings"

	"github.com/skip2/go-qrcode"
)

func generateQRBase64(url string) string {
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return ""
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
}

// categoryKeywords maps psychological categories to trigger keywords (no-diacritic Vietnamese + English).
var categoryKeywords = map[string][]string{
	"stress": {
		"stress", "cang thang", "ap luc", "kiem tra", "thi cu", "on thi",
		"thi dai hoc", "qua tai", "nhieu viec", "met moi", "dau dau",
		"so sanh", "truot", "bo me ky vong", "lich hoc", "bai kho",
	},
	"anxiety": {
		"lo lang", "anxiety", "lo au", "hoi hop", "so hai",
		"hoang loan", "panic", "ngu khong duoc", "mat ngu lo",
		"nguoi khac nghi", "bi phan xet", "bat an", "khong yen",
		"run", "tim dap nhanh", "kho tho",
	},
	"motivation": {
		"mat dong luc", "chan hoc", "khong muon hoc", "luoi", "tri hoan",
		"khong co muc tieu", "vo nghia", "game", "dien tu", "nan long",
		"bo cuoc", "that bai", "ghen ti", "procrastinat",
	},
	"focus": {
		"tap trung", "focus", "phan tam", "mat tap trung", "hay quen",
		"khong nho", "dien thoai", "mang xa hoi", "facebook", "tiktok",
		"lan man", "buon ngu khi hoc", "adhd", "tang dong", "khong hoan thanh",
	},
	"sleep": {
		"ngu", "sleep", "mat ngu", "kho ngu", "khong ngu duoc",
		"buon ngu", "ac mong", "thuc khuya", "day som", "giac ngu",
		"ngu khong ngon", "nghi nhieu truoc khi ngu",
	},
	"loneliness": {
		"co don", "le loi", "mot minh", "khong co ban", "ban be",
		"bi xa lanh", "bi bo roi", "chia tay", "mau thuan", "xung dot",
		"thay co", "bo me khong hieu", "khong ai hieu", "tinh yeu",
	},
	"self-esteem": {
		"tu ti", "kem coi", "ngoai hinh", "xau", "beo",
		"gay", "khong gioi", "dot", "vo dung", "khong xung dang",
		"tu trach", "tu phe binh", "diem thap",
	},
	"depression": {
		"buon", "tram cam", "depression", "sad", "trong rong", "vo cam",
		"mat hung", "khoc", "tuyet vong", "vo vong", "khong co hy vong",
		"tu tu", "tu lam hai", "chet", "khong con suc",
	},
}

type scoredCategory struct {
	category string
	score    int
}

// detectCategory scores the message against all category keywords and returns the best match.
func detectCategory(msg string) string {
	msg = strings.ToLower(msg)
	best := scoredCategory{}
	for category, keywords := range categoryKeywords {
		score := 0
		for _, kw := range keywords {
			if strings.Contains(msg, kw) {
				score++
			}
		}
		if score > best.score {
			best = scoredCategory{category, score}
		}
	}
	if best.score > 0 {
		return best.category
	}
	return ""
}

func buildAIResponse(message string, db *repository.DB) string {
	msg := strings.ToLower(message)

	// 1. Emergency check - self-harm keywords (highest priority)
	emergencyKws := []string{"tu tu", "tu lam hai", "muon chet", "khong muon song", "ket thuc tat ca"}
	for _, kw := range emergencyKws {
		if strings.Contains(msg, kw) {
			crisis, err := db.GetScenarioByKeyword("tu tu")
			if err == nil && crisis != nil {
				return crisis.Response + "\n\n Meo: " + crisis.Tips
			}
			return "Minh rat lo lang khi nghe dieu nay. Ban khong co don - co nguoi san sang lang nghe va giup ban ngay bay gio.\n\nDuong day ho tro khung hoang tam than Viet Nam: 1800 599 920 (mien phi, 24/7)\n\nHay goi ngay nhe. Minh o day ben ban."
		}
	}

	// 2. Detect psychological category from message content
	category := detectCategory(msg)

	// 3. If a category is detected, find the most relevant scenario in DB
	if category != "" {
		// Try to find a scenario that also matches a specific trigger keyword
		triggerKws := []string{
			"thi cu", "kiem tra", "bai tap",
			"bo me", "gia dinh", "ban be",
			"thay co", "dien thoai", "game",
			"ngu", "tap trung", "mat dong luc",
			"tu ti", "buon", "lo lang",
			"stress", "truot", "chan", "luoi",
		}
		for _, kw := range triggerKws {
			if strings.Contains(msg, kw) {
				scenario, err := db.GetScenarioByKeyword(kw)
				if err == nil && scenario != nil && scenario.Category == category {
					return scenario.Response + "\n\n Meo: " + scenario.Tips
				}
			}
		}
		// Fallback: random scenario from the detected category
		scenario, err := db.GetScenarioByCategory(category)
		if err == nil && scenario != nil {
			return scenario.Response + "\n\n Meo: " + scenario.Tips
		}
	}

	// 4. Simple greeting/keyword responses
	type kwr struct{ kw, resp string }
	greetings := []kwr{
		{"xin chao", "Xin chao! Minh la Buddy AI - nguoi ban dong hanh tam ly 24/7. Ban dang cam thay the nao?"},
		{"chao", "Xin chao! Minh la Buddy AI - nguoi ban dong hanh 24/7. Ban dang cam thay the nao hom nay?"},
		{"hello", "Hello! Minh o day de lang nghe ban. Hay chia se bat cu dieu gi ban muon nhe!"},
		{"hi ", "Hi! Buddy AI day. Ban can minh ho tro gi hom nay?"},
		{"hoc", "Hoc tap doi khi rat thu thach. Ban dang gap kho khan o diem nao?"},
		{"met", "Met moi la tin hieu co the can nghi ngoi. Ban dang met vi dieu gi?"},
		{"khoc", "Duoc khoc la dieu binh thuong. Minh o day ben ban. Chuyen gi dang xay ra vay?"},
		{"ap luc", "Ap luc co the rat nang ne. Hay chia se them de minh hieu ban dang doi mat voi gi nhe."},
		{"co don", "Cam giac co don rat pho bien. Ban khong he mot minh - minh luon o day lang nghe."},
	}
	for _, g := range greetings {
		if strings.Contains(msg, g.kw) {
			return g.resp
		}
	}

	// 5. Default response
	return "Cam on ban da chia se! Minh dang lang nghe. Ban co the ke them de minh hieu ro hon va ho tro ban tot hon khong?\n\nNgoai ra, ban co the thu:\n- Nghe am thanh song nao trong muc Audio\n- Xem video meo hoc tap\n- Quet ma QR de truy cap nhanh tai nguyen"
}

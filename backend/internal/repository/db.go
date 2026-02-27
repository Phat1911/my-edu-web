package repository

import (
	"context"
	"edu-web-backend/internal/models"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(dbURL string) (*DB, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	return &DB{pool: pool}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) Migrate(ctx context.Context) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS videos (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT DEFAULT '',
			drive_url TEXT NOT NULL,
			embed_url TEXT NOT NULL,
			thumbnail TEXT DEFAULT '',
			category VARCHAR(100) DEFAULT 'general',
			order_num INT DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS audios (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT DEFAULT '',
			drive_url TEXT NOT NULL,
			embed_url TEXT NOT NULL,
			category VARCHAR(100) DEFAULT 'general',
			duration VARCHAR(20) DEFAULT '',
			order_num INT DEFAULT 0,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS qrcodes (
			id SERIAL PRIMARY KEY,
			label VARCHAR(255) NOT NULL,
			target_url TEXT NOT NULL,
			type VARCHAR(50) DEFAULT 'general',
			qr_data TEXT DEFAULT '',
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS chat_messages (
			id SERIAL PRIMARY KEY,
			session_id VARCHAR(100) NOT NULL,
			role VARCHAR(20) NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS psych_scenarios (
			id SERIAL PRIMARY KEY,
			category VARCHAR(100) NOT NULL,
			trigger TEXT NOT NULL,
			response TEXT NOT NULL,
			tips TEXT DEFAULT ''
		)`,
	}
	for _, q := range queries {
		if _, err := db.pool.Exec(ctx, q); err != nil {
			return fmt.Errorf("migration error: %w", err)
		}
	}
	return nil
}

func (db *DB) SeedData(ctx context.Context) error {
	var count int
	db.pool.QueryRow(ctx, "SELECT COUNT(*) FROM videos").Scan(&count)
	if count > 0 {
		return nil
	}
 
	videos := []models.Video{
		{
			Title:       "Mẹo học tập hiệu quả - Phần 1",
			Description: "Các kỹ thuật học tập giúp tăng khả năng ghi nhớ và tập trung",
			DriveURL:    "https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8",
			EmbedURL:    "https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8",
			Thumbnail:   "",
			Category:    "learning-tips",
			Order:       1,
		},
		{
			Title:       "Kỹ thuật Pomodoro",
			Description: "Phương pháp quản lý thời gian học tập theo chu kỳ 25 phút",
			DriveURL:    "https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8",
			EmbedURL:    "https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8",
			Thumbnail:   "",
			Category:    "learning-tips",
			Order:       2,
		},
		{
			Title:       "Tư duy tích cực trong học tập",
			Description: "Xây dựng mindset phát triển để học tập hiệu quả hơn",
			DriveURL:    "https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8",
			EmbedURL:    "https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8",
			Thumbnail:   "",
			Category:    "psychology",
			Order:       3,
		},
	}

	for _, v := range videos {
		_, err := db.pool.Exec(ctx,
			`INSERT INTO videos (title, description, drive_url, embed_url, thumbnail, category, order_num) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
			v.Title, v.Description, v.DriveURL, v.EmbedURL, v.Thumbnail, v.Category, v.Order,
		)
		if err != nil {
			return err
		}
	}

	audios := []models.Audio{
		{
			Title:       "Sóng não Alpha - Tập trung học tập",
			Description: "Âm thanh sóng não Alpha 10Hz giúp tăng khả năng tập trung và học tập",
			DriveURL:    "https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q",
			EmbedURL:    "https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q",
			Category:    "brainwave",
			Duration:    "60:00",
			Order:       1,
		},
		{
			Title:       "Sóng não Theta - Sáng tạo và thư giãn",
			Description: "Âm thanh sóng não Theta 6Hz kích thích sáng tạo và thư giãn sâu",
			DriveURL:    "https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q",
			EmbedURL:    "https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q",
			Category:    "brainwave",
			Duration:    "45:00",
			Order:       2,
		},
		{
			Title:       "Sóng não Beta - Tăng cường trí nhớ",
			Description: "Âm thanh sóng não Beta 20Hz hỗ trợ ghi nhớ và xử lý thông tin",
			DriveURL:    "https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q",
			EmbedURL:    "https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q",
			Category:    "brainwave",
			Duration:    "30:00",
			Order:       3,
		},
	}

	for _, a := range audios {
		_, err := db.pool.Exec(ctx,
			`INSERT INTO audios (title, description, drive_url, embed_url, category, duration, order_num) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
			a.Title, a.Description, a.DriveURL, a.EmbedURL, a.Category, a.Duration, a.Order,
		)
		if err != nil {
			return err
		}
	}

	scenarios := []struct {
		category string
		trigger  string
		response string
		tips     string
	}{
		{"stress", "Tôi bị stress vì bài kiểm tra sắp tới", "Tôi hiểu bạn đang cảm thấy áp lực. Điều đó hoàn toàn bình thường. Hãy thử hít thở sâu 5 lần và chia nhỏ việc ôn tập thành các phần nhỏ hơn.", "Kỹ thuật Pomodoro: học 25 phút, nghỉ 5 phút"},
		{"anxiety", "Tôi lo lắng về tương lai", "Lo lắng về tương lai là điều nhiều bạn trải qua. Hãy tập trung vào những gì bạn có thể kiểm soát ngay hôm nay.", "Viết nhật ký cảm xúc mỗi ngày"},
		{"motivation", "Tôi không muốn học nữa", "Cảm giác mất động lực xảy ra với tất cả mọi người. Hãy nhớ lại lý do ban đầu bạn muốn học và những mục tiêu bạn đặt ra.", "Tặng thưởng cho bản thân sau mỗi mục tiêu nhỏ"},
		{"focus", "Tôi không tập trung được khi học", "Khó tập trung có thể do nhiều nguyên nhân. Thử nghe âm thanh sóng não Alpha trong khi học và tắt điện thoại.", "Tạo không gian học tập yên tĩnh, không có điện thoại"},
		{"sleep", "Tôi ngủ không được vì lo lắng", "Giấc ngủ rất quan trọng cho việc học. Thử nghe âm thanh sóng não Theta trước khi ngủ để thư giãn.", "Không dùng điện thoại 1 tiếng trước khi ngủ"},
	}

	for _, s := range scenarios {
		db.pool.Exec(ctx,
			`INSERT INTO psych_scenarios (category, trigger, response, tips) VALUES ($1,$2,$3,$4)`,
			s.category, s.trigger, s.response, s.tips,
		)
	}

	return nil
}

func (db *DB) GetAllVideos(ctx context.Context) ([]models.Video, error) {
	rows, err := db.pool.Query(ctx, `SELECT id, title, description, drive_url, embed_url, thumbnail, category, order_num, created_at FROM videos ORDER BY order_num ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var videos []models.Video
	for rows.Next() {
		var v models.Video
		if err := rows.Scan(&v.ID, &v.Title, &v.Description, &v.DriveURL, &v.EmbedURL, &v.Thumbnail, &v.Category, &v.Order, &v.CreatedAt); err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return videos, nil
}

func (db *DB) GetAllAudios(ctx context.Context) ([]models.Audio, error) {
	rows, err := db.pool.Query(ctx, `SELECT id, title, description, drive_url, embed_url, category, duration, order_num, created_at FROM audios ORDER BY order_num ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var audios []models.Audio
	for rows.Next() {
		var a models.Audio
		if err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.DriveURL, &a.EmbedURL, &a.Category, &a.Duration, &a.Order, &a.CreatedAt); err != nil {
			return nil, err
		}
		audios = append(audios, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return audios, nil
}

func (db *DB) GetAllQRCodes(ctx context.Context) ([]models.QRCode, error) {
	rows, err := db.pool.Query(ctx, `SELECT id, label, target_url, type, qr_data, created_at FROM qrcodes ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var qrs []models.QRCode
	for rows.Next() {
		var q models.QRCode
		if err := rows.Scan(&q.ID, &q.Label, &q.TargetURL, &q.Type, &q.QRData, &q.CreatedAt); err != nil {
			return nil, err
		}
		qrs = append(qrs, q)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return qrs, nil
}

func (db *DB) SaveQRCode(ctx context.Context, label, targetURL, qrType, qrData string) (*models.QRCode, error) {
	var q models.QRCode
	err := db.pool.QueryRow(ctx,
		`INSERT INTO qrcodes (label, target_url, type, qr_data) VALUES ($1,$2,$3,$4) RETURNING id, label, target_url, type, qr_data, created_at`,
		label, targetURL, qrType, qrData,
	).Scan(&q.ID, &q.Label, &q.TargetURL, &q.Type, &q.QRData, &q.CreatedAt)
	return &q, err
}

func (db *DB) SaveChatMessage(ctx context.Context, sessionID, role, content string) error {
	_, err := db.pool.Exec(ctx,
		`INSERT INTO chat_messages (session_id, role, content) VALUES ($1,$2,$3)`,
		sessionID, role, content,
	)
	return err
}

func (db *DB) GetChatHistory(ctx context.Context, sessionID string) ([]models.ChatMessage, error) {
	rows, err := db.pool.Query(ctx,
		`SELECT id, session_id, role, content, created_at FROM chat_messages WHERE session_id=$1 ORDER BY created_at ASC LIMIT 50`,
		sessionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var msgs []models.ChatMessage
	for rows.Next() {
		var m models.ChatMessage
		if err := rows.Scan(&m.ID, &m.SessionID, &m.Role, &m.Content, &m.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return msgs, nil
}

func (db *DB) GetScenarioByKeyword(ctx context.Context, keyword string) (*models.PsychScenario, error) {
	var s models.PsychScenario
	err := db.pool.QueryRow(ctx,
		`SELECT id, category, trigger, response, tips FROM psych_scenarios WHERE trigger ILIKE $1 LIMIT 1`,
		"%"+keyword+"%",
	).Scan(&s.ID, &s.Category, &s.Trigger, &s.Response, &s.Tips)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (db *DB) GetScenarioByCategory(ctx context.Context, category string) (*models.PsychScenario, error) {
	var s models.PsychScenario
	err := db.pool.QueryRow(ctx,
		`SELECT id, category, trigger, response, tips FROM psych_scenarios WHERE category = $1 ORDER BY RANDOM() LIMIT 1`,
		category,
	).Scan(&s.ID, &s.Category, &s.Trigger, &s.Response, &s.Tips)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &s, nil
}

func (db *DB) GetScenariosByCategory(ctx context.Context, category string) ([]models.PsychScenario, error) {
	rows, err := db.pool.Query(ctx,
		`SELECT id, category, trigger, response, tips FROM psych_scenarios WHERE category = $1`,
		category,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var scenarios []models.PsychScenario
	for rows.Next() {
		var s models.PsychScenario
		if err := rows.Scan(&s.ID, &s.Category, &s.Trigger, &s.Response, &s.Tips); err != nil {
			return nil, err
		}
		scenarios = append(scenarios, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return scenarios, nil
}

func (db *DB) SeedScenarios(ctx context.Context) error {
	var count int
	if err := db.pool.QueryRow(ctx, "SELECT COUNT(*) FROM psych_scenarios").Scan(&count); err != nil {
		return fmt.Errorf("seed scenarios count: %w", err)
	}
	if count >= 50 {
		return nil
	}

	db.pool.Exec(ctx, "DELETE FROM psych_scenarios")

	scenarios := []struct {
		category string
		trigger  string
		response string
		tips     string
	}{
		// ---- STRESS (10) ----
		{"stress", "stress thi cu kiem tra", "Bạn đang chịu áp lực thi cử - điều này rất phổ biến và hoàn toàn có thể vượt qua. Hãy chia nhỏ nội dung cần ôn thành các phần 25 phút (Pomodoro), nghỉ 5 phút giữa mỗi phần. Não bạn sẽ hấp thụ tốt hơn nhiều khi không bị nhồi nhét liên tục.", "Viết ra 3 chủ đề quan trọng nhất cần ôn, tập trung từng cái một"},
		{"stress", "stress ap luc gia dinh bo me", "Áp lực từ gia đình đôi khi nặng nề hơn cả bài vở. Bố mẹ thường kỳ vọng cao vì họ yêu thương bạn, nhưng điều đó không có nghĩa bạn phải gánh một mình. Hãy thử nói chuyện thẳng thắn với bố mẹ về cảm xúc của mình - nhiều bạn bất ngờ vì bố mẹ sẵn sàng lắng nghe hơn họ tưởng.", "Chọn một buổi tối yên tĩnh, chia sẻ cảm xúc bằng câu 'Con cảm thấy...' thay vì chỉ trích"},
		{"stress", "stress nhieu viec qua tai qua", "Khi mọi thứ dồn lại quá nhiều, não bạn bị quá tải và không thể hoạt động hiệu quả. Bước đầu tiên: dừng lại và thở. Hít vào 4 giây, giữ 4 giây, thở ra 6 giây - lặp 5 lần. Sau đó viết ra TẤT CẢ việc cần làm để đầu óc được giải phóng.", "Dùng ma trận Eisenhower: chia việc thành 'gấp-quan trọng', 'gấp-ít quan trọng', 'không gấp-quan trọng', 'bỏ qua'"},
		{"stress", "stress cang thang dau dau met moi", "Căng thẳng kéo dài biểu hiện qua cơ thể: đau đầu, mệt mỏi là tín hiệu cơ thể đang cần giúp đỡ. Đừng bỏ qua. Hãy uống đủ nước (não cần 2L/ngày), vận động nhẹ 15 phút, và đảm bảo ngủ đủ 7-8 tiếng tối nay.", "Massage nhẹ vùng thái dương và cổ gáy trong 2 phút để giảm đau đầu tức thì"},
		{"stress", "stress truoc ky thi lon dai hoc", "Kỳ thi đại học là áp lực thực sự lớn. Nhưng nhớ rằng: không có kỳ thi nào quyết định toàn bộ cuộc đời bạn. Hãy chuẩn bị tốt nhất có thể, nhưng cũng chấp nhận rằng kết quả không hoàn toàn trong tay bạn - và điều đó ổn thôi.", "3 ngày trước thi: ôn nhẹ, ngủ đủ giấc, ăn sáng đầy đủ - đây quan trọng hơn nhồi bài"},
		{"stress", "stress bi ban be ap luc dong loai", "Áp lực từ bạn bè và mạng xã hội (ai cũng có vẻ học giỏi, thành công hơn) rất độc hại. Thực tế, người ta chỉ đăng highlight của cuộc sống, không ai đăng lúc họ thất bại. Hãy tập trung vào hành trình của chính bạn.", "Giảm 30 phút lướt mạng xã hội mỗi ngày, thay bằng làm một việc bạn thích"},
		{"stress", "stress cong viec hoc nhieu qua khong xong", "Cảm giác bị chìm ngập trong công việc học tập. Hãy thử quy tắc '2 phút': nếu việc gì làm được trong 2 phút, làm ngay. Việc lớn hơn thì chia nhỏ - mỗi phần không quá 30 phút. Bắt đầu từ việc DỄ NHẤT để tạo đà.", "Dùng app Todoist hoặc viết tay danh sách, gạch bỏ khi hoàn thành - não rất thích cảm giác này"},
		{"stress", "stress lo ngai tuong lai khong biet lam gi", "Lo lắng về tương lai nghề nghiệp là hoàn toàn bình thường ở độ tuổi học sinh. Bạn không cần biết mình muốn làm gì cả đời ngay lúc này. Hãy tập trung khám phá: thử nhiều thứ, chú ý điều gì khiến bạn hứng thú và quên mất thời gian.", "Thử '5 câu hỏi tại sao': viết một điều bạn thích, hỏi 'tại sao' 5 lần để tìm ra giá trị thực sự"},
		{"stress", "stress thi truot hat thi truot mon", "Trượt môn không phải là thất bại cuối cùng - đó là thông tin để bạn học cách học hiệu quả hơn. Nhiều người thành công từng trượt nhiều lần. Hãy phân tích: trượt vì thiếu kiến thức, thiếu thời gian, hay thiếu phương pháp? Mỗi nguyên nhân có giải pháp khác nhau.", "Gặp thầy cô hỏi thẳng: 'Em cần cải thiện điểm gì để thi lại tốt hơn?'"},
		{"stress", "stress bi so sanh voi anh chi nguoi khac gioi hon", "Bị so sánh rất đau. Nhưng bạn đang được so sánh với người khác trong khi chỉ có thể trở thành phiên bản tốt hơn của chính mình. Anh/chị giỏi hơn không có nghĩa bạn kém - họ có lợi thế và hoàn cảnh khác nhau. Cuộc đua duy nhất có ý nghĩa là với bản thân bạn ngày hôm qua.", "Mỗi tối viết 1 điều bạn làm tốt hơn hôm qua, dù nhỏ"},

		// ---- ANXIETY / LO LANG (10) ----
		{"anxiety", "lo lang hoi hop truoc khi thi bai thuyet trinh", "Hồi hộp trước sự kiện quan trọng là phản ứng bình thường của cơ thể - đó là năng lượng, không phải yếu đuối. Hãy đổi góc nhìn: 'Tôi đang hứng khởi' thay vì 'Tôi đang lo'. Nghiên cứu cho thấy cách đặt tên cảm xúc này thực sự cải thiện hiệu suất.", "Thực hành 'power pose' - đứng thẳng, hai tay chống hông 2 phút trước khi vào phòng thi"},
		{"anxiety", "lo lang khong biet nguoi khac nghi gi ve minh", "Lo lắng về đánh giá của người khác (social anxiety) là một trong những nỗi lo phổ biến nhất ở tuổi học sinh. Sự thật: người khác đang bận lo cho bản thân họ hơn là để ý đến bạn. Hiệu ứng spotlight - bạn cảm thấy mình bị chú ý nhiều hơn thực tế.", "Khi lo người khác đánh giá, hỏi: 'Bằng chứng nào cho thấy họ đang phán xét tôi?'"},
		{"anxiety", "lo lang roi loan lo au cam giac kho thu", "Cảm giác lo âu liên tục, khó thở, tim đập nhanh - cơ thể đang ở chế độ 'chiến hay chạy'. Để tắt nó: hít thở theo kỹ thuật 4-7-8 (hít 4 giây, giữ 7 giây, thở ra 8 giây). Đặt tay lên ngực cảm nhận nhịp thở. Nói với bản thân: 'Tôi an toàn ngay lúc này.'", "Nghe âm thanh sóng não Alpha (có trong mục Audio) - đã được chứng minh giảm lo âu hiệu quả"},
		{"anxiety", "lo lang ve suc khoe co benh khong", "Lo lắng về sức khoẻ là tín hiệu bạn cần chú ý hơn đến cơ thể. Hãy kiểm tra: bạn đã ngủ đủ giấc chưa? Uống đủ nước? Ăn uống ổn không? Nếu triệu chứng kéo dài hơn 2 tuần, hãy đến gặp bác sĩ - đừng tự chẩn đoán trên mạng, thường chỉ khiến lo thêm.", "Ghi nhật ký triệu chứng: ghi lại khi nào xuất hiện, kéo dài bao lâu - giúp bác sĩ chẩn đoán chính xác hơn"},
		{"anxiety", "lo lang khong ngu duoc dem truoc thi", "Đêm trước thi mà không ngủ được? Đây là điều rất nhiều bạn gặp. Tin tốt: nghỉ nằm yên cũng giúp cơ thể phục hồi, dù không ngủ. Đừng cố ép bản thân ngủ - áp lực sẽ làm ngược lại. Thay vào đó, thử thả lỏng từng phần cơ thể từ chân lên đầu.", "Đọc sách nhàm chán (sách kỹ thuật, không phải tiểu thuyết) - não sẽ tìm cách ngủ để thoát khỏi nhàm chán"},
		{"anxiety", "lo lang bi tu choi bi phan xet bi chi trich", "Sợ bị phán xét hoặc từ chối là một nỗi sợ rất con người. Nhưng hãy nhớ: mỗi lần bị từ chối là bạn đang luyện tập khả năng chịu đựng và phục hồi. Người thành công nhất thường là người bị từ chối nhiều nhất và vẫn tiếp tục.", "Thử 'liệu pháp từ chối': mỗi ngày chủ động xin một điều nhỏ có khả năng bị từ chối - xin giảm giá, xin ưu tiên"},
		{"anxiety", "lo lang qua khong lam duoc gi cam giac te liet", "Khi lo âu làm tê liệt không làm được gì, đó gọi là 'analysis paralysis'. Cách thoát: thực hiện 'quy tắc 5 phút' - chỉ cần làm 5 phút, sau đó có thể dừng. Hầu hết mọi người thấy mình tiếp tục khi đã bắt đầu, vì bắt đầu là phần khó nhất.", "Đặt hẹn giờ 5 phút, làm bất kỳ phần nào nhỏ nhất của việc cần làm"},
		{"anxiety", "lo lang ve gia dinh bo me ca nnhau", "Lo lắng cho gia đình đang gặp khó khăn là gánh nặng không đáng có ở vai bạn. Bạn không thể giải quyết vấn đề của người lớn, nhưng bạn có thể kiểm soát phản ứng của mình. Hãy tập trung vào những gì trong tầm tay: học tốt, chăm sóc bản thân, hiện diện khi gia đình cần.", "Tìm một người lớn đáng tin cậy để nói chuyện - thầy cô tâm lý học đường có thể giúp"},
		{"anxiety", "lo lang thi truot dai hoc khong vao duoc", "Sợ trượt đại học là nỗi sợ có thật. Nhưng hãy nhìn rộng hơn: đại học là một con đường, không phải con đường duy nhất. Nhiều người thành công không học đại học hoặc vào trường không tên tuổi. Điều quan trọng hơn là bạn học gì và làm gì với nó.", "Lập kế hoạch B: nếu không vào trường mơ ước, mình sẽ làm gì? Có kế hoạch dự phòng giảm lo âu đáng kể"},
		{"anxiety", "lo lang mang xa hoi so sanh ban ban hoc gioi hon", "Mạng xã hội là highlight reel - bạn đang so sánh cuộc sống thực của mình với màn trình diễn tốt nhất của người khác. Đó là cuộc chiến không công bằng. Thử 'digital detox' 24 giờ mỗi tuần - nhiều bạn thấy mức lo âu giảm đáng kể chỉ sau vài tuần.", "Ẩn hoặc bỏ theo dõi tài khoản khiến bạn cảm thấy tệ về bản thân mình"},

		// ---- MOTIVATION / DONG LUC (8) ----
		{"motivation", "mat dong luc chán hoc khong muon hoc nua", "Mất động lực học tập thường xảy ra khi bạn không thấy kết nối giữa việc đang học và điều bạn thực sự quan tâm. Hãy thử 'kỹ thuật tại sao': viết ra lý do học môn này có ích gì cho mục tiêu của bạn. Nếu không tìm ra lý do, đó là tín hiệu cần xem lại định hướng.", "Tìm 1 ứng dụng thực tế của môn đang học trong cuộc sống - YouTube hay Google đều giúp được"},
		{"motivation", "khong co muc tieu khong biet muon gi", "Không biết mình muốn gì là trạng thái rất nhiều học sinh gặp - và đó không phải vấn đề, đó là cơ hội khám phá. Hãy thử nhiều thứ mới trong 3 tháng: tham gia câu lạc bộ, học kỹ năng mới, đọc sách nhiều thể loại. Sở thích không tự nhiên xuất hiện - chúng phát triển qua trải nghiệm.", "Thử 'thí nghiệm 30 ngày': mỗi tháng thử một điều mới hoàn toàn - nấu ăn, code, vẽ, nhạc cụ"},
		{"motivation", "lười biếng cứ trì hoãn không làm việc", "Trì hoãn thường không phải lười biếng - đó thường là sợ hãi (thất bại, hoàn hảo, bị phán xét). Não bạn đang né tránh cảm giác khó chịu. Giải pháp: làm cho bắt đầu dễ đến mức không thể từ chối - mở sách ra, chưa cần đọc. Ngồi vào bàn học, chưa cần làm gì.", "Quy tắc 2 phút: nếu việc gì làm được trong 2 phút, làm ngay. Hành động tạo ra động lực, không phải ngược lại."},
		{"motivation", "choi game nhieu qua bo hoc nghien game", "Nghiện game hoặc giải trí quá mức thường là cách não bộ tìm kiếm cảm giác thành tích và kiểm soát - những thứ mà học tập đôi khi không cho ngay lập tức. Thay vì cấm hoàn toàn (thường thất bại), hãy tạo quy tắc rõ ràng: game sau khi xong việc, có giới hạn thời gian.", "Thử 'gamification' việc học: đặt điểm, level, reward cho bản thân giống như trong game"},
		{"motivation", "cam thay vo nghia khong biet hoc de lam gi", "Cảm giác vô nghĩa trong học tập thường đến từ việc học theo yêu cầu người khác mà không kết nối với giá trị bản thân. Hỏi mình: điều gì khiến bạn tức giận với thế giới này? Điều gì bạn muốn thay đổi? Đó thường là manh mối cho nghề nghiệp và mục đích.", "Đọc sách 'Ikigai' hoặc xem TED Talk của Simon Sinek 'Start With Why' - 18 phút thay đổi cách nhìn"},
		{"motivation", "that bai qua nhieu nan long bo cuoc", "Thất bại nhiều lần dễ làm nản lòng. Nhưng mọi kỹ năng đều có đường cong học tập - ban đầu luôn khó. Thomas Edison thử 10.000 lần trước khi có bóng đèn. Hãy tách biệt 'thất bại trong việc này' khỏi 'tôi là kẻ thất bại' - đó là hai điều hoàn toàn khác nhau.", "Viết ra 3 thất bại lớn nhất và điều bạn học được từ mỗi cái - đây là tài sản quý giá"},
		{"motivation", "khong co nguoi ung ho cam giac mot minh", "Cảm giác thiếu sự ủng hộ và một mình trong hành trình học tập rất nặng nề. Hãy chủ động tìm cộng đồng: nhóm học tập, diễn đàn online, câu lạc bộ ở trường. Có những người đang đi cùng hướng với bạn - bạn chỉ cần tìm họ.", "Tham gia 1 nhóm học tập online (Discord, Facebook group) về lĩnh vực bạn quan tâm"},
		{"motivation", "ganh ty nguoi khac thanh cong hon cam giac kho chiu", "Ghen tị là cảm xúc rất con người - nó cho bạn biết điều bạn thực sự muốn. Thay vì xấu hổ về cảm xúc này, hãy khai thác nó: người bạn ghen tị có gì bạn muốn? Điều đó có thể thành hiện thực với bạn không? Nếu có, đó là hướng đi. Nếu không, hãy xem lại xem có phải đó thực sự là điều BẠN muốn.", "Biến người bạn ngưỡng mộ thành hình mẫu (role model) thay vì đối thủ"},

		// ---- TAP TRUNG / FOCUS (8) ----
		{"focus", "khong tap trung duoc trong lop hoc o truong", "Khó tập trung trong lớp có thể do nhiều nguyên nhân: mệt, đói, điện thoại, hoặc nội dung quá khó/dễ. Thử 'active listening': thay vì ngồi thụ động, đặt câu hỏi trong đầu về nội dung thầy cô đang dạy. Ghi chép tay (không phải gõ) cũng tăng đáng kể khả năng ghi nhớ.", "Ngồi bàn đầu hoặc gần thầy cô - không gian vật lý ảnh hưởng lớn đến sự tập trung"},
		{"focus", "hay bi phan tam boi dien thoai mang xa hoi", "Điện thoại được thiết kế để gây nghiện - đây là cuộc chiến không cân sức. Đừng cố 'tự kiểm soát', hãy thay đổi môi trường: để điện thoại ở phòng khác khi học. Khoảng cách vật lý hiệu quả hơn ý chí nhiều lần.", "App Forest hoặc Focus@Will: trồng cây ảo khi học, cây chết nếu mở điện thoại - gamification cho tập trung"},
		{"focus", "dau oc nghi nhieu suy nghi nhieu khi co gang hoc", "Tâm trí lang thang (mind wandering) xảy ra tới 47% thời gian thức - hoàn toàn bình thường. Kỹ thuật: khi nhận ra mình đang mơ màng, đừng tự trách, chỉ nhẹ nhàng đưa sự chú ý trở lại. Luyện tập này chính là thiền định, và nó tăng dần theo thời gian.", "Thực hành 'thở có ý thức' 3 phút trước khi học: đếm hơi thở từ 1 đến 10, lặp lại"},
		{"focus", "hoc duoc mot luc la quen ngay mat tap trung", "Trí nhớ ngắn hạn có dung lượng hạn chế (7±2 đơn vị). Khi đầy, thông tin mới bị đẩy ra. Giải pháp: ôn lại sau 10 phút, 1 ngày, 3 ngày, 1 tuần (spaced repetition). App Anki làm điều này tự động. Ghi chú ngay sau học, không đợi sau.", "Sau mỗi 25 phút học, dành 5 phút ghi lại điểm chính bằng lời của mình - không nhìn sách"},
		{"focus", "khong gian hoc tap on ao nhieu nguoi khong yên tinh", "Môi trường học tập ảnh hưởng trực tiếp đến hiệu suất. Nếu không có không gian yên tĩnh ở nhà, hãy thử: thư viện trường, quán cà phê yên tĩnh, tai nghe chống ồn với nhạc không lời. Âm nhạc không có lời (lofi hip-hop, classical) giúp nhiều người tập trung hơn.", "Tạo 'ritual' vào học: ngồi đúng chỗ, uống nước, đeo tai nghe - não sẽ học được: đây là lúc tập trung"},
		{"focus", "hay ngu gat buon ngu khi hoc bai", "Buồn ngủ khi học thường do: thiếu ngủ đêm trước, ăn quá no, hoặc học thụ động quá lâu. Giải pháp tức thì: đứng dậy đi lại 5 phút, uống nước lạnh, hít thở sâu. Về lâu dài: ngủ đủ 7-8 tiếng là nền tảng, không có gì thay thế được.", "Thử 'power nap' 20 phút sau bữa trưa - đặt báo thức 20 phút, không hơn (nếu hơn sẽ bị groggy)"},
		{"focus", "hoc nhieu mon cung mot luc khong biet uu tien gi", "Học nhiều môn đồng thời mà không ưu tiên dẫn đến 'task-switching' liên tục - tiêu hao năng lượng não rất nhiều. Nguyên tắc: làm XONG một nhiệm vụ trước khi chuyển sang cái khác. Ưu tiên theo: deadline gần nhất + độ quan trọng.", "Tạo lịch học theo khối: sáng môn khó, chiều môn dễ hơn, tối ôn lại - phù hợp với nhịp sinh học"},
		{"focus", "adhd kho tap trung chuan doan roi loan tang dong", "ADHD không phải yếu kém - đó là não bộ hoạt động khác biệt. Nhiều người ADHD rất thành công khi tìm được môi trường và phương pháp phù hợp. Thử: học trong khoảng ngắn hơn (15-20 phút), vận động giữa các phiên, dùng body doubling (học cùng người khác), và ghi chép màu sắc.", "Tham khảo bác sĩ hoặc chuyên gia tâm lý để được đánh giá và hỗ trợ chính thức nếu cần"},

		// ---- NGU / SLEEP (6) ----
		{"sleep", "ngu khong duoc mat ngu kho ngu", "Mất ngủ ảnh hưởng trực tiếp đến học tập - chỉ cần thiếu 1-2 tiếng, khả năng ghi nhớ và tập trung giảm đáng kể. Vệ sinh giấc ngủ: ngủ và thức dậy cùng giờ mỗi ngày (kể cả cuối tuần), tắt màn hình 1 tiếng trước khi ngủ, giữ phòng mát và tối.", "Nghe âm thanh sóng não Theta (trong mục Audio) - được thiết kế đặc biệt để hỗ trợ giấc ngủ"},
		{"sleep", "nghi dem qua lo lang kho di vao giac ngu", "Suy nghĩ quá nhiều khi nằm xuống là vòng lặp rất phổ biến. Thử 'worry dump': trước khi ngủ 30 phút, viết ra TẤT CẢ lo lắng đang có, kèm hành động cụ thể sẽ làm ngày mai. Não sẽ không cần 'nhắc nhở' bạn nữa vì đã được ghi lại.", "Kỹ thuật 4-7-8: hít 4 giây, giữ 7 giây, thở ra 8 giây - kích hoạt hệ thần kinh phó giao cảm"},
		{"sleep", "ngu qua nhieu van met ngu ngon nhung khong cam thay nghỉ ngu", "Ngủ nhiều mà vẫn mệt có thể do: chất lượng giấc ngủ kém (ngủ nông, hay tỉnh), thiếu sắt/vitamin D, trầm cảm nhẹ, hoặc ngủ sai giờ. Nếu kéo dài hơn 2 tuần, nên gặp bác sĩ để kiểm tra.", "Theo dõi giấc ngủ bằng app (Sleep Cycle, Google Fit) để xem thực sự ngủ bao nhiêu và chất lượng thế nào"},
		{"sleep", "thuc khuya quen thuc khuya kho di ngu som", "Thức khuya là thói quen - và thói quen có thể thay đổi. Nhưng cần thời gian, không thể đột ngột. Chiến lược: lùi giờ ngủ 15 phút mỗi tuần (không phải đột ngột 2-3 tiếng). Ánh sáng xanh từ màn hình ức chế melatonin - bật chế độ night mode từ 8pm.", "Tạo 'wind-down routine' 30 phút trước ngủ: đọc sách giấy, nghe nhạc nhẹ, tắm nước ấm"},
		{"sleep", "ngu ngay khi ve nha xong toi lai thuc khong ngu duoc", "Ngủ ngày nhiều làm lệch đồng hồ sinh học. Nếu cần ngủ trưa, giới hạn 20-30 phút trước 3pm. Hãy cố ngủ và thức đúng giờ ít nhất 5 ngày liên tiếp - não cần thời gian thiết lập lại nhịp sinh học.", "Tập thể dục buổi sáng 15-20 phút - ánh sáng mặt trời và vận động là 'đồng hồ sinh học' mạnh nhất"},
		{"sleep", "ac mong thường xuyen giac ngu khong yen", "Ác mộng thường xuyên là dấu hiệu stress hoặc lo âu cao. Chúng là cách não xử lý cảm xúc chưa được giải quyết ban ngày. Thử 'Image Rehearsal Therapy': viết lại kết thúc của giấc mơ theo hướng tích cực khi tỉnh dậy và đọc lại trước khi ngủ.", "Nếu ác mộng liên quan đến sự kiện traumatic cụ thể, nên tìm chuyên gia tâm lý hỗ trợ"},

		// ---- CO DON / RELATIONSHIP (6) ----
		{"loneliness", "co don khong co ban be cam giac bi loai tru", "Cô đơn không phải lỗi của bạn - nó thường là tín hiệu bạn cần kết nối sâu hơn, không chỉ nhiều hơn. Chất lượng quan trọng hơn số lượng. Thử tham gia hoạt động dựa trên sở thích chung - đây là nơi tốt nhất để tìm bạn thực sự.", "Bắt đầu với 'proximity friendship': người ngồi cạnh trong lớp, bạn cùng câu lạc bộ - quen mặt là bước đầu"},
		{"loneliness", "bi ban be xa la bo roi khong con thân", "Mất đi một tình bạn quan trọng đau không kém chia tay. Cho phép mình buồn - đây là mất mát thực sự. Nhưng nhớ rằng: mọi tình bạn đều có thời của nó. Bạn xứng đáng có những người bạn thực sự coi trọng bạn.", "Đừng cố giành lại tình bạn đã mất - hãy đầu tư năng lượng vào những người đang hiện diện"},
		{"loneliness", "cam thay khac biet khong ai hieu minh", "Cảm giác không ai hiểu mình thường đến từ việc chưa tìm được 'bộ lạc' của mình - những người chia sẻ giá trị và sở thích tương tự. Internet đã mở ra khả năng tìm kiếm rộng hơn nhiều. Có những cộng đồng cho hầu hết mọi sở thích và cách suy nghĩ.", "Thử diễn đàn Reddit, Discord server về sở thích của bạn - nhiều người bắt đầu từ đây"},
		{"loneliness", "yeu xa that bai tinh yeu dau khi chia tay", "Chia tay và mất đi người yêu là một trong những nỗi đau tâm lý mạnh nhất. Não trải qua phản ứng giống như cai nghiện về mặt sinh hóa. Cho phép mình đau trong một khoảng thời gian, nhưng đặt giới hạn: đừng xem lại ảnh, hạn chế theo dõi mạng xã hội của người cũ.", "Vận động thể chất giải phóng endorphin - chạy bộ, bơi lội, gym đặc biệt hiệu quả sau chia tay"},
		{"loneliness", "mau thuan voi thay co giao vien bat cong", "Xung đột với thầy cô có thể rất căng thẳng vì sự mất cân bằng quyền lực. Trước tiên, hãy thử hiểu góc nhìn của thầy cô. Nếu bạn tin mình bị đối xử không công bằng, hãy ghi chép sự kiện cụ thể và nói chuyện với phụ huynh hoặc cố vấn học đường.", "Tránh đối đầu trực tiếp trước lớp - chọn nói chuyện riêng, lịch sự nhưng rõ ràng"},
		{"loneliness", "mau thuan voi gia dinh bo me khong hieu", "Xung đột thế hệ với cha mẹ là phổ biến - họ lớn lên trong thế giới rất khác. Thay vì phán xét nhau, hãy tìm điểm chung: cả hai đều muốn bạn hạnh phúc và thành công, chỉ khác về phương pháp. Thử lắng nghe quan điểm của họ trước khi bảo vệ quan điểm của mình.", "Chọn thời điểm tốt để nói chuyện (không phải lúc ai đó đang mệt hay bực bội)"},

		// ---- TU TI / SELF-ESTEEM (6) ----
		{"self-esteem", "tu ti kem cam thay ban than kem coi khong gioi gi", "Tự ti thường xuất phát từ so sánh không công bằng với người khác hoặc tiêu chuẩn không thực tế. Hãy nhớ: bạn đang so sánh highlight của người khác với behind-the-scenes của mình. Mỗi người có điểm mạnh khác nhau - nhiệm vụ là tìm ra của bạn.", "Viết danh sách 10 điều bạn làm được tốt hơn 90% người xung quanh - ai cũng có"},
		{"self-esteem", "tu ti ngoai hinh body shame khong thich co the", "Không hài lòng với ngoại hình là một trong những nguồn tự ti phổ biến nhất, đặc biệt ở tuổi học sinh. Nhưng hãy nhớ: tiêu chuẩn 'đẹp' trên mạng là phi thực tế (filter, góc chụp, photoshop). Cơ thể bạn đang làm việc tuyệt vời để giữ bạn sống và hoạt động.", "Thực hành 'body gratitude': mỗi ngày cảm ơn 1 phần cơ thể vì đã làm tốt việc của nó"},
		{"self-esteem", "tu ti vi hoc kem hon ban be diem thap", "Điểm số không đo lường giá trị bạn như một con người. Chúng đo lường khả năng tái hiện thông tin trong một bối cảnh cụ thể. Nhiều người điểm số không cao nhưng rất thành công vì họ có kỹ năng khác - sáng tạo, giao tiếp, kiên trì.", "Tìm môn hoặc hoạt động bạn giỏi, đầu tư vào đó - thành công dù nhỏ xây dựng lại tự tin"},
		{"self-esteem", "cam thay minh khong xung dang voi tinh cam gia dinh", "Cảm giác không xứng đáng được yêu thương là dấu hiệu của tổn thương cảm xúc sâu. Đây không phải sự thật - đây là câu chuyện mà não bạn đã học từ những trải nghiệm đau trong quá khứ. Bạn xứng đáng được yêu thương chỉ vì bạn là con người.", "Liệu pháp nhận thức (CBT) rất hiệu quả cho vấn đề này - tìm chuyên gia tâm lý để được hỗ trợ"},
		{"self-esteem", "tu phe binh ban than qua khac nghiet hay tu tranh phac minh", "Tự phê bình quá mức là dấu hiệu của inner critic mạnh. Thử bài tập: khi bạn nói với bản thân điều gì đó khắc nghiệt, hỏi 'Tôi có nói điều này với người bạn thân không?' Nếu không, đừng nói với bản thân mình.", "Thực hành 'self-compassion': đối xử với mình như với người bạn thân nhất đang gặp khó khăn"},
		{"self-esteem", "cam thay vo dung khong co gi dat duoc", "Cảm giác vô dụng thường che giấu những kỳ vọng quá cao hoặc tiêu chuẩn không thực tế. Hãy nhìn lại: bạn đã đi được bao xa từ điểm xuất phát? So sánh với chính mình 1 năm trước, không phải với người khác.", "Tạo 'bảng thành tựu': ghi lại MỌI điều nhỏ bạn hoàn thành - từ hoàn thành bài tập đến giúp bạn bè"},

		// ---- TRAM CAM / DEPRESSION (6) ----
		{"depression", "buon khong ly do cam giac trong rong trong long", "Cảm giác buồn không lý do và trống rỗng kéo dài có thể là dấu hiệu trầm cảm nhẹ. Điều quan trọng: đừng chiến đấu một mình. Chia sẻ với một người bạn tin tưởng hoặc chuyên gia tâm lý. Cảm giác này có thể điều trị được.", "Vận động nhẹ 15 phút mỗi ngày - hiệu quả như thuốc chống trầm cảm nhẹ theo một số nghiên cứu"},
		{"depression", "khoc khong ro nguyen nhan hay khoc cam thay te", "Khóc không vì lý do rõ ràng là cách cơ thể giải phóng cảm xúc tích tụ. Đừng ngăn nước mắt - hãy để chúng chảy. Sau khi khóc xong, thử viết ra bất kỳ cảm xúc hoặc suy nghĩ nào xuất hiện - thường sẽ tìm ra nguyên nhân sâu xa.", "Nếu khóc không kiểm soát được và kéo dài nhiều tuần, hãy tìm chuyên gia tâm lý"},
		{"depression", "mat hang moi thu khong cam thay gi nua te liet", "Mất hứng thú với mọi thứ từng thích (anhedonia) là triệu chứng quan trọng của trầm cảm. Đây là tín hiệu cần được chú ý và hỗ trợ chuyên nghiệp. Bạn không yếu đuối - đây là vấn đề y tế có thể điều trị được.", "Hãy nói chuyện với người lớn đáng tin cậy ngay hôm nay - không cần phải đợi đến khi 'đủ tệ'"},
		{"depression", "suy nghi tieu cuc khong kiem soat duoc", "Suy nghĩ tiêu cực lặp đi lặp lại (rumination) có thể trở thành vòng lặp khó thoát. Kỹ thuật nhận thức: đặt câu hỏi với suy nghĩ tiêu cực: 'Bằng chứng nào ủng hộ suy nghĩ này? Bằng chứng nào chống lại?'. Suy nghĩ không phải sự thật - chỉ là suy nghĩ.", "Thiền mindfulness 10 phút/ngày (app Headspace hoặc Insight Timer) - giúp quan sát suy nghĩ mà không bị kéo đi"},
		{"depression", "nghi den tu tu tu thuong suy nghi ve cai chet", "Nếu bạn đang có suy nghĩ về tự làm hại bản thân hoặc tự tử, đây là tình huống khẩn cấp cần được hỗ trợ ngay. Bạn không phải đối mặt một mình. Hãy gọi ngay đường dây hỗ trợ sức khỏe tâm thần, hoặc nói với người lớn đáng tin cậy ngay bây giờ.", "Đường dây hỗ trợ khủng hoảng tâm thần Việt Nam: 1800 599 920 (miễn phí, 24/7)"},
		{"depression", "cam thay tuyet vong khong con hy vong gi nua", "Cảm giác tuyệt vọng và không thấy tương lai là triệu chứng nghiêm trọng cần được hỗ trợ chuyên nghiệp. Nhưng hãy nhớ: cảm giác tuyệt vọng là triệu chứng của trầm cảm, không phải phản ánh thực tế. Khi được điều trị, tương lai sẽ khác.", "Nói chuyện với chuyên gia tâm lý học đường hoặc gọi đường dây hỗ trợ ngay hôm nay"},
	}

	// Build single bulk INSERT for all scenarios (avoids N round trips to remote DB)
	if len(scenarios) == 0 {
		return nil
	}

	// Build parameterized bulk INSERT
	valueStrings := make([]string, 0, len(scenarios))
	valueArgs := make([]interface{}, 0, len(scenarios)*4)
	for i, s := range scenarios {
		base := i * 4
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d)", base+1, base+2, base+3, base+4))
		valueArgs = append(valueArgs, s.category, s.trigger, s.response, s.tips)
	}

	// Wrap DELETE + INSERT in a transaction so a failed INSERT does not leave an empty table.
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seed scenarios begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, "DELETE FROM psych_scenarios"); err != nil {
		return fmt.Errorf("seed scenarios delete: %w", err)
	}

	query := `INSERT INTO psych_scenarios (category, trigger, response, tips) VALUES ` + strings.Join(valueStrings, ",")
	if _, err := tx.Exec(ctx, query, valueArgs...); err != nil {
		return fmt.Errorf("bulk scenario insert failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("seed scenarios commit: %w", err)
	}
	return nil
}

func (db *DB) MigrateAuth(ctx context.Context) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(30) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			display_name VARCHAR(100) NOT NULL DEFAULT '',
			created_at TIMESTAMP DEFAULT NOW()
		)`,
		`CREATE TABLE IF NOT EXISTS direct_messages (
			id SERIAL PRIMARY KEY,
			sender_id INT NOT NULL REFERENCES users(id),
			receiver_id INT NOT NULL REFERENCES users(id),
			content TEXT NOT NULL,
			is_read BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT NOW()
		)`,
	}
	for _, q := range queries {
		if _, err := db.pool.Exec(ctx, q); err != nil {
			return fmt.Errorf("auth migration error: %w", err)
		}
	}
	return nil
}

func (db *DB) CreateUser(ctx context.Context, username, email, passwordHash, displayName string) (*models.User, error) {
	var u models.User
	err := db.pool.QueryRow(ctx,
		`INSERT INTO users (username, email, password_hash, display_name) VALUES ($1, $2, $3, $4) RETURNING id, username, email, password_hash, display_name, created_at`,
		username, email, passwordHash, displayName,
	).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.DisplayName, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *DB) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var u models.User
	err := db.pool.QueryRow(ctx,
		`SELECT id, username, email, password_hash, display_name, created_at FROM users WHERE username = $1`,
		username,
	).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.DisplayName, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (db *DB) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User
	err := db.pool.QueryRow(ctx,
		`SELECT id, username, email, password_hash, display_name, created_at FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.DisplayName, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (db *DB) SaveMessage(ctx context.Context, senderID, receiverID int, content string) (*models.DirectMessage, error) {
	var m models.DirectMessage
	err := db.pool.QueryRow(ctx,
		`INSERT INTO direct_messages (sender_id, receiver_id, content) VALUES ($1, $2, $3) RETURNING id, sender_id, receiver_id, content, is_read, created_at`,
		senderID, receiverID, content,
	).Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.IsRead, &m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (db *DB) GetConversation(ctx context.Context, userID, otherUserID int) ([]models.DirectMessage, error) {
	rows, err := db.pool.Query(ctx,
		`SELECT id, sender_id, receiver_id, content, is_read, created_at FROM direct_messages WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1) ORDER BY created_at ASC`,
		userID, otherUserID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var msgs []models.DirectMessage
	for rows.Next() {
		var m models.DirectMessage
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.IsRead, &m.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if msgs == nil {
		msgs = []models.DirectMessage{}
	}
	return msgs, nil
}

func (db *DB) GetUserList(ctx context.Context, excludeID int) ([]models.User, error) {
	rows, err := db.pool.Query(ctx,
		`SELECT id, username, email, display_name, created_at FROM users WHERE id != $1 ORDER BY username ASC`,
		excludeID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.DisplayName, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if users == nil {
		users = []models.User{}
	}
	return users, nil
}

/*
# UserのRepository層
  - Create, Read操作を定義
*/
package repository

// /*
// UserRepository Interface
//   - Mockするので抽象化
// */
// type UserRepository interface {
// 	SelectUser(input dto.SelectUser) (user *dto.User, err error)
// 	CreateUser(input dto.CreateUser) (err error)
// }

// // UserRepository 構造体
// type userRepository struct {
// 	db gorm.DB
// }

// // インスタンスの取得?
// func NewUserRepository(db gorm.DB) UserRepository {
// 	return &userRepository{db}
// }

// /*
// Read: UserをuserIDで指定して1件取得する (nameからuserID等を取得すべきかもしれない)
//   - input:
//   - UserID string `json:"userId"`
//   - return:
//   - User 1件
//   - Error:
//   - STATUS_SERVICE_UNAVAILABLE (503)
// */
// func (ur *userRepository) SelectUser(input dto.SelectUser) (*dto.User, error) {
// 	// string->intに変換し，userIDを取得
// 	userId, err := strconv.Atoi(input.UserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// userIDからレコードを取得
// 	record := entity.User{}

// 	query := "SELECT * FROM users WHERE id = ?"
// 	args := []interface{}{userId}

// 	// レコードを割り当てる
// 	result := ur.db.Raw(query, args...).Scan(&record)

// 	if result.Error == gorm.ErrRecordNotFound {
// 		// gormのエラーの種類でユーザーが存在するかどうかわかる
// 		// 意味ないが後学のための残しておく
// 		return nil, STATUS_SERVICE_UNAVAILABLE // 503
// 	} else if result.Error != nil {
// 		return nil, STATUS_SERVICE_UNAVAILABLE // 503
// 	}

// 	// レコードをjson形式のモデルに変換
// 	user := entity.UserFromEntity(&record)
// 	return user, nil
// }

// /*
// Create: Todoを作成する
//   - input:
//   - Name string `json:"name"`
//   - return:
//   - None
//   - Error:
//   - STATUS_SERVICE_UNAVAILABLE (503)
// */
// func (ur *userRepository) CreateUser(input dto.CreateUser) error {
// 	//  ユーザー名は必須にする
// 	if input.Name == "" {
// 		return STATUS_SERVICE_UNAVAILABLE // 503
// 	}

// 	query := `
// 	INSERT INTO users (name, created_at, updated_at)
// 	VALUES (?, ?, ?)
// 	`

// 	args := []interface{}{
// 		input.Name,
// 		time.Now(),
// 		time.Now(),
// 	}

// 	// レコードの作成
// 	result := ur.db.Exec(query, args...)

// 	if result.Error != nil {
// 		return STATUS_SERVICE_UNAVAILABLE // 503
// 	}

// 	return nil
// }

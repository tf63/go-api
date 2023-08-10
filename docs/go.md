## Go周りのメモ

### プロジェクトの作成
プロジェクトの初期化，`go.mod`が作成される
```
    go mod init github.com/tf63/go_api
```

`go.mod`に必要なパッケージを記載
```
    require (
        github.com/99designs/gqlgen v0.17.34
        github.com/gin-gonic/gin v1.9.1
        github.com/stretchr/testify v1.8.4
        github.com/vektah/gqlparser/v2 v2.5.6
        gorm.io/driver/postgres v1.5.2
        gorm.io/gorm v1.25.2
    )
```

`go.mod`からパッケージをインストール -> `go.sum`が作成される
```
    go mod tidy
```


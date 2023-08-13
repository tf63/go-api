### graphql周り

**gqlgenのインストール**
```
    // gqlgenの実行に必要
    go get github.com/99designs/gqlgen@latest
    // コマンドラインでgqlgenを使う (Makefileでインストールすれば良い)
    go install github.com/99designs/gqlgen@latest
```

**GraphQLプロジェクトの作成**

```
    gqlgen init
```
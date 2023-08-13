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

### 動作確認
```
    # ------------------------------------
    # Create User
    # ------------------------------------
    # mutation {
    #   createUser(input :{
    #     name: "graph"
    #   })
    #   {
    #     id,
    #     name,
    #     createdAt,
    #     updatedAt
    #   }
    # }

    # ------------------------------------
    # Read User
    # ------------------------------------
    # query{
    #   user(input: {userId: "1"}){
    #     id,
    #     name,
    #     createdAt,
    #     updatedAt
    #   }
    # }

    # ------------------------------------
    # Read Users
    # ------------------------------------
    # query{
    #   users {
    #     id,
    #     name,
    #     createdAt,
    #     updatedAt
    #   }
    # }


    # ------------------------------------
    # Update User
    # ------------------------------------
    # mutation {
    #   updateUser(input :{
    #     name: "updated graph"
    #   }, userId: 9)
    #   {
    #     id,
    #     name,
    #     createdAt,
    #     updatedAt
    #   }
    # }

    # ------------------------------------
    # Delete User
    # ------------------------------------
    # mutation {
    #   deleteUser(userId: 8)
    #   {
    #     id,
    #     name,
    #     createdAt,
    #     updatedAt
    #   }
    # }


    # ------------------------------------
    # Create Expense
    # ------------------------------------
    # mutation {
    #   createExpense(input :{
    #     title: "graph's payment"
    #     price: 100
    #     userId: "9"
    #   })
    #   {
    #     id,
    #     title,
    #     price,
    #     userId,
    #     createdAt,
    #     updatedAt
    #   }
    # }

    # ------------------------------------
    # Read Expense
    # ------------------------------------
    # query{
    #   expense(input: {userId: 9}, expenseId: 1){
    #     id,
    #     title,
    #     price,
    #     userId,
    #     createdAt,
    #     updatedAt
    #   }
    # }

    # ------------------------------------
    # Read Expenses
    # ------------------------------------
    query{
    expenses(input: {userId: 9}){
        id,
        title,
        price,
        userId,
        createdAt,
        updatedAt
    }
    }

    # ------------------------------------
    # Update Expense
    # ------------------------------------
    # mutation {
    #   updateExpense(input :{
    #     title: "updated graph's payment"
    #     userId: 9
    #   }, expenseId: 1)
    #   {
    #     id,
    #     title,
    #     price,
    #     userId,
    #     createdAt,
    #     updatedAt
    #   }
    # }

    # ------------------------------------
    # Delete Expense
    # ------------------------------------
    # mutation {
    #   deleteExpense(input :{
    #     userId: 9
    #   }, expenseId: 1)
    #   {
    #     id,
    #     title,
    #     price,
    #     userId,
    #     createdAt,
    #     updatedAt
    #   }
    # }

```
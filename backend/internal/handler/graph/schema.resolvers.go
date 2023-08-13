package handler

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/tf63/go_api/api/graph"
	"github.com/tf63/go_api/internal/entity"
)

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input graph.NewExpense) (*graph.Expense, error) {

	inputEntity := NewExpenseDTO(&input)

	expenseId, err := r.Er.CreateExpense(inputEntity)
	if err != nil {
		err = ExpenseErrorHandler("Create", err)
		return nil, err
	}

	findUser := entity.FindUser{ID: inputEntity.UserID}
	expenseEntity, err := r.Er.ReadExpense(findUser, expenseId)
	if err != nil {
		err = ExpenseErrorHandler("Read", err)
		return nil, err
	}

	expense := ExpenseDTO(&expenseEntity)

	return &expense, nil
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, input graph.NewExpense, expenseID uint) (*graph.Expense, error) {
	inputEntity := NewExpenseDTO(&input)

	err := r.Er.UpdateExpense(inputEntity, int(expenseID))
	if err != nil {
		err = ExpenseErrorHandler("Update", err)
		return nil, err
	}

	findUser := entity.FindUser{ID: inputEntity.UserID}
	expenseEntity, err := r.Er.ReadExpense(findUser, int(expenseID))
	if err != nil {
		err = ExpenseErrorHandler("Read", err)
		return nil, err
	}

	expense := ExpenseDTO(&expenseEntity)

	return &expense, nil
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, input graph.FindUser, expenseID uint) (*graph.Expense, error) {
	inputEntity := FindUserDTO(&input)

	// 削除するので先に取得する必要がある
	findUser := entity.FindUser{ID: inputEntity.ID}
	expenseEntity, err := r.Er.ReadExpense(findUser, int(expenseID))
	if err != nil {
		err = ExpenseErrorHandler("Read", err)
		return nil, err
	}

	err = r.Er.DeleteExpense(inputEntity, int(expenseID))
	if err != nil {
		err = ExpenseErrorHandler("Delete", err)
		return nil, err
	}
	expense := ExpenseDTO(&expenseEntity)

	return &expense, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input graph.NewUser) (*graph.User, error) {
	inputEntity := NewUserDTO(&input)

	userId, err := r.Ur.CreateUser(inputEntity)
	if err != nil {
		err = UserErrorHandler("Create", err)
		return nil, err
	}

	userEntity, err := r.Ur.ReadUser(userId)
	if err != nil {
		err = UserErrorHandler("Read", err)
		return nil, err
	}

	user := UserDTO(&userEntity)

	return &user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input graph.NewUser, userID uint) (*graph.User, error) {
	inputEntity := NewUserDTO(&input)

	err := r.Ur.UpdateUser(inputEntity, int(userID))
	if err != nil {
		err = UserErrorHandler("Update", err)
		return nil, err
	}

	userEntity, err := r.Ur.ReadUser(int(userID))
	if err != nil {
		err = UserErrorHandler("Read", err)
		return nil, err
	}

	user := UserDTO(&userEntity)

	return &user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID uint) (*graph.User, error) {
	// 削除するので先に取得する必要がある
	userEntity, err := r.Ur.ReadUser(int(userID))
	if err != nil {
		err = UserErrorHandler("Read", err)
		return nil, err
	}

	err = r.Ur.DeleteUser(int(userID))
	if err != nil {
		err = UserErrorHandler("Delete", err)
		return nil, err
	}

	user := UserDTO(&userEntity)

	return &user, nil
}

// Expense is the resolver for the expense field.
func (r *queryResolver) Expense(ctx context.Context, input graph.FindUser, expenseID uint) (*graph.Expense, error) {
	inputEntity := FindUserDTO(&input)

	expenseEntity, err := r.Er.ReadExpense(inputEntity, int(expenseID))
	if err != nil {
		err = ExpenseErrorHandler("Read", err)
		return nil, err
	}

	expense := ExpenseDTO(&expenseEntity)

	return &expense, nil
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context, input graph.FindUser) ([]*graph.Expense, error) {
	inputEntity := FindUserDTO(&input)

	expensesEntity, err := r.Er.ReadExpenses(inputEntity)
	if err != nil {
		err = ExpenseErrorHandler("Read", err)
		return nil, err
	}

	expenses := []*graph.Expense{}
	for _, expenseEntity := range expensesEntity {
		expense := ExpenseDTO(&expenseEntity)
		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, input graph.FindUser) (*graph.User, error) {

	userEntity, err := r.Ur.ReadUser(int(input.UserID))
	if err != nil {
		err = UserErrorHandler("Read", err)
		return nil, err
	}

	user := UserDTO(&userEntity)

	return &user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*graph.User, error) {

	usersEntity, err := r.Ur.ReadUsers()
	if err != nil {
		err = UserErrorHandler("Read", err)
		return nil, err
	}

	users := []*graph.User{}
	for _, userEntity := range usersEntity {
		user := UserDTO(&userEntity)
		users = append(users, &user)
	}

	return users, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

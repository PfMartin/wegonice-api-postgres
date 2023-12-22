package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	CreateRecipeTx(ctx context.Context, recipeArg CreateRecipeParams, ingredientsArg []CreateRecipeIngredientParams, stepsArg []CreateRecipeStepParams) (CreateRecipeTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// Executes a function within a database transaction
// Rolls back all changes if any action within the transaction fails
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
	}

	return tx.Commit()
}

type CreateRecipeTxResult struct {
	Recipe            Recipe             `json:"recipe"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients"`
	RecipeSteps       []RecipeStep       `json:"recipe_steps"`
}

func (store *SQLStore) CreateRecipeTx(ctx context.Context, recipeArg CreateRecipeParams, ingredientsArg []CreateRecipeIngredientParams, stepsArg []CreateRecipeStepParams) (CreateRecipeTxResult, error) {
	var result CreateRecipeTxResult

	fmt.Println("createRecipeTx started")

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Recipe, err = q.CreateRecipe(context.Background(), recipeArg)
		if err != nil {
			return err
		}

		fmt.Println(result.Recipe)

		for _, ingredient := range ingredientsArg {
			ingredient.RecipeID = result.Recipe.ID
			recipeIngredient, err := q.CreateRecipeIngredient(context.Background(), ingredient)
			if err != nil {
				return err
			}

			result.RecipeIngredients = append(result.RecipeIngredients, recipeIngredient)
		}

		for _, step := range stepsArg {
			step.RecipeID = result.Recipe.ID
			recipeStep, err := q.CreateRecipeStep(context.Background(), step)
			if err != nil {
				return err
			}

			result.RecipeSteps = append(result.RecipeSteps, recipeStep)
		}

		return nil
	})

	return result, err

}

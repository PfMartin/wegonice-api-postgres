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
	DeleteRecipeTx(ctx context.Context, recipeID int64) (DeleteRecipeTxResult, error)
	GetRecipeTx(ctx context.Context, recipeID int64) (GetRecipeTxResult, error)
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

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Recipe, err = q.CreateRecipe(context.Background(), recipeArg)
		if err != nil {
			return err
		}

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

type DeleteRecipeTxResult struct {
	Recipe            Recipe             `json:"recipe"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients"`
	RecipeSteps       []RecipeStep       `json:"recipe_steps"`
}

func (store *SQLStore) DeleteRecipeTx(ctx context.Context, recipeID int64) (DeleteRecipeTxResult, error) {
	var result DeleteRecipeTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.RecipeIngredients, err = q.DeleteRecipeIngredientsByRecipeId(ctx, recipeID)
		if err != nil {
			return err
		}

		result.RecipeSteps, err = q.DeleteRecipeStepsByRecipeId(ctx, recipeID)
		if err != nil {
			return err
		}

		result.Recipe, err = q.DeleteRecipeById(ctx, recipeID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type GetRecipeTxResult struct {
	Recipe            Recipe             `json:"recipe"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients"`
	RecipeSteps       []RecipeStep       `json:"recipe_steps"`
}

func (store *SQLStore) GetRecipeTx(ctx context.Context, recipeID int64) (GetRecipeTxResult, error) {
	var result GetRecipeTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.RecipeIngredients, err = q.ListRecipeIngredientsByRecipeId(ctx, recipeID)
		if err != nil {
			return err
		}

		result.RecipeSteps, err = q.ListRecipeStepsByRecipeId(ctx, recipeID)
		if err != nil {
			return err
		}

		result.Recipe, err = q.GetRecipe(ctx, recipeID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type UpdateRecipeTxResult struct {
	Recipe            Recipe             `json:"recipe"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients"`
	RecipeSteps       []RecipeStep       `json:"recipe_steps"`
}

// func (store *SQLStore) UpdateRecipeTx(ctx context.Context, recipeID int64, recipeArg UpdateRecipeByIdParams, ingredientsArg []UpdateRecipeIngredientByIdParams, stepsArg []UpdateRecipeStepByIdParams) (UpdateRecipeTxResult, error) {
// 	var result UpdateRecipeTxResult

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error {
// 			result.Recipe
// 		}
// 	})
// }

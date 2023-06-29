package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"
)

// Create todo
func (t *todoUseCase) Create(ctx context.Context, todo *models.Todo) (int64, error) {

	// user, err := utils.GetUserFromCtx(ctx)
	// if err != nil {
	// 	return nil, httpResponse.NewUnauthorizedError(errors.WithMessage(err, "newsUC.Create.GetUserFromCtx"))
	// }

	// news.AuthorID = user.UserID

	// if err = utils.ValidateStruct(ctx, news); err != nil {
	// 	return nil, httpResponse.NewBadRequestError(errors.WithMessage(err, "newsUC.Create.ValidateStruct"))
	// }

	res, err := t.todoRepo.Create(ctx, todo)
	if err != nil {
		return 0, err
	}

	return res, err
}

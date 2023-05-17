package service

import (
	"testing"

	"github.com/14mdzk/dev-store/internal/app/mocks"
	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_BrowseAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockRepo := mocks.NewMockICategoryRepository(mockCtrl)
	mockRepo.EXPECT().Browse().Return([]model.Category{
		{
			ID:          1,
			Name:        "Gadget",
			Description: "Collections of laptops, smartphone, etc.",
		},
		{
			ID:          2,
			Name:        "Gadget",
			Description: "Collections of laptops, smartphone, etc.",
		},
	}, nil)

	categoryService := NewCategoryService(mockRepo)
	categories, err := categoryService.BrowseAll()
	total := len(categories)
	assert.Equal(t, total, 1)
	assert.NoError(t, err)
}

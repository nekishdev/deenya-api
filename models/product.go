package models

import (
	"time"
)

// type Storefront struct {
// 	ID   *int64
// 	Name *string
// 	Featured []*int64

// 	Products []*Product `json:"product,omitempty"`
// }

type Inventory struct {
	ID           *int64
	ConsultantID *int64
	ProductID    *int64
	Stock        *int64

	Product Product
} //replace product with product model and use inventory instead?

type ProductModelData struct {
	ID          *int64    `json:"id,omitempty" db:"id" readonly:"true"`
	MediaIDs    []*int64  `json:"media_ids,omitempty" db:"media_ids"`
	Tags        []*string `json:"tags,omitempty" db:"tags"` //or just tags []string?
	Name        *string   `json:"name,omitempty" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	Price       *int64    `json:"price,omitempty" db:"price"`
}

type ProductModel struct {
	ProductModelData
	//add dropshipping option?
	Analytics *ProductAnalytics `json:"analytics,omitempty"`
	Medias    []*Media          `json:"medias,omitempty"`
	Reviews   []*ProductReview  `json:"reviews,omitempty"` //??
}

type ProductAnalytics struct {
	Rating *int64
}

type ProductData struct {
	ID             *int64 `json:"id,omitempty" db:"id" readonly:"true"`
	OwnerID        *int64 `json:"owner_id,omitempty" db:"owner_id"`
	ProductModelID *int64 `json:"product_model_id,omitempty" db:"product_model_id"`
	//MediaIDs []*int64 `json:"media_ids,omitempty" db:"media_ids"`
	//GroupIDs []*int64 `json:"group_id,omitempty" db:"group_ids"` //availability
	// LocationID *int64   `json:"location_id,omitempty" db:"location_id"`
	//Name   *string `json:"name,omitempty" db:"name"`
	Price *int64 `json:"price,omitempty" db:"price"`
	Stock *int64 `json:"stock,omitempty" db:"stock"`
	//Status *string `json:"status,omitempty" db:"status"`

	//IsFeatured  *bool      `json:"is_featured,omitempty" db:"is_featured"`
	//IsPublished *bool      `json:"is_published,omitempty" db:"is_published"`
	IsAvailable *bool      `json:"is_available,omitempty" db:"is_available"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at" readonly:"true"`
	//Medias      []Media    `json:"medias,omitempty"`
	// Location  *Location  `json:"location,omitempty"`
}

//Product - product
type Product struct {
	ProductData
	Model ProductModel
	//Reviews []ProductReview `json:"reviews,omitempty"`
	Owner *User `json:"owner,omitempty"`
}

type ProductReviewData struct {
	ID             *int64
	ProductModelID *int64
	Content        *string
	Rating         *int64
	CreatedAt      *time.Time
	OwnerID        *int64
}

//ProductReview product_review
type ProductReview struct {
	ProductReviewData
	Owner *User
	//Consultant *User
}

type ProductModelSearchRequest struct {
	Query string   `json:"query"`
	Tags  []string `json:"tags"`
}

type ProductModelSuggestRequest struct {
	Tags []string `json:"tags"`
}

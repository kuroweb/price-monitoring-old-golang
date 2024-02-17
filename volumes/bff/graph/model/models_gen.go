// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateProductResult interface {
	IsCreateProductResult()
}

type CreateProductResultErrors interface {
	IsCreateProductResultErrors()
}

type DeleteProductResult interface {
	IsDeleteProductResult()
}

type DeleteProductResultErrors interface {
	IsDeleteProductResultErrors()
}

type Node interface {
	IsNode()
	GetID() string
}

type ResultBase interface {
	IsResultBase()
	GetOk() bool
}

type UpdateProductResult interface {
	IsUpdateProductResult()
}

type UpdateProductResultErrors interface {
	IsUpdateProductResultErrors()
}

type UserError interface {
	IsUserError()
	GetCode() string
	GetMessage() string
}

type CreateProductInput struct {
	Name                     string                               `json:"name"`
	YahooAuctionCrawlSetting *CreateYahooAuctionCrawlSettingInput `json:"yahoo_auction_crawl_setting"`
}

type CreateProductResultError struct {
	Ok    bool                      `json:"ok"`
	Error CreateProductResultErrors `json:"error"`
}

func (CreateProductResultError) IsCreateProductResult() {}

func (CreateProductResultError) IsResultBase()    {}
func (this CreateProductResultError) GetOk() bool { return this.Ok }

type CreateProductResultSuccess struct {
	Ok      bool     `json:"ok"`
	Product *Product `json:"product"`
}

func (CreateProductResultSuccess) IsCreateProductResult() {}

func (CreateProductResultSuccess) IsResultBase()    {}
func (this CreateProductResultSuccess) GetOk() bool { return this.Ok }

type CreateProductResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (CreateProductResultValidationFailed) IsCreateProductResultErrors() {}

func (CreateProductResultValidationFailed) IsUserError()            {}
func (this CreateProductResultValidationFailed) GetCode() string    { return this.Code }
func (this CreateProductResultValidationFailed) GetMessage() string { return this.Message }

type CreateYahooAuctionCrawlSettingInput struct {
	Keyword    string `json:"keyword"`
	CategoryID *int   `json:"category_id,omitempty"`
	MinPrice   int    `json:"min_price"`
	MaxPrice   int    `json:"max_price"`
	Enabled    bool   `json:"enabled"`
}

type DeleteProductResultError struct {
	Ok    bool                      `json:"ok"`
	Error DeleteProductResultErrors `json:"error"`
}

func (DeleteProductResultError) IsDeleteProductResult() {}

func (DeleteProductResultError) IsResultBase()    {}
func (this DeleteProductResultError) GetOk() bool { return this.Ok }

type DeleteProductResultSuccess struct {
	Ok bool `json:"ok"`
}

func (DeleteProductResultSuccess) IsDeleteProductResult() {}

func (DeleteProductResultSuccess) IsResultBase()    {}
func (this DeleteProductResultSuccess) GetOk() bool { return this.Ok }

type DeleteProductResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (DeleteProductResultValidationFailed) IsDeleteProductResultErrors() {}

func (DeleteProductResultValidationFailed) IsUserError()            {}
func (this DeleteProductResultValidationFailed) GetCode() string    { return this.Code }
func (this DeleteProductResultValidationFailed) GetMessage() string { return this.Message }

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Product struct {
	ID                       string                    `json:"id"`
	Name                     string                    `json:"name"`
	YahooAuctionProducts     []*YahooAuctionProduct    `json:"yahooAuctionProducts"`
	YahooAuctionCrawlSetting *YahooAuctionCrawlSetting `json:"yahooAuctionCrawlSetting"`
}

func (Product) IsNode()            {}
func (this Product) GetID() string { return this.ID }

type UpdateProductInput struct {
	Name                     string                               `json:"name"`
	YahooAuctionCrawlSetting *UpdateYahooAuctionCrawlSettingInput `json:"yahoo_auction_crawl_setting"`
}

type UpdateProductResultError struct {
	Ok    bool                      `json:"ok"`
	Error UpdateProductResultErrors `json:"error"`
}

func (UpdateProductResultError) IsUpdateProductResult() {}

func (UpdateProductResultError) IsResultBase()    {}
func (this UpdateProductResultError) GetOk() bool { return this.Ok }

type UpdateProductResultSuccess struct {
	Ok      bool     `json:"ok"`
	Product *Product `json:"product"`
}

func (UpdateProductResultSuccess) IsUpdateProductResult() {}

func (UpdateProductResultSuccess) IsResultBase()    {}
func (this UpdateProductResultSuccess) GetOk() bool { return this.Ok }

type UpdateProductResultValidationFailed struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details []*ErrorDetail `json:"details"`
}

func (UpdateProductResultValidationFailed) IsUpdateProductResultErrors() {}

func (UpdateProductResultValidationFailed) IsUserError()            {}
func (this UpdateProductResultValidationFailed) GetCode() string    { return this.Code }
func (this UpdateProductResultValidationFailed) GetMessage() string { return this.Message }

type UpdateYahooAuctionCrawlSettingInput struct {
	Keyword    string `json:"keyword"`
	CategoryID *int   `json:"category_id,omitempty"`
	MinPrice   int    `json:"min_price"`
	MaxPrice   int    `json:"max_price"`
	Enabled    bool   `json:"enabled"`
}

type YahooAuctionCrawlSetting struct {
	ID         string `json:"id"`
	ProductID  int    `json:"productId"`
	Keyword    string `json:"keyword"`
	MinPrice   int    `json:"minPrice"`
	MaxPrice   int    `json:"maxPrice"`
	CategoryID *int   `json:"categoryId,omitempty"`
	Enabled    bool   `json:"enabled"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func (YahooAuctionCrawlSetting) IsNode()            {}
func (this YahooAuctionCrawlSetting) GetID() string { return this.ID }

type YahooAuctionProduct struct {
	ID             string  `json:"id"`
	ProductID      int     `json:"productId"`
	YahooAuctionID string  `json:"yahooAuctionId"`
	Name           string  `json:"name"`
	ThumbnailURL   string  `json:"thumbnailUrl"`
	Price          int     `json:"price"`
	Published      bool    `json:"published"`
	BoughtDate     *string `json:"boughtDate,omitempty"`
}

func (YahooAuctionProduct) IsNode()            {}
func (this YahooAuctionProduct) GetID() string { return this.ID }

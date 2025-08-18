package cartitemresponse

//	type CartItemDetailResponse struct {
//		CartID       uint     `json:"cart_id"`
//		UserID       uint     `json:"user_id"`
//		UserEmail    string   `json:"user_email"`
//		ProductID    *uint    `json:"product_id,omitempty"`
//		ProductName  *string  `json:"product_name,omitempty"`
//		ProductPrice *float64 `json:"product_price,omitempty"`
//		GroupID      *uint    `json:"group_id,omitempty"`
//		GroupName    *string  `json:"group_name,omitempty"`
//		GroupPrice   *float64 `json:"group_price,omitempty"`
//		Quantity     int      `json:"quantity"`
//	}

type CartItemDetailResponse struct {
	CartID       uint    `json:"cart_id"`
	UserID       uint    `json:"user_id"`
	UserEmail    string  `json:"user_email"`
	ProductID    *uint   `json:"product_id"`
	ProductType  string  `json:"product_type"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`

	GroupID    *uint    `json:"group_id"`
	GroupName  *string  `json:"group_name"`
	GroupPrice *float64 `json:"group_price"`

	Quantity int `json:"quantity"`

	// Book-specific
	BookSubject      *string `json:"book_subject,omitempty"`
	BookLearningArea *string `json:"book_learning_area,omitempty"`
	BookGrade        *string `json:"book_grade,omitempty"`
	BookPublisher    *string `json:"book_publisher,omitempty"`
	BookAuthor       *string `json:"book_author,omitempty"`
	BookISBN         *string `json:"book_isbn,omitempty"`

	// Learning Supply
	LearningBrand    *string `json:"learning_brand,omitempty"`
	LearningMaterial *string `json:"learning_material,omitempty"`

	// Office Supply
	OfficeColor *string `json:"office_color,omitempty"`
	OfficeSize  *string `json:"office_size,omitempty"`
}

// type CartItemDetailResponse struct {
// 	CartID       uint
// 	UserID       uint
// 	UserEmail    string
// 	ProductID    *uint
// 	ProductName  *string
// 	ProductPrice *float64
// 	GroupID      *uint
// 	GroupName    string
// 	GroupPrice   float64
// 	Quantity     int

// 	// Nested group detail
// 	GroupDetail GroupDetail `json:"group_detail,omitempty"`
// }

// type GroupProductDetail struct {
// 	ProductID    uint
// 	ProductName  string
// 	ProductPrice float64
// 	ProductType  string

// 	// Book detail (if type = Book)
// 	Subject      string
// 	LearningArea string
// 	Grade        string
// 	Publisher    string
// 	Author       string
// 	ISBN         string

// 	// LearningSupply detail
// 	Brand    string
// 	Material string

// 	// OfficeSupply detail
// 	Color string
// 	Size  string
// }

// type GroupDetail struct {
// 	GroupID    uint
// 	GroupName  string
// 	GroupPrice float64
// 	Products   []GroupProductDetail
// }

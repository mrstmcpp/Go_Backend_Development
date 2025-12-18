package models

type CreateUserRequestDTO struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type GetUserByIdResponseDTO struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

type UpdateUserByIdRequestDTO struct {
	Name string `json:"name" validate:"min=2,max=50"`
	DOB  string `json:"dob" validate:"datetime=2006-01-02"`
}

package handler

type IDOnlyParam struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

type IDListParam struct {
	IDS []uint `json:"ids" form:"ids" validate:"required"`
}

type EsIDOnlyParm struct {
	ID string `json:"id" form:"id" validate:"required"`
}

type EsIDListParm struct {
	IDS []string `json:"ids" form:"ids" validate:"required"`
}

type PaginateParam struct {
	Offset int `form:"offset" validate:"required,gte=0"`
	Count  int `form:"count" validate:"required,gte=0"`
}

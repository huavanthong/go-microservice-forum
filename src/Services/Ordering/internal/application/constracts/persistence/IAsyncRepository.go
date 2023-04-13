package constracts_persistence

type IAsyncRepository interface {
	GetAllAsync() ([]T, error)
	GetAsync(predicate func(T) bool) ([]T, error)
	GetAsyncWithOptions(predicate func(T) bool, orderBy func(db *gorm.DB) *gorm.DB, includeString string, disableTracking bool) ([]T, error)
	GetAsyncWithIncludes(predicate func(T) bool, orderBy func(db *gorm.DB) *gorm.DB, includes []string, disableTracking bool) ([]T, error)
	GetByIdAsync(id int) (*T, error)
	AddAsync(entity *T) error
	UpdateAsync(entity *T) error
	DeleteAsync(entity *T) error
}

package repositories

type RepositoryBase struct {
	dbContext *OrderContext
}

func NewRepositoryBase(dbContext *OrderContext) *RepositoryBase {
	if dbContext == nil {
		panic("dbContext cannot be nil")
	}

	return &RepositoryBase{dbContext}
}

func (r *RepositoryBase) GetAllAsync() ([]*EntityBase, error) {
	var entities []*EntityBase
	err := r.dbContext.Set(EntityBase{}).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryBase) GetAsync(predicate func(db *gorm.DB) *gorm.DB) ([]*EntityBase, error) {
	var entities []*EntityBase
	query := r.dbContext.Set(EntityBase{})
	query = predicate(query)
	err := query.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryBase) GetAsyncWithOptions(predicate func(db *gorm.DB) *gorm.DB, orderBy string, includes []string) ([]*EntityBase, error) {
	var entities []*EntityBase
	query := r.dbContext.Set(EntityBase{})
	query = predicate(query)
	if len(includes) > 0 {
		for _, include := range includes {
			query = query.Preload(include)
		}
	}
	if orderBy != "" {
		query = query.Order(orderBy)
	}
	err := query.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryBase) GetByIdAsync(id int) (*EntityBase, error) {
	var entity EntityBase
	err := r.dbContext.Set(EntityBase{}).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *RepositoryBase) AddAsync(entity *EntityBase) (*EntityBase, error) {
	err := r.dbContext.Set(EntityBase{}).Create(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *RepositoryBase) UpdateAsync(entity *EntityBase) error {
	err := r.dbContext.Save(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryBase) DeleteAsync(entity *EntityBase) error {
	err := r.dbContext.Delete(entity).Error
	if err != nil {
		return err
	}
	return nil
}

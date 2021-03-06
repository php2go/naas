package dao

import (
	"context"

	"github.com/nilorg/naas/internal/model"
	"github.com/nilorg/naas/internal/pkg/contexts"
	"gorm.io/gorm"
)

// ResourceWebRouter ...
type ResourceWebRouter interface {
	Insert(ctx context.Context, resourceWebRoute *model.ResourceWebRoute) (err error)
	Delete(ctx context.Context, id model.ID) (err error)
	DeleteInIDs(ctx context.Context, ids ...model.ID) (err error)
	Select(ctx context.Context, id model.ID) (resourceWebRoute *model.ResourceWebRoute, err error)
	Update(ctx context.Context, resourceWebRoute *model.ResourceWebRoute) (err error)
	ListByResourceID(ctx context.Context, resourceID model.ID, limit int) (list []*model.ResourceWebRoute, err error)
	ListInIDs(ctx context.Context, ids ...model.ID) (list []*model.ResourceWebRoute, err error)
	ListPaged(ctx context.Context, start, limit int) (list []*model.ResourceWebRoute, total int64, err error)
	ListPagedByResourceServerID(ctx context.Context, start, limit int, resourceServerID model.ID) (list []*model.ResourceWebRoute, total int64, err error)
}

type resourceWebRoute struct {
}

func (*resourceWebRoute) Insert(ctx context.Context, resourceWebRoute *model.ResourceWebRoute) (err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	err = gdb.Create(resourceWebRoute).Error
	return
}
func (*resourceWebRoute) Delete(ctx context.Context, id model.ID) (err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	err = gdb.Delete(&model.ResourceWebRoute{}, id).Error
	return
}
func (*resourceWebRoute) Select(ctx context.Context, id model.ID) (resourceWebRoute *model.ResourceWebRoute, err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	resourceWebRoute = new(model.ResourceWebRoute)
	err = gdb.Model(resourceWebRoute).Where("id = ?", id).Take(resourceWebRoute).Error
	if err != nil {
		resourceWebRoute = nil
	}
	return
}
func (*resourceWebRoute) Update(ctx context.Context, resourceWebRoute *model.ResourceWebRoute) (err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	err = gdb.Model(resourceWebRoute).Save(resourceWebRoute).Error
	return
}

func (*resourceWebRoute) ListByResourceID(ctx context.Context, resourceID model.ID, limit int) (list []*model.ResourceWebRoute, err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	exp := gdb.Model(&model.ResourceWebRoute{}).Where("resource_server_id = ?", resourceID)
	if limit > 0 {
		exp = exp.Offset(0).Limit(limit)
	}
	err = exp.Find(&list).Error
	return
}

func (*resourceWebRoute) ListInIDs(ctx context.Context, ids ...model.ID) (list []*model.ResourceWebRoute, err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	err = gdb.Model(&model.ResourceWebRoute{}).Where("id in ?", ids).Find(&list).Error
	return
}

func (*resourceWebRoute) DeleteInIDs(ctx context.Context, ids ...model.ID) (err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	err = gdb.Where("id in ?", ids).Delete(model.ResourceWebRoute{}).Error
	return
}

func (r *resourceWebRoute) ListPaged(ctx context.Context, start, limit int) (list []*model.ResourceWebRoute, total int64, err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	expression := gdb.Model(&model.ResourceWebRoute{})
	expression.Count(&total)
	err = expression.Offset(start).Limit(limit).Find(&list).Error
	return
}

func (r *resourceWebRoute) ListPagedByResourceServerID(ctx context.Context, start, limit int, resourceServerID model.ID) (list []*model.ResourceWebRoute, total int64, err error) {
	var gdb *gorm.DB
	gdb, err = contexts.FromGormContext(ctx)
	if err != nil {
		return
	}
	expression := gdb.Model(&model.ResourceWebRoute{}).Where("resource_server_id = ?", resourceServerID)
	expression.Count(&total)
	err = expression.Offset(start).Limit(limit).Find(&list).Error
	return
}

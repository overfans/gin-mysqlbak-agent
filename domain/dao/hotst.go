package dao

import (
	"backupAgent/proto/backupAgent/host"
	"context"
	"gorm.io/gorm"
)

type HostDatabase struct {
	Id         int64  `json:"host_id" gorm:"primary_key" description:"自增主键"`
	Host       string `json:"host" gorm:"column:host" description:"任务id"`
	User       string `json:"user"  gorm:"column:user" description:"是否发送钉钉消息"`
	Password   string `json:"password"  gorm:"column:password" description:"accessToken"`
	Content    string `json:"content" gorm:"column:content"`
	HostStatus int64  `json:"host_status" gorm:"column:host_status"`
	IsDeleted  int64  `json:"is_deleted" gorm:"column:is_deleted"`
}

func (h *HostDatabase) TableName() string {
	return "t_host"
}

func (h *HostDatabase) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(h).Error
}

func (h *HostDatabase) Find(ctx context.Context, tx *gorm.DB, search *HostDatabase) (*HostDatabase, error) {
	out := &HostDatabase{}
	err := tx.WithContext(ctx).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (h *HostDatabase) UpdatesStatus(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(h.TableName()).Where("host = ?", h.Host).Updates(map[string]interface{}{
		"host_status": h.HostStatus,
	}).Error
}

func (h *HostDatabase) FindAllHost(ctx context.Context, tx *gorm.DB) ([]HostDatabase, error) {
	var hostlist []HostDatabase
	if err := tx.WithContext(ctx).Where("is_deleted = 0").Find(&hostlist).Error; err != nil {
		return nil, err
	}
	return hostlist, nil
}

func (h *HostDatabase) PageList(ctx context.Context, tx *gorm.DB, params *host.HostListInput) ([]*HostDatabase, int64, error) {
	var total int64 = 0
	var list []*HostDatabase
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(ctx)
	query = query.Table(h.TableName()).Where("is_deleted=0")
	if params.Info != "" {
		query = query.Where("(host like ? )", "%"+params.Info+"%")
	}
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Find(&list).Count(&total)
	return list, total, nil
}

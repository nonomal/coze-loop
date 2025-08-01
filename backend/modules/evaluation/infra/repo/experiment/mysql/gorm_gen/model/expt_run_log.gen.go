// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExptRunLog = "expt_run_log"

// ExptRunLog expt_run_log
type ExptRunLog struct {
	ID            int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey;comment:id" json:"id"`                                                                                         // id
	SpaceID       int64          `gorm:"column:space_id;type:bigint(20) unsigned;not null;uniqueIndex:uk_expt_run,priority:1;index:idx_expt_run_item_turn,priority:1;comment:空间 id" json:"space_id"` // 空间 id
	CreatedBy     string         `gorm:"column:created_by;type:varchar(128);not null;comment:创建者 id" json:"created_by"`                                                                              // 创建者 id
	ExptID        int64          `gorm:"column:expt_id;type:bigint(20);not null;uniqueIndex:uk_expt_run,priority:2;index:idx_expt_run_item_turn,priority:2;comment:实验 id" json:"expt_id"`            // 实验 id
	ExptRunID     int64          `gorm:"column:expt_run_id;type:bigint(20);not null;uniqueIndex:uk_expt_run,priority:3;comment:运行 id" json:"expt_run_id"`                                            // 运行 id
	ItemIds       *[]byte        `gorm:"column:item_ids;type:blob binary;comment:组 ids" json:"item_ids"`                                                                                             // 组 ids
	Mode          *int32         `gorm:"column:mode;type:int(11);comment:模式" json:"mode"`                                                                                                            // 模式
	Status        *int64         `gorm:"column:status;type:bigint(20);comment:状态" json:"status"`                                                                                                     // 状态
	PendingCnt    int32          `gorm:"column:pending_cnt;type:int(11) unsigned;not null;comment:item 未执行数量" json:"pending_cnt"`                                                                    // item 未执行数量
	SuccessCnt    int32          `gorm:"column:success_cnt;type:int(11) unsigned;not null;comment:item 成功数量" json:"success_cnt"`                                                                     // item 成功数量
	FailCnt       int32          `gorm:"column:fail_cnt;type:int(11) unsigned;not null;comment:item 失败数量" json:"fail_cnt"`                                                                           // item 失败数量
	CreditCost    float64        `gorm:"column:credit_cost;type:decimal(15,2);not null;default:0.00;comment:credit 消耗" json:"credit_cost"`                                                           // credit 消耗
	TokenCost     *int64         `gorm:"column:token_cost;type:bigint(20);comment:token 消耗" json:"token_cost"`                                                                                       // token 消耗
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                                                         // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`                                                         // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:删除时间" json:"deleted_at"`                                                                                            // 删除时间
	StatusMessage *[]byte        `gorm:"column:status_message;type:blob binary;comment:提示信息" json:"status_message"`                                                                                  // 提示信息
	ProcessingCnt int32          `gorm:"column:processing_cnt;type:int(11);not null;comment:processing_cnt" json:"processing_cnt"`                                                                   // processing_cnt
	TerminatedCnt int32          `gorm:"column:terminated_cnt;type:int(11);not null;comment:terminated_cnt" json:"terminated_cnt"`                                                                   // terminated_cnt
}

// TableName ExptRunLog's table name
func (*ExptRunLog) TableName() string {
	return TableNameExptRunLog
}

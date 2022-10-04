package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Topics struct {
	// 话题 id
	TopicId uint64 `gorm:"column:topic_id;type:bigint(20) unsigned AUTO_INCREMENT;primaryKey;precision:20;scale:0;not null;autoIncrement;<-:false;index:idx_topic_mmstatus,priority:1,type:BTREE;index:idx_topic_status,priority:1,type:BTREE;index:idx_update_topic,priority:2,type:BTREE"`
	// 话题标题
	Title string `gorm:"column:title;type:varchar(64);size:64;not null;index:title,type:BTREE,unique"`
	// 话题描述
	TopicDesc sql.NullString `gorm:"column:topic_desc;type:varchar(1024);size:1024"`
	// 话题图片，NNCFeedElementPictureInfo pb数据
	TopicPicture sql.NullString `gorm:"column:topic_picture;type:mediumblob;size:16777215"`
	// 锁定状态(0:未锁定, 1:已锁定)
	LockStatus sql.NullInt32 `gorm:"column:lock_status;type:tinyint(4);default:0;precision:3;scale:0"`
	// 0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过
	Status sql.NullInt32 `gorm:"column:status;type:tinyint(4);default:0;precision:3;scale:0"`
	// 关注量
	AttentionCount sql.NullInt32 `gorm:"column:attention_count;type:int(11);default:0;precision:10;scale:0"`
	// 帖子量
	PostCount sql.NullInt32 `gorm:"column:post_count;type:int(11);default:0;precision:10;scale:0"`
	// 精华帖子量
	EssenceCount sql.NullInt32 `gorm:"column:essence_count;type:int(11);default:0;precision:10;scale:0"`
	// 创建牛牛号
	CreateUid sql.NullInt64 `gorm:"column:create_uid;type:bigint(20) unsigned;precision:20;scale:0;autoIncrement:false"`
	// 创建时间
	CreateTime sql.NullInt32 `gorm:"column:create_time;type:int(11);precision:10;scale:0"`
	// 最后编辑牛牛号
	LastEditUid sql.NullInt64 `gorm:"column:last_edit_uid;type:bigint(20) unsigned;precision:20;scale:0;autoIncrement:false"`
	// 最后编辑时间
	LastEditTime sql.NullInt32 `gorm:"column:last_edit_time;type:int(11);precision:10;scale:0"`
	// 操作人（管理员）(最新一次)
	Operator sql.NullString `gorm:"column:operator;type:varchar(128);size:128"`
	// 操作时间
	OperatTime sql.NullTime `gorm:"column:operat_time;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime"`
	//
	IsDeleted int8 `gorm:"column:is_deleted;type:tinyint(1);default:0;precision:3;scale:0;not null"`
	//
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime ON UPDATE CURRENT_TIMESTAMP;default:CURRENT_TIMESTAMP;not null;autoCreateTime;<-:create;index:idx_update_topic,priority:1,type:BTREE"`
	// 是否推荐标志, 0:不推荐 1:推荐
	IsRecommended sql.NullInt32 `gorm:"column:is_recommended;type:tinyint(4);default:0;precision:3;scale:0"`
	// moomoo审核状态(0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过)
	MmStatus sql.NullInt32 `gorm:"column:mm_status;type:tinyint(4);default:0;precision:3;scale:0"`
	// moomoo 精华帖子量
	MmEssenceCount sql.NullInt32 `gorm:"column:mm_essence_count;type:int(11);default:0;precision:10;scale:0"`
	// moomoo帖子量
	MmPostCount sql.NullInt32 `gorm:"column:mm_post_count;type:int(11);default:0;precision:10;scale:0"`
	// mm关注量
	MmAttentionCount sql.NullInt32 `gorm:"column:mm_attention_count;type:int(11);default:0;precision:10;scale:0"`
	// 创建用户来源, 1 牛牛, 2 MooMoo
	CreatorSource sql.NullInt32 `gorm:"column:creator_source;type:int(11);default:1;precision:10;scale:0"`
	// 创建标签员工ID
	EmployeeId sql.NullInt64 `gorm:"column:employee_id;type:bigint(20) unsigned;default:0;precision:20;scale:0;autoIncrement:false"`
}

type FeedTopicRelation struct {
	// 自增id
	Id int32 `gorm:"column:id;type:int(10);default:0;precision:10;scale:0;not null;autoIncrement:false"`
	// feed id
	FeedId uint64 `gorm:"column:feed_id;type:bigint(20) unsigned;primaryKey;precision:20;scale:0;not null;autoIncrement:false;index:idx_topic_feed,priority:2,type:BTREE;index:idx_topic_update,priority:3,type:BTREE"`
	// 话题 id
	TopicId uint64 `gorm:"column:topic_id;type:bigint(20) unsigned;primaryKey;precision:20;scale:0;not null;autoIncrement:false;index:idx_topic_feed,priority:1,type:BTREE;index:idx_topic_update,priority:1,type:BTREE"`
	// 创建牛牛号
	AuthorId sql.NullInt64 `gorm:"column:author_id;type:bigint(20) unsigned;precision:20;scale:0;autoIncrement:false"`
	// 1审核通过，3白名单自动通过 2审核不通过
	Status sql.NullInt32 `gorm:"column:status;type:tinyint(4);precision:3;scale:0"`
	// 是否精华(0: 否, 1: 是)
	IsEssence sql.NullInt32 `gorm:"column:is_essence;type:tinyint(4);default:0;precision:3;scale:0"`
	// 更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime ON UPDATE CURRENT_TIMESTAMP;default:CURRENT_TIMESTAMP;not null;autoCreateTime;<-:create;index:idx_topic_update,priority:2,type:BTREE;index:idx_update,type:BTREE"`
	// MooMoo审核状态 0未审核 1审核通过 2审核不通过 3白名单自动通过 4 被屏蔽 5 不可用
	MmStatus sql.NullInt32 `gorm:"column:mm_status;type:int(11);default:0;precision:10;scale:0"`
	// MooMoo精华标志 0非精华 1为精华
	MmIsEssence int32 `gorm:"column:mm_is_essence;type:int(11);default:0;precision:10;scale:0;not null"`
	// 用户所打标签: 0, 平台所打标签: 1
	TopicCreater int8 `gorm:"column:topic_creater;type:tinyint(4);default:0;precision:3;scale:0;not null"`
}

func (FeedTopicRelation) TableName() string {
	return "topic.feed_topic_relation"
}

func TestGetFeedAuthor(db *gorm.DB) {
	dsn := "root:root@tcp(172.24.31.95:13357)/nnq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("open db failed")
		return
	}
	feedId := 101270754164737
	var feedAuthor uint64
	result := db.Table("index_feeds").Select("author_id").Where("feed_id = ?", feedId).Find(&feedAuthor)
	if result.Error != nil {
		fmt.Printf("failed\n")
		return
	}
	fmt.Printf("successful author: %d\n", feedAuthor)
}

func TestBatchModifyTopicPostCount(db *gorm.DB) {
	postCount := "mm_post_count"
	topicIds := []uint64{1, 2, 3}
	incrVal := -1
	result := db.Model(Topics{}).Where("topic_id in (?)", topicIds).
		Update(postCount, gorm.Expr(fmt.Sprintf("%s + ?", postCount), incrVal))
	if result.Error != nil {
		fmt.Printf("failed\n")
		return
	}
	fmt.Printf("successful a: %d\n", result.RowsAffected)
}

func TestGetFeedTopicRelationByFeedId(db *gorm.DB) {
	feedId := 101270754164737
	var feedTopicRelations []FeedTopicRelation
	result := db.Table("feed_topic_relation").Where("feed_id = ?", feedId).Find(&feedTopicRelations)
	if result.Error != nil {
		fmt.Printf("failed\n")
		return
	}
	fmt.Printf("successful feed_topic_relation: %v\n", feedTopicRelations)
}

func TestAddFeedTopicRelation(db *gorm.DB) {
	var infos []FeedTopicRelation
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&infos)
	if result.Error != nil {
		fmt.Printf("failed reason: %s\n", result.Error.Error())
		return
	}
	fmt.Printf("successful a: %d\n", result.RowsAffected)
}

func TestDelFeedTopicRelation(db *gorm.DB) {
	var feedId uint64 = 123
	topicIds := []uint64{1, 2, 3}

	feedIdString := strconv.FormatUint(feedId, 10)
	sql := "DELETE FROM `topic`.`feed_topic_relation` WHERE ("
	sql += "(feed_id = " + feedIdString +
		" and topic_id = " + strconv.FormatUint(topicIds[0], 10) + ")"
	for i := 1; i < len(topicIds); i++ {
		sql += " or "
		sql += "(feed_id = " + feedIdString + " and topic_id = " + strconv.FormatUint(topicIds[i], 10) + ")"
	}
	sql += ");"

	result := db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("failed reason: %s\n", result.Error.Error())
		return
	}
	fmt.Printf("successful a: %d\n", result.RowsAffected)
}

func TestCreateTopic(db *gorm.DB) {
	// addTopic := Topics{
	// 	TopicId: 27841993459630096,
	// 	Title:   "0803测试03",
	// }
	// result := db.Model(&Topics{}).Create(map[string]interface{}{
	// 	"topic_id": 27841993459630096,
	// 	"title":    "0803测试04",
	// })
	sql := "insert into topics (topic_id, title, topic_desc, topic_picture, create_time, lock_status, status, mm_status" +
		", operat_time, employee_id, operator) values (27841993459630096, '0803测试05', '', '', 0, 0, 0, 0, 0, 0, 0);"
	result := db.Exec(sql)
	if result.Error != nil {
		fmt.Printf("failed reason: %s\n", result.Error.Error())
		return
	}
	fmt.Printf("successful a: %d\n", result.RowsAffected)
}

func main95() {
	dsn := "root:root@tcp(172.24.31.95:13357)/topic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("open db failed")
		return
	}
	// TestGetFeedAuthor(db)
	// TestBatchModifyTopicPostCount(db)
	// TestGetFeedTopicRelationByFeedId(db)
	// TestAddFeedTopicRelation(db)
	// TestDelFeedTopicRelation(db)
	TestCreateTopic(db)
}

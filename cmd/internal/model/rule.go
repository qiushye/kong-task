package model

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	cachec "github.com/zeromicro/go-zero/core/stores/cache"
)

var ruleCacheKey = "cache:Rule:" // 基于rule id 创建key

type (
	defaultVendorModel struct {
		mc *monc.Model
	}

	Rule struct {
		Id       primitive.ObjectID `bson:"_id"`
		Name     string             `bson:"name"`    // 规则名称
		Content  map[string]string  `bson:"content"` // 规则内容
		Project  string             `bson:"project"` // 所属项目名称，一个项目只有1个规则
		CreateAt int64              `json:"createAt"`
		UpdateAt int64              `json:"updateAt"`
	}
)

func NewRuleModel(url string, db string, c cachec.CacheConf) RuleModel {
	return &defaultVendorModel{
		mc: monc.MustNewModel(url, db, "rule", c),
	}
}

type RuleModel interface {
	Insert(ctx context.Context, rule *Rule) error
	Update(ctx context.Context, rule *Rule) error
	FindOne(ctx context.Context, query bson.M) (*Rule, error)
}

func (m *defaultVendorModel) Insert(ctx context.Context, rule *Rule) error {
	now := time.Now().Unix()
	rule.CreateAt = now
	rule.UpdateAt = now
	_, err := m.mc.InsertOneNoCache(ctx, rule)
	return err
}

func (m *defaultVendorModel) Update(ctx context.Context, rule *Rule) error {
	rule.UpdateAt = time.Now().Unix()
	_, err := m.mc.ReplaceOne(ctx, getRuleKey(rule.Id), bson.M{"_id": rule.Id}, rule)
	return err
}

func (m *defaultVendorModel) FindOne(ctx context.Context, query bson.M) (*Rule, error) {
	var rules []*Rule
	err := m.mc.Find(ctx, &rules, query)
	if err != nil {
		return nil, err
	}

	if len(rules) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return rules[0], nil
}

func getRuleKey(id primitive.ObjectID) string {
	return ruleCacheKey + id.String()
}

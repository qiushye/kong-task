package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	cachec "github.com/zeromicro/go-zero/core/stores/cache"
)

type (
	defaultRuleModel struct {
		mc *monc.Model
	}

	Rule struct {
		Id       primitive.ObjectID `bson:"_id"`
		Name     string             `bson:"name"`    // 规则名称
		Content  string             `bson:"content"` // 规则内容
		Project  string             `bson:"project"` // 所属项目名称，一个项目只有1个规则
		CreateAt int64              `bson:"createAt"`
		UpdateAt int64              `bson:"updateAt"`
	}
)

func (r Rule) RuleId() string {
	return r.Id.Hex()
}

func NewRuleModel(url string, db string, c cachec.CacheConf) RuleModel {
	return &defaultRuleModel{
		mc: monc.MustNewModel(url, db, "rule", c),
	}
}

type RuleModel interface {
	Insert(ctx context.Context, rule *Rule) error
	Update(ctx context.Context, rule *Rule) error
	FindOne(ctx context.Context, name string) (*Rule, error)
	FindByProject(ctx context.Context, project string) (*Rule, error)
}

func (m *defaultRuleModel) Insert(ctx context.Context, rule *Rule) error {
	if rule.Id.IsZero() {
		rule.Id = primitive.NewObjectID()
	}
	now := time.Now().Unix()
	rule.CreateAt = now
	rule.UpdateAt = now
	_, err := m.mc.InsertOneNoCache(ctx, rule)
	return err
}

func (m *defaultRuleModel) Update(ctx context.Context, rule *Rule) error {
	rule.UpdateAt = time.Now().Unix()
	_, err := m.mc.ReplaceOneNoCache(ctx, bson.M{"_id": rule.Id}, rule)
	return err
}

func (m *defaultRuleModel) FindOne(ctx context.Context, name string) (*Rule, error) {
	var rule Rule
	err := m.mc.FindOneNoCache(ctx, &rule, bson.M{"name": name})
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

func (m *defaultRuleModel) FindByProject(ctx context.Context, project string) (*Rule, error) {
	var rule Rule
	err := m.mc.FindOneNoCache(ctx, &rule, bson.M{"project": project})
	if err != nil {
		return nil, err
	}

	return &rule, nil
}

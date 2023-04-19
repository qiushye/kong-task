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
	defaultUserModel struct {
		mc *monc.Model
	}

	User struct {
		Id           primitive.ObjectID `bson:"_id"`
		Uid          uint32             `bson:"uid"`
		Name         string             `bson:"name"`         // 用户名
		Project      string             `bson:"project"`      // 所属项目名称
		Organization string             `bson:"organization"` // 所属组织
		IsAdmin      bool               `bson:"isAdmin"`      // 是否为管理员
		ActionType   uint               `bson:"actionType"`   // 权限
		CreateAt     int64              `bson:"createAt"`
		UpdateAt     int64              `bson:"updateAt"`
	}
)

func NewUserModel(url string, db string, c cachec.CacheConf) UserModel {
	return &defaultUserModel{
		mc: monc.MustNewModel(url, db, "user", c),
	}
}

type UserModel interface {
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindOne(ctx context.Context, uid uint32, project string) (*User, error)
	FindByProject(ctx context.Context, project string) ([]*User, error)
	FindByUidAndOrganization(ctx context.Context, uid uint32, organization string) (*User, error)
	Delete(ctx context.Context, uid uint32, project string) error
}

func (m *defaultUserModel) Insert(ctx context.Context, user *User) error {
	user.Id = primitive.NewObjectID()
	now := time.Now().Unix()
	user.CreateAt = now
	user.UpdateAt = now
	_, err := m.mc.InsertOneNoCache(ctx, user)
	return err
}

func (m *defaultUserModel) Update(ctx context.Context, user *User) error {
	user.UpdateAt = time.Now().Unix()
	_, err := m.mc.ReplaceOneNoCache(ctx, bson.M{"uid": user.Uid, "project": user.Project}, user)
	return err
}

func (m *defaultUserModel) FindByUidAndOrganization(ctx context.Context, uid uint32, organization string) (*User, error) {
	var user User
	err := m.mc.FindOneNoCache(ctx, &user, bson.M{"uid": uid, "organization": organization})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *defaultUserModel) FindOne(ctx context.Context, uid uint32, project string) (*User, error) {
	var user User
	err := m.mc.FindOneNoCache(ctx, &user, bson.M{"uid": uid, "project": project})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *defaultUserModel) FindByProject(ctx context.Context, project string) ([]*User, error) {
	var users []*User
	err := m.mc.Find(ctx, &users, bson.M{"project": project})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *defaultUserModel) Delete(ctx context.Context, uid uint32, project string) error {
	_, err := m.mc.DeleteOneNoCache(ctx, bson.M{"uid": uid, "project": project})
	return err
}

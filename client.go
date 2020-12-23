/*
 *
 * client.go
 * pie
 *
 * Created by lintao on 2020/8/9 11:02 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package pie

import (
	"context"

	"github.com/NSObjects/pie/schemas"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	CollectionNameForStruct(doc interface{}) (*schemas.Collection, error)
	CollectionNameForSlice(doc interface{}) (*schemas.Collection, error)
	Connect(ctx context.Context) (err error)
	Disconnect(ctx context.Context) error
	BulkWrite(ctx context.Context, docs interface{}) (*mongo.BulkWriteResult, error)
	Distinct(ctx context.Context, doc interface{}, columns string) ([]interface{}, error)
	ReplaceOne(ctx context.Context, doc interface{}) (*mongo.UpdateResult, error)
	FindOneAndReplace(ctx context.Context, doc interface{}) error
	FindOneAndUpdate(ctx context.Context, doc interface{}) (*mongo.SingleResult, error)
	FindAndDelete(ctx context.Context, doc interface{}) error
	FindOne(ctx context.Context, doc interface{}) error
	FindAll(ctx context.Context, docs interface{}) error
	RegexFilter(key, pattern string) Session
	Asc(colNames ...string) Session
	Eq(key string, value interface{}) Session
	Ne(key string, ne interface{}) Session
	Nin(key string, nin interface{}) Session
	Nor(c Condition) Session
	Exists(key string, exists bool, filter ...Condition) Session
	Type(key string, t interface{}) Session
	Expr(filter Condition) Session
	Regex(key string, value interface{}) Session
	DataBase() *mongo.Database
	Collection(name string) *mongo.Collection
	Ping() error
	Filter(key string, value interface{}) Session
	ID(id interface{}) Session
	Gt(key string, value interface{}) Session
	Gte(key string, value interface{}) Session
	Lt(key string, value interface{}) Session
	Lte(key string, value interface{}) Session
	In(key string, value interface{}) Session
	And(filter Condition) Session
	Not(key string, value interface{}) Session
	Or(filter Condition) Session
	InsertOne(ctx context.Context, v interface{}) (primitive.ObjectID, error)
	InsertMany(ctx context.Context, v interface{}) (*mongo.InsertManyResult, error)
	Limit(limit int64) Session
	Skip(skip int64) Session
	Count(i interface{}) (int64, error)
	Desc(s1 ...string) Session
	Update(ctx context.Context, bean interface{}) (*mongo.UpdateResult, error)
	//The following operation updates all of the documents with quantity value less than 50.
	UpdateMany(ctx context.Context, bean interface{}) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	SoftDeleteOne(ctx context.Context, filter interface{}) error
	SoftDeleteMany(ctx context.Context, filter interface{}) error
	FilterBy(object interface{}) Session
	DropAll(ctx context.Context, doc interface{}) error
	DropOne(ctx context.Context, doc interface{}, name string) error
	AddIndex(keys interface{}, opt ...*options.IndexOptions) Indexes
	NewIndexes() Indexes
	NewSession() Session
	Aggregate() Aggregate
	SetDatabase(string string) Client
}

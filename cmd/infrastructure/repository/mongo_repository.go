package repository

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoRepository(cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{cfg: cfg, db: db}
}

func (p *mongoRepository) CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.CreateProduct")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	_, err := collection.InsertOne(ctx, product, &options.InsertOneOptions{})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "InsertOne")
	}

	return product, nil
}

func (p *mongoRepository) GetProductById(ctx context.Context, uuid uuid.UUID) (*product.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetProductById")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	var product product.Product
	if err := collection.FindOne(ctx, bson.M{"_id": uuid.String()}).Decode(&product); err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Decode")
	}

	return &product, nil
}

func (p *mongoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}

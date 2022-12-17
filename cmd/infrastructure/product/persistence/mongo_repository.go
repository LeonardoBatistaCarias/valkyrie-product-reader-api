package persistence

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type mongoRepository struct {
	cfg *config.Config
	db  *mongo.Client
}

func NewMongoRepository(cfg *config.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{cfg: cfg, db: db}
}

func (p *mongoRepository) CreateProduct(ctx context.Context, product *product.Product) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.CreateProduct")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	if _, err := collection.InsertOne(ctx, product, &options.InsertOneOptions{}); err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "InsertOne")
	}

	return nil
}

func (p *mongoRepository) GetProductById(ctx context.Context, productID string) (*product.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetProductById")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	var product product.Product
	if err := collection.FindOne(ctx, bson.M{"productid": productID}).Decode(&product); err != nil {
		p.traceErr(span, err)
		log.Printf("Product with ID: %s doesn't exist")
		return nil, errors.Wrap(err, "Decode")
	}

	return &product, nil
}

func (p *mongoRepository) DeleteProductByID(ctx context.Context, productID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.DeleteProductByID")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	filter := bson.M{"productid": productID}
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "UpdateOne")
	}

	return nil
}

func (p *mongoRepository) DeactivateProductByID(ctx context.Context, productID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.DeactivateProductByID")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	filter := bson.M{"productid": productID}
	update := bson.D{{"$set", bson.D{{"active", false}}}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "UpdateOne")
	}

	return nil
}

func (p *mongoRepository) UpdateProductByID(ctx context.Context, product *product.Product) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateProductByID")
	defer span.Finish()

	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	if result := collection.FindOneAndUpdate(ctx, bson.M{"productid": product.ProductID}, bson.M{"$set": product}); result.Err() != nil {
		p.traceErr(span, result.Err())
		return errors.Wrap(result.Err(), "FindOneAndUpdate")
	}

	return nil
}

func (p *mongoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}

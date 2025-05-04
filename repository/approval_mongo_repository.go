package repository

import (
	"acl-casbin/constant"
	"acl-casbin/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

package repository

import (
"context"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"time"
"your_project/model"
)

type ApprovalMongoRepository struct {
	collection *mongo.Collection
}

func NewApprovalMongoRepository(collection *mongo.Collection) ApprovalRepository {
	return &ApprovalMongoRepository{collection: collection}
}

func (r *ApprovalMongoRepository) ApplyDecisionWithLogic(ctx context.Context, objectID, objectType string, stepID int, decision model.ApprovalDecision) error {
	filter := bson.M{
		"object_id":   objectID,
		"object_type": objectType,
		"steps.step_id": stepID,
	}

	update := bson.M{
		"$push": bson.M{
			"steps.$.decisions": decision,
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	res := r.collection.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return res.Err()
	}

	// تغییر وضعیت مرحله بر اساس تصمیم
	status := constant.ApprovalStatusPending
	if decision.Action == constant.ApprovalActionApprove {
		status = constant.ApprovalActionApprove
	} else if decision.Action == constant.ApprovalActionReject {
		status = constant.ApprovalActionReject
	}

	// بروزرسانی وضعیت مرحله در فلو
	_, err := r.collection.UpdateOne(ctx, bson.M{
		"object_id":   objectID,
		"object_type": objectType,
		"steps.step_id": stepID,
	}, bson.M{
		"$set": bson.M{
			"steps.$.status": status,
		},
	})
	return err
}
func (r *ApprovalMongoRepository) ListFlows(ctx context.Context, objectType string, status *int) ([]model.ApprovalFlow, error) {
	filter := bson.M{
		"object_type": objectType,
	}

	if status != nil {
		filter["status"] = *status
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var flows []model.ApprovalFlow
	for cursor.Next(ctx) {
		var flow model.ApprovalFlow
		if err := cursor.Decode(&flow); err != nil {
			return nil, err
		}
		flows = append(flows, flow)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return flows, nil
}

func (r *ApprovalMongoRepository) EnsureIndexes(ctx context.Context) error {
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "object_id", Value: 1},
				{Key: "object_type", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{Key: "status", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "steps.step_id", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "created_at", Value: -1},
			},
		},
		{
			Keys: bson.D{
				{Key: "updated_at", Value: -1},
			},
		},
	}
	_, err := r.collection.Indexes().CreateMany(ctx, indexModels)
	return err
}

func (r *ApprovalMongoRepository) UpdateFlow(ctx context.Context, flow *model.ApprovalFlow) error {
	filter := bson.M{
		"_id": flow.ID,
	}

	update := bson.M{
		"$set": bson.M{
			"steps":     flow.Steps,
			"status":    flow.Status,
			"updated_at": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *ApprovalMongoRepository) GetFlowByObject(ctx context.Context, objectID, objectType string) (*model.ApprovalFlow, error) {
	filter := bson.M{
		"object_id":   objectID,
		"object_type": objectType,
	}

	var flow model.ApprovalFlow
	err := r.collection.FindOne(ctx, filter).Decode(&flow)
	if err != nil {
		return nil, err
	}

	return &flow, nil
}

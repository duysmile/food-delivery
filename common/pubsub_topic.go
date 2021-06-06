package common

import "200lab/food-delivery/pubsub"

const (
	TopicUserLikeRestaurant   pubsub.Topic = "TopicUserLikeRestaurant"
	TopicUserUnLikeRestaurant pubsub.Topic = "TopicUserUnLikeRestaurant"
	TopicUserLikeFood         pubsub.Topic = "TopicUserLikeFood"
	TopicUserUnLikeFood       pubsub.Topic = "TopicUserUnLikeFood"
	TopicCreateOrder          pubsub.Topic = "TopicCreateOrder"
)

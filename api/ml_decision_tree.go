package main

import (
	// "fmt"
	"log"
	"strconv"
	"net/http"
	"math/rand"
	"github.com/gin-gonic/gin"
)

// TODO Decision Tree Algorithmを実装する(できればRandomForestかBoostingを実装する)

type CaliforniaHousingData struct {
	MedInc float64
	HouseAge float64
	AveRooms float64
	AveBedrms float64
	Population float64
	AveOccup float64
	Latitude float64
	Longitude float64
}

// type options func(*DecisionOptions)

// func DecisionOptions(v string) options {
// 	return func(o *DecisionOptions)
// }

func decisionTree(data []*CaliforniaHousingData, features_num int) {

	chosenFeature := rand.Int31n(features_num)
}
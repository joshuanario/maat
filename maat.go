package maat

import (
	"math"
	"math/rand"
	"time"
)

type Body struct {
	Weigth  int
	Pointer *interface{}
}

type Config struct {
	Alpha        int
	LambdaLevels int8
	Omega        int
}

func Judge(body *Body, config *Config) bool {
	if config.LambdaLevels <= 0 {
		panic("invalid lambda levels")
	}
	if body.Weigth >= config.Omega {
		return false
	}
	if body.Weigth <= config.Alpha {
		return true
	}
	return evaluateLambda(body, config)
}

func calculateLambda(body *Body, config *Config) int32 {
	lambda := int32(1)
	netWeight := body.Weigth - config.Alpha
	inc := int(math.Ceil(float64((config.Omega - config.Alpha)) / float64(config.LambdaLevels)))
	for int(lambda)*inc < netWeight {
		lambda++
	}
	return lambda
}

func evaluateLambda(body *Body, config *Config) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	lambda := calculateLambda(body, config)
	verdict := r.Int31n(int32(lambda))
	return verdict == 0
}

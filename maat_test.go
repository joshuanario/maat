package maat

import "testing"

var myheartwillgoonandon interface{} = "my heart will go on and on"
var itstearingupmyheart interface{} = "it's tearing up my heart"
var myheartisabattleground interface{} = "my heart is a battleground"

func TestJudge1(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  1400,
		Pointer: &myheartwillgoonandon,
	}
	output := Judge(body, config)
	if output != true {
		t.Fatal("expected true; outcome does not match")
	}
}

func TestJudge2(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  7007,
		Pointer: &itstearingupmyheart,
	}
	output := Judge(body, config)
	if output != false {
		t.Fatal("expected false; outcome does not match")
	}
}

func TestJudge3(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  1500,
		Pointer: &myheartwillgoonandon,
	}
	output := Judge(body, config)
	if output != true {
		t.Fatal("expected true; outcome does not match")
	}
}

func TestJudge4(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  3000,
		Pointer: &itstearingupmyheart,
	}
	output := Judge(body, config)
	if output != false {
		t.Fatal("expected false; outcome does not match")
	}
}

func TestCalculateLambda1(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  1501,
		Pointer: &myheartisabattleground,
	}
	output := calculateLambda(body, config)
	if output != 1 {
		t.Fatalf("expected lambda == 1; outcome %d;", output)
	}
}

func TestCalculateLambda2(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  2999,
		Pointer: &myheartisabattleground,
	}
	output := calculateLambda(body, config)
	if output != 3 {
		t.Fatalf("expected lambda == 3; outcome %d;", output)
	}
}

func TestCalculateLambda3(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  2500,
		Pointer: &myheartisabattleground,
	}
	output := calculateLambda(body, config)
	if output != 2 {
		t.Fatalf("expected lambda == 2; outcome %d;", output)
	}
}

func TestCalculateLambda4(t *testing.T) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	body := &Body{
		Weigth:  2001,
		Pointer: &myheartisabattleground,
	}
	output := calculateLambda(body, config)
	if output != 2 {
		t.Fatalf("expected lambda == 2; outcome %d;", output)
	}
}

func Benchmark(b *testing.B) {
	config := &Config{
		Alpha:        1500,
		LambdaLevels: 3,
		Omega:        3000,
	}
	for n := 0; n < b.N; n++ {
		Judge(&Body{
			Weigth:  2001,
			Pointer: &myheartisabattleground,
		}, config)
		Judge(&Body{
			Weigth:  3000,
			Pointer: &itstearingupmyheart,
		}, config)
		Judge(&Body{
			Weigth:  1400,
			Pointer: &myheartwillgoonandon,
		}, config)
	}
}

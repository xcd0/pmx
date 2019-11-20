package main

import (
	"fmt"
	//"math"
	"github.com/golang-collections/collections/stack"
	"math/rand"
)

func main() {
	rand.Seed(0)
	p := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 2, 1, 8, 7, 6, 9, 3},
	}
	for i, v := range p {
		fmt.Printf("p[%v] : %v\n", i, p[i])
	}
	p[0], p[1] = pmx(p[0], p[1])
	for i, v := range p {
		fmt.Printf("p[%v] : %v\n", i, p[i])
	}
}

func pmx(p1, p2 []int) ([]int, []int) {
	// p1とp2は同じ長さ
	l := len(p1)
	// 子供を2人生成 とりあえずコピー
	c1 := p1
	c2 := p2
	// 部分の区切りを生成
	r1 := rand.Intn(l - 2) // 0 ~ l - 2 までの値で乱数を得る
	// l-1までにするとr2 - r1 が一致する 一致したら入れ替えが起きないので無駄になる
	r2 := r1 + 1 + rand.Intn(l-r1) // r1 <= l - 1 までの値で乱数を得る

	fmt.Printf("r1 : %v, r2 : %v\n", r1, r2)

	// 部分入れ替え
	for i := r1; i < r2; i++ {
		c1[i], c2[i] = p2[i], p1[i]
	}
	c1 = crossOver(c1, p1, r1, r2) // c1を生成
	c2 = crossOver(c2, p2, r1, r2) // c2を生成
	return c1, c2
}

func crossOver(child, p1 []int, r1, r2 int) []int {
	// childを生成する。
	// childにはすでにp1がコピーされており、p1[r1:r2]はp2の[r1:r2]が入っている
	// child[r1:r2]にない数をchild[0~r1 or r2 ~ last]に詰めるとchildが完成する

	// child[r1:r2]にはすでにp2が入ってる つまりchild[r1:r2] == p2[r1:r2]
	// child[0~r1 or r2 ~ last] には child[r1:r2]にない数をp1から継承して入れる

	// 呼び出し時点ではchildには数がかぶって入っており、
	// p1[r1:r2]のなかの数字うち、いくつかが入っていない
	// かぶっている数と入っていない数の個数は一致する
	// かぶっている数のある場所でchlid[r1:r2]以外の場所に
	// child[r1:r2]にないものを p1[r1:r2]の中から詰める

	// p1[r1:r2]の中でchild[r1:r2]にない数のリスト
	// child[r1:r2]にない数をchild[0~r1 or r2 ~ last]に詰める際に使う
	s := stack.New()
	pp1 := p1[r1:r2]
	pchild := child[r1:r2]
	getUnique(s, pp1, pchild)

	for i := 0; i < len(p1); i++ {
		if i < r1 || r2 <= i { // 0 <= i < r1 || r2 <= i < len(p1)
			// listは部分
			for _, t := range child[r1:r2] {
				if t == p1[i] {
					// r1:r2の範囲でかぶらなかった個数と
					// npに入れられなかった個数は一致するはず
					child[i] = s.Pop()
				}
				// かぶらなかったらすでに入っているので
				// child[i]にp1[i]を上書きする必要はない
			}
		}
	}
	return child
}

func getUnique(s *stack.Stack, pp1, pp2 *[]int) { // {{{
	// pp1, pp2 はそれぞれp1[r1:r2], p2[r1:r2]である
	e := make([]int, 0)
	for _, t := range pp1 {
		// pp1全部舐める
		// tがpp2にあるか
		if func(int) bool {
			for _, c := range pp2 {
				// pp2全部舐める
				if c == t {
					// あった
					return false
				}
			}
			// なかった
			return true
		}(t) {
			// 一致するものがない
			// tを詰める
			s.Push(t)
		}
	}
} // }}}

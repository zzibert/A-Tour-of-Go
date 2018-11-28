package main

import "fmt"

func numJewelsInStones(J string, S string) int {
    num := 0
    
    for _, s := range S {
        for _, j := range J {
            if s == j {
                num++
                continue
            }
        }
    }
    
    return 0
}
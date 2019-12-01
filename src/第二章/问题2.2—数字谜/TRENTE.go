package main

import 	"fmt"

func main() {
	var (
		I int = 0
		N int
		R int
		IN int
		VINGT int
		CINQ int
		TRENTE int
		sum int
		E int = 1
		V int
		G int
		Q int
		T int
		C int
	)

	for V = 8; V <= 9; V++ {
		for I = 0; I <= 9; I++ {
			if I != V && I != T {
				for N = 0; N <= 9; N++ {
					if N != I && N != V && N != T {
						IN = I * 10 + N
						for G = 0; G <= 9; G++ {
							if G != N && G != V && G != T && G != I && G != R {
								for C = 2; C <= 9; C++ {
									if C != G && C != N && C != V && C != T && C != I && C != R {
										for Q = 2; Q <= 9; Q++ {
											if Q != C && Q != G && Q != N && Q != V && Q != T && Q != I && Q != R {
												for E = 3; E <= 9; E += 2 {
													if E != C && E != G && E != N && E != V && E != T && E != I && E != R && E != Q {
														TRENTE = ((((T*10+R)*10+E)*10+N)*10+T)*10+E
														VINGT = ((V*100+IN)*10+G)*10+T
														CINQ = (C*100+IN)*10+Q
														sum = VINGT + CINQ + CINQ
														if sum == TRENTE {
															fmt.Printf("VINGT: %d\n", VINGT)
															fmt.Printf("CINQ: %d\n", CINQ)
															fmt.Printf("TRENTE: %d\n", TRENTE)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
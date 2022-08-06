package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	x := make([]string, M)
	for i := 0; i < M; i++ {
		x[i] = strconv.Itoa(i + 1)
	}

	for a := 0; a < M; a++ {
		if a > M-N+1 {
			break
		}

		if N == 1 {
			fmt.Println(x[a])
		}

		for b := a + 1; b < M; b++ {
			if b > M-N+2 {
				break
			}

			if N == 2 {
				fmt.Println(x[a], x[b])
			}

			for c := b + 1; c < M; c++ {
				if c > M-N+3 {
					break
				}

				if N == 3 {
					fmt.Println(x[a], x[b], x[c])
				}

				for d := c + 1; d < M; d++ {
					if d > M-N+4 {
						break
					}

					if N == 4 {
						fmt.Println(x[a], x[b], x[c], x[d])
					}

					for e := d + 1; e < M; e++ {
						if e > M-N+5 {
							break
						}

						if N == 5 {
							fmt.Println(x[a], x[b], x[c], x[d], x[e])
						}

						for f := e + 1; f < M; f++ {
							if f > M-N+6 {
								break
							}

							if N == 6 {
								fmt.Println(x[a], x[b], x[c], x[d], x[e], x[f])
							}

							for g := f + 1; g < M; g++ {
								if g > M-N+7 {
									break
								}

								if N == 7 {
									fmt.Println(x[a], x[b], x[c], x[d], x[e], x[f], x[g])
								}

								for h := g + 1; h < M; h++ {
									if h > M-N+8 {
										break
									}

									if N == 8 {
										fmt.Println(x[a], x[b], x[c], x[d], x[e], x[f], x[g], x[h])
									}

									for i := h + 1; i < M; i++ {
										if i > M-N+9 {
											break
										}

										if N == 9 {
											fmt.Println(x[a], x[b], x[c], x[d], x[e], x[f], x[g], x[h], x[i])
										}

										for j := i + 1; j < M; j++ {
											if j > M-N+10 {
												break
											}

											if N == 10 {
												fmt.Println(x[a], x[b], x[c], x[d], x[e], x[f], x[g], x[h], x[i], x[j])
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

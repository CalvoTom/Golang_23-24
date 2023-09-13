package piscine

import "github.com/01-edu/z01"

func PrintCombn(n int) {
	switch {
	case n == 1:
		for i := 48; i < 58; i++ {
			z01.PrintRune(rune(i))
			if i == 57 {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	case n == 2:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				if i < j {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					if i == 56 && j == 57 {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	case n == 3:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					if i < j && j < k {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						if i == 55 && j == 56 {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	case n == 4:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					for l := 48; l < 58; l++ {
						if i < j && j < k && k < l {
							z01.PrintRune(rune(i))
							z01.PrintRune(rune(j))
							z01.PrintRune(rune(k))
							z01.PrintRune(rune(l))
							if i == 54 && j == 55 && k == 56 {
								break
							}
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	case n == 5:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					for l := 48; l < 58; l++ {
						for m := 48; m < 58; m++ {
							if i < j && j < k && k < l && l < m {
								z01.PrintRune(rune(i))
								z01.PrintRune(rune(j))
								z01.PrintRune(rune(k))
								z01.PrintRune(rune(l))
								z01.PrintRune(rune(m))
								if i == 53 && j == 54 && k == 55 && l == 56 {
									break
								}
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	case n == 6:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					for l := 48; l < 58; l++ {
						for m := 48; m < 58; m++ {
							for n := 48; n < 58; n++ {
								if i < j && j < k && k < l && l < m && m < n {
									z01.PrintRune(rune(i))
									z01.PrintRune(rune(j))
									z01.PrintRune(rune(k))
									z01.PrintRune(rune(l))
									z01.PrintRune(rune(m))
									z01.PrintRune(rune(n))
									if i == 52 && j == 53 && k == 54 && l == 55 && m == 56 {
										break
									}
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	case n == 7:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					for l := 48; l < 58; l++ {
						for m := 48; m < 58; m++ {
							for n := 48; n < 58; n++ {
								for o := 48; o < 58; o++ {
									if i < j && j < k && k < l && l < m && m < n && n < o {
										z01.PrintRune(rune(i))
										z01.PrintRune(rune(j))
										z01.PrintRune(rune(k))
										z01.PrintRune(rune(l))
										z01.PrintRune(rune(m))
										z01.PrintRune(rune(n))
										z01.PrintRune(rune(o))
										if i == 51 && j == 52 && k == 53 && l == 54 && m == 55 && n == 56 {
											break
										}
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	case n == 8:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					for l := 48; l < 58; l++ {
						for m := 48; m < 58; m++ {
							for n := 48; n < 58; n++ {
								for o := 48; o < 58; o++ {
									for p := 48; p < 58; p++ {
										if i < j && j < k && k < l && l < m && m < n && n < o && o < p {
											z01.PrintRune(rune(i))
											z01.PrintRune(rune(j))
											z01.PrintRune(rune(k))
											z01.PrintRune(rune(l))
											z01.PrintRune(rune(m))
											z01.PrintRune(rune(n))
											z01.PrintRune(rune(o))
											z01.PrintRune(rune(p))
											if i == 50 && j == 51 && k == 52 && l == 53 && m == 54 && n == 55 && o == 56 && p == 57 {
												break
											}
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case n == 9:
		for i := 48; i < 58; i++ {
			for j := 48; j < 58; j++ {
				for k := 48; k < 58; k++ {
					for l := 48; l < 58; l++ {
						for m := 48; m < 58; m++ {
							for n := 48; n < 58; n++ {
								for o := 48; o < 58; o++ {
									for p := 48; p < 58; p++ {
										for q := 48; q < 58; q++ {
											if i < j && j < k && k < l && l < m && m < n && n < o && o < p && p < q {
												z01.PrintRune(rune(i))
												z01.PrintRune(rune(j))
												z01.PrintRune(rune(k))
												z01.PrintRune(rune(l))
												z01.PrintRune(rune(m))
												z01.PrintRune(rune(n))
												z01.PrintRune(rune(o))
												z01.PrintRune(rune(p))
												z01.PrintRune(rune(q))
												if i == 49 && j == 50 && k == 51 && l == 52 && m == 53 && n == 54 && o == 55 && p == 56 {
													break
												}
												z01.PrintRune(',')
												z01.PrintRune(' ')
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
	z01.PrintRune('\n')
}

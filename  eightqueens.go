package piscine

import "text/tabwriter"

func EightQueens() {
	var tab [8]int
	for i:= 0; i<9;i++{
		for j:= 0; j<9;j++{
			if tab[i][j] != 1{
				tab[i][j] = 2
				
			}
	}
}
[[1;1;1;Q;1;1;1;1];
 [0;0;1;1;1;0;0;0];
 [0;1;0;1;0;1;0;0];
 [1;0;0;1;0;0;1;0];
 [0;0;0;1;0;0;0;1]; 
 [0;0;0;1;0;0;0;0];
 [0;0;0;1;0;0;0;0];
 [0;0;0;1;0;0;0;0]]
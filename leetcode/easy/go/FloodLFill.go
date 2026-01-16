package go_test

func floodFill(image [][]int, sr int, sc int, color int) [][]int {


	orig := image[sr][sc]

	if orig == color { 
		return image
	}

	m , n := len(image), len(image[0])

	var dfs func(r , c int)

	dfs = func(r, c int) {
		//bounds 
		if r <0 || r >= m || c < 0 || c >=n { 
			return 
		}
		
		//only doing maching colors
		if image[r][c] != orig { 
			return 
		}

		image[r][c] = color

		//call dfs on neibous 
		dfs(r+1, c) //up
		dfs(r-1, c) //down
		dfs(r, c+1)	//right
		dfs(r, c-1)	//left 
	}

	dfs(sr, sc)
	return image
}

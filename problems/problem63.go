// Unique Paths II
//
// You are given an m x n integer array grid. There is a robot initially located at the top-left corner (i.e., grid[0][0]).
// The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
//
// An obstacle and space are marked as 1 or 0 respectively in grid. A path that the robot takes cannot include any square that is an obstacle.
//
// Return the number of possible unique paths that the robot can take to reach the bottom-right corner.
//
// The testcases are generated so that the answer will be less than or equal to 2 * 109.

package problems

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

    // helper function
    var allowToVisit func (x, y int) int

    allowToVisit = func(x, y int) int{
        // 1 : allow to visit
        // 0 : can not visit due to obstacle
        return 1 - obstacleGrid[y][x]
    }

    h, w := len(obstacleGrid), len(obstacleGrid[0])

    if h * w == 0 || allowToVisit(0, 0) == 0 {
        // Quick response for invalid cases
        return 0
    }

    // update [0][0] as start point with one valid path
    obstacleGrid[0][0] = 1

    // base case: leftmost column
    for y := 1 ; y < h ; y++{
        obstacleGrid[y][0] = obstacleGrid[y-1][0] * allowToVisit(0, y)
    }

    // base case: top row
    for x := 1 ; x < w ; x++{
        obstacleGrid[0][x] = obstacleGrid[0][x-1] * allowToVisit(x, 0)
    }

    // general cases:
    for y := 1 ; y < h ; y++{
        for x := 1 ; x < w ; x ++{

            // update path count form left and top
            obstacleGrid[y][x] = ( obstacleGrid[y][x-1] + obstacleGrid[y-1][x] ) * allowToVisit(x, y)
        }
    }

    return obstacleGrid[h-1][w-1]
}

func Problem63() {
    grid := [][]int{
        {0, 0, 0},
        {0, 1, 0},
        {0, 0, 0},
    }

    fmt.Println("Number of possible unique paths:", uniquePathsWithObstacles(grid))
}

func init() {
    RegisterProblem(63, Problem63)
}

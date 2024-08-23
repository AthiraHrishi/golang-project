/*You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you
pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.
Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.*/

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Function to calculate the minimum cost to reach the top
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp0, dp1 := 0, 0

	for i := 2; i <= n; i++ {
		currentDP := min(dp1+cost[i-1], dp0+cost[i-2])
		dp0, dp1 = dp1, currentDP
	}

	return dp1
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Retrieve the cost array from PostgreSQL
func getCostArray(db *sql.DB) ([]int, error) {
	rows, err := db.Query("SELECT cost_value FROM cost_table ORDER BY step_index")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cost []int
	for rows.Next() {
		var value int
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		cost = append(cost, value)
	}
	return cost, nil
}

func main() {
	// PostgreSQL connection details
	connStr := "user=yourusername dbname=yourdbname sslmode=disable password=yourpassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Retrieve the cost array from the database
	cost, err := getCostArray(db)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the minimum cost
	result := minCostClimbingStairs(cost)
	fmt.Printf("Minimum cost to reach the top is: %d\n", result)
}

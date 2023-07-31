package strangePrinter

func strangePrinter(s string) int {
    n := len(s)

    dp := make([][]int, n)
    for i:=0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = n+1
        }
    }

    for i:=0; i<n;i++ {
        for j:=0; j < n-i; j++ {
            r := j+i
            k := j
            q := -1
            for ; k < r; k++ {
                if s[k] != s[r] && q == -1{
                    q = k
                } 
                if q > -1 { 
                    dp[j][r] = min(dp[j][r], 1 + dp[q][k] + dp[k+1][r])
                }
            }
            if q == -1 {
                dp[j][r] = 0
            }
            
        }
    }

    return dp[0][n-1]+1
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func canPartition(nums []int) bool {
    sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	if sum%2 == 1{
		return false
	}
	dp := make([][]bool,sum/2+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]bool,len(nums)+1)
	}
	for i:=0;i<=len(nums);i++{
		dp[0][i] = true
	}
	for i := 1;i<=sum/2;i++{
		for j:=1;j<=len(nums);j++{
			dp[i][j] = dp[i][j-1]
			if i>=nums[j-1]{
				dp[i][j] = dp[i][j] || dp[i-nums[j-1]][j-1]
			}
		}
	}
	return dp[sum/2][len(nums)]
}
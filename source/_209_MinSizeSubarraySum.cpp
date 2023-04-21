class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int n = nums.size() - 1;
        if (n < 0)
        {
            return 0;
        }
        int ans = INT_MAX;
        int start = 0, end = 0;
        int sum = 0;
        while (end <= n)
        {
            sum += nums[end];
            while (sum >= target)
            {
                int ans = min(ans, end - start + 1);
                sum -= nums[start];
                start++;
            }
            end++;   
        }
        return ans == INT_MAX ? 0 : ans;
    }
};
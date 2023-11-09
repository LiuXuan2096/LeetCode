class Solution
{
public:
    int search(vector<int>& nums, int target)
    {
        int left = 0; // 初始化区间左边界
        int right = nums.size() - 1; // 初始化区间右边界
        while (left <= right) // 区间为左闭右闭区间 所有是<=
        {
            /* code */
            int middle = (left + right) >> 1;
            if (nums[middle] < target)
            {
                /* code */
                left = middle + 1;
            }
            else if (nums[middle] > target)
            {
                right = middle - 1;
            }
            else 
            {
                return middle;
            }
            
        }
        return -1;

    }
};
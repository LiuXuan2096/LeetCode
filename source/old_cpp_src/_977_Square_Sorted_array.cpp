class Solution
{
public:
	vector<int> sortedSquares(vector<int>& nums)
	{
		int n = nums.size();
		vector<int> result(n);
	
		for (int i = 0,j = n - 1,pos = n - 1; i <= j;)
		{
			int square_left = nums[i] * nums[i];
			int square_right = nums[j] * nums[j];
			if (square_left >= square_right)
			{
				result[pos] = square_left;
				i++;
			}
			else
			{
				result[pos] = square_right;
				j--;
			}
			pos--;
		}
		return result;
	}
};
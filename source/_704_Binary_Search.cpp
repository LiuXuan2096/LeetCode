class Solution
{
public:
    int search(vector<int>& nums, int target)
    {
        int left = 0; // ��ʼ��������߽�
        int right = nums.size() - 1; // ��ʼ�������ұ߽�
        while (left <= right) // ����Ϊ����ұ����� ������<=
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
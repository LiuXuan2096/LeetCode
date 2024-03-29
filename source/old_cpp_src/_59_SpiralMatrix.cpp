class Solution
{
public:
	vector<vector<int>> generateMatrix(int n)
	{
		vector<vector<int>> res(n, vector<int>(n, 0)); //使用vector定义一个二维数组
		int startx = 0, starty = 0; // 定义每循环一个圈的起始位置
		int loop = n >> 1; // 确定每个圈循环几次
		int mid = n >> 1; // 确定矩阵的中间位置
		int count = 1; // 用来给矩阵每一个位置赋值
		int offset = 1; // 控制每一条边遍历的长度，每次循环右边界收缩一位
		int i, j;
		while (loop--)
		{
			i = startx;
			j = starty;

			// 以下四个for就是模拟转了一圈

			// 模拟填充上行从左到右(左闭右开）
			for (j = starty; j < n - offset; j++)
			{
				res[startx][j] = count++;
			}

			// 模拟填充右列从上到下（左闭右开）
			for (i = startx; i < n - offset; i++)
			{
				res[i][j] = count++;
			}

			// 模拟填充下行从右到左（左闭右开）
			for (;  j > starty; j--)
			{
				res[i][j] = count++;
			}

			// 模拟填充左列从下到上（左闭右开）
			for (; i > startx; i--)
			{
				res[i][j] = count++;
			}

			// 下一圈开始的时候，起始位置要各自+1
			startx++;
			starty++;

			// offset控制每一圈里每一条边遍历的长度
			offset += 1;
		}

		// 如果n为奇数的话，需要单独给矩阵最中间的位置赋值
		if (n % 2)
		{
			res[mid][mid] = count;
		}
		return res;
	}
};
class Solution
{
public:
	vector<vector<int>> generateMatrix(int n)
	{
		vector<vector<int>> res(n, vector<int>(n, 0)); //ʹ��vector����һ����ά����
		int startx = 0, starty = 0; // ����ÿѭ��һ��Ȧ����ʼλ��
		int loop = n >> 1; // ȷ��ÿ��Ȧѭ������
		int mid = n >> 1; // ȷ��������м�λ��
		int count = 1; // ����������ÿһ��λ�ø�ֵ
		int offset = 1; // ����ÿһ���߱����ĳ��ȣ�ÿ��ѭ���ұ߽�����һλ
		int i, j;
		while (loop--)
		{
			i = startx;
			j = starty;

			// �����ĸ�for����ģ��ת��һȦ

			// ģ��������д�����(����ҿ���
			for (j = starty; j < n - offset; j++)
			{
				res[startx][j] = count++;
			}

			// ģ��������д��ϵ��£�����ҿ���
			for (i = startx; i < n - offset; i++)
			{
				res[i][j] = count++;
			}

			// ģ��������д��ҵ�������ҿ���
			for (;  j > starty; j--)
			{
				res[i][j] = count++;
			}

			// ģ��������д��µ��ϣ�����ҿ���
			for (; i > startx; i--)
			{
				res[i][j] = count++;
			}

			// ��һȦ��ʼ��ʱ����ʼλ��Ҫ����+1
			startx++;
			starty++;

			// offset����ÿһȦ��ÿһ���߱����ĳ���
			offset += 1;
		}

		// ���nΪ�����Ļ�����Ҫ�������������м��λ�ø�ֵ
		if (n % 2)
		{
			res[mid][mid] = count;
		}
		return res;
	}
};
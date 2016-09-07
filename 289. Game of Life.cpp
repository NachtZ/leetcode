class Solution {
public:
	int getLive(vector<vector<int>> &board, int x, int y){
		int n = board.size();
		int m = board[0].size();
		int count = 0;
		for (int i = -1; i != 2; i++){
			for (int j = -1; j != 2; j++)
			{
				if (i == 0 && j == 0)continue;
				if (x + i >= 0 && x + i < n && y + j >= 0 && y + j < m && (board[x + i][y + j] == 1 || board[x + i][y + j] == 2))
					count++;
			}
		}
		return count;
	}
	//value 2 means change live to dead, value 3 means change dead to live
	void gameOfLife(vector<vector<int>>& board) {
		int n = board.size();
		if (n == 0)return;
		int m = board[0].size();
		int count = 0;
		for (int i = 0; i< n; i++){
			for (int j = 0; j<m; j++){
				count = getLive(board, i, j);
				if (board[i][j] == 1){
					if (count < 2 || count >3)board[i][j] = 2;
				}
				else{
					if (count == 3)board[i][j] = 3;
				}
			}

		}
		for (int i = 0; i<n; i++){
			for (int j = 0; j< m; j++){
				if (board[i][j] == 2)board[i][j] = 0;
				else if (board[i][j] == 3)board[i][j] = 1;
			}
		}
	}
};
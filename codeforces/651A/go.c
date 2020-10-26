#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int INFB = 2147483647;
int INFX = -2147483648;
void readIntArray(int* target, int N)
{
    for(int i = 0; i<N; i++)
    {
        scanf("%d", target+i);
    }
}
void printIntArray(int* target, int N)
{
    for(int i= 0;i<N;i++)
    {
        printf("%d ", *(target+i));
    }
}
int NUMS[3];
int DP[201][201];
int dp(int i, int j)
{
    if(i<j){
        return dp(j, i);
    }

    if(DP[i][j] != -1)
    {
        return DP[i][j];
    }
    if(j == 0)
    {
        DP[i][j] = 0;
        return 0;
    }
    
    DP[i][j] = dp(i-2, j+1) + 1;
    return DP[i][j];
    
}
int main()
{
    readIntArray(NUMS+1, 2);
    for(int i = 0; i != 201; i++)
    {
        for(int j = 0; j != 201; j++)
        {
            DP[i][j] = -1;
        }
    }
    DP[1][1] = 0;
    printf("%d\n", dp(NUMS[1], NUMS[2]));
}
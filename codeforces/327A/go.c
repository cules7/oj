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

int N;
int NUMS[101];
int wholeOne;
int DP[101];
int SUM;
int MIN;
int main()
{
    MIN = 1;
    SUM = 0;
    wholeOne = 0;
    readIntArray(&N, 1);
    for(int i = 1; i != N+1; i++)
    {
        readIntArray(NUMS+i, 1);
        if(NUMS[i]==1)
        {
            wholeOne++;
        }
    }

    for(int i = 1; i != N+1; i++)
    {
        if(NUMS[i] == 0)
        {
            if(DP[i-1] >= 0)
            {
                DP[i] = -1;
            }
            else
            {
                DP[i] = DP[i-1]-1;
            }
            if(DP[i]<=MIN)
            {
                MIN = DP[i];
            }
        }
        else
        {
            DP[i] = DP[i-1]+1;
        }
    }
    printf("%d\n", wholeOne-MIN);
}
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
int NUMS[100000+1];
int DP[100000+1];
int MAX;
int main()
{
    readIntArray(&N, 1);
    readIntArray(NUMS+1, N);
    if(N==1)
    {
        printf("%d\n", 1);
        return 0;
    }
    DP[1] = 1;
    MAX = 1;
    for(int i = 2; i != N+1; i++)
    {
        if(NUMS[i]>NUMS[i-1])
        {
            DP[i] = DP[i-1]+1;
        }
        else
        {
            DP[i] = 1;
        }
        if(DP[i]>MAX)
        {
            MAX = DP[i];
        }
    }
    printf("%d\n", MAX);
}
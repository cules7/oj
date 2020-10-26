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
long long dp(int index);
long long max(long long a, long long b);
int N;
long long NumsMap[100000+1];
long long DP[100000+1];
int ZAN;
int main()
{
    memset(NumsMap, 0, (100000+1)*sizeof(int));

    readIntArray(&N, 1);
    for(int i = 1; i != 100000+1; i++)
    {
        DP[i] = -1;
    }
    for(int i = 1; i != N+1; i++)
    {
        readIntArray(&ZAN, 1);
        NumsMap[ZAN]++;
    }

    printf("%I64d\n", dp(1));

}

long long dp(int index)
{
    if(index > 100000)
    {
        return 0;
    }
    if(DP[index] != -1)
    {
        return DP[index];
    }

    long long m = max(NumsMap[index]*index+dp(index+2), dp(index+1));
    DP[index] = m;
    return m;

}

long long max(long long a, long long b)
{
    return a>b?a:b;
}
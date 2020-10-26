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
int min3(int a, int b, int c)
{
    int ab = a<b?a:b;
    return ab<c?ab:c;
}
int min2(int a, int b)
{
    return a<b?a:b;
}
int N;
int DP[101][4];
int DAYS[101];
int main()
{
    // init DP
    for(int i = 1; i!=101; i++)
    {
        for(int j = 1; j != 4; j++)
        {
            DP[i][j] = 100000;
        }
    }

    readIntArray(&N, 1);
    readIntArray(DAYS+1, N);

    // start DP
    if(DAYS[N]==0)
    {
        DP[N][1] = 1;
    }
    else if(DAYS[N]==1)
    {
        DP[N][2] = 0;
        DP[N][1] = 1;
    }
    else if(DAYS[N]==2)
    {
        DP[N][1] = 1;
        DP[N][3] = 0;
    }
    else
    {
        DP[N][1] = 1;
        DP[N][2] = 0;
        DP[N][3] = 0;
    }

    for(int i = N-1; i!=0; i--)
    {
        if(DAYS[i]==0) // r
        {
            DP[i][1] = min3(DP[i+1][1], DP[i+1][2], DP[i+1][3]) + 1; // choose to rest
        }
        else if(DAYS[i]==1) // c2 r1
        {
            DP[i][2] = min2(DP[i+1][1], DP[i+1][3]);                // choose to contest
            DP[i][1] = min3(DP[i+1][1], DP[i+1][2], DP[i+1][3]) + 1;    // choose to rest
        }
        else if(DAYS[i]==2) // g3 r1
        {
            DP[i][1] = min3(DP[i+1][1], DP[i+1][2], DP[i+1][3]) + 1;
            DP[i][3] = min2(DP[i+1][1], DP[i+1][2]);
        }
        else                // c2 g3 r1
        {
            DP[i][1] = min3(DP[i+1][1], DP[i+1][2], DP[i+1][3]) + 1;
            DP[i][2] = min2(DP[i+1][1], DP[i+1][3]);
            DP[i][3] = min2(DP[i+1][1], DP[i+1][2]);
        }
    }
    printf("%d\n", min3(DP[1][1], DP[1][2], DP[1][3]));

}
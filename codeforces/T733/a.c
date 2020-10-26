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
void readStr(char* target){
    scanf("%s",target);
}
void printStr(char* target, int N){
    for(int i= 0;i<N;i++)
    {
        printf("%c ", *(target+i));
    }
}
char NUM[100];
int INIT = 0;
int MAX = 0;
int main()
{
    readStr(NUM);
    for(int i = 0; i<strlen(NUM); i++){
        if(NUM[i] == 'A' || NUM[i] == 'E' || NUM[i] == 'I' || NUM[i] == 'O' || NUM[i] == 'U' || NUM[i] == 'Y'){
            int now = i+1-INIT;
            if(now > MAX){
                MAX = now;
            }
            INIT = i+1;
        }
    }
    if(strlen(NUM) - INIT + 1 > MAX){   
        MAX = strlen(NUM) - INIT + 1;
    }
    printf("%d\n", MAX);
}
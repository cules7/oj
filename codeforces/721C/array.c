#include <stdio.h>
#include <sys/time.h>
int a[5000][5000];
int b[5000*5000];
int main(){
    struct timeval start, end, zaiend;
    gettimeofday(&start, 0);
    for(int i = 0; i<100; i++){
        for(int jia = 0; jia<5000; jia++)
        {
            for(int yi = 0; yi<5000; yi++)
            {
                a[jia][yi] = 3;
            }
        }
    }
    gettimeofday( &end, 0 );

    for(int i = 0; i<100; i++)
    {
        for(int jia = 0; jia<5000*5000; jia++)
        {
            b[jia] = 4;
        }
    }
    gettimeofday( &zaiend, 0 );
    printf("%ld . %d \n", start.tv_sec, start.tv_usec);
    printf("%ld . %d \n", end.tv_sec, end.tv_usec);
    printf("%ld . %d \n", zaiend.tv_sec, zaiend.tv_usec);
    return 0;
}
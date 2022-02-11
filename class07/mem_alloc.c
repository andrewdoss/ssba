// Demo for malloc/pmap watching
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main()
{
    int *p;
    while (1)
    {
        p = malloc(sizeof(int));
        printf("%ld\n", (long)p);
        // Sleep for 1 ms
        usleep(1000);
    }
    return 0;
}
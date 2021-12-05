#include <stdio.h>  /* printf, scanf, NULL */
#include <stdlib.h>  /* malloc, free, rand, system */

void malloc_buffer() {
    long i,n;
    char * buffer;

    i = 1000000000*sizeof(char);

    buffer = (char*)malloc(i);
    if(buffer==NULL) exit(1);  // 判断是否分配成功

    free(buffer);  // 释放内存空间
}

int main ()
{
    long i;

    for(i=0; i< 20000000000; i++)
        malloc_buffer();

    return 0;
}

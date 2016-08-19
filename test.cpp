#include <iostream>
#include <vector>
#include <stdio.h>
int main()
{
    std::vector< int > data;
    data.push_back( 1 );
    data.push_back( 2 );
    data.push_back( 3 );
    data.push_back( 4 );
    data.push_back( 5 );
    data.push_back( 6 );
    data.push_back( 7 );
    data.push_back( 8 );
    data.push_back( 9 );
    data.push_back( 10 );
    int front = 0;
    int end = 9;
    int middle = ( front + end ) / 2;
    int data1 = 10;

    while ( data[middle] != data1 && front < end )
    {
        if ( data[middle] > data1 )
            end = middle - 1;
        else if ( data[middle] < data1 )
            front = middle + 1;
        middle = ( front + end ) / 2;
    }
    printf( "point = %d\n", middle );
}

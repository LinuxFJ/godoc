// =====================================================================================
// 
//       Filename:  tree.c
// 
//    Description:  
// 
//        Version:  1.0
//        Created:  2015年12月24日 10时15分56秒
//       Revision:  none
//       Compiler:  g++
// 
//         Author:  fancy linuxfj@gmail.com
//        Company:  fantancy tec.
// 
// =====================================================================================

#include <math.h>
#include <stdio.h>
#include <stdlib.h>

#define PI 3.14159265359

float sx, sy;

float sdCircle(float px, float py, float r) {
    float dx = px - sx, dy = py - sy;
    return sqrtf(dx * dx + dy * dy) - r;
}

float opUnion(float d1, float d2) {
    return d1 < d2 ? d1 : d2;
}

#define T px + scale * r * cosf(theta), py + scale * r * sin(theta)

float f(float px, float py, float theta, float scale, int n) {
    float d = 0.0f;
    float r = 0.0f;
    for (; r < 0.8f; r += 0.02f)
        d = opUnion(d, sdCircle(T, 0.05f * scale * (0.95f - r)));

    if (n > 0) {
        int t = -1;
        for (; t <= 1; t += 2) {
            float tt = theta + t * 1.8f;
            float ss = scale * 0.9f;
            float r = 0.2f;
            for (; r < 0.8f; r += 0.1f) {
                d = opUnion(d, f(T, tt, ss * 0.5f, n - 1));
                ss *= 0.8f;
            }
        }
    }
    return d;
}

int main(int argc, char* argv[]) {
    int n = argc > 1 ? atoi(argv[1]) : 3;
    float sy = 0.8f;
    for (; sy > 0.0f; sy -= 0.02f, putchar('\n')) {
        float sx = -0.35f;
        for (; sx < 0.35f; sx += 0.01f)
            putchar(f(0, 0, PI * 0.5f, 1.0f, n) < 0 ? '*' : ' ');
    }
}

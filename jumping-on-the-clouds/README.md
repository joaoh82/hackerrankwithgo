**Jumping on the Clouds**

Link: https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem

Emma is playing a new mobile game involving n clouds numbered from 0 to n-1 . A player initially starts out on cloud c0, and they must jump to cloud c(n-1). In each step, she can jump from any cloud i to cloud i+1 or cloud i + 2.
There are two types of clouds, ordinary clouds and thunderclouds. The game ends if Emma jumps onto a thundercloud, but if she reaches the last cloud (i.e., c(n-1)), she wins the game!

Can you find the minimum number of jumps Emma must make to win the game? It is guaranteed that clouds c0 and c(n-1) are ordinary-clouds and it is always possible to win the game.

**Input Format**

The first line contains an integer,n (the total number of clouds). 
 The second line contains n space-separated binary integers describing clouds c0,c1,…c(n-1).

If ,c(i) = 0 the cloud i is an ordinary cloud.
If ,c(i) = 1 the cloud i is a thundercloud.

**Constraints**
```
2 <= n <= 100
c(i) is an element of {0,1}
c(0) = c(n-1) = 0;
```
**Output Format**

Print the minimum number of jumps needed to win the game.

Sample Input 0
```
7
0 0 1 0 0 1 0
```
Sample Output 0
```
4
```
Sample Input 1
```
6
0 0 0 0 1 0
```
Sample `utput 1
```
3
```
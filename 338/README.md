https://leetcode.com/problems/counting-bits/description/

## Note

One essence of DP is to reuse the answers that you have calculated as many as possible. 

So you need to find a way to convert the problem to a cumulative issue 
that the new answer is based on previous answers

One simple example is Fibonacci, which is so easy 
because you already know the cumulative way to solve this, f(n) = f(n-1) + f(n-2)

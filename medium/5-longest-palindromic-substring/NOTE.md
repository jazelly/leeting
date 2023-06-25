## Note

The DP way that I can think in the initial approach was to reuse sub-palindrome.

Every palindrom must contain a sub-palindrome, from 'aba' or 'bb'. 
So we scan from there, if they can expand wider and wider, then it's counted a bigger palindrome. 
Notice, for the case 'bbb', it's not originated from 'bb'. Instead, 'bbb' is the starting point.

~~There must be a better way~~

Ok, checked the solution. Looks like the O(n^2) is the best way?

One lesson I learned was stopping diving into O(n) solution in the first place. 
Been wasting a lot of time on that. Need to get a sense what is possible

Or, do the brute force and slowly optimise from there.
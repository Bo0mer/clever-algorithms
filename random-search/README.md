# Random Search

## Idea
The algorithm is simple. Generate a random candidate from the search space, try it out, and if it gives better result than the one which we currently have, save it. Repeat this as many iterations as needed.

## Pseudocode

```
Input: Iterations, ProblemSize, SearchSpace
Output: Best

Best <- nil
for i in Iterations do:
  candidate = RandomCandidate(ProblemSize, SearchSpace)
  if Cost(candidate) < Cost(Best):
    Best = candidate
  end
end

return Best

```

## Notes

* The worst case performance of the algorithm is worse than brute-force (enumeration) search.
* The algorithm works well for low problem dimensions. However, if the dimensions scales up, the algorithms does not scale well with it.
* The random candidates must be uniformly distributed. This may be a problem in some specific search spaces.

## The actual problem

The code in `rs.go` deals with actual problem. The problem is the following:

Given a function `f(x1, x2, ..., xn) = sum(xi^2)`, find `min(f)`, where `-5.0 < xi < 5.0, for i in [1, n]`.
We'll solve the problem for n = 2. As it is obvious, the best possible solution is f(0, 0) = 0.

Below is table with the output for the specified number of iterations.

| Iterations | x1                    | x2                     | f(x1, x2)             |
| ---------- | --------------------- | ---------------------- | --------------------- |
| 10         |  1.3177317309393572   | -0.24540872490426135   | 1.79664235698357      |
| 100        | -0.0756558912799985   |  0.34546763351919907   | 0.1250716996947266    |
| 1000       |  0.04016917665821751  |  0.04231913552663258   | 0.0034044719851205824 |
| 100000     | -0.044433306109305626 |  0.0023000501983876376 | 0.0019796089227183595 |

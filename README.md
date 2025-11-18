# Balanced Binary Tree Checker

## What This Project Does

This project checks whether a binary tree is balanced.

A binary tree is considered balanced when the height difference between the left subtree and the right subtree of every node is not more than 1.

- If the difference is 0 or 1 → the tree is balanced
- If the difference is 2 or more → the tree is unbalanced

## Methods Used

Two different approaches were explored for checking if a tree is balanced. Both are correct, but their efficiency is very different.

### 1. Slow Approach — Checking Each Node Separately

#### How It Works

- For every node, the program:
  - Calculates the height of the left subtree
  - Calculates the height of the right subtree
  - Compares the two
- Then it repeats the same process for every node in the tree.

#### Why This Is Slow

The heights of the same subtrees are calculated over and over again, resulting in a lot of unnecessary work.

#### Efficiency

- **Time Complexity:** `O(n^2)`
- **Space Complexity:** `O(h)` where `h` = height of the tree

This becomes very slow when the tree is large.

### 2. Fast Approach — Checking Balance While Computing Height

#### How It Works

This approach calculates height and checks balance at the same time:

1. Compute the height of the left subtree
2. Compute the height of the right subtree
3. If either subtree is unbalanced, return immediately
4. If both sides are valid, return the actual height

A special value (`-1`) is used to indicate an unbalanced subtree.

#### Why This Is Fast

- Each node is visited only once
- No repeated height calculations
- The function stops early when an imbalance is found

#### Efficiency

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(h)` where `h` = height of the tree

This is the optimal solution for this problem.

### Why the Fast Method Is Better

- It avoids repeated calculations
- It visits each node only one time
- It stops as soon as it detects an imbalance
- It provides the best possible performance for this problem

## Examples

### Balanced Tree Example

```
    1
   / \
  2   3
 / \
4   5
```

**Result:** `true`

### Unbalanced Tree Example

```
1
 \
  2
   \
    3
```

**Result:** `false`

## Summary

- A balanced tree requires the left and right subtree heights to differ by at most 1
- The slow method works but is inefficient (`O(n^2)`)
- The optimized method is efficient and clean (`O(n)`)
- The optimized method is the final and recommended solution

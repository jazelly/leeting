from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def insert(cur, node):
  if node.val == cur.val:
    return

  if node.val < cur.val:
    if cur.left == None:
      cur.left = node
      return
    else:
      insert(cur.left, node)
  else:
    if cur.right == None:
      cur.right = node
    else:
      insert(cur.right, node)

def search(root, val):
  if root == None:
     return

  if val == root.val:
    return val

  return search(root.left) if val < root.val else search(root.right)
      
class Solution:
    def rangeSumBST(self, root: Optional[TreeNode], low: int, high: int) -> int:
      sum = 0
      for i in range(low, high + 1):
        res = search(root, i)
        if res:
          sum += res
      return sum

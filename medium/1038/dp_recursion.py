# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def search(root, val):
  if root == None:
     return

  if val == root.val:
    return val

  return search(root.left) if val < root.val else search(root.right)

def post_order(root):
  if root == None:
     return
  post_order(root.right)
  print(root.val)
  post_order(root.left)

class Solution:
    def bstToGst(self, root: TreeNode) -> TreeNode:
      rightest = post_order(root)
      
      
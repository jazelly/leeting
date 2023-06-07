# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def post_order(root, res = []):
  if root.right != None:
    post_order(root.right, res)
  res.append(root) # print root.val
  if root.left != None:
    post_order(root.left, res)
  
# Convert a BST to array
# None on None node
def bfs(roots, res = []):
  newRoots = []
  buffer = []
  allNone = True
  for r in roots:
    buffer.append(r.val) if r != None else buffer.append(None)
    if r != None:
      allNone = False
      newRoots.append(r.left)
      newRoots.append(r.right)
    else:
       newRoots.append(None)
       newRoots.append(None)
  if allNone:
    return res
  else:
    res += buffer
    bfs(newRoots, res)
  return res

class Solution:
    def bstToGst(self, root: TreeNode):
        nodes = []
        post_order(root, nodes)
        c = 0
        for node in nodes:
          c += node.val
          node.val = c
      
        return root

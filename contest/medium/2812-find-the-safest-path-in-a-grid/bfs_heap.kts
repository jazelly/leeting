import java.util.PriorityQueue

class Solution {
    fun maximumSafenessFactor(grid: List<List<Int>>): Int {
        val INF = Integer.MAX_VALUE
        val directions = arrayOf(intArrayOf(0,1), intArrayOf(1,0), intArrayOf(0,-1), intArrayOf(-1,0))
        val n = grid.size

        if (grid[0][0] == 1 || grid[n - 1][n - 1] == 1) return 0
        val safe = Array(n) { IntArray(n) { -1 }}
        var queue = LinkedList<IntArray>()
        for (r in 0 until n) {
            for (c in 0 until n) {
                if (grid[r][c] == 1) {
                    queue.offer(intArrayOf(r, c))
                    safe[r][c] = 0
                }
            }
        }
        while (!queue.isEmpty()) {
            val temp = LinkedList<IntArray>()
            for (node in queue) {
                for (direction in directions) {
                    val newX = node[0] + direction[0]
                    val newY = node[1] + direction[1]
                    if (newX < 0 || newX >= n || newY < 0 || newY >= n || safe[newX][newY] != -1) continue
                    temp.offer(intArrayOf(newX, newY))
                    safe[newX][newY] = safe[node[0]][node[1]] + 1
                }
            }
            queue = temp
        }
        
        val heap = PriorityQueue<IntArray>() { e1, e2 ->
            e2[0] - e1[0]
        }
        heap.offer(intArrayOf(safe[0][0], 0, 0))
        val visit = Array(n) { BooleanArray(n) }
        visit[0][0] = true
        while (!heap.isEmpty()) {
            val node = heap.poll()
            if (node[1] == n - 1 && node[2] == n - 1) return node[0]
            for (direction in directions) {
                val newX = node[1] + direction[0]
                val newY = node[2] + direction[1]
                if (newX < 0 || newX >= n || newY < 0 || newY >= n || visit[newX][newY]) continue
                heap.offer(intArrayOf(Math.min(node[0], safe[newX][newY]), newX, newY))
                visit[newX][newY] = true
            }
        }
        return 0
    }
}
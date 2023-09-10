class Solution:
    def isReachableAtTime(self, sx: int, sy: int, fx: int, fy: int, t: int) -> bool:
        dx = abs(fx - sx)
        dy = abs(fy - sy)

        if sx == fx and sy == fy:
            if t == 1:
                return False
            return True

        initialStep = min(dx, dy)
        rx = dx - initialStep
        ry = dy - initialStep
        total = initialStep + max(rx, ry)

        if total <= t:
            return True
        return False
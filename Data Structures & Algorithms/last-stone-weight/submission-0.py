class Solution:
	def lastStoneWeight(self, stones: List[int]) -> int:
		heap = []
		for s in stones:
			heap.append(-s)
		heapq.heapify(heap)

		while len(heap) > 1:
			big1 = heapq.heappop(heap)
			big2 = heapq.heappop(heap)
			if big2 > big1:
				heapq.heappush(heap, big1 - big2)

		heap.append(0)
		return abs(heap[0])
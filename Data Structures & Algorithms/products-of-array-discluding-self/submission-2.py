class Solution:
	def productExceptSelf(self, nums: List[int]) -> List[int]:
		prefix = []
		suffix = [0] * len(nums)
		res = []

		for i in range(0, len(nums)):
			if i == 0:
				prefix.append(1)
			else:
				prev = prefix[i-1]
				prefix.append(prev * nums[i-1])

		for i in range(len(nums)-1, -1, -1):
			if i == len(nums)-1:
				suffix[i] = 1
			else:
				after = suffix[i+1]
				suffix[i] = nums[i+1]*after


		res = []
		for i in range(0, len(nums)):
			a = prefix[i]
			b = suffix[i]

			res.append(a*b)

		return res

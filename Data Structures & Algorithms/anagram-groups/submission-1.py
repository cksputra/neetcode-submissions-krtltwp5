class Solution:
	def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
		dic = {}
		
		
		for str in strs:
			arr = [0] * 26
			for char in str:
				arr[ord(char)-ord('a')]+=1

			if tuple(arr) not in dic:
				dic[tuple(arr)] = [str]
			else:
				dic[tuple(arr)].append(str)

		result = []
		for key in dic:
			result.append(dic[key])
		

		return result
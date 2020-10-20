/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

官方的思路，为防以后找不到，所以转存一下
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func lengthOfLongestSubstring(s string) int {
    n:=len(s)
    d:=map[byte]int{}
    ans,r:=0,-1						//result and right_point
    for i:=0;i<n;i++{
        if(i!=0){
            delete(d,s[i-1])
        }
        for r+1<n && d[s[r+1]]==0{
            d[s[r+1]]++
            r++
        }
        ans=max(ans,r-i+1)
    }
    return ans
}
func max(a int,b int) int{
    if(a<b){
        return b
    }
    return a
}
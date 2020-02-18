/*
https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
Given an integer number n, return the difference between the product of its digits and the sum of its digits.


Example 1:

Input: n = 234
Output: 15
Explanation:
Product of digits = 2 * 3 * 4 = 24
Sum of digits = 2 + 3 + 4 = 9
Result = 24 - 9 = 15
 */
/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function (n) {
    let str = "" + n
    let arr = str.split('')
    let product = 1
    let sum = 0
    arr.forEach((a, i) => {
        const num = Number(a)
        product *= num
        sum += num
    })
    return sum - product
};

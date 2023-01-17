# Go solutions to LeetCode problems

The purpose of this repository is to serve as a portfolio of my work and to provide a resource for others seeking clear and easy-to-understand solutions to these problems. Every day starting January 5, 2023, I will be adding my solutions to this repository.

> <p style = "color:red"><strong>Please note that:</strong> my solutions are not necessarily optimized for efficiency in terms of time complexity or memory consumption. Instead, my focus is on providing solutions that are clear, follow best practices, and are easy for others to understand and learn from.</p>

If you are seeking more clever or efficient solutions, I recommend checking out these repositories from <a href="https://github.com/halfrost/LeetCode-Go" target="_blank">halfrost</a> and <a href="https://github.com/aQuaYi/LeetCode-in-Go" target="_blank">aQuaYi</a>. They contain solutions that are highly optimized in terms of time and memory complexity.

I understand how challenging it can be to work through these problems and I hope that my solutions can provide guidance and direction for those struggling to find their own solutions.

## Completion Summary

[My LeetCode profile](https://leetcode.com/goldennovember/)

| Difficulty             | Solved Problems |
|------------------------|:---------------:|
| :green_circle: Easy    |       33        |
| :yellow_circle: Medium |        5        |
| :red_circle: Hard      |        1        |
| :black_circle: Total   |       39        |

## Repository Structure

I also add the test cases that I failed during my submissions, which is not included in the original problem's descriptions on the website.

```ascii
├── easy/
│   ├── word-pattern/                       
│   │   ├── word-pattern.go                 // Solution
│   │   ├── word-pattern_test.go            // Tests
│   │   └── README.md                       // Description
│   │
│   ├── contains-duplicate/
│   │   ├── contains-duplicate.go
│   │   ├── contains-duplicate_test.go
│   │   └── README.md
│   └── ...
│   
├── medium/
│   ├── group-anagrams/
│   │   ├── group-anagrams.go
│   │   ├── group-anagrams_test.go
│   │   └── README.md
│   └── ...
│  
│── hard/  
│
└── README.md                            
```

## Solutions


| No.  | Title                                                  |                                Solution                                | Acceptance | Difficulty |                             Tags                              |
|:----:|:-------------------------------------------------------|:----------------------------------------------------------------------:|:----------:|:----------:|:-------------------------------------------------------------:|
| 0001 | Two Sum                                                |                           [Go](easy/two-sum)                           |   49.2%    |    Easy    |                     `Array` `Hash Table`                      |
| 0009 | Palindrome Number                                      |                      [Go](easy/palindrome-number)                      |   53.1%    |    Easy    |                            `Math`                             |
| 0012 | Integer to Roman                                       |                     [Go](medium/integer-to-roman)                      |   61.7%    |   Medium   |                 `Hash Table` `Math` `String`                  |
| 0013 | Roman to Integer                                       |                      [Go](easy/roman-to-integer)                       |   58.2%    |    Easy    |                 `Hash Table` `Math` `String`                  |
| 0041 | First Missing Positive                                 |                   [Go](hard/first-missing-positive)                    |   36.6%    |    Hard    |                     `Array` `Hash Table`                      |
| 0049 | Group Anagrams                                         |                      [Go](medium/group-anagrams)                       |   66.6%    |   Medium   |            `Array` `Hash Table` `String` `Sorting`            |
| 0217 | Contains Duplicate                                     |                     [Go](easy/contains-duplicate)                      |   61.4%    |    Easy    |                `Array` `Hash Table` `Sorting`                 |
| 0242 | Valid Anagram                                          |                        [Go](easy/valid-anagram)                        |   62.9%    |    Easy    |                `Hash Table` `String` `Sorting`                |
| 0290 | Word Pattern                                           |                        [Go](easy/word-pattern)                         |   41.7%    |    Easy    |                     `Hash Table` `String`                     |
| 0349 | Intersection of Two Arrays                             |                 [Go](easy/intersection-of-two-arrays)                  |   70.6%    |    Easy    | `Array` `Hash Table` `Two Pointers` `Binary Search` `Sorting` |
| 0442 | Find All Duplicates in an Array                        |              [Go](medium/find-all-duplicates-in-an-array)              |   73.4%    |   Medium   |                     `Array` `Hash Table`                      |
| 0771 | Jewels and Stones                                      |                      [Go](easy/jewels-and-stones)                      |   88.1%    |    Easy    |                     `Hash Table` `String`                     |
| 1108 | Defanging an IP Address                                |                   [Go](easy/defanging-an-ip-address)                   |   89.3%    |    Easy    |                           `String`                            |
| 1281 | Subtract the Product and Sum of Digits of an Integer   |    [Go](easy/subtract-the-product-and-sum-of-digits-of-an-integer)     |   86.7%    |    Easy    |                            `Math`                             |
| 1282 | Group the People Given the Group Size They Belong To   |   [Go](medium/group-the-people-given-the-group-size-they-belong-to)    |   85.7%    |   Medium   |                     `Array` `Hash Table`                      |
| 1365 | How Many Numbers Are Smaller Than the Current Number   |    [Go](easy/how-many-numbers-are-smaller-than-the-current-number)     |   86.7%    |    Easy    |           `Array` `Hash Table` `Sorting` `Counting`           |
| 1393 | Capital Gain/Loss                                      |                   [SQL](database/capital-gain-loss)                    |   89.7%    |   Medium   |                          `Database`                           |
| 1431 | Kids With the Greatest Number of Candies               |          [Go](easy/kids-with-the-greatest-number-of-candies)           |   87.4%    |    Easy    |                            `Array`                            |
| 1470 | Shuffle the Array                                      |                      [Go](easy/shuffle-the-array)                      |   88.5%    |    Easy    |                            `Array`                            |
| 1480 | Running Sum of 1d Array                                |                   [Go](easy/running-sum-of-1d-array)                   |   88.7%    |    Easy    |                     `Array` `Prefix Sum`                      |
| 1484 | Group Sold Products By The Date                        |            [SQL](database/group-sold-products-by-the-date)             |   81.8%    |    Easy    |                          `Database`                           |
| 1512 | Number of Good Pairs                                   |                    [Go](easy/number-of-good-pairs)                     |   88.2%    |    Easy    |            `Array` `Hash Table` `Math` `Counting`             |
| 1528 | Shuffle String                                         |                       [Go](easy/shuffle-string)                        |   85.5%    |    Easy    |                       `Array` `String`                        |
| 1581 | Customer Who Visited but Did Not Make Any Transactions | [SQL](database/customer-who-visited-but-did-not-make-any-transactions) |   86.8%    |    Easy    |                          `Database`                           |
| 1587 | Bank Account Summary II                                |                [SQL](database/back-account-summary-ii)                 |   88.2%    |    Easy    |                          `Database`                           |
| 1672 | Richest Customer Wealth                                |                   [Go](easy/richest-customer-wealth)                   |   88.1%    |    Easy    |                       `Array` `Matrix`                        |
| 1693 | Daily Leads and Partners                               |                [SQL](database/daily-leads-and-partners)                |   89.3%    |    Easy    |                          `Database`                           |
| 1920 | Build Array from Permutation                           |                [Go](easy/build-array-from-permutation)                 |   91.0%    |    Easy    |                     `Array` `Simulation`                      |
| 1741 | Find Total Time Spent by Each Employee                 |         [SQL](database/find-total-time-spent-by-each-employee)         |   90.8%    |    Easy    |                          `Database`                           |
| 1757 | Recyclable and Low Fat Products                        |            [SQL](database/recyclable-and-low-fat-products)             |   93.2%    |    Easy    |                          `Database`                           |
| 1795 | Rearrange Products Table                               |                [SQL](database/rearrange-products-table)                |   88.1%    |    Easy    |                          `Database`                           |
| 1929 | Concatenation of Array                                 |                   [Go](easy/concatenation-of-array)                    |   90.9%    |    Easy    |                            `Array`                            |
| 2011 | Final Value of Variable After Performing Operations    |     [Go](easy/final-value-of-variable-after-performing-operations)     |   88.8%    |    Easy    |                 `Array` `String` `Simulation`                 |
| 2032 | Two Out of Three                                       |                      [Go](easy/two-out-of-three)                       |   73.0%    |    Easy    |                     `Array` `Hash Table`                      |
| 2114 | Maximum Number of Words Found in Sentences             |         [Go](easy/maximum-number-of-words-found-in-sentences)          |   87.8%    |    Easy    |                       `Array` `String`                        |
| 2235 | Add Two Integers                                       |                      [Go](easy/add-two-integers)                       |   89.2%    |    Easy    |                            `Math`                             |
| 2413 | Smallest Even Multiple                                 |                   [Go](easy/smallest-even-multiple)                    |   87.6%    |    Easy    |                    `Math` `Number Theory`                     |
| 2441 | Largest Positive Integer That Exists With Its Negative |   [Go](easy/largest-positive-integer-that-exists-with-its-negative )   |   67.8%    |    Easy    |                     `Array` `Hash Table`                      |
| 2469 | Convert the Temperature                                |                   [Go](easy/convert-the-temperature)                   |   89.8%    |    Easy    |                            `Math`                             |



## Resources

https://leetcode.com/
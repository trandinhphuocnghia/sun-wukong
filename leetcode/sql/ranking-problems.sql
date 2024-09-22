-- RANKING PROBLEMS, Using Window Functions like: RANK(), DENSE_RANK() + OVER ()

/*
* Problems: Show Rank Scores
* url: https://leetcode.com/problems/rank-scores/description/
* Tables:
 -- Scores:
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| score       | decimal |
+-------------+---------+
*/
--MYSQL
SELECT score, DENSE_RANK() OVER (ORDER BY score) as 'rank' FROM Scores;

/*
* Problems: Find Second Highest Salary
* url: https://leetcode.com/problems/second-highest-salary/description/
* Tables:
-- Employees:
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
*/
--MYSQL
-- Find in ranking list:
WITH Ranking as (SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) as 'rank' FROM Employees)
SELECT (SELECT DISTINCT salary FROM Ranking WHERE Ranking.rank = 2) AS SecondHighestSalary;
-- Find by MAX of (salary < MAX of salary):
SELECT MAX(salary) AS SecondHighestSalary FROM Employees WHERE salary < (SELECT MAX(salary) FROM Employees);
--TSQL, SQL Server:
WITH Ranking as (SELECT salary, DENSE_RANK() OVER (OVER BY salary DESC) as rank FROM Employees)
SELECT (SELECT salary FROM Ranking WHERE rank = 2) AS  SecondHighestSalary;

/*
* Problems: Find the nth Highest Salary.
* url: https://leetcode.com/problems/nth-highest-salary/description/
* Tables:
-- Employee:
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
* Approaches:
* this Problems have the input parameter => need to define an SQL function to handle the input parameter.
*/
-- MySQL
CREATE FUNCTION getNthHighestSalary(N INT) RETURN INT
BEGIN
RETURN(
    SELECT DISTINCT salary FROM (SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) AS 'rank' FROM Employee) AS Ranking WHERE Ranking.rank = N;
);
END
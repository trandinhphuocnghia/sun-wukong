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
-- TSQL, MySQL Server
CREATE FUNCTION getNthHighestSalary(@N INT) RETURN INT
BEGIN
RETURN(
    SELECT DISTINCT salary FROM (SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) as rank FROM Employee) AS Ranking WHERE rank = @N;
);
END


/*
* Problem: Department Highest Salary
* url: https://leetcode.com/problems/department-highest-salary/description/
* Tables:
--Employee
+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| id           | int     |
| name         | varchar |
| salary       | int     |
| departmentId | int     |
+--------------+---------+

--Department
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
*/
--MYSQL 
-- Partition:
WITH JoinedTable AS (SELECT E.name as Employee, E.salary as Salary, D.name as Department, Rank() OVER(PARTITION BY E.departmentId ORDER BY E.salary DESC) AS 'rank' FROM Employee AS E JOIN Department AS D ON E.departmentId = D.id)
SELECT Department,Employee,Salary FROM JoinedTable WHERE JoinedTable.rank = 1;
-- Group by:
SELECT
    d.name AS Department,
    e.name AS Employee,
    e.salary AS Salary
FROM
    Employee e
JOIN
    Department d ON e.departmentId = d.id
WHERE
    (e.departmentId, e.salary) IN (
        SELECT
            departmentId,
            MAX(salary)
        FROM
            Employee
        GROUP BY
            departmentId
    )
ORDER BY
    d.name, e.name


/*
* Problems: Department Top Three Salaries
* url: https://leetcode.com/problems/department-top-three-salaries/description/
* Tables: both above tables.
*/
--MYSQL
WITH JoinedTable AS (SELECT E.name as Employee, E.salary as Salary, D.name as Department, E.departmentId AS departmentId, DENSE_RANK() OVER (PARTITION BY E.departmentId ORDER BY E.salary DESC) AS 'rank' FROM Employee AS E JOIN Department AS D ON E.departmentId = D.id)
SELECT DISTINCT Department,Employee,Salary FROM JoinedTable WHERE JoinedTable.rank <= 3;
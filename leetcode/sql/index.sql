/*
* Problem: Delete duplicate email.
* url: https://leetcode.com/problems/delete-duplicate-emails
* Table:
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| email       | varchar |
+-------------+---------+
id is the primary key (column with unique values) for this table.
Each row of this table contains an email. The emails will not contain uppercase letters.
*/
/*GROUP BY*/
--MySQL
DELETE person from person JOIN (SELECT MIN(id) as id, email FROM person GROUP BY email) as emailGroups WHERE person.email = emailGroups.email AND person.id != emailGroups.id;
--SQL server
DELETE FROM person WHERE id NOT IN (SELECT MIN(id) from person GROUP BY email);
/*PARTITION*/
--MySQL
WITH CTE AS (SELECT id, email, ROW_NUMBER() OVER (PARTITION BY email ORDER BY id) as rn FROM person)
DELETE FROM person WHERE id NOT IN (SELECT id FROM CTE WHERE rn > 1);
--SQL server
WITH CTE AS (SELECT id, email, ROW_NUMBER() OVER (PARTITION BY email ORDER BY id) as rn FROM person)
DELETE FROM CTE WHERE rn>1;


/*
*Problem: Sale Persons.
* url: https://leetcode.com/problems/sales-person/description/
*Table:
- SalePersons:
+-----------------+---------+
| Column Name     | Type    |
+-----------------+---------+
| sales_id        | int     |
| name            | varchar |
| salary          | int     |
| commission_rate | int     |
| hire_date       | date    |
+-----------------+---------+
- Company:
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| com_id      | int     |
| name        | varchar |
| city        | varchar |
+-------------+---------+
- Orders:
+-------------+------+
| Column Name | Type |
+-------------+------+
| order_id    | int  |
| order_date  | date |
| com_id      | int  |
| sales_id    | int  |
| amount      | int  |
*/
--MySQL:
SELECT name FROM SalePersons WHERE id NOT IN (SELECT sales_id FROM Orders LEFT JOIN Company USING(com_id) RIGHT JOIN SalePersons USING(sales_id) WHERE Company.name='RED' )


/*
* Problems: Triangle Judgement
* url: https://leetcode.com/problems/triangle-judgement/
* Table:
+-------------+------+
| Column Name | Type |
+-------------+------+
| x           | int  |
| y           | int  |
| z           | int  |
+-------------+------+
*/
--MYSQL:
SELECT x,y,z, CAST(CASE WHEN x+y>z AND x+z>y AND y+z>x THEN 'Yes' ELSE 'No' END AS CHAR(3)) AS triangle FROM Triangle;
--TSQL, MySQL server:
SELECT *, CASE WHEN x+y>z AND x+z>y AND y+z>x THEN 'Yes' ELSE 'No' END AS triangle FROM Triangle;

/*
* Problems: Biggest Single Number
* url: https://leetcode.com/problems/biggest-single-number/description/
* Table: MyNumbers
+-------------+------+
| Column Name | Type |
+-------------+------+
| num         | int  |
+-------------+------+
*/
--MySQL,TSQL SQL server:
SELECT MAX(num) FROM MyNumbers WHERE num NOT IN (SELECT num FROM MyNumbers GROUP BY num HAVING COUNT(num) > 1);

/*
* Problem: Swap Salary
* url: https://leetcode.com/problems/swap-salary/description/
* Table: 
+-------------+----------+
| Column Name | Type     |
+-------------+----------+
| id          | int      |
| name        | varchar  |
| sex         | ENUM     |
| salary      | int      |
+-------------+----------+
*/
-- MYSQL
UPDATE SALARY SET sex CASE WHEN sex = 'm' THEN 'f' ELSE 'm' END;
-- SQL Server, TSQL
UPDATE SALARY SET sex = IIF(sex='m','f','m');

/*
* Project: Project Employees I
* url: https://leetcode.com/problems/project-employees-i/description/
* Table:
Project
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| project_id  | int     |
| employee_id | int     |
+-------------+---------+
Employee
+------------------+---------+
| Column Name      | Type    |
+------------------+---------+
| employee_id      | int     |
| name             | varchar |
| experience_years | int     |
+------------------+---------+
*/
SELECT project_id, ROUND(AVG(experience_years)) AS experience_years FROM Project JOIN Employee ON Project.employee_id = Employee.employee_id;

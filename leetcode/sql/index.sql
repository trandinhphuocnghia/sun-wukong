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




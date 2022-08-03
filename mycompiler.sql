CREATE TABLE emp (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  dep_id INTEGER NOT NULL,
  salary INTEGER NOT NULL
);

INSERT INTO emp VALUES (1, 'Agus', 1, 100);
INSERT INTO emp VALUES (2, 'Caca', 1, 200);
INSERT INTO emp VALUES (3, 'Eka', 1, 300);

INSERT INTO emp VALUES (4, 'Budi', 2, 50);
INSERT INTO emp VALUES (5, 'Dedi', 2, 100);
INSERT INTO emp VALUES (6, 'Feri', 2, 150);

-- SELECT * FROM emp WHERE id = 1;

CREATE TABLE dep (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL
);

INSERT INTO dep VALUES (1, 'IT');
INSERT INTO dep VALUES (2, 'BD');

WITH a as (
    select emp.id, max(salary)
    from emp join dep on emp.dep_id = dep.id
    where dep.name in ('IT', 'BD')
    group by dep_id
    order by salary DESC
)
SELECT MAX(emp.salary) - MIN(emp.salary) FROM emp join a on a.id = emp.id;

-- select ITEMS

-- cuma tau nama dep: IT, BD => selisih gaji terbesar

select ABS(a.salary - b.salary) from (
    select salary from emp 
    where dep_id = 1
    order by salary desc limit 1
) as a, (
    select salary from emp 
    where dep_id = 2
    order by salary desc limit 1
) as b;

SELECT (
       MAX(CASE WHEN dep.name = 'IT' THEN salary END)  
       -
       MAX(CASE WHEN dep.name = 'BD' THEN salary END)) AS salary_difference 
FROM emp
INNER JOIN  dep
ON emp.dep_id = dep.id;

# Write your MySQL query statement below

select b.name, sum(a.amount) as balance
from Transactions a left join Users b
on a.account=b.account
group by b.name
having balance >10000;
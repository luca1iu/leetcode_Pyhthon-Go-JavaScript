# Write your MySQL query statement below

select a.customer_id, count(case when b.amount is null then 1 end) as count_no_trans
from Visits a left join Transactions b
on a.visit_id = b.visit_id 
group by customer_id
having count_no_trans!=0;
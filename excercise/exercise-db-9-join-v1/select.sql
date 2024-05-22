-- TODO: answer here
select
reports.id as id,
students.fullname as fullname,
students.class as class,
students.status as status,
reports.study as study,
reports.score as score
from reports
inner join students
on students.id = reports.student_id
where reports.score < 70 and students.status='active' order by score asc;
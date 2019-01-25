INNER JOIN --> SELECT * FROM people ORDER BY FIRSTNAME, ID DESC LIMIT 4;
INNER JOIN --> SELECT people.firstname , contacts.email, contacts.mobile FROM people INNER JOIN contacts ON people.id = contacts.id limit 4;
	       SELECT p.firstname , c.email, c.mobile FROM people AS p INNER JOIN contacts AS c ON p.id = c.id; 
CROSS JOIN --> SELECT mobile, firstname, gender,age FROM people CROSS JOIN contacts limit 5;
LEFT OUTER JOIN --> SELECT firstname, age, city, state FROM people LEFT OUTER JOIN addresses ON addresses.id = people.id LIMIT 7;
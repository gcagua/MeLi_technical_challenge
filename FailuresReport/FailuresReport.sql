
	SELECT first_name || ' ' || last_name as customer, count (1) as failures
	FROM customers 
		INNER JOIN campaigns on customers.id = campaigns.customer_id
		INNER JOIN events on campaigns.id = events.campaign_id
	WHERE events.status = 'failure'
	GROUP BY customers.id, customers.first_name, customers.last_name
	HAVING count(1) > 3

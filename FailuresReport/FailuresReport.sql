-- Retrieve customers with more than 3 failed events
-- Concatenates first and last names as "customer" and counts the number of failures

SELECT 
    first_name || ' ' || last_name AS customer,  -- Full name of the customer
    COUNT(1) AS failures                         -- Total number of failure events
FROM customers
    -- Join campaigns to link customers to their campaigns
    INNER JOIN campaigns ON customers.id = campaigns.customer_id
    -- Join events to access the status of each campaign's events
    INNER JOIN events ON campaigns.id = events.campaign_id
WHERE events.status = 'failure'                 -- Only include events with failure status
GROUP BY 
    customers.id,                               -- Group by customer ID (required for aggregation)
    customers.first_name, 
    customers.last_name
HAVING COUNT(1) > 3;                    -- Only show customers with more than 3 failures



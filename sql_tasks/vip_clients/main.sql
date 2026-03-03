-- Напиши запрос, который выведет список пользователей (их id и name), которые соответствуют следующим условиям:
-- 1) Они совершили более 3-х заказов со статусом 'completed'.
-- 2) Общая сумма их покупок (сумма всех price * quantity в успешных заказах) превышает 10 000.
-- 3) Для каждого такого пользователя выведи дату его самого последнего успешного заказа и название самого дорогого товара, который он когда-либо покупал (независимо от того,  в каком заказе был этот товар).
WITH user_stats AS (
    SELECT 
        u.id AS user_id,
        u.name AS user_name,
        SUM(oi.price * oi.quantity) AS total_spent,
        MAX(o.created_at) AS last_order_date,
        COUNT(DISTINCT o.id) AS completed_orders_count
    FROM users u
    JOIN orders o ON u.id = o.user_id
    JOIN order_items oi ON o.id = oi.order_id
    WHERE o.status = 'completed'
    GROUP BY u.id, u.name
    HAVING COUNT(DISTINCT o.id) > 3 AND SUM(oi.price * oi.quantity) > 10000
),
product_ranking AS (
    SELECT 
        o.user_id,
        p.name AS product_name,
        ROW_NUMBER() OVER(PARTITION BY o.user_id ORDER BY oi.price DESC, p.id) as rn
    FROM order_items oi
    JOIN orders o ON oi.order_id = o.id
    JOIN products p ON oi.product_id = p.id
    WHERE o.status = 'completed'
)
SELECT 
    us.user_id,
    us.user_name,
    us.total_spent,
    us.last_order_date,
    pr.product_name AS most_expensive_product_name
FROM user_stats us
JOIN product_ranking pr ON us.user_id = pr.user_id
WHERE pr.rn = 1;
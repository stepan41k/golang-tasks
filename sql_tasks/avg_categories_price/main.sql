-- products (id, name, category_id, price)
-- categories (id, name)


-- Задание:
-- Выведи список названий категорий (category_name) и среднюю цену товара в этой категории (avg_price).
-- Условие: включи в результат только те категории, в которых есть хотя бы один товар дороже 5000.

SELECT c.id, c.name, AVG(p.price) AS avg_price
FROM categories c
INNER JOIN products p ON c.id = p.category_id
GROUP BY c.id, c.name
HAVING MAX(p.price) > 5000;